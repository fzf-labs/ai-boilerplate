package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1MallOrderService(
	logger log.Logger,
	mallOrderRepo *data.MallOrderRepo,
) *AdminV1MallOrderService {
	l := log.NewHelper(log.With(logger, "module", "service/mallOrder"))
	return &AdminV1MallOrderService{
		log:           l,
		mallOrderRepo: mallOrderRepo,
	}
}

type AdminV1MallOrderService struct {
	pb.UnimplementedMallOrderServer
	log           *log.Helper
	mallOrderRepo *data.MallOrderRepo
}
