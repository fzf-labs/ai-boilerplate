package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteUserMembership 用户会员关系表-删除一条数据
func (a *AdminV1UserMembershipService) DeleteUserMembership(ctx context.Context, req *pb.DeleteUserMembershipReq) (*pb.DeleteUserMembershipReply, error) {
	resp := &pb.DeleteUserMembershipReply{}
	err := a.userMembershipRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
