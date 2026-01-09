package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
)

// GetSmsChannelOperator 短信渠道-运营商
func (a *AdminV1SmsChannelService) GetSmsChannelOperator(_ context.Context, _ *pb.GetSmsChannelOperatorReq) (*pb.GetSmsChannelOperatorReply, error) {
	resp := &pb.GetSmsChannelOperatorReply{
		List: []*pb.SmsChannelOperator{},
	}
	for _, operator := range constant.SmsChannelCodeValues() {
		resp.List = append(resp.List, &pb.SmsChannelOperator{
			Name:     constant.SmsChannelCodeToName[operator.String()],
			Operator: operator.String(),
		})
	}
	return resp, nil
}
