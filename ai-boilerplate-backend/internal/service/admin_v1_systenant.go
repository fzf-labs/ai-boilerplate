package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1SysTenantService(
	logger log.Logger,
	commonRepo *data.CommonRepo,
	sysTenantRepo *data.SysTenantRepo,
	sysAdminRepo *data.SysAdminRepo,
) *AdminV1SysTenantService {
	l := log.NewHelper(log.With(logger, "module", "service/sysTenant"))
	return &AdminV1SysTenantService{
		log:           l,
		commonRepo:    commonRepo,
		sysTenantRepo: sysTenantRepo,
		sysAdminRepo:  sysAdminRepo,
	}
}

type AdminV1SysTenantService struct {
	pb.UnimplementedSysTenantServer
	log           *log.Helper
	commonRepo    *data.CommonRepo
	sysTenantRepo *data.SysTenantRepo
	sysAdminRepo  *data.SysAdminRepo
}
