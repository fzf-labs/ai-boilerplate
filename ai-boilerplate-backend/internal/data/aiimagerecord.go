package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAiImageRecordRepo(
	logger log.Logger,
	data *Data,
	aiImageRecordRepo *ai_boilerplate_repo.AiImageRecordRepo,
) *AiImageRecordRepo {
	l := log.NewHelper(log.With(logger, "module", "data/aiImageRecord"))
	return &AiImageRecordRepo{
		log:               l,
		data:              data,
		AiImageRecordRepo: aiImageRecordRepo,
	}
}

type AiImageRecordRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.AiImageRecordRepo
}
