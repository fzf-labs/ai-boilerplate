package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAppV1HelpCategoryService(
	logger log.Logger,
	helpCategoryRepo *data.HelpCategoryRepo,
) *AppV1HelpCategoryService {
	l := log.NewHelper(log.With(logger, "module", "service/helpCategory"))
	return &AppV1HelpCategoryService{
		log:              l,
		helpCategoryRepo: helpCategoryRepo,
	}
}

type AppV1HelpCategoryService struct {
	pb.UnimplementedHelpCategoryServer
	log              *log.Helper
	helpCategoryRepo *data.HelpCategoryRepo
}
