package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1MallProductService(
	logger log.Logger,
	mallProductRepo *data.MallProductRepo,
) *AdminV1MallProductService {
	l := log.NewHelper(log.With(logger, "module", "service/mallProduct"))
	return &AdminV1MallProductService{
		log:             l,
		mallProductRepo: mallProductRepo,
	}
}

type AdminV1MallProductService struct {
	pb.UnimplementedMallProductServer
	log             *log.Helper
	mallProductRepo *data.MallProductRepo
}
