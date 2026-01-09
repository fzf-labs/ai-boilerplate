package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewFileDatumRepo(
	logger log.Logger,
	data *Data,
	fileDatumRepo *ai_boilerplate_repo.FileDatumRepo,
) *FileDatumRepo {
	l := log.NewHelper(log.With(logger, "module", "data/fileDatum"))
	return &FileDatumRepo{
		log:           l,
		data:          data,
		FileDatumRepo: fileDatumRepo,
	}
}

type FileDatumRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.FileDatumRepo
}
