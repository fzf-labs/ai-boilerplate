package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/jsonutil"
)

// GetAiIndexChatConversationItem AI 聊天对话表-单条数据查询
func (a *AdminV1AiIndexChatService) GetAiIndexChatConversationItem(ctx context.Context, req *pb.GetAiIndexChatConversationItemReq) (*pb.GetAiIndexChatConversationItemReply, error) {
	resp := &pb.GetAiIndexChatConversationItemReply{
		Info: &pb.AiIndexChatConversationItem{},
	}
	data, err := a.aiChatConversationRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	promptSetting := &pb.AiIndexChatConversationItem_PromptSetting{}
	err = jsonutil.Unmarshal(data.PromptSetting, promptSetting)
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	modelSetting := &pb.AiIndexChatConversationItem_ModelSetting{}
	err = jsonutil.Unmarshal(data.ModelSetting, modelSetting)
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	knowledgeSetting := &pb.AiIndexChatConversationItem_KnowledgeSetting{}
	err = jsonutil.Unmarshal(data.KnowledgeSetting, knowledgeSetting)
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	mcpSetting := &pb.AiIndexChatConversationItem_McpSetting{}
	err = jsonutil.Unmarshal(data.McpSetting, mcpSetting)
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	resp.Info = &pb.AiIndexChatConversationItem{
		Id:               data.ID,
		AdminId:          data.AdminID,
		Title:            data.Title,
		Pinned:           data.Pinned,
		PinnedTime:       data.PinnedTime.Time.Format(time.RFC3339),
		PromptSetting:    promptSetting,
		ModelSetting:     modelSetting,
		KnowledgeSetting: knowledgeSetting,
		McpSetting:       mcpSetting,
		CreatedAt:        data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
