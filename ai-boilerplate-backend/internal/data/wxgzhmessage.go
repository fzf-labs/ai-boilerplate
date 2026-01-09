package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewWxGzhMessageRepo(
	logger log.Logger,
	data *Data,
	wxGzhMessageRepo *ai_boilerplate_repo.WxGzhMessageRepo,
) *WxGzhMessageRepo {
	l := log.NewHelper(log.With(logger, "module", "data/wxGzhMessage"))
	return &WxGzhMessageRepo{
		log:              l,
		data:             data,
		WxGzhMessageRepo: wxGzhMessageRepo,
	}
}

type WxGzhMessageRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.WxGzhMessageRepo
}
