package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteSmsChannel 短信渠道-删除一条数据
func (a *AdminV1SmsChannelService) DeleteSmsChannel(ctx context.Context, req *pb.DeleteSmsChannelReq) (*pb.DeleteSmsChannelReply, error) {
	resp := &pb.DeleteSmsChannelReply{}
	err := a.smsChannelRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
