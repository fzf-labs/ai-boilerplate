package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// ListActivities 内容-活动-列表数据查询
func (a *AppV1ActivityService) ListActivities(ctx context.Context, req *pb.ListActivitiesReq) (*pb.ListActivitiesReply, error) {
	resp := &pb.ListActivitiesReply{
		List:     []*pb.ActivityItem{},
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
	}

	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	list, total, err := a.activityRepo.ListForApp(ctx, data.AppActivityListFilter{
		TenantID: tenantID,
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Now:      time.Now(),
	})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = total

	for _, v := range list {
		resp.List = append(resp.List, &pb.ActivityItem{
			Id:       v.ID,
			Title:    v.Title,
			ImageURL: v.ImageURL,
			LinkURL:  v.LinkURL,
			LinkType: v.LinkType,
			Sort:     v.Sort,
		})
	}
	return resp, nil
}
