package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1WxGzhMenuService(
	logger log.Logger,
	wxGzhMenuRepo *data.WxGzhMenuRepo,
	wxGzhAccountRepo *data.WxGzhAccountRepo,
) *AdminV1WxGzhMenuService {
	l := log.NewHelper(log.With(logger, "module", "service/wxGzhMenu"))
	return &AdminV1WxGzhMenuService{
		log:              l,
		wxGzhMenuRepo:    wxGzhMenuRepo,
		wxGzhAccountRepo: wxGzhAccountRepo,
	}
}

type AdminV1WxGzhMenuService struct {
	pb.UnimplementedWxGzhMenuServer
	log              *log.Helper
	wxGzhMenuRepo    *data.WxGzhMenuRepo
	wxGzhAccountRepo *data.WxGzhAccountRepo
}
