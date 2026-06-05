package service

import (
	"context"
	"fmt"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// SysAuthLogout Auth-退出
func (a *AdminV1SysAuthService) SysAuthLogout(ctx context.Context, _ *pb.SysAuthLogoutReq) (*pb.SysAuthLogoutReply, error) {
	resp := &pb.SysAuthLogoutReply{}

	adminID := meta.GetMetadataFromClient(ctx, constant.XMdAdminID)
	if adminID != "" {
		if err := a.sysAdminRepo.JwtTokenClear(ctx, adminID); err != nil {
			return nil, pb.ErrorReasonTokenErr(pb.WithError(err))
		}
		return resp, nil
	}

	if tr, ok := http.RequestFromServerContext(ctx); ok {
		authorization := tr.Header.Get("Authorization")
		if authorization != "" {
			var token string
			if _, err := fmt.Sscanf(authorization, "Bearer %s", &token); err == nil && token != "" {
				claims, err := a.sysAdminRepo.CheckToken(ctx, token)
				if err == nil {
					if uid, ok := claims["uid"].(string); ok && uid != "" {
						if clearErr := a.sysAdminRepo.JwtTokenClear(ctx, uid); clearErr != nil {
							return nil, pb.ErrorReasonTokenErr(pb.WithError(clearErr))
						}
					}
				}
			}
		}
	}
	return resp, nil
}
