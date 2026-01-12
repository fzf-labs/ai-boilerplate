package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewBannerRepo(
	logger log.Logger,
	data *Data,
	bannerRepo *ai_boilerplate_repo.BannerRepo,
) *BannerRepo {
	l := log.NewHelper(log.With(logger, "module", "data/banner"))
	return &BannerRepo{
		log:        l,
		data:       data,
		BannerRepo: bannerRepo,
	}
}

type BannerRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.BannerRepo
}
