package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewSysNoticeRepo(
	logger log.Logger,
	data *Data,
	sysNoticeRepo *ai_boilerplate_repo.SysNoticeRepo,
) *SysNoticeRepo {
	l := log.NewHelper(log.With(logger, "module", "data/sysNotice"))
	return &SysNoticeRepo{
		log:           l,
		data:          data,
		SysNoticeRepo: sysNoticeRepo,
	}
}

type SysNoticeRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.SysNoticeRepo
}
