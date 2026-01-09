package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateAiVideoRecordStatus AI 视频表-更新状态
func (a *AdminV1AiVideoRecordService) UpdateAiVideoRecordStatus(ctx context.Context, req *pb.UpdateAiVideoRecordStatusReq) (*pb.UpdateAiVideoRecordStatusReply, error) {
	resp := &pb.UpdateAiVideoRecordStatusReply{}
	data, err := a.aiVideoRecordRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.aiVideoRecordRepo.DeepCopy(data)
	data.Status = req.GetStatus()
	err = a.aiVideoRecordRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
