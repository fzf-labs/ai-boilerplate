package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAiVideoRecordRepo(
	logger log.Logger,
	data *Data,
	aiVideoRecordRepo *ai_boilerplate_repo.AiVideoRecordRepo,
) *AiVideoRecordRepo {
	l := log.NewHelper(log.With(logger, "module", "data/aiVideoRecord"))
	return &AiVideoRecordRepo{
		log:               l,
		data:              data,
		AiVideoRecordRepo: aiVideoRecordRepo,
	}
}

type AiVideoRecordRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.AiVideoRecordRepo
}
