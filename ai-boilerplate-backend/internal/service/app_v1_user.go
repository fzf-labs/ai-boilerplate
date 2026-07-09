package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAppV1UserService(
	logger log.Logger,
	commonRepo *data.CommonRepo,
	smsCodeRepo *data.SmsCodeRepo,
	userRepo *data.UserRepo,
	userMembershipRepo *data.UserMembershipRepo,
	userMembershipChangeRepo *data.UserMembershipChangeRepo,
	userMessageRepo *data.UserMessageRepo,
	helpFeedbackRepo *data.HelpFeedbackRepo,
	userNotificationSettingRepo *data.UserNotificationSettingRepo,
	userBindDeviceRepo *data.UserBindDeviceRepo,
	mallOrderRepo *data.MallOrderRepo,
	mallActivationCodeRepo *data.MallActivationCodeRepo,
	smsLogRepo *data.SmsLogRepo,
	smsTemplateRepo *data.SmsTemplateRepo,
	smsChannelRepo *data.SmsChannelRepo,
) *AppV1UserService {
	l := log.NewHelper(log.With(logger, "module", "service/user"))
	return &AppV1UserService{
		log:                         l,
		commonRepo:                  commonRepo,
		smsCodeRepo:                 smsCodeRepo,
		userRepo:                    userRepo,
		userMembershipRepo:          userMembershipRepo,
		userMembershipChangeRepo:    userMembershipChangeRepo,
		userMessageRepo:             userMessageRepo,
		helpFeedbackRepo:            helpFeedbackRepo,
		userNotificationSettingRepo: userNotificationSettingRepo,
		userBindDeviceRepo:          userBindDeviceRepo,
		mallOrderRepo:               mallOrderRepo,
		mallActivationCodeRepo:      mallActivationCodeRepo,
		smsLogRepo:                  smsLogRepo,
		smsTemplateRepo:             smsTemplateRepo,
		smsChannelRepo:              smsChannelRepo,
	}
}

type AppV1UserService struct {
	pb.UnimplementedUserServer
	log                         *log.Helper
	commonRepo                  *data.CommonRepo
	smsCodeRepo                 *data.SmsCodeRepo
	userRepo                    *data.UserRepo
	userMembershipRepo          *data.UserMembershipRepo
	userMembershipChangeRepo    *data.UserMembershipChangeRepo
	userMessageRepo             *data.UserMessageRepo
	helpFeedbackRepo            *data.HelpFeedbackRepo
	userNotificationSettingRepo *data.UserNotificationSettingRepo
	userBindDeviceRepo          *data.UserBindDeviceRepo
	mallOrderRepo               *data.MallOrderRepo
	mallActivationCodeRepo      *data.MallActivationCodeRepo
	smsLogRepo                  *data.SmsLogRepo
	smsTemplateRepo             *data.SmsTemplateRepo
	smsChannelRepo              *data.SmsChannelRepo
}
