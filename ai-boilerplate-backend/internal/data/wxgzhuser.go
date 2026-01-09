package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewWxGzhUserRepo(
	logger log.Logger,
	data *Data,
	wxGzhUserRepo *ai_boilerplate_repo.WxGzhUserRepo,
) *WxGzhUserRepo {
	l := log.NewHelper(log.With(logger, "module", "data/wxGzhUser"))
	return &WxGzhUserRepo{
		log:           l,
		data:          data,
		WxGzhUserRepo: wxGzhUserRepo,
	}
}

type WxGzhUserRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.WxGzhUserRepo
}
