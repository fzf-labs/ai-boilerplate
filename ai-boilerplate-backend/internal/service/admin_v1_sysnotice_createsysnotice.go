package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// CreateSysNotice 系统-公告-创建一条数据
func (a *AdminV1SysNoticeService) CreateSysNotice(ctx context.Context, req *pb.CreateSysNoticeReq) (*pb.CreateSysNoticeReply, error) {
	resp := &pb.CreateSysNoticeReply{}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	data := a.sysNoticeRepo.NewData()
	data.TenantID = tenantID
	data.Type = req.GetType()
	data.Title = req.GetTitle()
	data.Content = req.GetContent()
	data.Status = int16(req.GetStatus())
	err := a.sysNoticeRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
