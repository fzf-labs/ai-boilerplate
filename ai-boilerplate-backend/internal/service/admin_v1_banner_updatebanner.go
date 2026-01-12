package service

import (
	"context"

	"github.com/dromara/carbon/v2"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/goutil/timeutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// UpdateBanner 通用-轮播图-更新一条数据
func (a *AdminV1BannerService) UpdateBanner(ctx context.Context, req *pb.UpdateBannerReq) (*pb.UpdateBannerReply, error) {
	resp := &pb.UpdateBannerReply{}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	data, err := a.bannerRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	if data.TenantID != tenantID {
		return nil, pb.ErrorReasonAccountNoDataPermission()
	}
	oldData := a.bannerRepo.DeepCopy(data)
	data.Title = req.GetTitle()
	data.ImageURL = req.GetImageURL()
	data.LinkURL = req.GetLinkURL()
	data.LinkType = req.GetLinkType()
	data.Position = req.GetPosition()
	data.Platform = req.GetPlatform()
	data.Sort = req.GetSort()
	data.Status = int16(req.GetStatus())
	if req.GetStartTime() != "" {
		data.StartTime = timeutil.TimeToSQLNullTime(carbon.Parse(req.GetStartTime()).StdTime())
	}
	if req.GetEndTime() != "" {
		data.EndTime = timeutil.TimeToSQLNullTime(carbon.Parse(req.GetEndTime()).StdTime())
	}
	err = a.bannerRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
