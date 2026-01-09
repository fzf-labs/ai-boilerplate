package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// SysAuthCheckToken Auth-检查token
func (a *AdminV1SysAuthService) SysAuthCheckToken(ctx context.Context, req *pb.SysAuthCheckTokenReq) (*pb.SysAuthCheckTokenReply, error) {
	resp := &pb.SysAuthCheckTokenReply{
		AdminId: "",
	}
	uid, err := a.sysAdminRepo.CheckToken(ctx, req.GetToken())
	if err != nil {
		return nil, pb.ErrorReasonTokenInvalidErr(pb.WithError(err))
	}
	resp.AdminId = uid["uid"].(string)
	resp.Nickname = uid["nickname"].(string)
	resp.TenantId = uid["tenant_id"].(string)
	return resp, nil
}
