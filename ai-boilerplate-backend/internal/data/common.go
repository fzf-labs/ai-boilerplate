package data

import (
	"context"
	"time"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_dao"
	"github.com/fzf-labs/godb/cache/rueidiscache"
	conf "github.com/fzf-labs/kratos-contrib/api/conf/v1"
	"github.com/go-kratos/kratos/v2/log"
)

func NewCommonRepo(
	logger log.Logger,
	cfg *conf.Bootstrap,
	data *Data,
) *CommonRepo {
	l := log.NewHelper(log.With(logger, "module", "data/common"), log.WithMessageKey("message"))
	return &CommonRepo{
		log:  l,
		cfg:  cfg,
		data: data,
	}
}

type CommonRepo struct {
	log  *log.Helper
	cfg  *conf.Bootstrap
	data *Data
}

// LockOnce 锁一次
func (a *CommonRepo) LockOnce(ctx context.Context, key string, ttl time.Duration, fn func() error) error {
	return rueidiscache.NewLocker(rueidiscache.NewDefaultLockerOption(a.data.rueidis)).LockOnce(ctx, key, ttl, fn)
}

// LockRetry 锁重试
func (a *CommonRepo) LockRetry(ctx context.Context, key string, ttl time.Duration, fn func() error) error {
	return rueidiscache.NewLocker(rueidiscache.NewDefaultLockerOption(a.data.rueidis)).LockRetry(ctx, key, ttl, fn)
}

// LockOnceNotRelease 锁一次不释放
func (a *CommonRepo) LockOnceNotRelease(ctx context.Context, key string, ttl time.Duration, fn func() error) error {
	return rueidiscache.NewLocker(rueidiscache.NewDefaultLockerOption(a.data.rueidis)).LockOnceNotRelease(ctx, key, ttl, fn)
}

// ClearRedisCache 清除缓存
func (a *CommonRepo) ClearRedisCache(ctx context.Context) error {
	return a.data.rueidis.Do(ctx, a.data.rueidis.B().Flushdb().Async().Build()).Error()
}

// Transaction 事务处理
func (a *CommonRepo) Transaction(_ context.Context, fn func(tx *ai_boilerplate_dao.Query) error) error {
	err := ai_boilerplate_dao.Use(a.data.gorm).Transaction(fn)
	if err != nil {
		return err
	}
	return nil
}
