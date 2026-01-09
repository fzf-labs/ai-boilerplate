package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteMembershipBenefit 会员权益配置表-删除一条数据
func (a *AdminV1MembershipBenefitService) DeleteMembershipBenefit(ctx context.Context, req *pb.DeleteMembershipBenefitReq) (*pb.DeleteMembershipBenefitReply, error) {
	resp := &pb.DeleteMembershipBenefitReply{}
	err := a.membershipBenefitRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
