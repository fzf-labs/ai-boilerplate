package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GenerateParentTestToken 用户表-生成测试 Token
func (a *AdminV1UserService) GenerateParentTestToken(ctx context.Context, req *pb.GenerateParentTestTokenReq) (*pb.GenerateParentTestTokenReply, error) {
	resp := &pb.GenerateParentTestTokenReply{}
	// 查询用户
	user, err := a.userRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if user == nil || user.ID == "" {
		return nil, pb.ErrorReasonAccountNotFound()
	}
	// 生成token
	token, err := a.userRepo.GenerateToken(ctx, user.ID, user.WxGzhUserID, user.WxGzhXcxID)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Token = token.Token
	return resp, nil
}
