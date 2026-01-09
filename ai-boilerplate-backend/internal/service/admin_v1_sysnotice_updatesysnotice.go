package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// UpdateSysNotice 系统-公告-更新一条数据
func (a *AdminV1SysNoticeService) UpdateSysNotice(ctx context.Context, req *pb.UpdateSysNoticeReq) (*pb.UpdateSysNoticeReply, error) {
	resp := &pb.UpdateSysNoticeReply{}
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
	data.Type = req.GetType()
	data.Title = req.GetTitle()
	data.Content = req.GetContent()
	data.Status = int16(req.GetStatus())
	err = a.sysNoticeRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
