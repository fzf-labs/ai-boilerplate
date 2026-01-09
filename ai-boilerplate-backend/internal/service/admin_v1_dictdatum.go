package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1DictDatumService(
	logger log.Logger,
	dictDatumRepo *data.DictDatumRepo,
) *AdminV1DictDatumService {
	l := log.NewHelper(log.With(logger, "module", "service/dictDatum"))
	return &AdminV1DictDatumService{
		log:           l,
		dictDatumRepo: dictDatumRepo,
	}
}

type AdminV1DictDatumService struct {
	pb.UnimplementedDictDatumServer
	log           *log.Helper
	dictDatumRepo *data.DictDatumRepo
}
