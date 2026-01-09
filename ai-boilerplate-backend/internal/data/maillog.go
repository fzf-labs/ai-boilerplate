package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewMailLogRepo(
	logger log.Logger,
	data *Data,
	mailLogRepo *ai_boilerplate_repo.MailLogRepo,
) *MailLogRepo {
	l := log.NewHelper(log.With(logger, "module", "data/mailLog"))
	return &MailLogRepo{
		log:         l,
		data:        data,
		MailLogRepo: mailLogRepo,
	}
}

type MailLogRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.MailLogRepo
}
