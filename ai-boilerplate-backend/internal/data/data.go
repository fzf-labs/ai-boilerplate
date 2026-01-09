package data

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/fzf-labs/godb/cache/rueidiscache"
	"github.com/fzf-labs/godb/orm/dbcache"
	"github.com/fzf-labs/godb/orm/dbcache/rueidisdbcache"
	"github.com/fzf-labs/godb/orm/encoding"
	"github.com/fzf-labs/godb/orm/gen/config"
	"github.com/fzf-labs/godb/orm/gormx"
	"github.com/fzf-labs/goutil/httputil"
	conf "github.com/fzf-labs/kratos-contrib/api/conf/v1"
	"github.com/fzf-labs/kratos-contrib/pkg/mq"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/hibiken/asynq"
	"github.com/redis/rueidis"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewGorm,
	NewRueidis,
	NewDBCache,
	NewCommonRepo,
	NewConfigRepo,
	NewAsynqClient,
	NewHTTPClient,
	NewDeviceHeartbeatRepo,
	NewAiAudioRecordRepo,
	NewAiChatConversationRepo,
	NewAiChatMessageRepo,
	NewAiImageRecordRepo,
	NewAiPromptRepo,
	NewAiProviderModelRepo,
	NewAiProviderPlatformRepo,
	NewAiVideoRecordRepo,
	NewAiWriteRecordRepo,
	NewConfigDatumRepo,
	NewDeviceRepo,
	NewDictDatumRepo,
	NewDictTypeRepo,
	NewFileConfigRepo,
	NewFileDatumRepo,
	NewHelpCategoryRepo,
	NewHelpFaqRepo,
	NewHelpFeedbackRepo,
	NewMailAccountRepo,
	NewMailLogRepo,
	NewMailTemplateRepo,
	NewMallActivationCodeRepo,
	NewMallOrderRepo,
	NewMallPaymentRecordRepo,
	NewMallProductRepo,
	NewMembershipBenefitRepo,
	NewMembershipRepo,
	NewSelfAppReleaseRepo,
	NewSelfAppRepo,
	NewSensitiveWordRepo,
	NewSmsChannelRepo,
	NewSmsLogRepo,
	NewSmsTemplateRepo,
	NewSysAPIRepo,
	NewSysAdminRepo,
	NewSysDeptRepo,
	NewSysMenuRepo,
	NewSysNoticeRepo,
	NewSysNotifyMessageRepo,
	NewSysOperateLogRepo,
	NewSysPostRepo,
	NewSysRoleRepo,
	NewSysTenantRepo,
	NewUserBindDeviceRepo,
	NewUserMembershipRepo,
	NewUserNotificationSettingRepo,
	NewUserRepo,
	NewWxGzhAccountRepo,
	NewWxGzhAutoReplyRepo,
	NewWxGzhMaterialRepo,
	NewWxGzhMenuRepo,
	NewWxGzhMessageRepo,
	NewWxGzhTagRepo,
	NewWxGzhUserRepo,
	NewWxXcxUserRepo,
	ai_boilerplate_repo.NewAiAudioRecordRepo,
	ai_boilerplate_repo.NewAiChatConversationRepo,
	ai_boilerplate_repo.NewAiChatMessageRepo,
	ai_boilerplate_repo.NewAiImageRecordRepo,
	ai_boilerplate_repo.NewAiPromptRepo,
	ai_boilerplate_repo.NewAiProviderModelRepo,
	ai_boilerplate_repo.NewAiProviderPlatformRepo,
	ai_boilerplate_repo.NewAiVideoRecordRepo,
	ai_boilerplate_repo.NewAiWriteRecordRepo,
	ai_boilerplate_repo.NewConfigDatumRepo,
	ai_boilerplate_repo.NewDeviceRepo,
	ai_boilerplate_repo.NewDictDatumRepo,
	ai_boilerplate_repo.NewDictTypeRepo,
	ai_boilerplate_repo.NewFileConfigRepo,
	ai_boilerplate_repo.NewFileDatumRepo,
	ai_boilerplate_repo.NewHelpCategoryRepo,
	ai_boilerplate_repo.NewHelpFaqRepo,
	ai_boilerplate_repo.NewHelpFeedbackRepo,
	ai_boilerplate_repo.NewMailAccountRepo,
	ai_boilerplate_repo.NewMailLogRepo,
	ai_boilerplate_repo.NewMailTemplateRepo,
	ai_boilerplate_repo.NewMallActivationCodeRepo,
	ai_boilerplate_repo.NewMallOrderRepo,
	ai_boilerplate_repo.NewMallPaymentRecordRepo,
	ai_boilerplate_repo.NewMallProductRepo,
	ai_boilerplate_repo.NewMembershipBenefitRepo,
	ai_boilerplate_repo.NewMembershipRepo,
	ai_boilerplate_repo.NewSelfAppReleaseRepo,
	ai_boilerplate_repo.NewSelfAppRepo,
	ai_boilerplate_repo.NewSensitiveWordRepo,
	ai_boilerplate_repo.NewSmsChannelRepo,
	ai_boilerplate_repo.NewSmsLogRepo,
	ai_boilerplate_repo.NewSmsTemplateRepo,
	ai_boilerplate_repo.NewSysAPIRepo,
	ai_boilerplate_repo.NewSysAdminRepo,
	ai_boilerplate_repo.NewSysDeptRepo,
	ai_boilerplate_repo.NewSysMenuRepo,
	ai_boilerplate_repo.NewSysNoticeRepo,
	ai_boilerplate_repo.NewSysNotifyMessageRepo,
	ai_boilerplate_repo.NewSysOperateLogRepo,
	ai_boilerplate_repo.NewSysPostRepo,
	ai_boilerplate_repo.NewSysRoleRepo,
	ai_boilerplate_repo.NewSysTenantRepo,
	ai_boilerplate_repo.NewUserBindDeviceRepo,
	ai_boilerplate_repo.NewUserMembershipRepo,
	ai_boilerplate_repo.NewUserNotificationSettingRepo,
	ai_boilerplate_repo.NewUserRepo,
	ai_boilerplate_repo.NewWxGzhAccountRepo,
	ai_boilerplate_repo.NewWxGzhAutoReplyRepo,
	ai_boilerplate_repo.NewWxGzhMaterialRepo,
	ai_boilerplate_repo.NewWxGzhMenuRepo,
	ai_boilerplate_repo.NewWxGzhMessageRepo,
	ai_boilerplate_repo.NewWxGzhTagRepo,
	ai_boilerplate_repo.NewWxGzhUserRepo,
	ai_boilerplate_repo.NewWxXcxUserRepo,
)

// Data .
type Data struct {
	cfg      *conf.Bootstrap
	logger   *log.Helper
	gorm     *gorm.DB
	rueidis  rueidis.Client
	DBCache  dbcache.IDBCache
	MQClient mq.Client
}

// NewData .
func NewData(
	c *conf.Bootstrap,
	logger log.Logger,
	gorm *gorm.DB,
	rueidis rueidis.Client,
	dbCache dbcache.IDBCache,
	mqClient mq.Client,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	d := Data{
		cfg:      c,
		logger:   &log.Helper{},
		gorm:     gorm,
		rueidis:  rueidis,
		DBCache:  dbCache,
		MQClient: mqClient,
	}
	return &d, cleanup, nil
}

// NewGorm 创建gorm实例
func NewGorm(c *conf.Bootstrap) *gorm.DB {
	gorm, err := gormx.NewPostgresGormClient(&gormx.ClientConfig{
		Driver:          c.GetData().GetGorm().GetDriver(),
		DataSourceName:  c.GetData().GetGorm().GetDataSourceName(),
		MaxIdleConn:     int(c.GetData().GetGorm().GetMaxIdleConn()),
		MaxOpenConn:     int(c.GetData().GetGorm().GetMaxOpenConn()),
		ConnMaxIdleTime: c.GetData().GetGorm().GetConnMaxLifeTime().AsDuration(),
		ConnMaxLifeTime: c.GetData().GetGorm().GetConnMaxLifeTime().AsDuration(),
		ShowLog:         c.GetData().GetGorm().GetShowLog(),
		Tracing:         c.GetData().GetGorm().GetTracing(),
	})
	if err != nil {
		panic(err)
	}
	return gorm
}

// NewRueidis 创建rueidis实例
func NewRueidis(c *conf.Bootstrap) rueidis.Client {
	rueidis, err := rueidiscache.NewRueidisClient(&rueidis.ClientOption{
		InitAddress: []string{c.GetData().GetRedis().GetAddr()},
		Password:    c.GetData().GetRedis().GetPassword(),
	})
	if err != nil {
		panic(err)
	}
	return rueidis
}

// NewDBCache 创建dbcache实例
func NewDBCache(c *conf.Bootstrap, rueidis rueidis.Client) dbcache.IDBCache {
	return rueidisdbcache.NewRueidisDBCache(rueidis, rueidisdbcache.WithName(c.GetName()))
}

// NewConfigRepo 创建config实例
func NewConfigRepo(_ *conf.Bootstrap, gorm *gorm.DB, dbCache dbcache.IDBCache) *config.Repo {
	encoding := encoding.NewSonic()
	return &config.Repo{
		DB:       gorm,
		Cache:    dbCache,
		Encoding: encoding,
	}
}

// NewAsynqClient 创建asynq客户端
func NewAsynqClient(c *conf.Bootstrap, logger log.Logger) mq.Client {
	redisClientOpt := asynq.RedisClientOpt{
		Addr:     c.Data.Redis.Addr,
		Password: c.Data.Redis.Password,
		DB:       int(c.Data.Redis.Db),
	}
	return mq.NewAsynqClient(logger, redisClientOpt)
}

func NewHTTPClient(_ *conf.Bootstrap) *httputil.Client {
	return httputil.NewClient()
}
