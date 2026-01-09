package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAiChatMessageRepo(
	logger log.Logger,
	data *Data,
	aiChatMessageRepo *ai_boilerplate_repo.AiChatMessageRepo,
) *AiChatMessageRepo {
	l := log.NewHelper(log.With(logger, "module", "data/aiChatMessage"))
	return &AiChatMessageRepo{
		log:               l,
		data:              data,
		AiChatMessageRepo: aiChatMessageRepo,
	}
}

type AiChatMessageRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.AiChatMessageRepo
}
