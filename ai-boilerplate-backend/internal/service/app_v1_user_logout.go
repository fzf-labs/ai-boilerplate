package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
)

// Logout 退出登录
func (a *AppV1UserService) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutReply, error) {
	// 从context中获取用户ID，通过中间件注入
	userID, ok := ctx.Value("userId").(string)
	if !ok || userID == "" {
		return nil, pb.ErrorReasonUnauthorized()
	}

	// 清理用户的所有token
	if err := a.userRepo.JwtTokenClear(ctx, userID); err != nil {
		return nil, pb.ErrorReasonTokenErr(pb.WithError(err))
	}

	return &pb.LogoutReply{}, nil
}
