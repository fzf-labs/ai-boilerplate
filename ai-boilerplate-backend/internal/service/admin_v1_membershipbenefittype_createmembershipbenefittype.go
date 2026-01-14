package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// CreateMembershipBenefitType 会员权益类型表-创建一条数据
func (a *AdminV1MembershipBenefitTypeService) CreateMembershipBenefitType(ctx context.Context, req *pb.CreateMembershipBenefitTypeReq) (*pb.CreateMembershipBenefitTypeReply, error) {
	resp := &pb.CreateMembershipBenefitTypeReply{}
	data := a.membershipBenefitTypeRepo.NewData()
	data.BenefitKey = req.GetBenefitKey()
	data.BenefitName = req.GetBenefitName()
	data.BenefitIcon = req.GetBenefitIcon()
	data.BenefitDesc = req.GetBenefitDesc()
	data.Sort = req.GetSort()
	data.Status = req.GetStatus()
	err := a.membershipBenefitTypeRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
