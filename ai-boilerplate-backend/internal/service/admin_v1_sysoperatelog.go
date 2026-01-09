package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1SysOperateLogService(
	logger log.Logger,
	sysOperateLogRepo *data.SysOperateLogRepo,
	sysAdminRepo *data.SysAdminRepo,
) *AdminV1SysOperateLogService {
	l := log.NewHelper(log.With(logger, "module", "service/sysOperateLog"))
	return &AdminV1SysOperateLogService{
		log:               l,
		sysOperateLogRepo: sysOperateLogRepo,
		sysAdminRepo:      sysAdminRepo,
	}
}

type AdminV1SysOperateLogService struct {
	pb.UnimplementedSysOperateLogServer
	log               *log.Helper
	sysOperateLogRepo *data.SysOperateLogRepo
	sysAdminRepo      *data.SysAdminRepo
}
