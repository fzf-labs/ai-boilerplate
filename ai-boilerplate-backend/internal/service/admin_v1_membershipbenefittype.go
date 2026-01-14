package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1MembershipBenefitTypeService(
	logger log.Logger,
	membershipBenefitTypeRepo *data.MembershipBenefitTypeRepo,
) *AdminV1MembershipBenefitTypeService {
	l := log.NewHelper(log.With(logger, "module", "service/membershipBenefitType"))
	return &AdminV1MembershipBenefitTypeService{
		log:                       l,
		membershipBenefitTypeRepo: membershipBenefitTypeRepo,
	}
}

type AdminV1MembershipBenefitTypeService struct {
	pb.UnimplementedMembershipBenefitTypeServer
	log                       *log.Helper
	membershipBenefitTypeRepo *data.MembershipBenefitTypeRepo
}
