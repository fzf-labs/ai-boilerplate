package service

import (
	"context"

	"github.com/dromara/carbon/v2"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/goutil/timeutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// UpdateActivity 内容-活动-更新一条数据
func (a *AdminV1ActivityService) UpdateActivity(ctx context.Context, req *pb.UpdateActivityReq) (*pb.UpdateActivityReply, error) {
	resp := &pb.UpdateActivityReply{}
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

	data.Title = req.GetTitle()
	data.ImageURL = req.GetImageURL()
	data.LinkURL = req.GetLinkURL()
	data.LinkType = req.GetLinkType()
	data.Sort = req.GetSort()
	data.Status = int16(req.GetStatus())

	if req.GetStartTime() != "" {
		data.StartTime = timeutil.TimeToSQLNullTime(carbon.Parse(req.GetStartTime()).StdTime())
	} else {
		data.StartTime.Valid = false
	}
	if req.GetEndTime() != "" {
		data.EndTime = timeutil.TimeToSQLNullTime(carbon.Parse(req.GetEndTime()).StdTime())
	} else {
		data.EndTime.Valid = false
	}

	if err := a.activityRepo.Update(ctx, data); err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
