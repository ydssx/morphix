package biz

import (
	"github.com/google/wire"

	goredis "github.com/redis/go-redis/v9"
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/client/redis"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewSmsUseCase, NewSmsRedisClient)

func NewSmsRedisClient(c *conf.Bootstrap) (*goredis.Client, error) {
	redisConf := c.ServiceSet.Sms.Data.Redis
	return redis.NewRedis(&goredis.Options{
		Addr:         redisConf.Addr,
		Password:     redisConf.Password,
		Username:     redisConf.Username,
		ReadTimeout:  redisConf.ReadTimeout.AsDuration(),
		WriteTimeout: redisConf.WriteTimeout.AsDuration(),
		DialTimeout:  redisConf.WriteTimeout.AsDuration(),
	})
}
