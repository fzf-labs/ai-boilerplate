package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1SysPostService(
	logger log.Logger,
	sysPostRepo *data.SysPostRepo,
) *AdminV1SysPostService {
	l := log.NewHelper(log.With(logger, "module", "service/sysPost"))
	return &AdminV1SysPostService{
		log:         l,
		sysPostRepo: sysPostRepo,
	}
}

type AdminV1SysPostService struct {
	pb.UnimplementedSysPostServer
	log         *log.Helper
	sysPostRepo *data.SysPostRepo
}
