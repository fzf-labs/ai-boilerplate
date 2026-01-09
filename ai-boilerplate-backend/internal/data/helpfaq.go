package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewHelpFaqRepo(
	logger log.Logger,
	data *Data,
	helpFaqRepo *ai_boilerplate_repo.HelpFaqRepo,
) *HelpFaqRepo {
	l := log.NewHelper(log.With(logger, "module", "data/helpFaq"))
	return &HelpFaqRepo{
		log:         l,
		data:        data,
		HelpFaqRepo: helpFaqRepo,
	}
}

type HelpFaqRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.HelpFaqRepo
}
