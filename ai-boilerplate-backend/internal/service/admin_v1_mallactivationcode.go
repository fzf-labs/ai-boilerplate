package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1MallActivationCodeService(
	logger log.Logger,
	mallActivationCodeRepo *data.MallActivationCodeRepo,
	mallProductRepo *data.MallProductRepo,
	userRepo *data.UserRepo,
) *AdminV1MallActivationCodeService {
	l := log.NewHelper(log.With(logger, "module", "service/mallActivationCode"))
	return &AdminV1MallActivationCodeService{
		log:                    l,
		mallActivationCodeRepo: mallActivationCodeRepo,
		mallProductRepo:        mallProductRepo,
		userRepo:               userRepo,
	}
}

type AdminV1MallActivationCodeService struct {
	pb.UnimplementedMallActivationCodeServer
	log                    *log.Helper
	mallActivationCodeRepo *data.MallActivationCodeRepo
	mallProductRepo        *data.MallProductRepo
	userRepo               *data.UserRepo
}
