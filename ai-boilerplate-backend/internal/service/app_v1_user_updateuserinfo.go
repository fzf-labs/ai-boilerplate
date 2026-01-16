package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// UpdateUserInfo 更新用户信息
func (a *AppV1UserService) UpdateUserInfo(ctx context.Context, req *pb.UpdateUserInfoReq) (*pb.UpdateUserInfoReply, error) {
	resp := &pb.UpdateUserInfoReply{}
	userID := meta.GetMetadataFromClient(ctx, constant.XMdUserID)
	data, err := a.userRepo.FindOneCacheByID(ctx, userID)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.userRepo.DeepCopy(data)
	data.Nickname = req.GetNickname()
	data.Avatar = req.GetAvatar()
	data.Gender = req.GetGender()
	data.Profile = req.GetProfile()
	err = a.userRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
