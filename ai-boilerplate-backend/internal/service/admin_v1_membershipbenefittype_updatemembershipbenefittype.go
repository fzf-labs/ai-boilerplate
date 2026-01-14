package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateMembershipBenefitType 会员权益类型表-更新一条数据
func (a *AdminV1MembershipBenefitTypeService) UpdateMembershipBenefitType(ctx context.Context, req *pb.UpdateMembershipBenefitTypeReq) (*pb.UpdateMembershipBenefitTypeReply, error) {
	resp := &pb.UpdateMembershipBenefitTypeReply{}
	data, err := a.membershipBenefitTypeRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.membershipBenefitTypeRepo.DeepCopy(data)
	data.BenefitKey = req.GetBenefitKey()
	data.BenefitName = req.GetBenefitName()
	data.BenefitIcon = req.GetBenefitIcon()
	data.BenefitDesc = req.GetBenefitDesc()
	data.Sort = req.GetSort()
	data.Status = req.GetStatus()
	err = a.membershipBenefitTypeRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
