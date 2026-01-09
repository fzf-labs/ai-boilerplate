package server

import (
	"context"
	"fmt"

	"github.com/dromara/carbon/v2"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	conf "github.com/fzf-labs/kratos-contrib/api/conf/v1"
	"github.com/fzf-labs/kratos-contrib/pkg/mq"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
)

func NewMQServer(
	c *conf.Bootstrap,
	logger log.Logger,
) mq.Server {
	redisClientOpt := asynq.RedisClientOpt{
		Addr:     c.Data.Redis.Addr,
		Password: c.Data.Redis.Password,
		DB:       int(c.Data.Redis.Db),
	}
	srv := mq.NewAsynqServer(logger, redisClientOpt, mq.NwDefaultAsynqConfig(), mq.NewDefaultSchedulerOpts(logger))
	srv.ConsumerCronRegister(constant.MQTest, test, "@every 5s") // 每5秒执行一次
	return srv
}

func test(_ context.Context, _ []byte) error {
	fmt.Println(carbon.Now().ToDateTimeString())
	return nil
}
