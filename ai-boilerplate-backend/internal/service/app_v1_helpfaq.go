package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAppV1HelpFaqService(
	logger log.Logger,
	helpFaqRepo *data.HelpFaqRepo,
) *AppV1HelpFaqService {
	l := log.NewHelper(log.With(logger, "module", "service/helpFaq"))
	return &AppV1HelpFaqService{
		log:         l,
		helpFaqRepo: helpFaqRepo,
	}
}

type AppV1HelpFaqService struct {
	pb.UnimplementedHelpFaqServer
	log         *log.Helper
	helpFaqRepo *data.HelpFaqRepo
}
