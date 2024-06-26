package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	goredis "github.com/redis/go-redis/v9"
	"github.com/ydssx/morphix/app/sms/internal/biz"
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/cache"
	"github.com/ydssx/morphix/pkg/client/mysql"
	"github.com/ydssx/morphix/pkg/client/redis"
	"github.com/ydssx/morphix/pkg/client/tencentcloud"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewRedisCLient,
	NewMysqlDB,
	// NewTransaction,
	tencentcloud.New,
	NewSmsRepo,
)

// Data .
type Data struct {
	db *gorm.DB
}

type contextTxKey struct{}

// NewData .
func NewData(logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

func (d *Data) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}

func NewTransaction(d *Data) biz.Transaction {
	return d
}

func NewRedisCLient(c *conf.Bootstrap) (*goredis.Client, error) {
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

func NewMysqlDB(c *conf.Bootstrap) (*gorm.DB, error) {
	return mysql.NewDB(c.ServiceSet.Sms.Data.Database.Source)
}

func NewRedisCache(client *goredis.Client) cache.Cache {
	return cache.NewRedisCache(client)
}
