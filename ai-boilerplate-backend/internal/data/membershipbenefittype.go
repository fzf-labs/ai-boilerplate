package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewMembershipBenefitTypeRepo(
	logger log.Logger,
	data *Data,
	membershipBenefitTypeRepo *ai_boilerplate_repo.MembershipBenefitTypeRepo,
) *MembershipBenefitTypeRepo {
	l := log.NewHelper(log.With(logger, "module", "data/membershipBenefitType"))
	return &MembershipBenefitTypeRepo{
		log:                       l,
		data:                      data,
		MembershipBenefitTypeRepo: membershipBenefitTypeRepo,
	}
}

type MembershipBenefitTypeRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.MembershipBenefitTypeRepo
}
