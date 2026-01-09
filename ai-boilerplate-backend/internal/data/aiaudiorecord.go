package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAiAudioRecordRepo(
	logger log.Logger,
	data *Data,
	aiAudioRecordRepo *ai_boilerplate_repo.AiAudioRecordRepo,
) *AiAudioRecordRepo {
	l := log.NewHelper(log.With(logger, "module", "data/aiAudioRecord"))
	return &AiAudioRecordRepo{
		log:               l,
		data:              data,
		AiAudioRecordRepo: aiAudioRecordRepo,
	}
}

type AiAudioRecordRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.AiAudioRecordRepo
}
