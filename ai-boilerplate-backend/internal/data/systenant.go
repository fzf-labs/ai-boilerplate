package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewSysTenantRepo(
	logger log.Logger,
	data *Data,
	sysTenantRepo *ai_boilerplate_repo.SysTenantRepo,
) *SysTenantRepo {
	l := log.NewHelper(log.With(logger, "module", "data/sysTenant"))
	return &SysTenantRepo{
		log:           l,
		data:          data,
		SysTenantRepo: sysTenantRepo,
	}
}

type SysTenantRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.SysTenantRepo
}
