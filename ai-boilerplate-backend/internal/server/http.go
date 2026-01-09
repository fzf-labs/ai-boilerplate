package server

import (
	adminv1 "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	appv1 "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/middleware/auth"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/service"
	conf "github.com/fzf-labs/kratos-contrib/api/conf/v1"
	"github.com/fzf-labs/kratos-contrib/bootstrap"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	c *conf.Bootstrap,
	logger log.Logger,
	// admin
	adminV1SysAuthService *service.AdminV1SysAuthService,
	adminV1SysTenantService *service.AdminV1SysTenantService,
	adminV1SysAdminService *service.AdminV1SysAdminService,
	adminV1SysMenuService *service.AdminV1SysMenuService,
	adminV1SysRoleService *service.AdminV1SysRoleService,
	adminV1SysDeptService *service.AdminV1SysDeptService,
	adminV1SysPostService *service.AdminV1SysPostService,
	adminV1SysAPIService *service.AdminV1SysAPIService,
	adminV1SysOperateLogService *service.AdminV1SysOperateLogService,
	adminV1DictTypeService *service.AdminV1DictTypeService,
	adminV1DictDatumService *service.AdminV1DictDatumService,
	adminV1SysNotifyMessageService *service.AdminV1SysNotifyMessageService,
	adminV1SysNoticeService *service.AdminV1SysNoticeService,
	adminV1SmsChannelService *service.AdminV1SmsChannelService,
	adminV1SmsTemplateService *service.AdminV1SmsTemplateService,
	adminV1SmsLogService *service.AdminV1SmsLogService,
	adminV1MailAccountService *service.AdminV1MailAccountService,
	adminV1MailTemplateService *service.AdminV1MailTemplateService,
	adminV1MailLogService *service.AdminV1MailLogService,
	adminV1ConfigDatumService *service.AdminV1ConfigDatumService,
	adminV1FileConfigService *service.AdminV1FileConfigService,
	adminV1FileDatumService *service.AdminV1FileDatumService,
	adminV1WxGzhAccountService *service.AdminV1WxGzhAccountService,
	adminV1WxGzhAutoReplyService *service.AdminV1WxGzhAutoReplyService,
	adminV1WxGzhMaterialService *service.AdminV1WxGzhMaterialService,
	adminV1WxGzhMenuService *service.AdminV1WxGzhMenuService,
	adminV1WxGzhMessageService *service.AdminV1WxGzhMessageService,
	adminV1WxGzhTagService *service.AdminV1WxGzhTagService,
	adminV1WxGzhUserService *service.AdminV1WxGzhUserService,
	adminV1WxXcxUserService *service.AdminV1WxXcxUserService,
	adminV1DeviceService *service.AdminV1DeviceService,
	adminV1SensitiveWordService *service.AdminV1SensitiveWordService,
	adminV1UserService *service.AdminV1UserService,
	adminV1UserMembershipService *service.AdminV1UserMembershipService,
	adminV1MembershipService *service.AdminV1MembershipService,
	adminV1MembershipBenefitService *service.AdminV1MembershipBenefitService,
	adminV1SelfAppService *service.AdminV1SelfAppService,
	adminV1SelfAppReleaseService *service.AdminV1SelfAppReleaseService,
	adminV1MallActivationCodeService *service.AdminV1MallActivationCodeService,
	adminV1MallOrderService *service.AdminV1MallOrderService,
	adminV1MallPaymentRecordService *service.AdminV1MallPaymentRecordService,
	adminV1MallProductService *service.AdminV1MallProductService,
	// AI 创作平台
	adminV1AiProviderModelService *service.AdminV1AiProviderModelService,
	adminV1AiProviderPlatformService *service.AdminV1AiProviderPlatformService,
	adminV1AiPromptService *service.AdminV1AiPromptService,
	adminV1AiChatConversationService *service.AdminV1AiChatConversationService,
	adminV1AiChatMessageService *service.AdminV1AiChatMessageService,
	adminV1AiImageRecordService *service.AdminV1AiImageRecordService,
	adminV1AiAudioRecordService *service.AdminV1AiAudioRecordService,
	adminV1AiVideoRecordService *service.AdminV1AiVideoRecordService,
	adminV1AiWriteRecordService *service.AdminV1AiWriteRecordService,
	adminV1AiIndexPromptService *service.AdminV1AiIndexPromptService,
	adminV1AiIndexChatService *service.AdminV1AiIndexChatService,
	// App
	appV1HomeService *service.AppV1HomeService,
	appV1UserService *service.AppV1UserService,
	appV1HelpFeedbackService *service.AppV1HelpFeedbackService,
	appV1HelpFaqService *service.AppV1HelpFaqService,
	appV1HelpCategoryService *service.AppV1HelpCategoryService,
) *http.Server {
	srv := bootstrap.NewHTTPServer(
		c,
		logger,
		auth.AdminAuthSelectorMiddleware(adminV1SysAuthService, adminV1SysOperateLogService),
		auth.AppAuthSelectorMiddleware(appV1UserService),
	)
	// Admin v1 服务注册
	adminv1.RegisterSysAuthHTTPServer(srv, adminV1SysAuthService)
	adminv1.RegisterSysTenantHTTPServer(srv, adminV1SysTenantService)
	adminv1.RegisterSysAdminHTTPServer(srv, adminV1SysAdminService)
	adminv1.RegisterSysMenuHTTPServer(srv, adminV1SysMenuService)
	adminv1.RegisterSysRoleHTTPServer(srv, adminV1SysRoleService)
	adminv1.RegisterSysDeptHTTPServer(srv, adminV1SysDeptService)
	adminv1.RegisterSysPostHTTPServer(srv, adminV1SysPostService)
	adminv1.RegisterSysAPIHTTPServer(srv, adminV1SysAPIService)
	adminv1.RegisterSysOperateLogHTTPServer(srv, adminV1SysOperateLogService)
	adminv1.RegisterDictTypeHTTPServer(srv, adminV1DictTypeService)
	adminv1.RegisterDictDatumHTTPServer(srv, adminV1DictDatumService)
	adminv1.RegisterSysNotifyMessageHTTPServer(srv, adminV1SysNotifyMessageService)
	adminv1.RegisterSysNoticeHTTPServer(srv, adminV1SysNoticeService)
	adminv1.RegisterSmsChannelHTTPServer(srv, adminV1SmsChannelService)
	adminv1.RegisterSmsTemplateHTTPServer(srv, adminV1SmsTemplateService)
	adminv1.RegisterSmsLogHTTPServer(srv, adminV1SmsLogService)
	adminv1.RegisterMailAccountHTTPServer(srv, adminV1MailAccountService)
	adminv1.RegisterMailTemplateHTTPServer(srv, adminV1MailTemplateService)
	adminv1.RegisterMailLogHTTPServer(srv, adminV1MailLogService)
	adminv1.RegisterConfigDatumHTTPServer(srv, adminV1ConfigDatumService)
	adminv1.RegisterFileConfigHTTPServer(srv, adminV1FileConfigService)
	adminv1.RegisterFileDatumHTTPServer(srv, adminV1FileDatumService)
	adminv1.RegisterWxGzhAccountHTTPServer(srv, adminV1WxGzhAccountService)
	adminv1.RegisterWxGzhAutoReplyHTTPServer(srv, adminV1WxGzhAutoReplyService)
	adminv1.RegisterWxGzhMaterialHTTPServer(srv, adminV1WxGzhMaterialService)
	adminv1.RegisterWxGzhMenuHTTPServer(srv, adminV1WxGzhMenuService)
	adminv1.RegisterWxGzhMessageHTTPServer(srv, adminV1WxGzhMessageService)
	adminv1.RegisterWxGzhTagHTTPServer(srv, adminV1WxGzhTagService)
	adminv1.RegisterWxGzhUserHTTPServer(srv, adminV1WxGzhUserService)
	adminv1.RegisterWxXcxUserHTTPServer(srv, adminV1WxXcxUserService)
	adminv1.RegisterDeviceHTTPServer(srv, adminV1DeviceService)
	adminv1.RegisterSensitiveWordHTTPServer(srv, adminV1SensitiveWordService)
	adminv1.RegisterUserHTTPServer(srv, adminV1UserService)
	adminv1.RegisterUserMembershipHTTPServer(srv, adminV1UserMembershipService)
	adminv1.RegisterMembershipHTTPServer(srv, adminV1MembershipService)
	adminv1.RegisterMembershipBenefitHTTPServer(srv, adminV1MembershipBenefitService)
	adminv1.RegisterSelfAppHTTPServer(srv, adminV1SelfAppService)
	adminv1.RegisterSelfAppReleaseHTTPServer(srv, adminV1SelfAppReleaseService)
	adminv1.RegisterAiProviderPlatformHTTPServer(srv, adminV1AiProviderPlatformService)
	adminv1.RegisterAiPromptHTTPServer(srv, adminV1AiPromptService)
	adminv1.RegisterAiIndexPromptHTTPServer(srv, adminV1AiIndexPromptService)
	adminv1.RegisterAiChatConversationHTTPServer(srv, adminV1AiChatConversationService)
	adminv1.RegisterAiChatMessageHTTPServer(srv, adminV1AiChatMessageService)
	adminv1.RegisterAiProviderModelHTTPServer(srv, adminV1AiProviderModelService)
	adminv1.RegisterAiImageRecordHTTPServer(srv, adminV1AiImageRecordService)
	adminv1.RegisterAiAudioRecordHTTPServer(srv, adminV1AiAudioRecordService)
	adminv1.RegisterAiVideoRecordHTTPServer(srv, adminV1AiVideoRecordService)
	adminv1.RegisterAiWriteRecordHTTPServer(srv, adminV1AiWriteRecordService)
	adminv1.RegisterMallActivationCodeHTTPServer(srv, adminV1MallActivationCodeService)
	adminv1.RegisterMallOrderHTTPServer(srv, adminV1MallOrderService)
	adminv1.RegisterMallPaymentRecordHTTPServer(srv, adminV1MallPaymentRecordService)
	adminv1.RegisterMallProductHTTPServer(srv, adminV1MallProductService)
	adminv1.RegisterAiIndexChatHTTPServer(srv, adminV1AiIndexChatService)
	// App v1 服务注册
	appv1.RegisterHomeHTTPServer(srv, appV1HomeService)
	appv1.RegisterUserHTTPServer(srv, appV1UserService)
	appv1.RegisterHelpFeedbackHTTPServer(srv, appV1HelpFeedbackService)
	appv1.RegisterHelpFaqHTTPServer(srv, appV1HelpFaqService)
	appv1.RegisterHelpCategoryHTTPServer(srv, appV1HelpCategoryService)
	// 自定义路由
	adminRoute := srv.Route("/admin")
	adminRoute.POST("/v1/ai_index_chat/completions", adminV1AiIndexChatService.AiIndexChatCompletionsHandler) // AI 聊天-聊天 ChatCompletions格式 (SSE 流式返回)
	adminRoute.POST("/v1/wx_gzh_material/upload", adminV1WxGzhMaterialService.UploadWxGzhMaterialHandler)     // 上传素材
	srv.HandleFunc("/wx_gzh_account/callback", adminV1WxGzhAccountService.OfficialAccountCallback)            // 公众号回调

	return srv
}
