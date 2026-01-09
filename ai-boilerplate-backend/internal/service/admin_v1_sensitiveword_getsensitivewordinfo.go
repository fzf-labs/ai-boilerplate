package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetSensitiveWordInfo 敏感词-单条数据查询
func (a *AdminV1SensitiveWordService) GetSensitiveWordInfo(ctx context.Context, req *pb.GetSensitiveWordInfoReq) (*pb.GetSensitiveWordInfoReply, error) {
	resp := &pb.GetSensitiveWordInfoReply{}
	data, err := a.sensitiveWordRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.SensitiveWordInfo{
		Id:        data.ID,
		Word:      data.Word,
		Lab:       data.Lab,
		CreatedAt: data.CreatedAt.Format(time.RFC3339),
		UpdatedAt: data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
