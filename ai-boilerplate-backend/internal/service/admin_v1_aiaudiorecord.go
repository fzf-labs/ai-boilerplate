package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1AiAudioRecordService(
	logger log.Logger,
	aiAudioRecordRepo *data.AiAudioRecordRepo,
) *AdminV1AiAudioRecordService {
	l := log.NewHelper(log.With(logger, "module", "service/aiAudioRecord"))
	return &AdminV1AiAudioRecordService{
		log:               l,
		aiAudioRecordRepo: aiAudioRecordRepo,
	}
}

type AdminV1AiAudioRecordService struct {
	pb.UnimplementedAiAudioRecordServer
	log               *log.Helper
	aiAudioRecordRepo *data.AiAudioRecordRepo
}
