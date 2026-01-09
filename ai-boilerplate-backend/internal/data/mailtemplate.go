package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewMailTemplateRepo(
	logger log.Logger,
	data *Data,
	mailTemplateRepo *ai_boilerplate_repo.MailTemplateRepo,
) *MailTemplateRepo {
	l := log.NewHelper(log.With(logger, "module", "data/mailTemplate"))
	return &MailTemplateRepo{
		log:              l,
		data:             data,
		MailTemplateRepo: mailTemplateRepo,
	}
}

type MailTemplateRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.MailTemplateRepo
}
