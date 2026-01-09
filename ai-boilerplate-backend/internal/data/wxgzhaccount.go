package data

import (
	"context"
	"strings"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewWxGzhAccountRepo(
	logger log.Logger,
	data *Data,
	wxGzhAccountRepo *ai_boilerplate_repo.WxGzhAccountRepo,
) *WxGzhAccountRepo {
	l := log.NewHelper(log.With(logger, "module", "data/wxGzhAccount"))
	return &WxGzhAccountRepo{
		log:              l,
		data:             data,
		WxGzhAccountRepo: wxGzhAccountRepo,
	}
}

type WxGzhAccountRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.WxGzhAccountRepo
	defaultGzhAccount *officialAccount.OfficialAccount
}

// 创建公众号客户端
func (r *WxGzhAccountRepo) NewOfficialAccountClient(appID string, appSecret string, token string, aesKey string) (*officialAccount.OfficialAccount, error) {
	userConfig := &officialAccount.UserConfig{
		AppID:     appID,     // 公众号appid
		Secret:    appSecret, // 公众号app secret
		Token:     token,     // 公众号token
		AESKey:    aesKey,    // 公众号aesKey
		HttpDebug: false,
		Cache: kernel.NewRedisClient(&kernel.UniversalOptions{
			ClientName: r.data.cfg.Name,
			Addrs:      []string{r.data.cfg.Data.Redis.Addr},
			Password:   r.data.cfg.Data.Redis.Password,
			DB:         int(r.data.cfg.Data.Redis.Db),
		}),
		Log: officialAccount.Log{
			Stdout: true,
		},
	}
	// 不是线上环境，则开启调试
	if r.data.cfg.Env != "production" {
		userConfig.HttpDebug = true
	}
	m, err := officialAccount.NewOfficialAccount(userConfig)
	if err != nil {
		return nil, err
	}
	m.AccessToken.SetCacheKey(strings.Join([]string{"wx_access_token", appID}, ":"))
	return m, nil
}

// 获取默认公众号账号
func (r *WxGzhAccountRepo) GetDefaultGzhAccount() string {
	return r.data.cfg.GetBusiness()["wx"].GetFields()["defaultGzhAppId"].GetStringValue()
}

// 获取默认小程序账号
func (r *WxGzhAccountRepo) GetDefaultXcxAccount() string {
	return r.data.cfg.GetBusiness()["wx"].GetFields()["defaultXcxAppId"].GetStringValue()
}

// 创建默认公众号客户端
func (r *WxGzhAccountRepo) NewDefaultGzhAccountClient(ctx context.Context) (*officialAccount.OfficialAccount, error) {
	defaultGzhAccount, err := r.FindOneCacheByAppID(ctx, r.GetDefaultGzhAccount())
	if err != nil {
		return nil, err
	}
	officialAccount, err := r.NewOfficialAccountClient(defaultGzhAccount.AppID, defaultGzhAccount.AppSecret, defaultGzhAccount.Token, defaultGzhAccount.EncodingAesKey)
	if err != nil {
		return nil, err
	}
	return officialAccount, nil
}

// 获取默认公众号客户端
func (r *WxGzhAccountRepo) GetDefaultGzhAccountClient(ctx context.Context) (*officialAccount.OfficialAccount, error) {
	if r.defaultGzhAccount == nil {
		officialAccount, err := r.NewDefaultGzhAccountClient(ctx)
		if err != nil {
			return nil, err
		}
		r.defaultGzhAccount = officialAccount
	}
	return r.defaultGzhAccount, nil
}
