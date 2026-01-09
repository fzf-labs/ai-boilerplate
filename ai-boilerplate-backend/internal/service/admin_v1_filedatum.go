package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1FileDatumService(
	logger log.Logger,
	fileConfigRepo *data.FileConfigRepo,
	fileDatumRepo *data.FileDatumRepo,
) *AdminV1FileDatumService {
	l := log.NewHelper(log.With(logger, "module", "service/fileDatum"))
	return &AdminV1FileDatumService{
		log:            l,
		fileConfigRepo: fileConfigRepo,
		fileDatumRepo:  fileDatumRepo,
	}
}

type AdminV1FileDatumService struct {
	pb.UnimplementedFileDatumServer
	log            *log.Helper
	fileConfigRepo *data.FileConfigRepo
	fileDatumRepo  *data.FileDatumRepo
}
