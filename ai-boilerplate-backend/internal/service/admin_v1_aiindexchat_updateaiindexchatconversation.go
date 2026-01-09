package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// UpdateAiIndexChatConversation AI 聊天对话表-更新一条数据
func (a *AdminV1AiIndexChatService) UpdateAiIndexChatConversation(ctx context.Context, req *pb.UpdateAiIndexChatConversationReq) (*pb.UpdateAiIndexChatConversationReply, error) {
	resp := &pb.UpdateAiIndexChatConversationReply{}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	adminID := meta.GetMetadataFromClient(ctx, constant.XMdAdminID)
	data, err := a.aiChatConversationRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
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
	oldData := a.aiChatConversationRepo.DeepCopy(data)
	data.TenantID = tenantID
	data.AdminID = adminID
	data.Title = req.GetTitle()
	data.PromptSetting = promptSetting
	data.ModelSetting = modelSetting
	data.KnowledgeSetting = knowledgeSetting
	data.McpSetting = mcpSetting
	err = a.aiChatConversationRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
