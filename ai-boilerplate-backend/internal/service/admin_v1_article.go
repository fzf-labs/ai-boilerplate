package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1ArticleService(
	logger log.Logger,
	articleRepo *data.ArticleRepo,
) *AdminV1ArticleService {
	l := log.NewHelper(log.With(logger, "module", "service/article"))
	return &AdminV1ArticleService{
		log:         l,
		articleRepo: articleRepo,
	}
}

type AdminV1ArticleService struct {
	pb.UnimplementedArticleServer
	log         *log.Helper
	articleRepo *data.ArticleRepo
}
