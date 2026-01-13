package service

import (
	"io"
	"strings"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/schema"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
	"github.com/fzf-labs/kratos-contrib/pkg/sse"
	"github.com/go-kratos/kratos/v2/transport/http"
	"golang.org/x/net/context"
)

type aiIndexWriteCompletionsReq struct {
	Prompt          string `json:"prompt"`
	ModelId         string `json:"modelId"`
	Type            int32  `json:"type"`
	OriginalContent string `json:"originalContent"`
	Length          int32  `json:"length"`
	Format          int32  `json:"format"`
	Tone            int32  `json:"tone"`
	Language        int32  `json:"language"`
}

type aiIndexWriteCompletionsReply struct {
	Content string `json:"content"`
}

// AiIndexWriteCompletionsHandler AI 写作-流式返回
func (a *AdminV1AiWriteRecordService) AiIndexWriteCompletionsHandler(ctx http.Context) error {
	var in aiIndexWriteCompletionsReq
	if err := ctx.Bind(&in); err != nil {
		return err
	}
	http.SetOperation(ctx, "/admin.v1.AiWriteRecord/AiIndexWriteCompletions")
	h := ctx.Middleware(func(ctx context.Context, _ interface{}) (interface{}, error) {
		sseWriter, streamCtx, err := sse.NewWriter(ctx)
		if err != nil {
			a.log.Errorf("create sse writer failed: %v", err)
			return nil, err
		}
		if in.Prompt == "" || in.ModelId == "" {
			return nil, pb.ErrorReasonParamError()
		}

		modelRecord, err := a.aiProviderModelRepo.FindOneCacheByID(streamCtx, in.ModelId)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		if modelRecord == nil || modelRecord.ID == "" {
			return nil, pb.ErrorReasonDataRecordNotFound()
		}
		platformRecord, err := a.aiProviderPlatformRepo.FindOneCacheByID(streamCtx, modelRecord.PlatformID)
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

		config := &ark.ChatModelConfig{
			Model:  modelRecord.ModelID,
			APIKey: platformRecord.APIKey,
		}
		if platformRecord.APIURL != "" {
			config.BaseURL = platformRecord.APIURL
		}

		model, err := ark.NewChatModel(streamCtx, config)
		if err != nil {
			a.log.Errorf("create model failed: %v", err)
			_ = sseWriter.WriteError(err)
			return nil, err
		}

		messages := []*schema.Message{schema.UserMessage(in.Prompt)}
		streamResult, err := model.Stream(streamCtx, messages)
		if err != nil {
			a.log.Errorf("generate response failed: %v", err)
			_ = sseWriter.WriteError(err)
			return nil, err
		}

		var fullContent strings.Builder
		for {
			chunk, err := streamResult.Recv()
			if err == io.EOF {
				if fullContent.Len() > 0 {
					tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
					adminID := meta.GetMetadataFromClient(ctx, constant.XMdAdminID)
					data := a.aiWriteRecordRepo.NewData()
					data.TenantID = tenantID
					data.AdminID = adminID
					data.Type = in.Type
					data.Platform = platformRecord.Platform
					data.ModelID = modelRecord.ID
					if modelRecord.ModelName != "" {
						data.Model = modelRecord.ModelName
					} else {
						data.Model = modelRecord.ModelID
					}
					data.Prompt = in.Prompt
					data.GeneratedContent = fullContent.String()
					data.OriginalContent = in.OriginalContent
					data.Length = in.Length
					data.Format = in.Format
					data.Tone = in.Tone
					data.Language = in.Language
					if err := a.aiWriteRecordRepo.CreateOneCache(streamCtx, data); err != nil {
						return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
					}
				}

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
			reply := &aiIndexWriteCompletionsReply{
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
