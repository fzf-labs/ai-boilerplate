package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1WxXcxUserService(
	logger log.Logger,
	wxXcxUserRepo *data.WxXcxUserRepo,
) *AdminV1WxXcxUserService {
	l := log.NewHelper(log.With(logger, "module", "service/wxXcxUser"))
	return &AdminV1WxXcxUserService{
		log:           l,
		wxXcxUserRepo: wxXcxUserRepo,
	}
}

type AdminV1WxXcxUserService struct {
	pb.UnimplementedWxXcxUserServer
	log           *log.Helper
	wxXcxUserRepo *data.WxXcxUserRepo
}
