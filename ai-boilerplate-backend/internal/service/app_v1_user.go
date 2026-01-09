package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAppV1UserService(
	logger log.Logger,
	userRepo *data.UserRepo,
) *AppV1UserService {
	l := log.NewHelper(log.With(logger, "module", "service/user"))
	return &AppV1UserService{
		log:      l,
		userRepo: userRepo,
	}
}

type AppV1UserService struct {
	pb.UnimplementedUserServer
	log      *log.Helper
	userRepo *data.UserRepo
}
