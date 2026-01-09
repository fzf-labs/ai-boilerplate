package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1WxGzhUserService(
	logger log.Logger,
	wxGzhUserRepo *data.WxGzhUserRepo,
) *AdminV1WxGzhUserService {
	l := log.NewHelper(log.With(logger, "module", "service/wxGzhUser"))
	return &AdminV1WxGzhUserService{
		log:           l,
		wxGzhUserRepo: wxGzhUserRepo,
	}
}

type AdminV1WxGzhUserService struct {
	pb.UnimplementedWxGzhUserServer
	log           *log.Helper
	wxGzhUserRepo *data.WxGzhUserRepo
}
