package data

import (
	"context"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/security"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAiProviderPlatformRepo(
	logger log.Logger,
	data *Data,
	aiProviderPlatformRepo *ai_boilerplate_repo.AiProviderPlatformRepo,
) *AiProviderPlatformRepo {
	l := log.NewHelper(log.With(logger, "module", "data/aiProviderPlatform"))
	return &AiProviderPlatformRepo{
		log:                    l,
		data:                   data,
		AiProviderPlatformRepo: aiProviderPlatformRepo,
	}
}

type AiProviderPlatformRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.AiProviderPlatformRepo
}

func (r *AiProviderPlatformRepo) CreateOneCache(ctx context.Context, data *ai_boilerplate_model.AiProviderPlatform) error {
	if err := encryptAiProviderPlatformSecrets(data); err != nil {
		return err
	}
	return r.AiProviderPlatformRepo.CreateOneCache(ctx, data)
}

func (r *AiProviderPlatformRepo) UpdateOneCacheWithZero(ctx context.Context, newData *ai_boilerplate_model.AiProviderPlatform, oldData *ai_boilerplate_model.AiProviderPlatform) error {
	if err := prepareAiProviderPlatformSecretsForUpdate(newData, oldData); err != nil {
		return err
	}
	return r.AiProviderPlatformRepo.UpdateOneCacheWithZero(ctx, newData, oldData)
}

func encryptAiProviderPlatformSecrets(data *ai_boilerplate_model.AiProviderPlatform) error {
	if data == nil {
		return nil
	}
	apiKey, err := security.EncryptSecret(data.APIKey)
	if err != nil {
		return err
	}
	data.APIKey = apiKey
	return nil
}

func prepareAiProviderPlatformSecretsForUpdate(newData *ai_boilerplate_model.AiProviderPlatform, oldData *ai_boilerplate_model.AiProviderPlatform) error {
	if newData == nil {
		return nil
	}
	oldAPIKey := ""
	if oldData != nil {
		oldAPIKey = oldData.APIKey
	}
	apiKey, err := security.PrepareSecretForUpdate(newData.APIKey, oldAPIKey)
	if err != nil {
		return err
	}
	newData.APIKey = apiKey
	return nil
}
