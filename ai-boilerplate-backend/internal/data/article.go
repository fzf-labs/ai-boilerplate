package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewArticleRepo(
	logger log.Logger,
	data *Data,
	articleRepo *ai_boilerplate_repo.ArticleRepo,
) *ArticleRepo {
	l := log.NewHelper(log.With(logger, "module", "data/article"))
	return &ArticleRepo{
		log:         l,
		data:        data,
		ArticleRepo: articleRepo,
	}
}

type ArticleRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.ArticleRepo
}
