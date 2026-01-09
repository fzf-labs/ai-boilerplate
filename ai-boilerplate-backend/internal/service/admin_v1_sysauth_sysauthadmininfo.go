package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/goutil/timeutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// SysAuthAdminInfo Auth-查询用户信息
func (a *AdminV1SysAuthService) SysAuthAdminInfo(ctx context.Context, req *pb.SysAuthAdminInfoReq) (*pb.SysAuthAdminInfoReply, error) {
	_ = req

	resp := &pb.SysAuthAdminInfoReply{
		Info: &pb.SysAdminInfo{},
	}
	adminID := meta.GetMetadataFromClient(ctx, constant.XMdAdminID)
	admin, err := a.sysAdminRepo.FindOneCacheByID(ctx, adminID)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	roleNameMap, err := a.sysRoleRepo.RoleIDToName(ctx, []string{admin.RoleID})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	deptNameMap, err := a.sysDeptRepo.DeptIDToName(ctx, []string{admin.DeptID})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	postNameMap, err := a.sysPostRepo.PostIDToName(ctx, []string{admin.PostID})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.SysAdminInfo{
		Id:        admin.ID,
		Username:  admin.Username,
		Nickname:  admin.Nickname,
		Avatar:    admin.Avatar,
		Sex:       int32(admin.Sex),
		Email:     admin.Email,
		Mobile:    admin.Mobile,
		RoleId:    admin.RoleID,
		PostId:    admin.PostID,
		DeptId:    admin.DeptID,
		Status:    int32(admin.Status),
		CreatedAt: timeutil.RFC3339(admin.CreatedAt),
		UpdatedAt: timeutil.RFC3339(admin.UpdatedAt),
		RoleName:  roleNameMap[admin.RoleID],
		DeptName:  deptNameMap[admin.DeptID],
		PostName:  postNameMap[admin.PostID],
	}
	return resp, nil
}
