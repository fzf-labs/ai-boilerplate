package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateSysRoleStatus 系统-角色-更新状态
func (a *AdminV1SysRoleService) UpdateSysRoleStatus(ctx context.Context, req *pb.UpdateSysRoleStatusReq) (*pb.UpdateSysRoleStatusReply, error) {
	resp := &pb.UpdateSysRoleStatusReply{}
	data, err := a.sysRoleRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.sysRoleRepo.DeepCopy(data)
	data.Status = int16(req.GetStatus())
	err = a.sysRoleRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
