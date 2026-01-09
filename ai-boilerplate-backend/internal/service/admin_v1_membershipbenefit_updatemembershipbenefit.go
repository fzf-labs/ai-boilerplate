package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateMembershipBenefit 会员权益配置表-更新一条数据
func (a *AdminV1MembershipBenefitService) UpdateMembershipBenefit(ctx context.Context, req *pb.UpdateMembershipBenefitReq) (*pb.UpdateMembershipBenefitReply, error) {
	resp := &pb.UpdateMembershipBenefitReply{}
	data, err := a.membershipBenefitRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.membershipBenefitRepo.DeepCopy(data)
	data.MembershipType = req.GetMembershipType()
	data.BenefitKey = req.GetBenefitKey()
	data.BenefitName = req.GetBenefitName()
	data.BenefitDesc = req.GetBenefitDesc()
	data.BenefitValue = req.GetBenefitValue()
	data.BenefitNum = req.GetBenefitNum()
	data.Sort = req.GetSort()
	data.Status = req.GetStatus()
	err = a.membershipBenefitRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
