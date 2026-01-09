package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetSysAdminInfo 系统-用户-单条数据查询
func (a *AdminV1SysAdminService) GetSysAdminInfo(ctx context.Context, req *pb.GetSysAdminInfoReq) (*pb.GetSysAdminInfoReply, error) {
	resp := &pb.GetSysAdminInfoReply{}
	data, err := a.sysAdminRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	roleNameMap, err := a.sysRoleRepo.RoleIDToName(ctx, []string{data.RoleID})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	deptNameMap, err := a.sysDeptRepo.DeptIDToName(ctx, []string{data.DeptID})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	postNameMap, err := a.sysPostRepo.PostIDToName(ctx, []string{data.PostID})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.SysAdminInfo{
		Id:        data.ID,
		Username:  data.Username,
		Nickname:  data.Nickname,
		Avatar:    data.Avatar,
		Sex:       int32(data.Sex),
		Email:     data.Email,
		Mobile:    data.Mobile,
		RoleId:    data.RoleID,
		PostId:    data.PostID,
		DeptId:    data.DeptID,
		Status:    int32(data.Status),
		CreatedAt: timeutil.RFC3339(data.CreatedAt),
		UpdatedAt: timeutil.RFC3339(data.UpdatedAt),
		RoleName:  roleNameMap[data.RoleID],
		DeptName:  deptNameMap[data.DeptID],
		PostName:  postNameMap[data.PostID],
	}
	return resp, nil
}
