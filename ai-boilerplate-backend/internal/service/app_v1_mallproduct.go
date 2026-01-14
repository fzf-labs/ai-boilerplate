package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAppV1MallProductService(
	logger log.Logger,
	mallProductRepo *data.MallProductRepo,
) *AppV1MallProductService {
	l := log.NewHelper(log.With(logger, "module", "service/mallProduct"))
	return &AppV1MallProductService{
		log:             l,
		mallProductRepo: mallProductRepo,
	}
}

type AppV1MallProductService struct {
	pb.UnimplementedMallProductServer
	log             *log.Helper
	mallProductRepo *data.MallProductRepo
}
