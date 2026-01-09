package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewDictDatumRepo(
	logger log.Logger,
	data *Data,
	dictDatumRepo *ai_boilerplate_repo.DictDatumRepo,
) *DictDatumRepo {
	l := log.NewHelper(log.With(logger, "module", "data/dictDatum"))
	return &DictDatumRepo{
		log:           l,
		data:          data,
		DictDatumRepo: dictDatumRepo,
	}
}

type DictDatumRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.DictDatumRepo
}
