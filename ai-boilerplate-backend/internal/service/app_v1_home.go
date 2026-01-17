package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAppV1HomeService(
	logger log.Logger,
	bannerRepo *data.BannerRepo,
	articleRepo *data.ArticleRepo,
) *AppV1HomeService {
	l := log.NewHelper(log.With(logger, "module", "service/app_home"))
	return &AppV1HomeService{
		log:         l,
		bannerRepo:  bannerRepo,
		articleRepo: articleRepo,
	}
}

type AppV1HomeService struct {
	pb.UnimplementedHomeServer
	log         *log.Helper
	bannerRepo  *data.BannerRepo
	articleRepo *data.ArticleRepo
}
