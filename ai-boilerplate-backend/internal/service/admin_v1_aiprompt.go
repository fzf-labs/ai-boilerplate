package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1AiPromptService(
	logger log.Logger,
	aiPromptRepo *data.AiPromptRepo,
) *AdminV1AiPromptService {
	l := log.NewHelper(log.With(logger, "module", "service/aiPrompt"))
	return &AdminV1AiPromptService{
		log:          l,
		aiPromptRepo: aiPromptRepo,
	}
}

type AdminV1AiPromptService struct {
	pb.UnimplementedAiPromptServer
	log          *log.Helper
	aiPromptRepo *data.AiPromptRepo
}
