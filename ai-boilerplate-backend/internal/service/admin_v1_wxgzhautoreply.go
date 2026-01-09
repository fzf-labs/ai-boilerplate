package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1WxGzhAutoReplyService(
	logger log.Logger,
	wxGzhAutoReplyRepo *data.WxGzhAutoReplyRepo,
	wxGzhAccountRepo *data.WxGzhAccountRepo,
) *AdminV1WxGzhAutoReplyService {
	l := log.NewHelper(log.With(logger, "module", "service/wxGzhAutoReply"))
	return &AdminV1WxGzhAutoReplyService{
		log:                l,
		wxGzhAutoReplyRepo: wxGzhAutoReplyRepo,
		wxGzhAccountRepo:   wxGzhAccountRepo,
	}
}

type AdminV1WxGzhAutoReplyService struct {
	pb.UnimplementedWxGzhAutoReplyServer
	log                *log.Helper
	wxGzhAutoReplyRepo *data.WxGzhAutoReplyRepo
	wxGzhAccountRepo   *data.WxGzhAccountRepo
}
