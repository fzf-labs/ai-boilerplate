package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1MembershipBenefitService(
	logger log.Logger,
	membershipBenefitRepo *data.MembershipBenefitRepo,
) *AdminV1MembershipBenefitService {
	l := log.NewHelper(log.With(logger, "module", "service/membershipBenefit"))
	return &AdminV1MembershipBenefitService{
		log:                   l,
		membershipBenefitRepo: membershipBenefitRepo,
	}
}

type AdminV1MembershipBenefitService struct {
	pb.UnimplementedMembershipBenefitServer
	log                   *log.Helper
	membershipBenefitRepo *data.MembershipBenefitRepo
}
