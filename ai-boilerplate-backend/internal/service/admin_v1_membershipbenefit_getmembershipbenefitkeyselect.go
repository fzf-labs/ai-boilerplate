package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
)

// GetMembershipBenefitKeySelect 会员权益配置表-获取权益标识选择器
func (a *AdminV1MembershipBenefitService) GetMembershipBenefitKeySelect(_ context.Context, _ *pb.GetMembershipBenefitKeySelectReq) (*pb.GetMembershipBenefitKeySelectReply, error) {
	resp := &pb.GetMembershipBenefitKeySelectReply{
		List: []*pb.MembershipBenefitKeySelect{},
	}
	for _, v := range constant.MembershipBenefitKeyValues() {
		resp.List = append(resp.List, &pb.MembershipBenefitKeySelect{
			Key:  v.String(),
			Name: constant.MembershipBenefitKeyToName[v.String()],
		})
	}
	return resp, nil
}
