package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// CreateMembershipBenefit 会员权益配置表-创建一条数据
func (a *AdminV1MembershipBenefitService) CreateMembershipBenefit(ctx context.Context, req *pb.CreateMembershipBenefitReq) (*pb.CreateMembershipBenefitReply, error) {
	resp := &pb.CreateMembershipBenefitReply{}
	data := a.membershipBenefitRepo.NewData()
	data.MembershipType = req.GetMembershipType()
	data.BenefitKey = req.GetBenefitKey()
	data.BenefitName = req.GetBenefitName()
	data.BenefitDesc = req.GetBenefitDesc()
	data.BenefitValue = req.GetBenefitValue()
	data.BenefitNum = req.GetBenefitNum()
	data.Sort = req.GetSort()
	data.Status = req.GetStatus()
	err := a.membershipBenefitRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
