package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteSysMenu 菜单和权限规则表-删除一条数据
func (a *AdminV1SysMenuService) DeleteSysMenu(ctx context.Context, req *pb.DeleteSysMenuReq) (*pb.DeleteSysMenuReply, error) {
	resp := &pb.DeleteSysMenuReply{}
	// 查询是否有下级
	children, err := a.sysMenuRepo.FindMultiCacheByPid(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if len(children) > 0 {
		return nil, pb.ErrorReasonMenuOperationFailed(pb.WithFmtMsg("有下级菜单"))
	}
	// 删除菜单
	err = a.sysMenuRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
