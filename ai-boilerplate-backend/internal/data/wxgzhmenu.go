package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewWxGzhMenuRepo(
	logger log.Logger,
	data *Data,
	wxGzhMenuRepo *ai_boilerplate_repo.WxGzhMenuRepo,
) *WxGzhMenuRepo {
	l := log.NewHelper(log.With(logger, "module", "data/wxGzhMenu"))
	return &WxGzhMenuRepo{
		log:           l,
		data:          data,
		WxGzhMenuRepo: wxGzhMenuRepo,
	}
}

type WxGzhMenuRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.WxGzhMenuRepo
}
