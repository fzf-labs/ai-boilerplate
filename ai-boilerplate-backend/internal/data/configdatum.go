package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewConfigDatumRepo(
	logger log.Logger,
	data *Data,
	configDatumRepo *ai_boilerplate_repo.ConfigDatumRepo,
) *ConfigDatumRepo {
	l := log.NewHelper(log.With(logger, "module", "data/configDatum"))
	return &ConfigDatumRepo{
		log:             l,
		data:            data,
		ConfigDatumRepo: configDatumRepo,
	}
}

type ConfigDatumRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.ConfigDatumRepo
}
