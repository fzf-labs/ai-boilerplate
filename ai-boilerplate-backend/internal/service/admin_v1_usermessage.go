package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1UserMessageService(
	logger log.Logger,
	userMessageRepo *data.UserMessageRepo,
	userRepo *data.UserRepo,
	userMembershipRepo *data.UserMembershipRepo,
) *AdminV1UserMessageService {
	l := log.NewHelper(log.With(logger, "module", "service/userMessage"))
	return &AdminV1UserMessageService{
		log:                l,
		userMessageRepo:    userMessageRepo,
		userRepo:           userRepo,
		userMembershipRepo: userMembershipRepo,
	}
}

type AdminV1UserMessageService struct {
	pb.UnimplementedUserMessageServer
	log                *log.Helper
	userMessageRepo    *data.UserMessageRepo
	userRepo           *data.UserRepo
	userMembershipRepo *data.UserMembershipRepo
}
