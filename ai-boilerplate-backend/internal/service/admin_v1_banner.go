package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1BannerService(
	logger log.Logger,
	bannerRepo *data.BannerRepo,
) *AdminV1BannerService {
	l := log.NewHelper(log.With(logger, "module", "service/banner"))
	return &AdminV1BannerService{
		log:        l,
		bannerRepo: bannerRepo,
	}
}

type AdminV1BannerService struct {
	pb.UnimplementedBannerServer
	log        *log.Helper
	bannerRepo *data.BannerRepo
}
