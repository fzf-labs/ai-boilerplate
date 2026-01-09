package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1UserMembershipService(
	logger log.Logger,
	userMembershipRepo *data.UserMembershipRepo,
) *AdminV1UserMembershipService {
	l := log.NewHelper(log.With(logger, "module", "service/userMembership"))
	return &AdminV1UserMembershipService{
		log:                l,
		userMembershipRepo: userMembershipRepo,
	}
}

type AdminV1UserMembershipService struct {
	pb.UnimplementedUserMembershipServer
	log                *log.Helper
	userMembershipRepo *data.UserMembershipRepo
}
