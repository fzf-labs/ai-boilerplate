package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
)

// DeleteAccount 注销账号
func (a *AppV1UserService) DeleteAccount(_ context.Context, _ *pb.DeleteAccountReq) (*pb.DeleteAccountReply, error) {
	resp := &pb.DeleteAccountReply{}
	// TODO
	return resp, nil
}
