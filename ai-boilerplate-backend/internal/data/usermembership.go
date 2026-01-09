package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewUserMembershipRepo(
	logger log.Logger,
	data *Data,
	userMembershipRepo *ai_boilerplate_repo.UserMembershipRepo,
) *UserMembershipRepo {
	l := log.NewHelper(log.With(logger, "module", "data/userMembership"))
	return &UserMembershipRepo{
		log:                l,
		data:               data,
		UserMembershipRepo: userMembershipRepo,
	}
}

type UserMembershipRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.UserMembershipRepo
}
