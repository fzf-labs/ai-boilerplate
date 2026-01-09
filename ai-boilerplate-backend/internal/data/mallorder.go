package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewMallOrderRepo(
	logger log.Logger,
	data *Data,
	mallOrderRepo *ai_boilerplate_repo.MallOrderRepo,
) *MallOrderRepo {
	l := log.NewHelper(log.With(logger, "module", "data/mallOrder"))
	return &MallOrderRepo{
		log:           l,
		data:          data,
		MallOrderRepo: mallOrderRepo,
	}
}

type MallOrderRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.MallOrderRepo
}
