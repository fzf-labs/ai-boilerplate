package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1SysAuthService(
	logger log.Logger,
	sysAdminRepo *data.SysAdminRepo,
	sysMenuRepo *data.SysMenuRepo,
	sysRoleRepo *data.SysRoleRepo,
	sysDeptRepo *data.SysDeptRepo,
	sysPostRepo *data.SysPostRepo,
) *AdminV1SysAuthService {
	l := log.NewHelper(log.With(logger, "module", "service/sysAuth"))
	return &AdminV1SysAuthService{
		log:          l,
		sysAdminRepo: sysAdminRepo,
		sysMenuRepo:  sysMenuRepo,
		sysRoleRepo:  sysRoleRepo,
		sysDeptRepo:  sysDeptRepo,
		sysPostRepo:  sysPostRepo,
	}
}

type AdminV1SysAuthService struct {
	pb.UnimplementedSysAuthServer
	log          *log.Helper
	sysAdminRepo *data.SysAdminRepo
	sysMenuRepo  *data.SysMenuRepo
	sysRoleRepo  *data.SysRoleRepo
	sysDeptRepo  *data.SysDeptRepo
	sysPostRepo  *data.SysPostRepo
}
