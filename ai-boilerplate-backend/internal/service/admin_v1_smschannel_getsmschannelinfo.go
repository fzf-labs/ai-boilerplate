package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetSmsChannelInfo 短信渠道-单条数据查询
func (a *AdminV1SmsChannelService) GetSmsChannelInfo(ctx context.Context, req *pb.GetSmsChannelInfoReq) (*pb.GetSmsChannelInfoReply, error) {
	resp := &pb.GetSmsChannelInfoReply{}
	data, err := a.smsChannelRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = smsChannelInfoFromModel(data)
	return resp, nil
}
