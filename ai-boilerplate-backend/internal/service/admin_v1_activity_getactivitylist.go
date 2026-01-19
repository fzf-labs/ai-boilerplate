package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/goutil/timeutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetActivityList 内容-活动-列表数据查询
func (a *AdminV1ActivityService) GetActivityList(ctx context.Context, req *pb.GetActivityListReq) (*pb.GetActivityListReply, error) {
	resp := &pb.GetActivityListReply{
		List: []*pb.ActivityInfo{},
	}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)

	list, total, err := a.activityRepo.ListForAdmin(ctx, data.ActivityListFilter{
		TenantID: tenantID,
		Keyword:  req.GetKeyword(),
		Status:   int16(req.GetStatus()),
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
	})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = int32(total)

	for _, v := range list {
		startTime := ""
		if v.StartTime.Valid {
			startTime = timeutil.RFC3339(v.StartTime.Time)
		}
		endTime := ""
		if v.EndTime.Valid {
			endTime = timeutil.RFC3339(v.EndTime.Time)
		}
		resp.List = append(resp.List, &pb.ActivityInfo{
			Id:        v.ID,
			TenantId:  v.TenantID,
			Title:     v.Title,
			ImageURL:  v.ImageURL,
			LinkURL:   v.LinkURL,
			LinkType:  v.LinkType,
			Sort:      v.Sort,
			Status:    int32(v.Status),
			StartTime: startTime,
			EndTime:   endTime,
			CreatedAt: timeutil.RFC3339(v.CreatedAt),
			UpdatedAt: timeutil.RFC3339(v.UpdatedAt),
		})
	}
	return resp, nil
}
