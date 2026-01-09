package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAiWriteRecordRepo(
	logger log.Logger,
	data *Data,
	aiWriteRecordRepo *ai_boilerplate_repo.AiWriteRecordRepo,
) *AiWriteRecordRepo {
	l := log.NewHelper(log.With(logger, "module", "data/aiWriteRecord"))
	return &AiWriteRecordRepo{
		log:               l,
		data:              data,
		AiWriteRecordRepo: aiWriteRecordRepo,
	}
}

type AiWriteRecordRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.AiWriteRecordRepo
}
