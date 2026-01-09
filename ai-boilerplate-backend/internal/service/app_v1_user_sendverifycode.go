package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
)

// SendVerifyCode 发送验证码
func (a *AppV1UserService) SendVerifyCode(_ context.Context, _ *pb.SendVerifyCodeReq) (*pb.SendVerifyCodeReply, error) {
	resp := &pb.SendVerifyCodeReply{}
	// TODO
	return resp, nil
}
