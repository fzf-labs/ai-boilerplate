package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1MembershipService(
	logger log.Logger,
	membershipRepo *data.MembershipRepo,
) *AdminV1MembershipService {
	l := log.NewHelper(log.With(logger, "module", "service/membership"))
	return &AdminV1MembershipService{
		log:            l,
		membershipRepo: membershipRepo,
	}
}

type AdminV1MembershipService struct {
	pb.UnimplementedMembershipServer
	log            *log.Helper
	membershipRepo *data.MembershipRepo
}
