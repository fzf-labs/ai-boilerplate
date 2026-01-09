package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1AiProviderPlatformService(
	logger log.Logger,
	aiProviderPlatformRepo *data.AiProviderPlatformRepo,
) *AdminV1AiProviderPlatformService {
	l := log.NewHelper(log.With(logger, "module", "service/aiProviderPlatform"))
	return &AdminV1AiProviderPlatformService{
		log:                    l,
		aiProviderPlatformRepo: aiProviderPlatformRepo,
	}
}

type AdminV1AiProviderPlatformService struct {
	pb.UnimplementedAiProviderPlatformServer
	log                    *log.Helper
	aiProviderPlatformRepo *data.AiProviderPlatformRepo
}
