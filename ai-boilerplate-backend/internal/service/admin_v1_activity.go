package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1ActivityService(
	logger log.Logger,
	activityRepo *data.ActivityRepo,
) *AdminV1ActivityService {
	l := log.NewHelper(log.With(logger, "module", "service/activity"))
	return &AdminV1ActivityService{
		log:          l,
		activityRepo: activityRepo,
	}
}

type AdminV1ActivityService struct {
	pb.UnimplementedActivityServer
	log          *log.Helper
	activityRepo *data.ActivityRepo
}
