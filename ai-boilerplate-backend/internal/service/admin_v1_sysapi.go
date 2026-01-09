package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1SysAPIService(
	logger log.Logger,
	sysAPIRepo *data.SysAPIRepo,
) *AdminV1SysAPIService {
	l := log.NewHelper(log.With(logger, "module", "service/sysAPI"))
	return &AdminV1SysAPIService{
		log:        l,
		sysAPIRepo: sysAPIRepo,
	}
}

type AdminV1SysAPIService struct {
	pb.UnimplementedSysAPIServer
	log        *log.Helper
	sysAPIRepo *data.SysAPIRepo
}
