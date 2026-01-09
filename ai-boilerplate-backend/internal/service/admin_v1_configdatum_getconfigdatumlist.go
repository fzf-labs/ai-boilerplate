package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetConfigDatumList 配置管理-列表数据查询
func (a *AdminV1ConfigDatumService) GetConfigDatumList(ctx context.Context, req *pb.GetConfigDatumListReq) (*pb.GetConfigDatumListReply, error) {
	resp := &pb.GetConfigDatumListReply{
		Total: 0,
		List:  []*pb.ConfigDatumInfo{},
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
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetKey() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "key",
			Value: "%" + req.GetKey() + "%",
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if len(req.GetCreatedAt()) > 0 {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "created_at",
			Value: req.GetCreatedAt()[0],
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "created_at",
			Value: req.GetCreatedAt()[1],
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	list, p, err := a.configDatumRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.ConfigDatumInfo{
				Id:        v.ID,
				Name:      v.Name,
				Key:       v.Key,
				Value:     string(v.Value),
				Remark:    v.Remark,
				Status:    v.Status,
				CreatedAt: v.CreatedAt.Format(time.RFC3339),
				UpdatedAt: v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
