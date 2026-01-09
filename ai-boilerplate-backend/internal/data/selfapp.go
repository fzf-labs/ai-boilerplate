package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewSelfAppRepo(
	logger log.Logger,
	data *Data,
	selfAppRepo *ai_boilerplate_repo.SelfAppRepo,
) *SelfAppRepo {
	l := log.NewHelper(log.With(logger, "module", "data/selfApp"))
	return &SelfAppRepo{
		log:         l,
		data:        data,
		SelfAppRepo: selfAppRepo,
	}
}

type SelfAppRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.SelfAppRepo
}
