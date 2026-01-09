package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1AiIndexPromptService(
	logger log.Logger,
	aiPromptRepo *data.AiPromptRepo,
) *AdminV1AiIndexPromptService {
	l := log.NewHelper(log.With(logger, "module", "service/aiIndexPrompt"))
	return &AdminV1AiIndexPromptService{
		log:          l,
		aiPromptRepo: aiPromptRepo,
	}
}

type AdminV1AiIndexPromptService struct {
	pb.UnimplementedAiIndexPromptServer
	log          *log.Helper
	aiPromptRepo *data.AiPromptRepo
}
