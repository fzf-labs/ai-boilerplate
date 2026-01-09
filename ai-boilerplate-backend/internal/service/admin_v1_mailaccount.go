package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1MailAccountService(
	logger log.Logger,
	mailAccountRepo *data.MailAccountRepo,
) *AdminV1MailAccountService {
	l := log.NewHelper(log.With(logger, "module", "service/mailAccount"))
	return &AdminV1MailAccountService{
		log:             l,
		mailAccountRepo: mailAccountRepo,
	}
}

type AdminV1MailAccountService struct {
	pb.UnimplementedMailAccountServer
	log             *log.Helper
	mailAccountRepo *data.MailAccountRepo
}
