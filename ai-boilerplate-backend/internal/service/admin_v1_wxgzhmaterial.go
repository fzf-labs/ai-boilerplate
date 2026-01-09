package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1WxGzhMaterialService(
	logger log.Logger,
	wxGzhMaterialRepo *data.WxGzhMaterialRepo,
	wxGzhAccountRepo *data.WxGzhAccountRepo,
) *AdminV1WxGzhMaterialService {
	l := log.NewHelper(log.With(logger, "module", "service/wxGzhMaterial"))
	return &AdminV1WxGzhMaterialService{
		log:               l,
		wxGzhMaterialRepo: wxGzhMaterialRepo,
		wxGzhAccountRepo:  wxGzhAccountRepo,
	}
}

type AdminV1WxGzhMaterialService struct {
	pb.UnimplementedWxGzhMaterialServer
	log               *log.Helper
	wxGzhMaterialRepo *data.WxGzhMaterialRepo
	wxGzhAccountRepo  *data.WxGzhAccountRepo
}
