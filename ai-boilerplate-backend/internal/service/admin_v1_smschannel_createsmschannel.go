package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// CreateSmsChannel 短信渠道-创建一条数据
func (a *AdminV1SmsChannelService) CreateSmsChannel(ctx context.Context, req *pb.CreateSmsChannelReq) (*pb.CreateSmsChannelReply, error) {
	resp := &pb.CreateSmsChannelReply{}
	data := a.smsChannelRepo.NewData()
	data.Name = req.GetName()
	data.Operator = req.GetOperator()
	data.Remark = req.GetRemark()
	data.APIKey = req.GetAPIKey()
	data.APISecret = req.GetAPISecret()
	data.CallbackURL = req.GetCallbackURL()
	data.Status = int16(req.GetStatus())
	err := a.smsChannelRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
