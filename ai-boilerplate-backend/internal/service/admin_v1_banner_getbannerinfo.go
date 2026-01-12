package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetBannerInfo 通用-轮播图-单条数据查询
func (a *AdminV1BannerService) GetBannerInfo(ctx context.Context, req *pb.GetBannerInfoReq) (*pb.GetBannerInfoReply, error) {
	resp := &pb.GetBannerInfoReply{}
	data, err := a.bannerRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	startTime := ""
	if data.StartTime.Valid {
		startTime = timeutil.RFC3339(data.StartTime.Time)
	}
	endTime := ""
	if data.EndTime.Valid {
		endTime = timeutil.RFC3339(data.EndTime.Time)
	}
	resp.Info = &pb.BannerInfo{
		Id:        data.ID,
		TenantId:  data.TenantID,
		Title:     data.Title,
		ImageURL:  data.ImageURL,
		LinkURL:   data.LinkURL,
		LinkType:  data.LinkType,
		Position:  data.Position,
		Platform:  data.Platform,
		Sort:      data.Sort,
		Status:    int32(data.Status),
		StartTime: startTime,
		EndTime:   endTime,
		CreatedAt: timeutil.RFC3339(data.CreatedAt),
		UpdatedAt: timeutil.RFC3339(data.UpdatedAt),
	}
	return resp, nil
}
