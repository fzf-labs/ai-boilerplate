package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// ListBanners 通用-轮播图-列表数据查询
func (a *AppV1BannerService) ListBanners(ctx context.Context, req *pb.ListBannersReq) (*pb.ListBannersReply, error) {
	resp := &pb.ListBannersReply{
		List: []*pb.BannerItem{},
	}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	param := &condition.Req{
		Page:     1,
		PageSize: 1000,
		Query: []*condition.QueryParam{
			{
				Field: "status",
				Value: int16(1),
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
			{
				Field: "position",
				Value: req.GetPosition(),
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
	if tenantID != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "tenant_id",
			Value: tenantID,
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetPlatform() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "platform",
			Value: []string{req.GetPlatform(), "all"},
			Exp:   condition.IN,
			Logic: condition.AND,
		})
	}
	list, _, err := a.bannerRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	now := time.Now()
	for _, v := range list {
		if v.StartTime.Valid && v.StartTime.Time.After(now) {
			continue
		}
		if v.EndTime.Valid && v.EndTime.Time.Before(now) {
			continue
		}
		resp.List = append(resp.List, &pb.BannerItem{
			Id:       v.ID,
			Title:    v.Title,
			ImageURL: v.ImageURL,
			LinkURL:  v.LinkURL,
			LinkType: v.LinkType,
			Position: v.Position,
			Platform: v.Platform,
			Sort:     v.Sort,
		})
	}
	return resp, nil
}
