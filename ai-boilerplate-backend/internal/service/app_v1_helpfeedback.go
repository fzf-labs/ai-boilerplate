package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAppV1HelpFeedbackService(
	logger log.Logger,
	helpFeedbackRepo *data.HelpFeedbackRepo,
) *AppV1HelpFeedbackService {
	l := log.NewHelper(log.With(logger, "module", "service/helpFeedback"))
	return &AppV1HelpFeedbackService{
		log:              l,
		helpFeedbackRepo: helpFeedbackRepo,
	}
}

type AppV1HelpFeedbackService struct {
	pb.UnimplementedHelpFeedbackServer
	log              *log.Helper
	helpFeedbackRepo *data.HelpFeedbackRepo
}
