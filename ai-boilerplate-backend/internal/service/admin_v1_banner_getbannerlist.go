package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/timeutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetBannerList 通用-轮播图-列表数据查询
func (a *AdminV1BannerService) GetBannerList(ctx context.Context, req *pb.GetBannerListReq) (*pb.GetBannerListReply, error) {
	resp := &pb.GetBannerListReply{
		Total: 0,
		List:  []*pb.BannerInfo{},
	}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query:    []*condition.QueryParam{},
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
	if req.GetTitle() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "title",
			Value: "%" + req.GetTitle() + "%",
			Exp:   condition.LIKE,
			Logic: condition.AND,
		})
	}
	if req.GetPosition() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "position",
			Value: req.GetPosition(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetPlatform() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "platform",
			Value: req.GetPlatform(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetLinkType() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "link_type",
			Value: req.GetLinkType(),
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
	list, p, err := a.bannerRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = int32(p.Total)
	if len(list) > 0 {
		for _, v := range list {
			startTime := ""
			if v.StartTime.Valid {
				startTime = timeutil.RFC3339(v.StartTime.Time)
			}
			endTime := ""
			if v.EndTime.Valid {
				endTime = timeutil.RFC3339(v.EndTime.Time)
			}
			resp.List = append(resp.List, &pb.BannerInfo{
				Id:        v.ID,
				TenantId:  v.TenantID,
				Title:     v.Title,
				ImageURL:  v.ImageURL,
				LinkURL:   v.LinkURL,
				LinkType:  v.LinkType,
				Position:  v.Position,
				Platform:  v.Platform,
				Sort:      v.Sort,
				Status:    int32(v.Status),
				StartTime: startTime,
				EndTime:   endTime,
				CreatedAt: timeutil.RFC3339(v.CreatedAt),
				UpdatedAt: timeutil.RFC3339(v.UpdatedAt),
			})
		}
	}
	return resp, nil
}
