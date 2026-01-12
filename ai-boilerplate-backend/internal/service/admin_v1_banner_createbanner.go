package service

import (
	"context"

	"github.com/dromara/carbon/v2"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/goutil/timeutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// CreateBanner 通用-轮播图-创建一条数据
func (a *AdminV1BannerService) CreateBanner(ctx context.Context, req *pb.CreateBannerReq) (*pb.CreateBannerReply, error) {
	resp := &pb.CreateBannerReply{}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	data := a.bannerRepo.NewData()
	data.TenantID = tenantID
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
	err := a.bannerRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
