package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewUserMessageRepo(
	logger log.Logger,
	data *Data,
	userMessageRepo *ai_boilerplate_repo.UserMessageRepo,
) *UserMessageRepo {
	l := log.NewHelper(log.With(logger, "module", "data/userMessage"))
	return &UserMessageRepo{
		log:             l,
		data:            data,
		UserMessageRepo: userMessageRepo,
	}
}

type UserMessageRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.UserMessageRepo
}
