package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetMembershipBenefitKeySelect 会员权益配置表-获取权益标识选择器
func (a *AdminV1MembershipBenefitService) GetMembershipBenefitKeySelect(ctx context.Context, _ *pb.GetMembershipBenefitKeySelectReq) (*pb.GetMembershipBenefitKeySelectReply, error) {
	resp := &pb.GetMembershipBenefitKeySelectReply{
		List: []*pb.MembershipBenefitKeySelect{},
	}
	// 从 membership_benefit_type 表读取权益类型
	param := &condition.Req{
		Page:     1,
		PageSize: 1000,
		Query: []*condition.QueryParam{
			{
				Field: "status",
				Value: int32(constant.StatusEnable),
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
		},
		Order: []*condition.OrderParam{
			{
				Field: "sort",
				Order: condition.ASC,
			},
		},
	}
	list, _, err := a.membershipBenefitTypeRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	for _, v := range list {
		resp.List = append(resp.List, &pb.MembershipBenefitKeySelect{
			Key:  v.BenefitKey,
			Name: v.BenefitName,
		})
	}
	return resp, nil
}
