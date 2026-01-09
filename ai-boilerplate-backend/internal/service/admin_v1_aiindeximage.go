package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1AiIndexImageService(
	logger log.Logger,
	aiImageRecordRepo *data.AiImageRecordRepo,
) *AdminV1AiIndexImageService {
	l := log.NewHelper(log.With(logger, "module", "service/aiIndexImage"))
	return &AdminV1AiIndexImageService{
		log:               l,
		aiImageRecordRepo: aiImageRecordRepo,
	}
}

type AdminV1AiIndexImageService struct {
	pb.UnimplementedAiIndexImageServer
	log               *log.Helper
	aiImageRecordRepo *data.AiImageRecordRepo
}
