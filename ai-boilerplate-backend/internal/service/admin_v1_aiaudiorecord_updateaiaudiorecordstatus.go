package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateAiAudioRecordStatus AI 音乐表-更新状态
func (a *AdminV1AiAudioRecordService) UpdateAiAudioRecordStatus(ctx context.Context, req *pb.UpdateAiAudioRecordStatusReq) (*pb.UpdateAiAudioRecordStatusReply, error) {
	resp := &pb.UpdateAiAudioRecordStatusReply{}
	data, err := a.aiAudioRecordRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.aiAudioRecordRepo.DeepCopy(data)
	data.Status = req.GetStatus()
	err = a.aiAudioRecordRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
