package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAppV1MembershipService(
	logger log.Logger,
	userMembershipRepo *data.UserMembershipRepo,
	membershipRepo *data.MembershipRepo,
	membershipBenefitRepo *data.MembershipBenefitRepo,
	membershipBenefitTypeRepo *data.MembershipBenefitTypeRepo,
) *AppV1MembershipService {
	l := log.NewHelper(log.With(logger, "module", "service/app_v1_membership"))
	return &AppV1MembershipService{
		log:                       l,
		userMembershipRepo:        userMembershipRepo,
		membershipRepo:            membershipRepo,
		membershipBenefitRepo:     membershipBenefitRepo,
		membershipBenefitTypeRepo: membershipBenefitTypeRepo,
	}
}

type AppV1MembershipService struct {
	pb.UnimplementedMembershipServer
	log                       *log.Helper
	userMembershipRepo        *data.UserMembershipRepo
	membershipRepo            *data.MembershipRepo
	membershipBenefitRepo     *data.MembershipBenefitRepo
	membershipBenefitTypeRepo *data.MembershipBenefitTypeRepo
}
