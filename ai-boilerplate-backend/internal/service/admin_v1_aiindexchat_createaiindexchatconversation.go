package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// CreateAiIndexChatConversation AI 聊天对话表-创建一条数据
func (a *AdminV1AiIndexChatService) CreateAiIndexChatConversation(ctx context.Context, req *pb.CreateAiIndexChatConversationReq) (*pb.CreateAiIndexChatConversationReply, error) {
	resp := &pb.CreateAiIndexChatConversationReply{}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	adminID := meta.GetMetadataFromClient(ctx, constant.XMdAdminID)
	promptSetting, err := jsonutil.Marshal(req.GetPromptSetting())
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	modelSetting, err := jsonutil.Marshal(req.GetModelSetting())
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	knowledgeSetting, err := jsonutil.Marshal(req.GetKnowledgeSetting())
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	mcpSetting, err := jsonutil.Marshal(req.GetMcpSetting())
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	data := a.aiChatConversationRepo.NewData()
	data.TenantID = tenantID
	data.AdminID = adminID
	data.Title = req.GetTitle()
	data.PromptSetting = promptSetting
	data.ModelSetting = modelSetting
	data.KnowledgeSetting = knowledgeSetting
	data.McpSetting = mcpSetting
	err = a.aiChatConversationRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
