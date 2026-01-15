package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewUserMembershipChangeRepo(
	logger log.Logger,
	data *Data,
	userMembershipChangeRepo *ai_boilerplate_repo.UserMembershipChangeRepo,
) *UserMembershipChangeRepo {
	l := log.NewHelper(log.With(logger, "module", "data/userMembershipChange"))
	return &UserMembershipChangeRepo{
		log:                      l,
		data:                     data,
		UserMembershipChangeRepo: userMembershipChangeRepo,
	}
}

type UserMembershipChangeRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.UserMembershipChangeRepo
}
