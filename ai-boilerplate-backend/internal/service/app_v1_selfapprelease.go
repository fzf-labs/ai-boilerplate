package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAppV1SelfAppReleaseService(
	logger log.Logger,
	selfAppReleaseRepo *data.SelfAppReleaseRepo,
) *AppV1SelfAppReleaseService {
	l := log.NewHelper(log.With(logger, "module", "service/selfAppRelease"))
	return &AppV1SelfAppReleaseService{
		log:                l,
		selfAppReleaseRepo: selfAppReleaseRepo,
	}
}

type AppV1SelfAppReleaseService struct {
	pb.UnimplementedSelfAppReleaseServer
	log                *log.Helper
	selfAppReleaseRepo *data.SelfAppReleaseRepo
}
