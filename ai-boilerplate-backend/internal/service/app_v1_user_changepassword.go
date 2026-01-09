package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
)

// ChangePassword 修改密码
func (a *AppV1UserService) ChangePassword(_ context.Context, _ *pb.ChangePasswordReq) (*pb.ChangePasswordReply, error) {
	resp := &pb.ChangePasswordReply{}
	// TODO
	return resp, nil
}
