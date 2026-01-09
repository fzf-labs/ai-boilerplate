package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// CreateSysMenu 菜单和权限规则表-创建一条数据
func (a *AdminV1SysMenuService) CreateSysMenu(ctx context.Context, req *pb.CreateSysMenuReq) (*pb.CreateSysMenuReply, error) {
	resp := &pb.CreateSysMenuReply{}
	data := a.sysMenuRepo.NewData()
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
	err := a.sysMenuRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
