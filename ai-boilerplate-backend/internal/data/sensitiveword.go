package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewSensitiveWordRepo(
	logger log.Logger,
	data *Data,
	sensitiveWordRepo *ai_boilerplate_repo.SensitiveWordRepo,
) *SensitiveWordRepo {
	l := log.NewHelper(log.With(logger, "module", "data/sensitiveWord"))
	return &SensitiveWordRepo{
		log:               l,
		data:              data,
		SensitiveWordRepo: sensitiveWordRepo,
	}
}

type SensitiveWordRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.SensitiveWordRepo
}
