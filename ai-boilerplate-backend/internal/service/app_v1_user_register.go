package service

import (
	"context"
	"strings"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
)

// Register 注册
func (a *AppV1UserService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	resp := &pb.RegisterReply{}

	phone := strings.TrimSpace(req.GetPhone())
	if phone == "" {
		return nil, pb.ErrorReasonParamError()
	}
	password := req.GetPassword()
	confirmPassword := req.GetConfirmPassword()
	if password != confirmPassword {
		return nil, pb.ErrorReasonParamError(pb.WithFmtMsg("两次输入的密码不一致"))
	}

	existUser, err := a.userRepo.FindOneCacheByPhone(ctx, phone)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if existUser != nil && existUser.ID != "" {
		return nil, pb.ErrorReasonAccountAlreadyExists()
	}

	salt := a.userRepo.GenerateSalt()
	passwordHash, err := a.userRepo.GeneratePassword(salt, password)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	data := a.userRepo.NewData()
	data.Phone = phone
	data.Password = passwordHash
	data.Salt = salt
	data.Nickname = a.userRepo.GenerateNicknameByPhone(phone)
	data.Status = int32(constant.StatusEnable)

	if nickname := strings.TrimSpace(req.GetNickname()); nickname != "" {
		data.Nickname = nickname
	}

	if err := a.userRepo.CreateOneCache(ctx, data); err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	resp.Id = data.ID
	return resp, nil
}
