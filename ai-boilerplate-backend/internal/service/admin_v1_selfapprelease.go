package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1SelfAppReleaseService(
	logger log.Logger,
	selfAppReleaseRepo *data.SelfAppReleaseRepo,
) *AdminV1SelfAppReleaseService {
	l := log.NewHelper(log.With(logger, "module", "service/selfAppRelease"))
	return &AdminV1SelfAppReleaseService{
		log:                l,
		selfAppReleaseRepo: selfAppReleaseRepo,
	}
}

type AdminV1SelfAppReleaseService struct {
	pb.UnimplementedSelfAppReleaseServer
	log                *log.Helper
	selfAppReleaseRepo *data.SelfAppReleaseRepo
}
