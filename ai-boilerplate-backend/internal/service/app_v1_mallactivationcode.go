package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAppV1MallActivationCodeService(
	logger log.Logger,
	commonRepo *data.CommonRepo,
	userMembershipRepo *data.UserMembershipRepo,
	userMembershipChangeRepo *data.UserMembershipChangeRepo,
	mallActivationCodeRepo *data.MallActivationCodeRepo,
	mallProductRepo *data.MallProductRepo,
) *AppV1MallActivationCodeService {
	l := log.NewHelper(log.With(logger, "module", "service/app_v1_mallactivationcode"))
	return &AppV1MallActivationCodeService{
		log:                      l,
		commonRepo:               commonRepo,
		userMembershipRepo:       userMembershipRepo,
		userMembershipChangeRepo: userMembershipChangeRepo,
		mallActivationCodeRepo:   mallActivationCodeRepo,
		mallProductRepo:          mallProductRepo,
	}
}

type AppV1MallActivationCodeService struct {
	pb.UnimplementedMallActivationCodeServer
	log                      *log.Helper
	commonRepo               *data.CommonRepo
	userMembershipRepo       *data.UserMembershipRepo
	userMembershipChangeRepo *data.UserMembershipChangeRepo
	mallActivationCodeRepo   *data.MallActivationCodeRepo
	mallProductRepo          *data.MallProductRepo
}
