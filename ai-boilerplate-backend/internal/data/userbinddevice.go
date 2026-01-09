package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewUserBindDeviceRepo(
	logger log.Logger,
	data *Data,
	userBindDeviceRepo *ai_boilerplate_repo.UserBindDeviceRepo,
) *UserBindDeviceRepo {
	l := log.NewHelper(log.With(logger, "module", "data/userBindDevice"))
	return &UserBindDeviceRepo{
		log:                l,
		data:               data,
		UserBindDeviceRepo: userBindDeviceRepo,
	}
}

type UserBindDeviceRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.UserBindDeviceRepo
}
