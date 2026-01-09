package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewUserNotificationSettingRepo(
	logger log.Logger,
	data *Data,
	userNotificationSettingRepo *ai_boilerplate_repo.UserNotificationSettingRepo,
) *UserNotificationSettingRepo {
	l := log.NewHelper(log.With(logger, "module", "data/userNotificationSetting"))
	return &UserNotificationSettingRepo{
		log:                         l,
		data:                        data,
		UserNotificationSettingRepo: userNotificationSettingRepo,
	}
}

type UserNotificationSettingRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.UserNotificationSettingRepo
}
