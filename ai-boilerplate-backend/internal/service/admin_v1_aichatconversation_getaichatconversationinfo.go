package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/jsonutil"
)

// GetAiChatConversationInfo AI 聊天对话表-单条数据查询
func (a *AdminV1AiChatConversationService) GetAiChatConversationInfo(ctx context.Context, req *pb.GetAiChatConversationInfoReq) (*pb.GetAiChatConversationInfoReply, error) {
	resp := &pb.GetAiChatConversationInfoReply{}
	data, err := a.aiChatConversationRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	promptSetting := &pb.AiChatConversationInfo_PromptSetting{}
	err = jsonutil.Unmarshal(data.PromptSetting, promptSetting)
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	modelSetting := &pb.AiChatConversationInfo_ModelSetting{}
	err = jsonutil.Unmarshal(data.ModelSetting, modelSetting)
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	knowledgeSetting := &pb.AiChatConversationInfo_KnowledgeSetting{}
	err = jsonutil.Unmarshal(data.KnowledgeSetting, knowledgeSetting)
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	mcpSetting := &pb.AiChatConversationInfo_McpSetting{}
	err = jsonutil.Unmarshal(data.McpSetting, mcpSetting)
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	resp.Info = &pb.AiChatConversationInfo{
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
