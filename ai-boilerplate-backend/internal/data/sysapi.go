package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewSysAPIRepo(
	logger log.Logger,
	data *Data,
	sysAPIRepo *ai_boilerplate_repo.SysAPIRepo,
) *SysAPIRepo {
	l := log.NewHelper(log.With(logger, "module", "data/sysAPI"))
	return &SysAPIRepo{
		log:        l,
		data:       data,
		SysAPIRepo: sysAPIRepo,
	}
}

type SysAPIRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.SysAPIRepo
}
