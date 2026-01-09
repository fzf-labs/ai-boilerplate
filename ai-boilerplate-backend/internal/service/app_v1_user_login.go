package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
)

// Login 登录
func (a *AppV1UserService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	resp := &pb.LoginReply{}
	// 查询用户(默认 username=phone)
	user, err := a.userRepo.FindOneCacheByPhone(ctx, req.GetUsername())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if user == nil || user.ID == "" {
		return nil, pb.ErrorReasonAccountNotFound()
	}
	// 验证密码
	if !a.userRepo.VerifyPassword(user.Salt, req.GetPassword(), user.Password) {
		return nil, pb.ErrorReasonAccountPasswordError()
	}
	// 生成token
	token, err := a.userRepo.GenerateToken(ctx, user.ID, user.WxGzhUserID, user.WxGzhXcxID)
	if err != nil {
		return nil, pb.ErrorReasonTokenErr(pb.WithError(err))
	}
	resp.Token = token.Token
	resp.ExpiredAt = token.ExpiredAt
	resp.RefreshAt = token.RefreshAt
	return resp, nil
}
