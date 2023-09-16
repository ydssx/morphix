package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	goredis "github.com/redis/go-redis/v9"
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/cache"
	"github.com/ydssx/morphix/pkg/mysql"
	"github.com/ydssx/morphix/pkg/redis"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepoCacheDecorator, NewUserRepo, NewRedisCLient, NewRedisCache, NewMysqlDB)

// Data .
type Data struct {
	rdb *goredis.Client
	db  *gorm.DB
}

// NewData .
func NewData(logger log.Logger, rdb *goredis.Client, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{rdb: rdb, db: db}, cleanup, nil
}

func NewRedisCLient(c *conf.Bootstrap) *goredis.Client {
	redisConf := c.ServiceSet.User.Data.Redis
	return redis.NewRedis(&goredis.Options{
		Addr:         redisConf.Addr,
		Password:     redisConf.Password,
		Username:     redisConf.Username,
		ReadTimeout:  redisConf.ReadTimeout.AsDuration(),
		WriteTimeout: redisConf.WriteTimeout.AsDuration(),
		DialTimeout:  redisConf.WriteTimeout.AsDuration(),
	})
}

func NewMysqlDB(c *conf.Bootstrap) *gorm.DB {
	return mysql.NewDB(c.ServiceSet.User.Data.Database.Source)
}

func NewRedisCache(client *goredis.Client) cache.Cache {
	return cache.NewRedisCache(client)
}
