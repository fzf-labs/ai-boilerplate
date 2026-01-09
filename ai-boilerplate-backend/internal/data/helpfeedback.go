package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewHelpFeedbackRepo(
	logger log.Logger,
	data *Data,
	helpFeedbackRepo *ai_boilerplate_repo.HelpFeedbackRepo,
) *HelpFeedbackRepo {
	l := log.NewHelper(log.With(logger, "module", "data/helpFeedback"))
	return &HelpFeedbackRepo{
		log:              l,
		data:             data,
		HelpFeedbackRepo: helpFeedbackRepo,
	}
}

type HelpFeedbackRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.HelpFeedbackRepo
}
