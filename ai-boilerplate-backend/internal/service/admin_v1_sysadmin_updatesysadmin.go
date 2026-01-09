package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateSysAdmin 系统-用户-更新一条数据
func (a *AdminV1SysAdminService) UpdateSysAdmin(ctx context.Context, req *pb.UpdateSysAdminReq) (*pb.UpdateSysAdminReply, error) {
	resp := &pb.UpdateSysAdminReply{}
	data, err := a.sysAdminRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.sysAdminRepo.DeepCopy(data)
	data.Username = req.GetUsername()
	data.Nickname = req.GetNickname()
	data.Avatar = req.GetAvatar()
	data.Sex = int16(req.GetSex())
	data.Email = req.GetEmail()
	data.Mobile = req.GetMobile()
	data.RoleID = req.GetRoleId()
	data.PostID = req.GetPostId()
	data.DeptID = req.GetDeptId()
	data.Status = int16(req.GetStatus())
	err = a.sysAdminRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
