package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1FileConfigService(
	logger log.Logger,
	fileConfigRepo *data.FileConfigRepo,
) *AdminV1FileConfigService {
	l := log.NewHelper(log.With(logger, "module", "service/fileConfig"))
	return &AdminV1FileConfigService{
		log:            l,
		fileConfigRepo: fileConfigRepo,
	}
}

type AdminV1FileConfigService struct {
	pb.UnimplementedFileConfigServer
	log            *log.Helper
	fileConfigRepo *data.FileConfigRepo
}
