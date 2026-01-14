package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetMembershipBenefitTypeInfo 会员权益类型表-单条数据查询
func (a *AdminV1MembershipBenefitTypeService) GetMembershipBenefitTypeInfo(ctx context.Context, req *pb.GetMembershipBenefitTypeInfoReq) (*pb.GetMembershipBenefitTypeInfoReply, error) {
	resp := &pb.GetMembershipBenefitTypeInfoReply{}
	data, err := a.membershipBenefitTypeRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	resp.Info = &pb.MembershipBenefitTypeInfo{
		Id:          data.ID,
		BenefitKey:  data.BenefitKey,
		BenefitName: data.BenefitName,
		BenefitIcon: data.BenefitIcon,
		BenefitDesc: data.BenefitDesc,
		Sort:        data.Sort,
		Status:      data.Status,
		CreatedAt:   data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
