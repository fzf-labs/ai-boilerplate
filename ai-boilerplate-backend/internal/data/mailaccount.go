package data

import (
	"context"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/security"
	"github.com/go-kratos/kratos/v2/log"
)

func NewMailAccountRepo(
	logger log.Logger,
	data *Data,
	mailAccountRepo *ai_boilerplate_repo.MailAccountRepo,
) *MailAccountRepo {
	l := log.NewHelper(log.With(logger, "module", "data/mailAccount"))
	return &MailAccountRepo{
		log:             l,
		data:            data,
		MailAccountRepo: mailAccountRepo,
	}
}

type MailAccountRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.MailAccountRepo
}

func (r *MailAccountRepo) CreateOneCache(ctx context.Context, data *ai_boilerplate_model.MailAccount) error {
	if err := encryptMailAccountSecrets(data); err != nil {
		return err
	}
	return r.MailAccountRepo.CreateOneCache(ctx, data)
}

func (r *MailAccountRepo) UpdateOneCacheWithZero(ctx context.Context, newData *ai_boilerplate_model.MailAccount, oldData *ai_boilerplate_model.MailAccount) error {
	if err := prepareMailAccountSecretsForUpdate(newData, oldData); err != nil {
		return err
	}
	return r.MailAccountRepo.UpdateOneCacheWithZero(ctx, newData, oldData)
}

func encryptMailAccountSecrets(data *ai_boilerplate_model.MailAccount) error {
	if data == nil {
		return nil
	}
	password, err := security.EncryptSecret(data.Password)
	if err != nil {
		return err
	}
	data.Password = password
	return nil
}

func prepareMailAccountSecretsForUpdate(newData *ai_boilerplate_model.MailAccount, oldData *ai_boilerplate_model.MailAccount) error {
	if newData == nil {
		return nil
	}
	oldPassword := ""
	if oldData != nil {
		oldPassword = oldData.Password
	}
	password, err := security.PrepareSecretForUpdate(newData.Password, oldPassword)
	if err != nil {
		return err
	}
	newData.Password = password
	return nil
}
