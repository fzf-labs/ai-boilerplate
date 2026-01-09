package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/goutil/cryptutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// SysAuthUpdateAdminPassword Auth-更新密码
func (a *AdminV1SysAuthService) SysAuthUpdateAdminPassword(ctx context.Context, req *pb.SysAuthUpdateAdminPasswordReq) (*pb.SysAuthUpdateAdminPasswordReply, error) {
	resp := &pb.SysAuthUpdateAdminPasswordReply{}
	adminID := meta.GetMetadataFromClient(ctx, constant.XMdAdminID)
	admin, err := a.sysAdminRepo.FindOneCacheByID(ctx, adminID)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if admin == nil || admin.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	// 验证旧密码
	if compareErr := cryptutil.Compare(admin.Password, req.GetOldPassword()); compareErr != nil {
		return nil, pb.ErrorReasonAccountPasswordError()
	}
	// 加密新密码
	oldData := a.sysAdminRepo.DeepCopy(admin)
	password, err := cryptutil.Encrypt(req.GetNewPassword())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	admin.Password = password
	err = a.sysAdminRepo.UpdateOneCacheWithZero(ctx, admin, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
