package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// CreateAiIndexPrompt AI 提示词-创建一条数据
func (a *AdminV1AiIndexPromptService) CreateAiIndexPrompt(ctx context.Context, req *pb.CreateAiIndexPromptReq) (*pb.CreateAiIndexPromptReply, error) {
	resp := &pb.CreateAiIndexPromptReply{}
	data := a.aiPromptRepo.NewData()
	data.Name = req.GetName()
	data.Desc = req.GetDesc()
	data.Prompt = req.GetPrompt()
	data.Sort = req.GetSort()
	err := a.aiPromptRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
