package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewSysNotifyMessageRepo(
	logger log.Logger,
	data *Data,
	sysNotifyMessageRepo *ai_boilerplate_repo.SysNotifyMessageRepo,
) *SysNotifyMessageRepo {
	l := log.NewHelper(log.With(logger, "module", "data/sysNotifyMessage"))
	return &SysNotifyMessageRepo{
		log:                  l,
		data:                 data,
		SysNotifyMessageRepo: sysNotifyMessageRepo,
	}
}

type SysNotifyMessageRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.SysNotifyMessageRepo
}
