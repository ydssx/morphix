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
	// redisConf, err := goredis.ParseURL(c.Sms.Data.Redis.Addr)
	// if err != nil {
	// 	log.Fatalf("redis address invalid: %v", err)
	// }

	return redis.NewRedis(&goredis.Options{
		Addr: c.Sms.Data.Redis.Addr,
		// Password:     redisConf.Password,
		// Username:     redisConf.Username,
		ReadTimeout:  time.Duration(c.Sms.Data.Redis.ReadTimeout),
		WriteTimeout: time.Duration(c.Sms.Data.Redis.WriteTimeout),
	})
}
