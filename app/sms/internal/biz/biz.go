package biz

import (
	"time"

	"github.com/google/wire"

	goredis "github.com/redis/go-redis/v9"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/pkg/redis"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewSmsUseCase, NewSmsRedisClient)

func NewSmsRedisClient(c *common.Config) *goredis.Client {
	redisConf := c.Sms.Data.Redis
	return redis.NewRedis(&goredis.Options{
		Addr: redisConf.Addr,
		Password:     redisConf.Password,
		Username:     redisConf.Username,
		ReadTimeout:  time.Duration(redisConf.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(redisConf.WriteTimeout) * time.Second,
		DialTimeout:  time.Duration(redisConf.WriteTimeout) * time.Second,
	})
}
