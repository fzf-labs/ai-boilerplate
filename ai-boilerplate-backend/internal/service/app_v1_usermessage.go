package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAppV1UserMessageService(
	logger log.Logger,
	userMessageRepo *data.UserMessageRepo,
) *AppV1UserMessageService {
	l := log.NewHelper(log.With(logger, "module", "service/userMessage"))
	return &AppV1UserMessageService{
		log:             l,
		userMessageRepo: userMessageRepo,
	}
}

type AppV1UserMessageService struct {
	pb.UnimplementedUserMessageServer
	log             *log.Helper
	userMessageRepo *data.UserMessageRepo
}
