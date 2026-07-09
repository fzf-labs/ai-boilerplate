package data

import (
	"context"
	"errors"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/security"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/datatypes"
)

func NewFileConfigRepo(
	logger log.Logger,
	data *Data,
	fileConfigRepo *ai_boilerplate_repo.FileConfigRepo,
) *FileConfigRepo {
	l := log.NewHelper(log.With(logger, "module", "data/fileConfig"))
	return &FileConfigRepo{
		log:            l,
		data:           data,
		FileConfigRepo: fileConfigRepo,
	}
}

type FileConfigRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.FileConfigRepo
}

func (f *FileConfigRepo) CreateOneCache(ctx context.Context, data *ai_boilerplate_model.FileConfig) error {
	if err := encryptFileConfigSecrets(data); err != nil {
		return err
	}
	return f.FileConfigRepo.CreateOneCache(ctx, data)
}

func (f *FileConfigRepo) UpdateOneCacheWithZero(ctx context.Context, newData *ai_boilerplate_model.FileConfig, oldData *ai_boilerplate_model.FileConfig) error {
	if err := prepareFileConfigSecretsForUpdate(newData, oldData); err != nil {
		return err
	}
	return f.FileConfigRepo.UpdateOneCacheWithZero(ctx, newData, oldData)
}

// FindMasterConfig 查询主配置
func (f *FileConfigRepo) FindMasterConfig(ctx context.Context) (*ai_boilerplate_model.FileConfig, error) {
	result, err := f.FindMultiCacheByMaster(ctx, true)
	if err != nil {
		return nil, err
	}
	if len(result) != 1 {
		return nil, errors.New("主配置不存在或不唯一")
	}
	return result[0], nil
}

func encryptFileConfigSecrets(data *ai_boilerplate_model.FileConfig) error {
	if data == nil {
		return nil
	}
	config, err := security.EncryptJSONSecrets(data.Config)
	if err != nil {
		return err
	}
	data.Config = datatypes.JSON(config)
	return nil
}

func prepareFileConfigSecretsForUpdate(newData *ai_boilerplate_model.FileConfig, oldData *ai_boilerplate_model.FileConfig) error {
	if newData == nil {
		return nil
	}
	var oldConfig []byte
	if oldData != nil {
		oldConfig = oldData.Config
	}
	config, err := security.PrepareJSONSecretsForUpdate(newData.Config, oldConfig)
	if err != nil {
		return err
	}
	newData.Config = datatypes.JSON(config)
	return nil
}
