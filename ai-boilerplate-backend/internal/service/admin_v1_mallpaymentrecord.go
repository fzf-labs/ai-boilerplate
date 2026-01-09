package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1MallPaymentRecordService(
	logger log.Logger,
	mallPaymentRecordRepo *data.MallPaymentRecordRepo,
) *AdminV1MallPaymentRecordService {
	l := log.NewHelper(log.With(logger, "module", "service/mallPaymentRecord"))
	return &AdminV1MallPaymentRecordService{
		log:                   l,
		mallPaymentRecordRepo: mallPaymentRecordRepo,
	}
}

type AdminV1MallPaymentRecordService struct {
	pb.UnimplementedMallPaymentRecordServer
	log                   *log.Helper
	mallPaymentRecordRepo *data.MallPaymentRecordRepo
}
