package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
)

// CheckToken 检查token
func (a *AppV1UserService) CheckToken(ctx context.Context, req *pb.CheckTokenReq) (*pb.CheckTokenReply, error) {
	resp := &pb.CheckTokenReply{
		UserId: "",
	}
	claims, err := a.userRepo.CheckToken(ctx, req.GetToken())
	if err != nil {
		return nil, pb.ErrorReasonTokenInvalidErr(pb.WithError(err))
	}
	uid, ok := claims["uid"].(string)
	if !ok || uid == "" {
		return nil, pb.ErrorReasonTokenInvalidErr()
	}
	resp.UserId = uid
	if wxGzhUserID, ok := claims["wxGzhUserId"].(string); ok {
		resp.WxGzhUserId = wxGzhUserID
	}
	if wxGzhXcxID, ok := claims["wxGzhXcxId"].(string); ok {
		resp.WxGzhXcxId = wxGzhXcxID
	}
	return resp, nil
}
