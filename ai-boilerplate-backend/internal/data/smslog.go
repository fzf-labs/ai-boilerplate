package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewSmsLogRepo(
	logger log.Logger,
	data *Data,
	smsLogRepo *ai_boilerplate_repo.SmsLogRepo,
) *SmsLogRepo {
	l := log.NewHelper(log.With(logger, "module", "data/smsLog"))
	return &SmsLogRepo{
		log:        l,
		data:       data,
		SmsLogRepo: smsLogRepo,
	}
}

type SmsLogRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.SmsLogRepo
}
