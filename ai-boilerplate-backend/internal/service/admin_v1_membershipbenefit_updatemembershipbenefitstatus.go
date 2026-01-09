package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateMembershipBenefitStatus 会员权益配置表-更新状态
func (a *AdminV1MembershipBenefitService) UpdateMembershipBenefitStatus(ctx context.Context, req *pb.UpdateMembershipBenefitStatusReq) (*pb.UpdateMembershipBenefitStatusReply, error) {
	resp := &pb.UpdateMembershipBenefitStatusReply{}
	data, err := a.membershipBenefitRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.membershipBenefitRepo.DeepCopy(data)
	data.Status = req.GetStatus()
	err = a.membershipBenefitRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
