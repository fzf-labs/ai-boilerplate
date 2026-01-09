package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/jsonutil"
)

// GetWxGzhMaterialList 公众号素材表-列表数据查询
func (a *AdminV1WxGzhMaterialService) GetWxGzhMaterialList(ctx context.Context, req *pb.GetWxGzhMaterialListReq) (*pb.GetWxGzhMaterialListReply, error) {
	resp := &pb.GetWxGzhMaterialListReply{
		Total: 0,
		List:  []*pb.WxGzhMaterialInfo{},
	}
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query: []*condition.QueryParam{
			{
				Field: "app_id",
				Value: req.GetAppId(),
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
			{
				Field: "type",
				Value: req.GetType(),
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
	if req.GetName() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "name",
			Value: "%" + req.GetName() + "%",
			Exp:   condition.LIKE,
			Logic: condition.AND,
		})
	}
	list, p, err := a.wxGzhMaterialRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			tags := make([]string, 0)
			if v.Tags.String() != "" {
				if err := jsonutil.Unmarshal(v.Tags, &tags); err != nil {
					return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
				}
			}
			resp.List = append(resp.List, &pb.WxGzhMaterialInfo{
				Id:          v.ID,
				AppId:       v.AppID,
				Type:        v.Type,
				MediaId:     v.MediaID,
				Tags:        tags,
				UpdateTime:  v.UpdateTime.Format(time.RFC3339),
				Name:        v.Name,
				URL:         v.URL,
				CoverURL:    v.CoverURL,
				Description: v.Description,
				Newcat:      v.Newcat,
				Newsubcat:   v.Newsubcat,
				Vid:         v.Vid,
				CreatedAt:   v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:   v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
