package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1AiChatMessageService(
	logger log.Logger,
	aiChatMessageRepo *data.AiChatMessageRepo,
) *AdminV1AiChatMessageService {
	l := log.NewHelper(log.With(logger, "module", "service/aiChatMessage"))
	return &AdminV1AiChatMessageService{
		log:               l,
		aiChatMessageRepo: aiChatMessageRepo,
	}
}

type AdminV1AiChatMessageService struct {
	pb.UnimplementedAiChatMessageServer
	log               *log.Helper
	aiChatMessageRepo *data.AiChatMessageRepo
}
