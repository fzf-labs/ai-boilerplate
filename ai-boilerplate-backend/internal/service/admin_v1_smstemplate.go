package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1SmsTemplateService(
	logger log.Logger,
	smsTemplateRepo *data.SmsTemplateRepo,
	smsChannelRepo *data.SmsChannelRepo,
) *AdminV1SmsTemplateService {
	l := log.NewHelper(log.With(logger, "module", "service/smsTemplate"))
	return &AdminV1SmsTemplateService{
		log:             l,
		smsTemplateRepo: smsTemplateRepo,
		smsChannelRepo:  smsChannelRepo,
	}
}

type AdminV1SmsTemplateService struct {
	pb.UnimplementedSmsTemplateServer
	log             *log.Helper
	smsTemplateRepo *data.SmsTemplateRepo
	smsChannelRepo  *data.SmsChannelRepo
}
