package service

import (
	"context"
	"strings"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// ChangePassword 修改密码
func (a *AppV1UserService) ChangePassword(ctx context.Context, req *pb.ChangePasswordReq) (*pb.ChangePasswordReply, error) {
	resp := &pb.ChangePasswordReply{}
	userID := meta.GetMetadataFromClient(ctx, constant.XMdUserID)
	data, err := a.userRepo.FindOneCacheByID(ctx, userID)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonAccountNotFound()
	}
	if !a.userRepo.VerifyPassword(data.Salt, req.GetOldPassword(), data.Password) {
		return nil, pb.ErrorReasonAccountPasswordError()
	}

	newPassword := strings.TrimSpace(req.GetNewPassword())
	confirmPassword := strings.TrimSpace(req.GetConfirmPassword())
	if newPassword == "" || confirmPassword == "" || newPassword != confirmPassword {
		return nil, pb.ErrorReasonParamError(pb.WithFmtMsg("两次输入的新密码不一致"))
	}

	oldData := a.userRepo.DeepCopy(data)
	data.Salt = a.userRepo.GenerateSalt()
	data.Password, err = a.userRepo.GeneratePassword(data.Salt, newPassword)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if err := a.userRepo.UpdateOneCacheWithZero(ctx, data, oldData); err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
