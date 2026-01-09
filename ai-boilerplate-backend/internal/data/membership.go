package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewMembershipRepo(
	logger log.Logger,
	data *Data,
	membershipRepo *ai_boilerplate_repo.MembershipRepo,
) *MembershipRepo {
	l := log.NewHelper(log.With(logger, "module", "data/membership"))
	return &MembershipRepo{
		log:            l,
		data:           data,
		MembershipRepo: membershipRepo,
	}
}

type MembershipRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.MembershipRepo
}
