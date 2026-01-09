package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAiChatConversationRepo(
	logger log.Logger,
	data *Data,
	aiChatConversationRepo *ai_boilerplate_repo.AiChatConversationRepo,
) *AiChatConversationRepo {
	l := log.NewHelper(log.With(logger, "module", "data/aiChatConversation"))
	return &AiChatConversationRepo{
		log:                    l,
		data:                   data,
		AiChatConversationRepo: aiChatConversationRepo,
	}
}

type AiChatConversationRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.AiChatConversationRepo
}
