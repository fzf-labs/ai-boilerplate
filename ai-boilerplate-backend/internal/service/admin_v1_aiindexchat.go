package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1AiIndexChatService(
	logger log.Logger,
	aiChatConversationRepo *data.AiChatConversationRepo,
	aiChatMessageRepo *data.AiChatMessageRepo,
) *AdminV1AiIndexChatService {
	l := log.NewHelper(log.With(logger, "module", "service/aiIndexChat"))
	return &AdminV1AiIndexChatService{
		log:                    l,
		aiChatConversationRepo: aiChatConversationRepo,
		aiChatMessageRepo:      aiChatMessageRepo,
	}
}

type AdminV1AiIndexChatService struct {
	pb.UnimplementedAiIndexChatServer
	log                    *log.Helper
	aiChatConversationRepo *data.AiChatConversationRepo
	aiChatMessageRepo      *data.AiChatMessageRepo
}
