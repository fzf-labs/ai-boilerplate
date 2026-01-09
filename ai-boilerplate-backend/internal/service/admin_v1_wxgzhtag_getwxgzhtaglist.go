package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetWxGzhTagList 公众号标签表-列表数据查询
func (a *AdminV1WxGzhTagService) GetWxGzhTagList(ctx context.Context, req *pb.GetWxGzhTagListReq) (*pb.GetWxGzhTagListReply, error) {
	resp := &pb.GetWxGzhTagListReply{
		Total: 0,
		List:  []*pb.WxGzhTagInfo{},
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
	if req.GetAppId() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "app_id",
			Value: req.GetAppId(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	list, p, err := a.wxGzhTagRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.WxGzhTagInfo{
				Id:        v.ID,
				AppId:     v.AppID,
				TagId:     v.TagID,
				Name:      v.Name,
				Count:     v.Count,
				CreatedAt: v.CreatedAt.Format(time.RFC3339),
				UpdatedAt: v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
