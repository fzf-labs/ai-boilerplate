package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewWxXcxUserRepo(
	logger log.Logger,
	data *Data,
	wxXcxUserRepo *ai_boilerplate_repo.WxXcxUserRepo,
) *WxXcxUserRepo {
	l := log.NewHelper(log.With(logger, "module", "data/wxXcxUser"))
	return &WxXcxUserRepo{
		log:           l,
		data:          data,
		WxXcxUserRepo: wxXcxUserRepo,
	}
}

type WxXcxUserRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.WxXcxUserRepo
}
