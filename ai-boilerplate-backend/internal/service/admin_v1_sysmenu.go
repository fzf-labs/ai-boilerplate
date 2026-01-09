package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1SysMenuService(
	logger log.Logger,
	sysMenuRepo *data.SysMenuRepo,
) *AdminV1SysMenuService {
	l := log.NewHelper(log.With(logger, "module", "service/sysMenu"))
	return &AdminV1SysMenuService{
		log:         l,
		sysMenuRepo: sysMenuRepo,
	}
}

type AdminV1SysMenuService struct {
	pb.UnimplementedSysMenuServer
	log         *log.Helper
	sysMenuRepo *data.SysMenuRepo
}
