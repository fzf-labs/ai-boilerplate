package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetSysPostList 系统-工作岗位-列表数据查询
func (a *AdminV1SysPostService) GetSysPostList(ctx context.Context, req *pb.GetSysPostListReq) (*pb.GetSysPostListReply, error) {
	resp := &pb.GetSysPostListReply{
		Total: 0,
		List:  []*pb.SysPostInfo{},
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
	if req.GetCode() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "code",
			Value: req.GetCode(),
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
		param.Query = append(param.Query,
			&condition.QueryParam{
				Field: "created_at",
				Value: req.GetCreatedAt()[0],
				Exp:   condition.GTE,
				Logic: condition.AND,
			}, &condition.QueryParam{
				Field: "created_at",
				Value: req.GetCreatedAt()[1],
				Exp:   condition.LTE,
				Logic: condition.AND,
			})
	}
	list, p, err := a.sysPostRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.SysPostInfo{
				Id:        v.ID,
				TenantId:  v.TenantID,
				Name:      v.Name,
				Code:      v.Code,
				Remark:    v.Remark,
				Sort:      int32(v.Sort),
				Status:    int32(v.Status),
				CreatedAt: v.CreatedAt.Format(time.RFC3339),
				UpdatedAt: v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
