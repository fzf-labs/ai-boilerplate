package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1SelfAppService(
	logger log.Logger,
	selfAppRepo *data.SelfAppRepo,
) *AdminV1SelfAppService {
	l := log.NewHelper(log.With(logger, "module", "service/selfApp"))
	return &AdminV1SelfAppService{
		log:         l,
		selfAppRepo: selfAppRepo,
	}
}

type AdminV1SelfAppService struct {
	pb.UnimplementedSelfAppServer
	log         *log.Helper
	selfAppRepo *data.SelfAppRepo
}
