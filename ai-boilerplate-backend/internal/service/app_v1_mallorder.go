package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAppV1MallOrderService(
	logger log.Logger,
	commonRepo *data.CommonRepo,
	mallOrderRepo *data.MallOrderRepo,
	mallPaymentRecordRepo *data.MallPaymentRecordRepo,
	mallProductRepo *data.MallProductRepo,
	userMembershipRepo *data.UserMembershipRepo,
	userMembershipChangeRepo *data.UserMembershipChangeRepo,
) *AppV1MallOrderService {
	l := log.NewHelper(log.With(logger, "module", "service/mallOrder"))
	return &AppV1MallOrderService{
		log:                      l,
		commonRepo:               commonRepo,
		mallOrderRepo:            mallOrderRepo,
		mallPaymentRecordRepo:    mallPaymentRecordRepo,
		mallProductRepo:          mallProductRepo,
		userMembershipRepo:       userMembershipRepo,
		userMembershipChangeRepo: userMembershipChangeRepo,
	}
}

type AppV1MallOrderService struct {
	pb.UnimplementedMallOrderServer
	log                      *log.Helper
	commonRepo               *data.CommonRepo
	mallOrderRepo            *data.MallOrderRepo
	mallPaymentRecordRepo    *data.MallPaymentRecordRepo
	mallProductRepo          *data.MallProductRepo
	userMembershipRepo       *data.UserMembershipRepo
	userMembershipChangeRepo *data.UserMembershipChangeRepo
}
