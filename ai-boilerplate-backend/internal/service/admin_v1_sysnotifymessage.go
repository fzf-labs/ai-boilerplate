package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1SysNotifyMessageService(
	logger log.Logger,
	sysNotifyMessageRepo *data.SysNotifyMessageRepo,
	sysAdminRepo *data.SysAdminRepo,
) *AdminV1SysNotifyMessageService {
	l := log.NewHelper(log.With(logger, "module", "service/sysNotifyMessage"))
	return &AdminV1SysNotifyMessageService{
		log:                  l,
		sysNotifyMessageRepo: sysNotifyMessageRepo,
		sysAdminRepo:         sysAdminRepo,
	}
}

type AdminV1SysNotifyMessageService struct {
	pb.UnimplementedSysNotifyMessageServer
	log                  *log.Helper
	sysNotifyMessageRepo *data.SysNotifyMessageRepo
	sysAdminRepo         *data.SysAdminRepo
}
