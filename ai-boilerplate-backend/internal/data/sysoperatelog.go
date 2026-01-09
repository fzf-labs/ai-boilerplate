package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewSysOperateLogRepo(
	logger log.Logger,
	data *Data,
	sysOperateLogRepo *ai_boilerplate_repo.SysOperateLogRepo,
) *SysOperateLogRepo {
	l := log.NewHelper(log.With(logger, "module", "data/sysOperateLog"))
	return &SysOperateLogRepo{
		log:               l,
		data:              data,
		SysOperateLogRepo: sysOperateLogRepo,
	}
}

type SysOperateLogRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.SysOperateLogRepo
}
