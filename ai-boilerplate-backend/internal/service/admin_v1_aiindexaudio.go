package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1AiIndexAudioService(
	logger log.Logger,
	aiAudioRecordRepo *data.AiAudioRecordRepo,
) *AdminV1AiIndexAudioService {
	l := log.NewHelper(log.With(logger, "module", "service/aiIndexAudio"))
	return &AdminV1AiIndexAudioService{
		log:               l,
		aiAudioRecordRepo: aiAudioRecordRepo,
	}
}

type AdminV1AiIndexAudioService struct {
	pb.UnimplementedAiIndexAudioServer
	log               *log.Helper
	aiAudioRecordRepo *data.AiAudioRecordRepo
}
