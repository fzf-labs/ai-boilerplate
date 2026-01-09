package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1AiProviderModelService(
	logger log.Logger,
	aiProviderModelRepo *data.AiProviderModelRepo,
) *AdminV1AiProviderModelService {
	l := log.NewHelper(log.With(logger, "module", "service/aiProviderModel"))
	return &AdminV1AiProviderModelService{
		log:                 l,
		aiProviderModelRepo: aiProviderModelRepo,
	}
}

type AdminV1AiProviderModelService struct {
	pb.UnimplementedAiProviderModelServer
	log                 *log.Helper
	aiProviderModelRepo *data.AiProviderModelRepo
}
