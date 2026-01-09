package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
)

// BindPhone 绑定手机号
func (a *AppV1UserService) BindPhone(_ context.Context, _ *pb.BindPhoneReq) (*pb.BindPhoneReply, error) {
	resp := &pb.BindPhoneReply{}
	// TODO
	return resp, nil
}
