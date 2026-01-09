package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetMembershipList 会员类型配置表-列表数据查询
func (a *AdminV1MembershipService) GetMembershipList(ctx context.Context, req *pb.GetMembershipListReq) (*pb.GetMembershipListReply, error) {
	resp := &pb.GetMembershipListReply{
		Total: 0,
		List:  []*pb.MembershipInfo{},
	}
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query:    []*condition.QueryParam{},
		Order: []*condition.OrderParam{
			{
				Field: "created_at",
				Order: condition.DESC,
			},
		},
	}
	if req.GetName() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "name",
			Value: "%" + req.GetName() + "%",
			Exp:   condition.LIKE,
			Logic: condition.AND,
		})
	}
	if req.GetType() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "type",
			Value: req.GetType(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetStatus() != 0 {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "status",
			Value: req.GetStatus(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if len(req.GetCreatedAt()) > 0 {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "created_at",
			Value: req.GetCreatedAt()[0],
			Exp:   condition.GTE,
			Logic: condition.AND,
		})
	}
	if len(req.GetCreatedAt()) > 0 {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "created_at",
			Value: req.GetCreatedAt()[1],
			Exp:   condition.LTE,
			Logic: condition.AND,
		})
	}
	list, p, err := a.membershipRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.MembershipInfo{
				Id:          v.ID,
				Name:        v.Name,
				Type:        v.Type,
				Description: v.Description,
				Sort:        v.Sort,
				Status:      v.Status,
				CreatedAt:   v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:   v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
