package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewWxGzhTagRepo(
	logger log.Logger,
	data *Data,
	wxGzhTagRepo *ai_boilerplate_repo.WxGzhTagRepo,
) *WxGzhTagRepo {
	l := log.NewHelper(log.With(logger, "module", "data/wxGzhTag"))
	return &WxGzhTagRepo{
		log:          l,
		data:         data,
		WxGzhTagRepo: wxGzhTagRepo,
	}
}

type WxGzhTagRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.WxGzhTagRepo
}
