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
	list, p, err := a.membershipBenefitRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total

	// 获取权益类型映射表
	benefitTypeMap := make(map[string]struct {
		Name string
		Desc string
	})
	if len(list) > 0 {
		// 收集所有 benefit_key
		keys := make([]string, 0, len(list))
		for _, v := range list {
			keys = append(keys, v.BenefitKey)
		}
		// 查询权益类型表
		typeParam := &condition.Req{
			Page:     1,
			PageSize: 1000,
			Query: []*condition.QueryParam{
				{
					Field: "benefit_key",
					Value: keys,
					Exp:   condition.IN,
					Logic: condition.AND,
				},
			},
		}
		typeList, _, err := a.membershipBenefitTypeRepo.FindMultiCacheByCondition(ctx, typeParam)
		if err == nil {
			for _, t := range typeList {
				benefitTypeMap[t.BenefitKey] = struct {
					Name string
					Desc string
				}{Name: t.BenefitName, Desc: t.BenefitDesc}
			}
		}

		for _, v := range list {
			info := &pb.MembershipBenefitInfo{
				Id:             v.ID,
				MembershipType: v.MembershipType,
				BenefitKey:     v.BenefitKey,
				BenefitValue:   v.BenefitValue,
				BenefitNum:     v.BenefitNum,
				Sort:           v.Sort,
				Status:         v.Status,
				CreatedAt:      v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:      v.UpdatedAt.Format(time.RFC3339),
			}
			// 从权益类型表获取名称和描述
			if typeInfo, ok := benefitTypeMap[v.BenefitKey]; ok {
				info.BenefitName = typeInfo.Name
				info.BenefitDesc = typeInfo.Desc
			}
			resp.List = append(resp.List, info)
		}
	}
	return resp, nil
}
