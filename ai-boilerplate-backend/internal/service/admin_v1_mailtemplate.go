package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1MailTemplateService(
	logger log.Logger,
	mailTemplateRepo *data.MailTemplateRepo,
) *AdminV1MailTemplateService {
	l := log.NewHelper(log.With(logger, "module", "service/mailTemplate"))
	return &AdminV1MailTemplateService{
		log:              l,
		mailTemplateRepo: mailTemplateRepo,
	}
}

type AdminV1MailTemplateService struct {
	pb.UnimplementedMailTemplateServer
	log              *log.Helper
	mailTemplateRepo *data.MailTemplateRepo
}
