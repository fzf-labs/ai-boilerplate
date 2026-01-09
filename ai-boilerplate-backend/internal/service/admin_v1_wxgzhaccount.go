package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1WxGzhAccountService(
	logger log.Logger,
	wxGzhAccountRepo *data.WxGzhAccountRepo,
	wxGzhUserRepo *data.WxGzhUserRepo,
) *AdminV1WxGzhAccountService {
	l := log.NewHelper(log.With(logger, "module", "service/wxGzhAccount"))
	return &AdminV1WxGzhAccountService{
		log:              l,
		wxGzhAccountRepo: wxGzhAccountRepo,
		wxGzhUserRepo:    wxGzhUserRepo,
	}
}

type AdminV1WxGzhAccountService struct {
	pb.UnimplementedWxGzhAccountServer
	log              *log.Helper
	wxGzhAccountRepo *data.WxGzhAccountRepo
	wxGzhUserRepo    *data.WxGzhUserRepo
}
