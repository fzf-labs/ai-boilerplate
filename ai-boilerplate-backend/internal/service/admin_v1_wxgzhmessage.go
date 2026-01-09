package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1WxGzhMessageService(
	logger log.Logger,
	wxGzhMessageRepo *data.WxGzhMessageRepo,
) *AdminV1WxGzhMessageService {
	l := log.NewHelper(log.With(logger, "module", "service/wxGzhMessage"))
	return &AdminV1WxGzhMessageService{
		log:              l,
		wxGzhMessageRepo: wxGzhMessageRepo,
	}
}

type AdminV1WxGzhMessageService struct {
	pb.UnimplementedWxGzhMessageServer
	log              *log.Helper
	wxGzhMessageRepo *data.WxGzhMessageRepo
}
