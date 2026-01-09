package service

import (
	"io"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
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

		// TODO: 从请求中获取实际的对话信息
		// 1. 获取或创建对话
		// conversationID := req.ConversationId
		// conversation, err := a.aiChatConversationRepo.Get(ctx, conversationID)

		// 2. 构建历史消息
		// messages, err := a.buildChatMessages(ctx, conversation, req)

		// 临时示例：构建简单的提示模板
		template := prompt.FromMessages(schema.FString,
			schema.SystemMessage("你是一个{role}。你需要用{style}的语气回答问题。你的目标是帮助程序员保持积极乐观的心态，提供技术建议的同时也要关注他们的心理健康。"),
			schema.MessagesPlaceholder("chat_history", true),
			schema.UserMessage("问题: {question}"),
		)

		// 使用模板生成消息
		messages, err := template.Format(streamCtx, map[string]any{
			"role":     "程序员鼓励师",
			"style":    "积极、温暖且专业",
			"question": "我的代码一直报错，感觉好沮丧，该怎么办？",
			"chat_history": []*schema.Message{
				schema.UserMessage("你好"),
				schema.AssistantMessage("嘿！我是你的程序员鼓励师！记住，每个优秀的程序员都是从 Debug 中成长起来的。有什么我可以帮你的吗？", nil),
				schema.UserMessage("我觉得自己写的代码太烂了"),
				schema.AssistantMessage("每个程序员都经历过这个阶段！重要的是你在不断学习和进步。让我们一起看看代码，我相信通过重构和优化，它会变得更好。记住，Rome wasn't built in a day，代码质量是通过持续改进来提升的。", nil),
			},
		})
		if err != nil {
			a.log.Errorf("format template failed: %v", err)
			_ = sseWriter.WriteError(err)
			return nil, err
		}

		// TODO: 从配置或数据库中获取模型配置
		// modelConfig := conversation.ModelSetting
		// 临时使用硬编码配置
		model, err := ark.NewChatModel(streamCtx, &ark.ChatModelConfig{
			Model:  "doubao-seed-1-6-251015",
			APIKey: "cb3074a7-058c-4a86-a1c2-3f3e2b570768",
		})
		if err != nil {
			a.log.Errorf("create model failed: %v", err)
			_ = sseWriter.WriteError(err)
			return nil, err
		}

		// 流式生成回答
		streamResult, err := model.Stream(streamCtx, messages)
		if err != nil {
			a.log.Errorf("generate response failed: %v", err)
			_ = sseWriter.WriteError(err)
			return nil, err
		}

		// TODO: 准备保存消息到数据库
		// var fullContent strings.Builder

		// 流式发送每个 chunk (SSE 格式)
		for {
			chunk, err := streamResult.Recv()
			if err == io.EOF {
				// TODO: 保存完整的消息到数据库
				// a.aiChatMessageRepo.Create(ctx, &model.AiChatMessage{
				//     ConversationId: conversationID,
				//     Content: fullContent.String(),
				//     ...
				// })

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

			// TODO: 累积完整内容
			// fullContent.WriteString(chunk.Content)

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
