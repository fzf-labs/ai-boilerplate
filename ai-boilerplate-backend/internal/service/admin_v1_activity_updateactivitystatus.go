package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// UpdateActivityStatus 内容-活动-更新状态
func (a *AdminV1ActivityService) UpdateActivityStatus(ctx context.Context, req *pb.UpdateActivityStatusReq) (*pb.UpdateActivityStatusReply, error) {
	resp := &pb.UpdateActivityStatusReply{}
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

	data.Status = int16(req.GetStatus())
	if err := a.activityRepo.Update(ctx, data); err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
