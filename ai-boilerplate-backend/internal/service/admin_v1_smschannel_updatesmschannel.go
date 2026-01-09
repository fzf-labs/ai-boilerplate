package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateSmsChannel 短信渠道-更新一条数据
func (a *AdminV1SmsChannelService) UpdateSmsChannel(ctx context.Context, req *pb.UpdateSmsChannelReq) (*pb.UpdateSmsChannelReply, error) {
	resp := &pb.UpdateSmsChannelReply{}
	data, err := a.smsChannelRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.smsChannelRepo.DeepCopy(data)
	data.Name = req.GetName()
	data.Operator = req.GetOperator()
	data.Remark = req.GetRemark()
	data.APIKey = req.GetAPIKey()
	data.APISecret = req.GetAPISecret()
	data.CallbackURL = req.GetCallbackURL()
	data.Status = int16(req.GetStatus())
	err = a.smsChannelRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
