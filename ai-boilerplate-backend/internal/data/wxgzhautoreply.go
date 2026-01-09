package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewWxGzhAutoReplyRepo(
	logger log.Logger,
	data *Data,
	wxGzhAutoReplyRepo *ai_boilerplate_repo.WxGzhAutoReplyRepo,
) *WxGzhAutoReplyRepo {
	l := log.NewHelper(log.With(logger, "module", "data/wxGzhAutoReply"))
	return &WxGzhAutoReplyRepo{
		log:                l,
		data:               data,
		WxGzhAutoReplyRepo: wxGzhAutoReplyRepo,
	}
}

type WxGzhAutoReplyRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.WxGzhAutoReplyRepo
}
