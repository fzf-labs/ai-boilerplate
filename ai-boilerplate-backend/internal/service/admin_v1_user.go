package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1UserService(
	logger log.Logger,
	userRepo *data.UserRepo,
	userMembershipRepo *data.UserMembershipRepo,
) *AdminV1UserService {
	l := log.NewHelper(log.With(logger, "module", "service/user"))
	return &AdminV1UserService{
		log:                l,
		userRepo:           userRepo,
		userMembershipRepo: userMembershipRepo,
	}
}

type AdminV1UserService struct {
	pb.UnimplementedUserServer
	log                *log.Helper
	userRepo           *data.UserRepo
	userMembershipRepo *data.UserMembershipRepo
}
