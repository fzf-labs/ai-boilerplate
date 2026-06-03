package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAppV1UserService(
	logger log.Logger,
	commonRepo *data.CommonRepo,
	smsCodeRepo *data.SmsCodeRepo,
	userRepo *data.UserRepo,
	userMembershipRepo *data.UserMembershipRepo,
) *AppV1UserService {
	l := log.NewHelper(log.With(logger, "module", "service/user"))
	return &AppV1UserService{
		log:                l,
		commonRepo:         commonRepo,
		smsCodeRepo:        smsCodeRepo,
		userRepo:           userRepo,
		userMembershipRepo: userMembershipRepo,
	}
}

type AppV1UserService struct {
	pb.UnimplementedUserServer
	log                *log.Helper
	commonRepo         *data.CommonRepo
	smsCodeRepo        *data.SmsCodeRepo
	userRepo           *data.UserRepo
	userMembershipRepo *data.UserMembershipRepo
}
