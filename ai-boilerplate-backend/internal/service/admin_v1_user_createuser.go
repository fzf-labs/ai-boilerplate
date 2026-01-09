package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// CreateUser 用户表-创建一条数据
func (a *AdminV1UserService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserReply, error) {
	resp := &pb.CreateUserReply{}
	salt := a.userRepo.GenerateSalt()
	password, err := a.userRepo.GeneratePassword(salt, req.GetPhone())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	data := a.userRepo.NewData()
	data.Phone = req.GetPhone()
	data.Password = password
	data.Salt = salt
	data.Nickname = a.userRepo.GenerateNicknameByPhone(req.GetPhone())
	data.Gender = req.GetGender()
	data.Avatar = req.GetAvatar()
	data.Profile = req.GetProfile()
	data.Status = req.GetStatus()
	err = a.userRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
