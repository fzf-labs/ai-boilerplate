package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
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
	user, err := a.userRepo.FindOneCacheByID(ctx, uid)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if user == nil || user.ID == "" || user.Status != int32(constant.StatusEnable) {
		return nil, pb.ErrorReasonUnauthorized()
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
