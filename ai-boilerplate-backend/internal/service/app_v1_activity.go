package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAppV1ActivityService(
	logger log.Logger,
	activityRepo *data.ActivityRepo,
) *AppV1ActivityService {
	l := log.NewHelper(log.With(logger, "module", "service/app-activity"))
	return &AppV1ActivityService{
		log:          l,
		activityRepo: activityRepo,
	}
}

type AppV1ActivityService struct {
	pb.UnimplementedActivityServer
	log          *log.Helper
	activityRepo *data.ActivityRepo
}
