package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteMembershipBenefitType 会员权益类型表-删除一条数据
func (a *AdminV1MembershipBenefitTypeService) DeleteMembershipBenefitType(ctx context.Context, req *pb.DeleteMembershipBenefitTypeReq) (*pb.DeleteMembershipBenefitTypeReply, error) {
	resp := &pb.DeleteMembershipBenefitTypeReply{}
	err := a.membershipBenefitTypeRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
