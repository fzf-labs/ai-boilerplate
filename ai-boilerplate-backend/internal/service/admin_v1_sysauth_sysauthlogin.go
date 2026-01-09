package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/cryptutil"
)

// SysAuthLogin Auth-登录
func (a *AdminV1SysAuthService) SysAuthLogin(ctx context.Context, req *pb.SysAuthLoginReq) (*pb.SysAuthLoginReply, error) {
	resp := &pb.SysAuthLoginReply{
		Token:     "",
		ExpiredAt: 0,
		RefreshAt: 0,
	}
	// 查询用户
	sysAdmin, err := a.sysAdminRepo.FindOneCacheByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if sysAdmin == nil || sysAdmin.ID == "" {
		return nil, pb.ErrorReasonAccountNotFound()
	}
	// 验证密码
	if compareErr := cryptutil.Compare(sysAdmin.Password, req.GetPassword()); compareErr != nil {
		return nil, pb.ErrorReasonAccountPasswordError()
	}
	// 生成token
	token, err := a.sysAdminRepo.GenerateToken(ctx, sysAdmin)
	if err != nil {
		return nil, pb.ErrorReasonTokenErr(pb.WithError(err))
	}
	resp.Token = token.Token
	resp.ExpiredAt = token.ExpiredAt
	resp.RefreshAt = token.RefreshAt
	return resp, nil
}
