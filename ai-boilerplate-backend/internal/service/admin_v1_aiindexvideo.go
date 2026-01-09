package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1AiIndexVideoService(
	logger log.Logger,
	aiVideoRecordRepo *data.AiVideoRecordRepo,
) *AdminV1AiIndexVideoService {
	l := log.NewHelper(log.With(logger, "module", "service/aiIndexVideo"))
	return &AdminV1AiIndexVideoService{
		log:               l,
		aiVideoRecordRepo: aiVideoRecordRepo,
	}
}

type AdminV1AiIndexVideoService struct {
	pb.UnimplementedAiIndexVideoServer
	log               *log.Helper
	aiVideoRecordRepo *data.AiVideoRecordRepo
}
