package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1SysAdminService(
	logger log.Logger,
	sysAdminRepo *data.SysAdminRepo,
	sysRoleRepo *data.SysRoleRepo,
	sysDeptRepo *data.SysDeptRepo,
	sysPostRepo *data.SysPostRepo,
) *AdminV1SysAdminService {
	l := log.NewHelper(log.With(logger, "module", "service/sysAdmin"))
	return &AdminV1SysAdminService{
		log:          l,
		sysAdminRepo: sysAdminRepo,
		sysRoleRepo:  sysRoleRepo,
		sysDeptRepo:  sysDeptRepo,
		sysPostRepo:  sysPostRepo,
	}
}

type AdminV1SysAdminService struct {
	pb.UnimplementedSysAdminServer
	log          *log.Helper
	sysAdminRepo *data.SysAdminRepo
	sysRoleRepo  *data.SysRoleRepo
	sysDeptRepo  *data.SysDeptRepo
	sysPostRepo  *data.SysPostRepo
}
