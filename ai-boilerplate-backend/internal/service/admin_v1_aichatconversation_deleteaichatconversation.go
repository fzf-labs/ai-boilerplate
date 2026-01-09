package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteAiChatConversation AI 聊天对话表-删除一条数据
func (a *AdminV1AiChatConversationService) DeleteAiChatConversation(ctx context.Context, req *pb.DeleteAiChatConversationReq) (*pb.DeleteAiChatConversationReply, error) {
	resp := &pb.DeleteAiChatConversationReply{}
	err := a.aiChatConversationRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
