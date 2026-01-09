package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateSysMenuStatus 菜单和权限规则表-更新状态
func (a *AdminV1SysMenuService) UpdateSysMenuStatus(ctx context.Context, req *pb.UpdateSysMenuStatusReq) (*pb.UpdateSysMenuStatusReply, error) {
	resp := &pb.UpdateSysMenuStatusReply{}
	data, err := a.sysMenuRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.sysMenuRepo.DeepCopy(data)
	data.Status = int16(req.GetStatus())
	err = a.sysMenuRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
