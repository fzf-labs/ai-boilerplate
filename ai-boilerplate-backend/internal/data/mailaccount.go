package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewMailAccountRepo(
	logger log.Logger,
	data *Data,
	mailAccountRepo *ai_boilerplate_repo.MailAccountRepo,
) *MailAccountRepo {
	l := log.NewHelper(log.With(logger, "module", "data/mailAccount"))
	return &MailAccountRepo{
		log:             l,
		data:            data,
		MailAccountRepo: mailAccountRepo,
	}
}

type MailAccountRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.MailAccountRepo
}
