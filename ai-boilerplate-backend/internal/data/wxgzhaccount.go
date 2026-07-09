package data

import (
	"context"
	"strings"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/security"
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

func (r *WxGzhAccountRepo) CreateOneCache(ctx context.Context, data *ai_boilerplate_model.WxGzhAccount) error {
	if err := encryptWxGzhAccountSecrets(data); err != nil {
		return err
	}
	return r.WxGzhAccountRepo.CreateOneCache(ctx, data)
}

func (r *WxGzhAccountRepo) UpdateOneCacheWithZero(ctx context.Context, newData *ai_boilerplate_model.WxGzhAccount, oldData *ai_boilerplate_model.WxGzhAccount) error {
	if err := prepareWxGzhAccountSecretsForUpdate(newData, oldData); err != nil {
		return err
	}
	return r.WxGzhAccountRepo.UpdateOneCacheWithZero(ctx, newData, oldData)
}

// 创建公众号客户端
func (r *WxGzhAccountRepo) NewOfficialAccountClient(appID string, appSecret string, token string, aesKey string) (*officialAccount.OfficialAccount, error) {
	appSecret, err := security.DecryptSecret(appSecret)
	if err != nil {
		return nil, err
	}
	token, err = security.DecryptSecret(token)
	if err != nil {
		return nil, err
	}
	aesKey, err = security.DecryptSecret(aesKey)
	if err != nil {
		return nil, err
	}
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

func encryptWxGzhAccountSecrets(data *ai_boilerplate_model.WxGzhAccount) error {
	if data == nil {
		return nil
	}
	appSecret, err := security.EncryptSecret(data.AppSecret)
	if err != nil {
		return err
	}
	token, err := security.EncryptSecret(data.Token)
	if err != nil {
		return err
	}
	encodingAesKey, err := security.EncryptSecret(data.EncodingAesKey)
	if err != nil {
		return err
	}
	data.AppSecret = appSecret
	data.Token = token
	data.EncodingAesKey = encodingAesKey
	return nil
}

func prepareWxGzhAccountSecretsForUpdate(newData *ai_boilerplate_model.WxGzhAccount, oldData *ai_boilerplate_model.WxGzhAccount) error {
	if newData == nil {
		return nil
	}
	oldAppSecret := ""
	oldToken := ""
	oldEncodingAesKey := ""
	if oldData != nil {
		oldAppSecret = oldData.AppSecret
		oldToken = oldData.Token
		oldEncodingAesKey = oldData.EncodingAesKey
	}
	appSecret, err := security.PrepareSecretForUpdate(newData.AppSecret, oldAppSecret)
	if err != nil {
		return err
	}
	token, err := security.PrepareSecretForUpdate(newData.Token, oldToken)
	if err != nil {
		return err
	}
	encodingAesKey, err := security.PrepareSecretForUpdate(newData.EncodingAesKey, oldEncodingAesKey)
	if err != nil {
		return err
	}
	newData.AppSecret = appSecret
	newData.Token = token
	newData.EncodingAesKey = encodingAesKey
	return nil
}
