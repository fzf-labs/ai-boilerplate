package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1SysDeptService(
	logger log.Logger,
	sysDeptRepo *data.SysDeptRepo,
	sysAdminRepo *data.SysAdminRepo,
) *AdminV1SysDeptService {
	l := log.NewHelper(log.With(logger, "module", "service/sysDept"))
	return &AdminV1SysDeptService{
		log:          l,
		sysDeptRepo:  sysDeptRepo,
		sysAdminRepo: sysAdminRepo,
	}
}

type AdminV1SysDeptService struct {
	pb.UnimplementedSysDeptServer
	log          *log.Helper
	sysDeptRepo  *data.SysDeptRepo
	sysAdminRepo *data.SysAdminRepo
}
