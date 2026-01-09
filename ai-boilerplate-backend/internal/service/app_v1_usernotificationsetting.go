package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAppV1UserNotificationSettingService(
	logger log.Logger,
	userNotificationSettingRepo *data.UserNotificationSettingRepo,
) *AppV1UserNotificationSettingService {
	l := log.NewHelper(log.With(logger, "module", "service/userNotificationSetting"))
	return &AppV1UserNotificationSettingService{
		log:                         l,
		userNotificationSettingRepo: userNotificationSettingRepo,
	}
}

type AppV1UserNotificationSettingService struct {
	pb.UnimplementedUserNotificationSettingServer
	log                         *log.Helper
	userNotificationSettingRepo *data.UserNotificationSettingRepo
}
