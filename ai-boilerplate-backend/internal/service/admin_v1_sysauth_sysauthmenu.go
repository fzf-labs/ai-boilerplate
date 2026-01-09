package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// SysAuthMenu Auth-获取菜单
func (a *AdminV1SysAuthService) SysAuthMenu(ctx context.Context, req *pb.SysAuthMenuReq) (*pb.SysAuthMenuReply, error) {
	_ = req

	resp := &pb.SysAuthMenuReply{}
	// 管理员 ID
	adminID := meta.GetMetadataFromClient(ctx, constant.XMdAdminID)
	// 查询管理员的角色
	admin, err := a.sysAdminRepo.FindOneCacheByID(ctx, adminID)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if admin == nil || admin.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	// 查询管理员的角色
	role, err := a.sysRoleRepo.FindOneCacheByID(ctx, admin.RoleID)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if role == nil || role.ID == "" {
		return resp, nil
	}
	menuIDs := make([]string, 0)
	if role.MenuIds.String() != "" {
		if unmarshalErr := jsonutil.Unmarshal(role.MenuIds, &menuIDs); unmarshalErr != nil {
			return nil, pb.ErrorReasonDataFormattingError(pb.WithError(unmarshalErr))
		}
	}
	if len(menuIDs) == 0 {
		return resp, nil
	}
	menus, err := a.sysMenuRepo.FindMultiCacheByIDS(ctx, menuIDs)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	menuTree, err := a.sysMenuRepo.TraverseMenuTree(ctx, menus)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Menu = menuTree
	return resp, nil
}
