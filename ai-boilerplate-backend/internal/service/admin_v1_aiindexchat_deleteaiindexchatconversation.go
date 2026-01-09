package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteAiIndexChatConversation AI 聊天对话表-删除一条数据
func (a *AdminV1AiIndexChatService) DeleteAiIndexChatConversation(ctx context.Context, req *pb.DeleteAiIndexChatConversationReq) (*pb.DeleteAiIndexChatConversationReply, error) {
	resp := &pb.DeleteAiIndexChatConversationReply{}
	err := a.aiChatConversationRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
