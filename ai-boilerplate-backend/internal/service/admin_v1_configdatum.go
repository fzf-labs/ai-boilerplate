package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1ConfigDatumService(
	logger log.Logger,
	configDatumRepo *data.ConfigDatumRepo,
) *AdminV1ConfigDatumService {
	l := log.NewHelper(log.With(logger, "module", "service/configDatum"))
	return &AdminV1ConfigDatumService{
		log:             l,
		configDatumRepo: configDatumRepo,
	}
}

type AdminV1ConfigDatumService struct {
	pb.UnimplementedConfigDatumServer
	log             *log.Helper
	configDatumRepo *data.ConfigDatumRepo
}
