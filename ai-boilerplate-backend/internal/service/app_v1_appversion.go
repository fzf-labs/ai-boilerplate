package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAppV1AppVersionService(
	logger log.Logger,
	selfAppReleaseRepo *data.SelfAppReleaseRepo,
) *AppV1AppVersionService {
	l := log.NewHelper(log.With(logger, "module", "service/appVersion"))
	return &AppV1AppVersionService{
		log:                l,
		selfAppReleaseRepo: selfAppReleaseRepo,
	}
}

type AppV1AppVersionService struct {
	pb.UnimplementedAppVersionServer
	log                *log.Helper
	selfAppReleaseRepo *data.SelfAppReleaseRepo
}
