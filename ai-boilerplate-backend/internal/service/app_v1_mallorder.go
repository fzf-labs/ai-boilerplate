package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAppV1MallOrderService(
	logger log.Logger,
	mallOrderRepo *data.MallOrderRepo,
	mallProductRepo *data.MallProductRepo,
) *AppV1MallOrderService {
	l := log.NewHelper(log.With(logger, "module", "service/mallOrder"))
	return &AppV1MallOrderService{
		log:             l,
		mallOrderRepo:   mallOrderRepo,
		mallProductRepo: mallProductRepo,
	}
}

type AppV1MallOrderService struct {
	pb.UnimplementedMallOrderServer
	log             *log.Helper
	mallOrderRepo   *data.MallOrderRepo
	mallProductRepo *data.MallProductRepo
}
