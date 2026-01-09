package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteMembership 会员类型配置表-删除一条数据
func (a *AdminV1MembershipService) DeleteMembership(ctx context.Context, req *pb.DeleteMembershipReq) (*pb.DeleteMembershipReply, error) {
	resp := &pb.DeleteMembershipReply{}
	err := a.membershipRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
