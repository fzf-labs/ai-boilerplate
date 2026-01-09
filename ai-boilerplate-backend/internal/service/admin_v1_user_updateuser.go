package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateUser 用户表-更新一条数据
func (a *AdminV1UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserReply, error) {
	resp := &pb.UpdateUserReply{}
	data, err := a.userRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.userRepo.DeepCopy(data)
	data.Nickname = req.GetNickname()
	data.Gender = req.GetGender()
	data.Avatar = req.GetAvatar()
	data.Profile = req.GetProfile()
	data.Status = req.GetStatus()
	err = a.userRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
