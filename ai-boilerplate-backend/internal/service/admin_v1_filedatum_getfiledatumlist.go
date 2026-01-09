package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetFileDatumList 文件表-列表数据查询
func (a *AdminV1FileDatumService) GetFileDatumList(ctx context.Context, req *pb.GetFileDatumListReq) (*pb.GetFileDatumListReply, error) {
	resp := &pb.GetFileDatumListReply{
		Total: 0,
		List:  []*pb.FileDatumInfo{},
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
	if req.GetConfigId() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "config_id",
			Value: req.GetConfigId(),
			Exp:   condition.EQ,
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
	if req.GetPath() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "path",
			Value: "%" + req.GetPath() + "%",
			Exp:   condition.LIKE,
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
	list, p, err := a.fileDatumRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.FileDatumInfo{
				Id:        v.ID,
				ConfigId:  v.ConfigID,
				Name:      v.Name,
				Path:      v.Path,
				URL:       v.URL,
				Ext:       v.Ext,
				Size:      v.Size,
				Status:    v.Status,
				CreatedAt: timeutil.RFC3339(v.CreatedAt),
				UpdatedAt: timeutil.RFC3339(v.UpdatedAt),
			})
		}
	}
	return resp, nil
}
