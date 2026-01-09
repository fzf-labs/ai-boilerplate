package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1SmsLogService(
	logger log.Logger,
	smsLogRepo *data.SmsLogRepo,
) *AdminV1SmsLogService {
	l := log.NewHelper(log.With(logger, "module", "service/smsLog"))
	return &AdminV1SmsLogService{
		log:        l,
		smsLogRepo: smsLogRepo,
	}
}

type AdminV1SmsLogService struct {
	pb.UnimplementedSmsLogServer
	log        *log.Helper
	smsLogRepo *data.SmsLogRepo
}
