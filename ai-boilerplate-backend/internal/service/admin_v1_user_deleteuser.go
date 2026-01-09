package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteUser 用户表-删除一条数据
func (a *AdminV1UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.DeleteUserReply, error) {
	resp := &pb.DeleteUserReply{}
	err := a.userRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
