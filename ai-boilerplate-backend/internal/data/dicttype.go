package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewDictTypeRepo(
	logger log.Logger,
	data *Data,
	dictTypeRepo *ai_boilerplate_repo.DictTypeRepo,
) *DictTypeRepo {
	l := log.NewHelper(log.With(logger, "module", "data/dictType"))
	return &DictTypeRepo{
		log:          l,
		data:         data,
		DictTypeRepo: dictTypeRepo,
	}
}

type DictTypeRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.DictTypeRepo
}
