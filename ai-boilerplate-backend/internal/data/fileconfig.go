package data

import (
	"context"
	"errors"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
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
