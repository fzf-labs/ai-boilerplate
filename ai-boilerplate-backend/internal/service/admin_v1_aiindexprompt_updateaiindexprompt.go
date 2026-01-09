package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateAiIndexPrompt AI 提示词-更新一条数据
func (a *AdminV1AiIndexPromptService) UpdateAiIndexPrompt(ctx context.Context, req *pb.UpdateAiIndexPromptReq) (*pb.UpdateAiIndexPromptReply, error) {
	resp := &pb.UpdateAiIndexPromptReply{}
	data, err := a.aiPromptRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.aiPromptRepo.DeepCopy(data)
	data.Name = req.GetName()
	data.Desc = req.GetDesc()
	data.Prompt = req.GetPrompt()
	data.Sort = req.GetSort()
	err = a.aiPromptRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
