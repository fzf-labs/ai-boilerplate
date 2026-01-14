package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateMembershipBenefitTypeStatus 会员权益类型表-更新状态
func (a *AdminV1MembershipBenefitTypeService) UpdateMembershipBenefitTypeStatus(ctx context.Context, req *pb.UpdateMembershipBenefitTypeStatusReq) (*pb.UpdateMembershipBenefitTypeStatusReply, error) {
	resp := &pb.UpdateMembershipBenefitTypeStatusReply{}
	data, err := a.membershipBenefitTypeRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.membershipBenefitTypeRepo.DeepCopy(data)
	data.Status = req.GetStatus()
	err = a.membershipBenefitTypeRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
