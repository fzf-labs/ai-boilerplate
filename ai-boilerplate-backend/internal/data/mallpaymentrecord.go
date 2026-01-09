package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewMallPaymentRecordRepo(
	logger log.Logger,
	data *Data,
	mallPaymentRecordRepo *ai_boilerplate_repo.MallPaymentRecordRepo,
) *MallPaymentRecordRepo {
	l := log.NewHelper(log.With(logger, "module", "data/mallPaymentRecord"))
	return &MallPaymentRecordRepo{
		log:                   l,
		data:                  data,
		MallPaymentRecordRepo: mallPaymentRecordRepo,
	}
}

type MallPaymentRecordRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.MallPaymentRecordRepo
}
