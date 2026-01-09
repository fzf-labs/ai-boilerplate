package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// SysAuthLogout Auth-退出
func (a *AdminV1SysAuthService) SysAuthLogout(_ context.Context, _ *pb.SysAuthLogoutReq) (*pb.SysAuthLogoutReply, error) {
	resp := &pb.SysAuthLogoutReply{}
	// TODO
	return resp, nil
}
