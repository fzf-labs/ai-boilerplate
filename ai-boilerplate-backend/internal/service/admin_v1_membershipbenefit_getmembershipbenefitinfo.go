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
		BenefitValue:   data.BenefitValue,
		BenefitNum:     data.BenefitNum,
		Sort:           data.Sort,
		Status:         data.Status,
		CreatedAt:      data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      data.UpdatedAt.Format(time.RFC3339),
	}
	// 从权益类型表获取名称和描述
	typeData, err := a.membershipBenefitTypeRepo.FindOneCacheByBenefitKey(ctx, data.BenefitKey)
	if err == nil && typeData != nil {
		resp.Info.BenefitName = typeData.BenefitName
		resp.Info.BenefitDesc = typeData.BenefitDesc
	}
	return resp, nil
}
