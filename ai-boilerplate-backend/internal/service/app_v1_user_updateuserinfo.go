package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
)

// UpdateUserInfo 更新用户信息
func (a *AppV1UserService) UpdateUserInfo(_ context.Context, _ *pb.UpdateUserInfoReq) (*pb.UpdateUserInfoReply, error) {
	resp := &pb.UpdateUserInfoReply{}
	// TODO
	return resp, nil
}
