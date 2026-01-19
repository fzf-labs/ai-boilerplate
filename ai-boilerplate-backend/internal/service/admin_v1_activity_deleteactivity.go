package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// DeleteActivity 内容-活动-删除一条数据
func (a *AdminV1ActivityService) DeleteActivity(ctx context.Context, req *pb.DeleteActivityReq) (*pb.DeleteActivityReply, error) {
	resp := &pb.DeleteActivityReply{}
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

	if err := a.activityRepo.Delete(ctx, req.GetId()); err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
