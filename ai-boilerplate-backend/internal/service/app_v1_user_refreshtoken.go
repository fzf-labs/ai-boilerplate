package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
)

// RefreshToken 刷新Token
func (a *AppV1UserService) RefreshToken(ctx context.Context, req *pb.RefreshTokenReq) (*pb.RefreshTokenReply, error) {
	resp := &pb.RefreshTokenReply{}

	// 刷新token
	token, err := a.userRepo.RefreshToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, pb.ErrorReasonTokenErr(pb.WithError(err))
	}

	resp.Token = token.Token
	resp.ExpiredAt = token.ExpiredAt
	resp.RefreshAt = token.RefreshAt
	return resp, nil
}
