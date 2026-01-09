package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateSmsChannelStatus 短信渠道-更新状态
func (a *AdminV1SmsChannelService) UpdateSmsChannelStatus(ctx context.Context, req *pb.UpdateSmsChannelStatusReq) (*pb.UpdateSmsChannelStatusReply, error) {
	resp := &pb.UpdateSmsChannelStatusReply{}
	data, err := a.smsChannelRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.smsChannelRepo.DeepCopy(data)
	data.Status = int16(req.GetStatus())
	err = a.smsChannelRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
