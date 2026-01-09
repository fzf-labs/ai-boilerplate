package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1SysRoleService(
	logger log.Logger,
	sysRoleRepo *data.SysRoleRepo,
) *AdminV1SysRoleService {
	l := log.NewHelper(log.With(logger, "module", "service/sysRole"))
	return &AdminV1SysRoleService{
		log:         l,
		sysRoleRepo: sysRoleRepo,
	}
}

type AdminV1SysRoleService struct {
	pb.UnimplementedSysRoleServer
	log         *log.Helper
	sysRoleRepo *data.SysRoleRepo
}
