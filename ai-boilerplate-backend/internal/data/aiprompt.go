package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAiPromptRepo(
	logger log.Logger,
	data *Data,
	aiPromptRepo *ai_boilerplate_repo.AiPromptRepo,
) *AiPromptRepo {
	l := log.NewHelper(log.With(logger, "module", "data/aiPrompt"))
	return &AiPromptRepo{
		log:          l,
		data:         data,
		AiPromptRepo: aiPromptRepo,
	}
}

type AiPromptRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.AiPromptRepo
}
