package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1AiImageRecordService(
	logger log.Logger,
	aiImageRecordRepo *data.AiImageRecordRepo,
) *AdminV1AiImageRecordService {
	l := log.NewHelper(log.With(logger, "module", "service/aiImageRecord"))
	return &AdminV1AiImageRecordService{
		log:               l,
		aiImageRecordRepo: aiImageRecordRepo,
	}
}

type AdminV1AiImageRecordService struct {
	pb.UnimplementedAiImageRecordServer
	log               *log.Helper
	aiImageRecordRepo *data.AiImageRecordRepo
}
