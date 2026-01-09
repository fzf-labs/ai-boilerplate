package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewMembershipBenefitRepo(
	logger log.Logger,
	data *Data,
	membershipBenefitRepo *ai_boilerplate_repo.MembershipBenefitRepo,
) *MembershipBenefitRepo {
	l := log.NewHelper(log.With(logger, "module", "data/membershipBenefit"))
	return &MembershipBenefitRepo{
		log:                   l,
		data:                  data,
		MembershipBenefitRepo: membershipBenefitRepo,
	}
}

type MembershipBenefitRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.MembershipBenefitRepo
}
