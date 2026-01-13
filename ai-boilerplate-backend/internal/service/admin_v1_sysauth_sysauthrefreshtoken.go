package service

import (
	"context"
	"errors"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/gopkg/jwt"
)

// SysAuthRefreshToken Auth-刷新token
func (a *AdminV1SysAuthService) SysAuthRefreshToken(ctx context.Context, req *pb.SysAuthRefreshTokenReq) (*pb.SysAuthRefreshTokenReply, error) {
	resp := &pb.SysAuthRefreshTokenReply{
		Token:     "",
		ExpiredAt: 0,
		RefreshAt: 0,
	}
	token, err := a.sysAdminRepo.RefreshToken(ctx, req.GetRefreshToken())
	if err != nil {
		switch {
		case errors.Is(err, jwt.TokenExpired):
			return nil, pb.ErrorReasonTokenExpiredErr(pb.WithError(err))
		case errors.Is(err, jwt.TokenInvalid):
			return nil, pb.ErrorReasonTokenInvalidErr(pb.WithError(err))
		case errors.Is(err, jwt.TokenCanNotRefresh):
			return nil, pb.ErrorReasonTokenErr(pb.WithError(err))
		default:
			return nil, pb.ErrorReasonTokenErr(pb.WithError(err))
		}
	}
	resp.Token = token.Token
	resp.ExpiredAt = token.ExpiredAt
	resp.RefreshAt = token.RefreshAt
	return resp, nil
}
