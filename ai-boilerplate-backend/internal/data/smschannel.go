package data

import (
	"context"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/security"
	"github.com/go-kratos/kratos/v2/log"
)

func NewSmsChannelRepo(
	logger log.Logger,
	data *Data,
	smsChannelRepo *ai_boilerplate_repo.SmsChannelRepo,
) *SmsChannelRepo {
	l := log.NewHelper(log.With(logger, "module", "data/smsChannel"))
	return &SmsChannelRepo{
		log:            l,
		data:           data,
		SmsChannelRepo: smsChannelRepo,
	}
}

type SmsChannelRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.SmsChannelRepo
}

func (r *SmsChannelRepo) CreateOneCache(ctx context.Context, data *ai_boilerplate_model.SmsChannel) error {
	if err := encryptSmsChannelSecrets(data); err != nil {
		return err
	}
	return r.SmsChannelRepo.CreateOneCache(ctx, data)
}

func (r *SmsChannelRepo) UpdateOneCacheWithZero(ctx context.Context, newData *ai_boilerplate_model.SmsChannel, oldData *ai_boilerplate_model.SmsChannel) error {
	if err := prepareSmsChannelSecretsForUpdate(newData, oldData); err != nil {
		return err
	}
	return r.SmsChannelRepo.UpdateOneCacheWithZero(ctx, newData, oldData)
}

func (r *SmsChannelRepo) IDToName(ctx context.Context, ids []string) (map[string]string, error) {
	resp := make(map[string]string)
	list, err := r.FindMultiCacheByIDS(ctx, ids)
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		resp[v.ID] = v.Name
	}
	return resp, nil
}

func encryptSmsChannelSecrets(data *ai_boilerplate_model.SmsChannel) error {
	if data == nil {
		return nil
	}
	apiKey, err := security.EncryptSecret(data.APIKey)
	if err != nil {
		return err
	}
	apiSecret, err := security.EncryptSecret(data.APISecret)
	if err != nil {
		return err
	}
	data.APIKey = apiKey
	data.APISecret = apiSecret
	return nil
}

func prepareSmsChannelSecretsForUpdate(newData *ai_boilerplate_model.SmsChannel, oldData *ai_boilerplate_model.SmsChannel) error {
	if newData == nil {
		return nil
	}
	oldAPIKey := ""
	oldAPISecret := ""
	if oldData != nil {
		oldAPIKey = oldData.APIKey
		oldAPISecret = oldData.APISecret
	}
	apiKey, err := security.PrepareSecretForUpdate(newData.APIKey, oldAPIKey)
	if err != nil {
		return err
	}
	apiSecret, err := security.PrepareSecretForUpdate(newData.APISecret, oldAPISecret)
	if err != nil {
		return err
	}
	newData.APIKey = apiKey
	newData.APISecret = apiSecret
	return nil
}
