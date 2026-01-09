package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAiProviderModelRepo(
	logger log.Logger,
	data *Data,
	aiProviderModelRepo *ai_boilerplate_repo.AiProviderModelRepo,
) *AiProviderModelRepo {
	l := log.NewHelper(log.With(logger, "module", "data/aiProviderModel"))
	return &AiProviderModelRepo{
		log:                 l,
		data:                data,
		AiProviderModelRepo: aiProviderModelRepo,
	}
}

type AiProviderModelRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.AiProviderModelRepo
}
