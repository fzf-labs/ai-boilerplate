package service

import (
	"context"

	"github.com/dromara/carbon/v2"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/goutil/timeutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// CreateActivity 内容-活动-创建一条数据
func (a *AdminV1ActivityService) CreateActivity(ctx context.Context, req *pb.CreateActivityReq) (*pb.CreateActivityReply, error) {
	resp := &pb.CreateActivityReply{}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)

	data := a.activityRepo.NewData()
	data.TenantID = tenantID
	data.Title = req.GetTitle()
	data.ImageURL = req.GetImageURL()
	data.LinkURL = req.GetLinkURL()
	data.LinkType = req.GetLinkType()
	data.Sort = req.GetSort()
	data.Status = int16(req.GetStatus())

	if req.GetStartTime() != "" {
		data.StartTime = timeutil.TimeToSQLNullTime(carbon.Parse(req.GetStartTime()).StdTime())
	}
	if req.GetEndTime() != "" {
		data.EndTime = timeutil.TimeToSQLNullTime(carbon.Parse(req.GetEndTime()).StdTime())
	}

	if err := a.activityRepo.Create(ctx, data); err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
