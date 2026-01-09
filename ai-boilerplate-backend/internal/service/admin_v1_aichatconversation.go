package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1AiChatConversationService(
	logger log.Logger,
	aiChatConversationRepo *data.AiChatConversationRepo,
) *AdminV1AiChatConversationService {
	l := log.NewHelper(log.With(logger, "module", "service/aiChatConversation"))
	return &AdminV1AiChatConversationService{
		log:                    l,
		aiChatConversationRepo: aiChatConversationRepo,
	}
}

type AdminV1AiChatConversationService struct {
	pb.UnimplementedAiChatConversationServer
	log                    *log.Helper
	aiChatConversationRepo *data.AiChatConversationRepo
}
