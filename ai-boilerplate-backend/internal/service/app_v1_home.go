package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAppV1HomeService(
	logger log.Logger,
) *AppV1HomeService {
	l := log.NewHelper(log.With(logger, "module", "service/app_home"))
	return &AppV1HomeService{
		log: l,
	}
}

type AppV1HomeService struct {
	pb.UnimplementedHomeServer
	log *log.Helper
}
