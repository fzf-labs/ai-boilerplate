package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteSysRole 系统-角色-删除一条数据
func (a *AdminV1SysRoleService) DeleteSysRole(ctx context.Context, req *pb.DeleteSysRoleReq) (*pb.DeleteSysRoleReply, error) {
	resp := &pb.DeleteSysRoleReply{}
	err := a.sysRoleRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
