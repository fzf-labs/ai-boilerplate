package service

import (
	"io"
	"strings"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/schema"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/fzf-labs/kratos-contrib/meta"
	"github.com/fzf-labs/kratos-contrib/pkg/sse"
	"github.com/go-kratos/kratos/v2/transport/http"
	"golang.org/x/net/context"
)

// AiIndexChatCompletionsHandler AI 聊天-聊天 ChatCompletions格式 (SSE 流式返回)
func (a *AdminV1AiIndexChatService) AiIndexChatCompletionsHandler(ctx http.Context) error {
	var in pb.AiIndexChatCompletionsReq
	if err := ctx.Bind(&in); err != nil {
		return err
	}
	http.SetOperation(ctx, "/admin.v1.AiIndexChat/AiIndexChatCompletions")
	h := ctx.Middleware(func(ctx context.Context, _ interface{}) (interface{}, error) {
		// 创建 SSE Writer
		sseWriter, streamCtx, err := sse.NewWriter(ctx)
		if err != nil {
			a.log.Errorf("create sse writer failed: %v", err)
			return nil, err
		}

		if in.GetConversationId() == "" {
			return nil, pb.ErrorReasonParamError()
		}

		conversation, err := a.aiChatConversationRepo.FindOneCacheByID(streamCtx, in.GetConversationId())
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		if conversation == nil || conversation.ID == "" {
			return nil, pb.ErrorReasonDataRecordNotFound()
		}
		tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
		adminID := meta.GetMetadataFromClient(ctx, constant.XMdAdminID)
		if tenantID != "" && conversation.TenantID != tenantID {
			return nil, pb.ErrorReasonDataRecordNotFound()
		}
		if adminID != "" && conversation.AdminID != adminID {
			return nil, pb.ErrorReasonDataRecordNotFound()
		}

		promptSetting := &pb.AiIndexChatConversationItem_PromptSetting{}
		if len(conversation.PromptSetting) > 0 {
			if err := jsonutil.Unmarshal(conversation.PromptSetting, promptSetting); err != nil {
				return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
			}
		}
		modelSetting := &pb.AiIndexChatConversationItem_ModelSetting{}
		if len(conversation.ModelSetting) > 0 {
			if err := jsonutil.Unmarshal(conversation.ModelSetting, modelSetting); err != nil {
				return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
			}
		}

		modelRecordID := modelSetting.GetModelId()
		if modelRecordID == "" {
			return nil, pb.ErrorReasonParamError()
		}
		modelRecord, err := a.aiProviderModelRepo.FindOneCacheByID(streamCtx, modelRecordID)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		if modelRecord == nil || modelRecord.ID == "" {
			return nil, pb.ErrorReasonDataRecordNotFound()
		}
		platformID := modelRecord.PlatformID
		if platformID == "" {
			platformID = modelSetting.GetPlatformId()
		}
		if platformID == "" {
			return nil, pb.ErrorReasonParamError()
		}
		platformRecord, err := a.aiProviderPlatformRepo.FindOneCacheByID(streamCtx, platformID)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		if platformRecord == nil || platformRecord.ID == "" {
			return nil, pb.ErrorReasonDataRecordNotFound()
		}
		if platformRecord.APIKey == "" {
			return nil, pb.ErrorReasonParamError()
		}
		if strings.ToLower(platformRecord.Platform) != "ark" {
			return nil, pb.ErrorReasonParamError()
		}

		systemPrompt := strings.TrimSpace(promptSetting.GetPrompt())
		if systemPrompt == "" {
			systemPrompt = strings.TrimSpace(strings.Join([]string{
				strings.TrimSpace(promptSetting.GetName()),
				strings.TrimSpace(promptSetting.GetDesc()),
			}, "\n"))
		}

		messages := buildChatMessages(systemPrompt, in.GetMessages())
		if len(messages) == 0 {
			return nil, pb.ErrorReasonParamError()
		}

		config := &ark.ChatModelConfig{
			Model:  modelRecord.ModelID,
			APIKey: platformRecord.APIKey,
		}
		if platformRecord.APIURL != "" {
			config.BaseURL = platformRecord.APIURL
		}
		if modelSetting.GetTemperature() >= 0 {
			temperature := float32(modelSetting.GetTemperature())
			config.Temperature = &temperature
		}
		if modelSetting.GetTopP() >= 0 {
			topP := float32(modelSetting.GetTopP())
			config.TopP = &topP
		}
		if modelSetting.GetMaxTokens() > 0 {
			maxTokens := int(modelSetting.GetMaxTokens())
			config.MaxTokens = &maxTokens
		}

		model, err := ark.NewChatModel(streamCtx, config)
		if err != nil {
			a.log.Errorf("create model failed: %v", err)
			_ = sseWriter.WriteError(err)
			return nil, err
		}

		userContent := lastUserContent(in.GetMessages())
		if userContent != "" {
			messageData := a.aiChatMessageRepo.NewData()
			messageData.TenantID = conversation.TenantID
			messageData.AdminID = conversation.AdminID
			messageData.ConversationID = conversation.ID
			messageData.Type = "user"
			messageData.ModelID = modelRecord.ID
			if modelRecord.ModelName != "" {
				messageData.Model = modelRecord.ModelName
			} else {
				messageData.Model = modelRecord.ModelID
			}
			messageData.Content = userContent
			messageData.UseContext = len(in.GetMessages()) > 1
			if err := a.aiChatMessageRepo.CreateOneCache(streamCtx, messageData); err != nil {
				return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
			}
		}

		// 流式生成回答
		streamResult, err := model.Stream(streamCtx, messages)
		if err != nil {
			a.log.Errorf("generate response failed: %v", err)
			_ = sseWriter.WriteError(err)
			return nil, err
		}

		var fullContent strings.Builder
		// 流式发送每个 chunk (SSE 格式)
		for {
			chunk, err := streamResult.Recv()
			if err == io.EOF {
				if fullContent.Len() > 0 {
					messageData := a.aiChatMessageRepo.NewData()
					messageData.TenantID = conversation.TenantID
					messageData.AdminID = conversation.AdminID
					messageData.ConversationID = conversation.ID
					messageData.Type = "assistant"
					messageData.ModelID = modelRecord.ID
					if modelRecord.ModelName != "" {
						messageData.Model = modelRecord.ModelName
					} else {
						messageData.Model = modelRecord.ModelID
					}
					messageData.Content = fullContent.String()
					messageData.UseContext = len(in.GetMessages()) > 1
					if err := a.aiChatMessageRepo.CreateOneCache(streamCtx, messageData); err != nil {
						return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
					}
				}

				// 发送结束标记
				if writeErr := sseWriter.WriteDone(); writeErr != nil {
					a.log.Errorf("write done failed: %v", writeErr)
				}
				return nil, nil
			}
			if err != nil {
				a.log.Errorf("receive response failed: %v", err)
				_ = sseWriter.WriteError(err)
				return nil, err
			}

			if chunk.Content != "" {
				fullContent.WriteString(chunk.Content)
			}

			// 构造 SSE 响应数据并发送
			reply := &pb.AiIndexChatCompletionsReply{
				Content: chunk.Content,
			}
			if err := sseWriter.WriteEvent(reply); err != nil {
				a.log.Errorf("write event failed: %v", err)
				return nil, err
			}
		}
	})
	_, err := h(ctx, &in)
	if err != nil {
		return err
	}
	return nil
}

func buildChatMessages(systemPrompt string, reqMessages []*pb.AiIndexChatCompletionsReq_Message) []*schema.Message {
	messages := make([]*schema.Message, 0, len(reqMessages)+1)
	if strings.TrimSpace(systemPrompt) != "" {
		messages = append(messages, schema.SystemMessage(systemPrompt))
	}
	for _, item := range reqMessages {
		content := strings.TrimSpace(item.GetText())
		if content == "" {
			continue
		}
		switch strings.ToLower(item.GetType()) {
		case "system":
			messages = append(messages, schema.SystemMessage(content))
		case "assistant":
			messages = append(messages, schema.AssistantMessage(content, nil))
		default:
			messages = append(messages, schema.UserMessage(content))
		}
	}
	return messages
}

func lastUserContent(reqMessages []*pb.AiIndexChatCompletionsReq_Message) string {
	for i := len(reqMessages) - 1; i >= 0; i-- {
		item := reqMessages[i]
		if strings.ToLower(item.GetType()) == "user" {
			return strings.TrimSpace(item.GetText())
		}
	}
	if len(reqMessages) > 0 {
		return strings.TrimSpace(reqMessages[len(reqMessages)-1].GetText())
	}
	return ""
}
