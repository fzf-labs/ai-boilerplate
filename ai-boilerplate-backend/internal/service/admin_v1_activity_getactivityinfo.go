package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/goutil/timeutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetActivityInfo 内容-活动-单条数据查询
func (a *AdminV1ActivityService) GetActivityInfo(ctx context.Context, req *pb.GetActivityInfoReq) (*pb.GetActivityInfoReply, error) {
	resp := &pb.GetActivityInfoReply{}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)

	data, err := a.activityRepo.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	if data.TenantID != tenantID {
		return nil, pb.ErrorReasonAccountNoDataPermission()
	}

	startTime := ""
	if data.StartTime.Valid {
		startTime = timeutil.RFC3339(data.StartTime.Time)
	}
	endTime := ""
	if data.EndTime.Valid {
		endTime = timeutil.RFC3339(data.EndTime.Time)
	}
	resp.Info = &pb.ActivityInfo{
		Id:        data.ID,
		TenantId:  data.TenantID,
		Title:     data.Title,
		ImageURL:  data.ImageURL,
		LinkURL:   data.LinkURL,
		LinkType:  data.LinkType,
		Sort:      data.Sort,
		Status:    int32(data.Status),
		StartTime: startTime,
		EndTime:   endTime,
		CreatedAt: timeutil.RFC3339(data.CreatedAt),
		UpdatedAt: timeutil.RFC3339(data.UpdatedAt),
	}
	return resp, nil
}
