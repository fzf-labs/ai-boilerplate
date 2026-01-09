package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteAiIndexPrompt AI 提示词-删除一条数据
func (a *AdminV1AiIndexPromptService) DeleteAiIndexPrompt(ctx context.Context, req *pb.DeleteAiIndexPromptReq) (*pb.DeleteAiIndexPromptReply, error) {
	resp := &pb.DeleteAiIndexPromptReply{}
	err := a.aiPromptRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
