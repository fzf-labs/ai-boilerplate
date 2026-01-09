package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetSelfAppList 自应用信息表-列表数据查询
func (a *AdminV1SelfAppService) GetSelfAppList(ctx context.Context, req *pb.GetSelfAppListReq) (*pb.GetSelfAppListReply, error) {
	resp := &pb.GetSelfAppListReply{
		Total: 0,
		List:  []*pb.SelfAppInfo{},
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
	if req.GetPackageName() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "package_name",
			Value: "%" + req.GetPackageName() + "%",
			Exp:   condition.LIKE,
			Logic: condition.AND,
		})
	}
	if req.GetName() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "name",
			Value: "%" + req.GetName() + "%",
			Exp:   condition.LIKE,
			Logic: condition.AND,
		})
	}
	if req.GetPlatform() != 0 {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "platform",
			Value: req.GetPlatform(),
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
	list, p, err := a.selfAppRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.SelfAppInfo{
				Id:          v.ID,
				PackageName: v.PackageName,
				Name:        v.Name,
				Description: v.Description,
				Status:      v.Status,
				CreatedAt:   v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:   v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
