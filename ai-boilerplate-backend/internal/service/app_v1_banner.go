package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAppV1BannerService(
	logger log.Logger,
	bannerRepo *data.BannerRepo,
) *AppV1BannerService {
	l := log.NewHelper(log.With(logger, "module", "service/banner"))
	return &AppV1BannerService{
		log:        l,
		bannerRepo: bannerRepo,
	}
}

type AppV1BannerService struct {
	pb.UnimplementedBannerServer
	log        *log.Helper
	bannerRepo *data.BannerRepo
}
