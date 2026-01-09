package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// UpdateSysNoticeStatus 系统-公告-更新状态
func (a *AdminV1SysNoticeService) UpdateSysNoticeStatus(ctx context.Context, req *pb.UpdateSysNoticeStatusReq) (*pb.UpdateSysNoticeStatusReply, error) {
	resp := &pb.UpdateSysNoticeStatusReply{}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	data, err := a.sysNoticeRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	// 判断是否是当前租户的公告
	if data.TenantID != tenantID {
		return nil, pb.ErrorReasonAccountNoDataPermission()
	}
	oldData := a.sysNoticeRepo.DeepCopy(data)
	data.Status = int16(req.GetStatus())
	err = a.sysNoticeRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
