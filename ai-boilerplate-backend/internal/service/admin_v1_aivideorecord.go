package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1AiVideoRecordService(
	logger log.Logger,
	aiVideoRecordRepo *data.AiVideoRecordRepo,
) *AdminV1AiVideoRecordService {
	l := log.NewHelper(log.With(logger, "module", "service/aiVideoRecord"))
	return &AdminV1AiVideoRecordService{
		log:               l,
		aiVideoRecordRepo: aiVideoRecordRepo,
	}
}

type AdminV1AiVideoRecordService struct {
	pb.UnimplementedAiVideoRecordServer
	log               *log.Helper
	aiVideoRecordRepo *data.AiVideoRecordRepo
}
