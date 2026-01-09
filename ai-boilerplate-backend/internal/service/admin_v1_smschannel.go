package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1SmsChannelService(
	logger log.Logger,
	smsChannelRepo *data.SmsChannelRepo,
) *AdminV1SmsChannelService {
	l := log.NewHelper(log.With(logger, "module", "service/smsChannel"))
	return &AdminV1SmsChannelService{
		log:            l,
		smsChannelRepo: smsChannelRepo,
	}
}

type AdminV1SmsChannelService struct {
	pb.UnimplementedSmsChannelServer
	log            *log.Helper
	smsChannelRepo *data.SmsChannelRepo
}
