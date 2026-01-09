package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetMembershipBenefitList 会员权益配置表-列表数据查询
func (a *AdminV1MembershipBenefitService) GetMembershipBenefitList(ctx context.Context, req *pb.GetMembershipBenefitListReq) (*pb.GetMembershipBenefitListReply, error) {
	resp := &pb.GetMembershipBenefitListReply{
		Total: 0,
		List:  []*pb.MembershipBenefitInfo{},
	}
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query: []*condition.QueryParam{
			{
				Field: "membership_type",
				Value: req.GetMembershipType(),
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
		},
		Order: []*condition.OrderParam{
			{
				Field: "created_at",
				Order: condition.DESC,
			},
		},
	}
	if req.GetBenefitKey() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "benefit_key",
			Value: req.GetBenefitKey(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetBenefitName() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "benefit_name",
			Value: "%" + req.GetBenefitName() + "%",
			Exp:   condition.LIKE,
			Logic: condition.AND,
		})
	}
	list, p, err := a.membershipBenefitRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.MembershipBenefitInfo{
				Id:             v.ID,
				MembershipType: v.MembershipType,
				BenefitKey:     v.BenefitKey,
				BenefitName:    v.BenefitName,
				BenefitDesc:    v.BenefitDesc,
				BenefitValue:   v.BenefitValue,
				BenefitNum:     v.BenefitNum,
				Sort:           v.Sort,
				Status:         v.Status,
				CreatedAt:      v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:      v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
