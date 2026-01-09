package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/jsonutil"
)

// GetFileConfigList 文件配置表-列表数据查询
func (a *AdminV1FileConfigService) GetFileConfigList(ctx context.Context, req *pb.GetFileConfigListReq) (*pb.GetFileConfigListReply, error) {
	resp := &pb.GetFileConfigListReply{
		Total: 0,
		List:  []*pb.FileConfigInfo{},
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
			Value: req.GetName(),
			Exp:   condition.LIKE,
			Logic: condition.AND,
		})
	}
	if req.GetStorage() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "storage",
			Value: req.GetStorage(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	list, p, err := a.fileConfigRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			config := &pb.StorageConfig{}
			err = jsonutil.Unmarshal(v.Config, config)
			if err != nil {
				return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
			}
			resp.List = append(resp.List, &pb.FileConfigInfo{
				Id:        v.ID,
				Name:      v.Name,
				Storage:   v.Storage,
				Remark:    v.Remark,
				Master:    v.Master,
				Config:    config,
				CreatedAt: v.CreatedAt.Format(time.RFC3339),
				UpdatedAt: v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
