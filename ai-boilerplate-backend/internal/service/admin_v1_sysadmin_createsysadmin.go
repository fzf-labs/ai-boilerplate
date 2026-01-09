package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/goutil/cryptutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// CreateSysAdmin 系统-用户-创建一条数据
func (a *AdminV1SysAdminService) CreateSysAdmin(ctx context.Context, req *pb.CreateSysAdminReq) (*pb.CreateSysAdminReply, error) {
	resp := &pb.CreateSysAdminReply{}
	// 查询用户名是否存在
	sysAdmin, err := a.sysAdminRepo.FindOneCacheByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if sysAdmin != nil && sysAdmin.ID != "" {
		return nil, pb.ErrorReasonAccountAlreadyExists()
	}
	password, err := cryptutil.Encrypt(req.GetPassword())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	data := a.sysAdminRepo.NewData()
	data.TenantID = meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	data.Username = req.GetUsername()
	data.Nickname = req.GetNickname()
	data.Password = password
	data.Avatar = req.GetAvatar()
	data.Sex = int16(req.GetSex())
	data.Email = req.GetEmail()
	data.Mobile = req.GetMobile()
	data.RoleID = req.GetRoleId()
	data.PostID = req.GetPostId()
	data.DeptID = req.GetDeptId()
	data.Status = int16(req.GetStatus())
	err = a.sysAdminRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
