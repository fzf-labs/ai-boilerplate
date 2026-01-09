package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1DictTypeService(
	logger log.Logger,
	dictTypeRepo *data.DictTypeRepo,
) *AdminV1DictTypeService {
	l := log.NewHelper(log.With(logger, "module", "service/dictType"))
	return &AdminV1DictTypeService{
		log:          l,
		dictTypeRepo: dictTypeRepo,
	}
}

type AdminV1DictTypeService struct {
	pb.UnimplementedDictTypeServer
	log          *log.Helper
	dictTypeRepo *data.DictTypeRepo
}
