package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateAiImageRecordStatus AI 绘画表-更新状态
func (a *AdminV1AiImageRecordService) UpdateAiImageRecordStatus(ctx context.Context, req *pb.UpdateAiImageRecordStatusReq) (*pb.UpdateAiImageRecordStatusReply, error) {
	resp := &pb.UpdateAiImageRecordStatusReply{}
	data, err := a.aiImageRecordRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.aiImageRecordRepo.DeepCopy(data)
	data.Status = req.GetStatus()
	err = a.aiImageRecordRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
