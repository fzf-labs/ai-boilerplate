package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAiProviderPlatformRepo(
	logger log.Logger,
	data *Data,
	aiProviderPlatformRepo *ai_boilerplate_repo.AiProviderPlatformRepo,
) *AiProviderPlatformRepo {
	l := log.NewHelper(log.With(logger, "module", "data/aiProviderPlatform"))
	return &AiProviderPlatformRepo{
		log:                    l,
		data:                   data,
		AiProviderPlatformRepo: aiProviderPlatformRepo,
	}
}

type AiProviderPlatformRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.AiProviderPlatformRepo
}
