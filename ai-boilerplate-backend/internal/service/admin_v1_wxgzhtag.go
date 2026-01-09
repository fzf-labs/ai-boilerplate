package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1WxGzhTagService(
	logger log.Logger,
	wxGzhTagRepo *data.WxGzhTagRepo,
	wxGzhAccountRepo *data.WxGzhAccountRepo,
) *AdminV1WxGzhTagService {
	l := log.NewHelper(log.With(logger, "module", "service/wxGzhTag"))
	return &AdminV1WxGzhTagService{
		log:              l,
		wxGzhTagRepo:     wxGzhTagRepo,
		wxGzhAccountRepo: wxGzhAccountRepo,
	}
}

type AdminV1WxGzhTagService struct {
	pb.UnimplementedWxGzhTagServer
	log              *log.Helper
	wxGzhTagRepo     *data.WxGzhTagRepo
	wxGzhAccountRepo *data.WxGzhAccountRepo
}
