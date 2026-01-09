package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetMembershipBenefitInfo 会员权益配置表-单条数据查询
func (a *AdminV1MembershipBenefitService) GetMembershipBenefitInfo(ctx context.Context, req *pb.GetMembershipBenefitInfoReq) (*pb.GetMembershipBenefitInfoReply, error) {
	resp := &pb.GetMembershipBenefitInfoReply{}
	data, err := a.membershipBenefitRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	resp.Info = &pb.MembershipBenefitInfo{
		Id:             data.ID,
		MembershipType: data.MembershipType,
		BenefitKey:     data.BenefitKey,
		BenefitName:    data.BenefitName,
		BenefitDesc:    data.BenefitDesc,
		BenefitValue:   data.BenefitValue,
		BenefitNum:     data.BenefitNum,
		Sort:           data.Sort,
		Status:         data.Status,
		CreatedAt:      data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
