package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1MailLogService(
	logger log.Logger,
	mailLogRepo *data.MailLogRepo,
) *AdminV1MailLogService {
	l := log.NewHelper(log.With(logger, "module", "service/mailLog"))
	return &AdminV1MailLogService{
		log:         l,
		mailLogRepo: mailLogRepo,
	}
}

type AdminV1MailLogService struct {
	pb.UnimplementedMailLogServer
	log         *log.Helper
	mailLogRepo *data.MailLogRepo
}
