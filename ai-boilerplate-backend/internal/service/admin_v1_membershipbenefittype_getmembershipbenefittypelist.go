package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetMembershipBenefitTypeList 会员权益类型表-列表数据查询
func (a *AdminV1MembershipBenefitTypeService) GetMembershipBenefitTypeList(ctx context.Context, req *pb.GetMembershipBenefitTypeListReq) (*pb.GetMembershipBenefitTypeListReply, error) {
	resp := &pb.GetMembershipBenefitTypeListReply{
		Total: 0,
		List:  []*pb.MembershipBenefitTypeInfo{},
	}
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query:    []*condition.QueryParam{},
		Order: []*condition.OrderParam{
			{
				Field: "sort",
				Order: condition.ASC,
			},
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
	list, p, err := a.membershipBenefitTypeRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.MembershipBenefitTypeInfo{
				Id:          v.ID,
				BenefitKey:  v.BenefitKey,
				BenefitName: v.BenefitName,
				BenefitIcon: v.BenefitIcon,
				BenefitDesc: v.BenefitDesc,
				Sort:        v.Sort,
				Status:      v.Status,
				CreatedAt:   v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:   v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
