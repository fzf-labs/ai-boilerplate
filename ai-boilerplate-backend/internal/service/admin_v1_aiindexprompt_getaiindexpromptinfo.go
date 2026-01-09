package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetAiIndexPromptInfo AI 提示词-单条数据查询
func (a *AdminV1AiIndexPromptService) GetAiIndexPromptInfo(ctx context.Context, req *pb.GetAiIndexPromptInfoReq) (*pb.GetAiIndexPromptInfoReply, error) {
	resp := &pb.GetAiIndexPromptInfoReply{}
	data, err := a.aiPromptRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.AiIndexPromptInfo{
		Id:        data.ID,
		TenantId:  data.TenantID,
		AdminId:   data.AdminID,
		Name:      data.Name,
		Desc:      data.Desc,
		Prompt:    data.Prompt,
		Sort:      data.Sort,
		CreatedAt: data.CreatedAt.Format(time.RFC3339),
		UpdatedAt: data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
