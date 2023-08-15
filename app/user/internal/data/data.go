package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	goredis "github.com/redis/go-redis/v9"
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/cache"
	"github.com/ydssx/morphix/pkg/redis"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepoCacheDecorator, NewUserRepo, NewRedisCLient, cache.NewRedisCache)

// Data .
type Data struct {
	rdb *goredis.Client
}

// NewData .
func NewData(logger log.Logger, rdb *goredis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{rdb: rdb}, cleanup, nil
}

func NewRedisCLient(c *conf.Bootstrap) *goredis.Client {
	redisConf := c.User.Data.Redis
	return redis.NewRedis(&goredis.Options{
		Addr:         redisConf.Addr,
		Password:     redisConf.Password,
		Username:     redisConf.Username,
		ReadTimeout:  redisConf.ReadTimeout.AsDuration(),
		WriteTimeout: redisConf.WriteTimeout.AsDuration(),
		DialTimeout:  redisConf.WriteTimeout.AsDuration(),
	})
}
