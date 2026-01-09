package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateSysMenu 菜单和权限规则表-更新一条数据
func (a *AdminV1SysMenuService) UpdateSysMenu(ctx context.Context, req *pb.UpdateSysMenuReq) (*pb.UpdateSysMenuReply, error) {
	resp := &pb.UpdateSysMenuReply{}
	data, err := a.sysMenuRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.sysMenuRepo.DeepCopy(data)
	data.Pid = req.GetPid()
	data.Name = req.GetName()
	data.Type = req.GetType()
	data.Path = req.GetPath()
	data.Permission = req.GetPermission()
	data.Icon = req.GetIcon()
	data.Component = req.GetComponent()
	data.ComponentName = req.GetComponentName()
	data.Sort = int64(req.GetSort())
	data.Status = int16(req.GetStatus())
	err = a.sysMenuRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
