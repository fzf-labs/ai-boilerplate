package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1SysNoticeService(
	logger log.Logger,
	sysNoticeRepo *data.SysNoticeRepo,
) *AdminV1SysNoticeService {
	l := log.NewHelper(log.With(logger, "module", "service/sysNotice"))
	return &AdminV1SysNoticeService{
		log:           l,
		sysNoticeRepo: sysNoticeRepo,
	}
}

type AdminV1SysNoticeService struct {
	pb.UnimplementedSysNoticeServer
	log           *log.Helper
	sysNoticeRepo *data.SysNoticeRepo
}
