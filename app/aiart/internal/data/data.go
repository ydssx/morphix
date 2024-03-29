package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	goredis "github.com/redis/go-redis/v9"
	"github.com/ydssx/morphix/app/aiart/internal/biz"
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/cache"
	"github.com/ydssx/morphix/pkg/client/mysql"
	"github.com/ydssx/morphix/pkg/client/redis"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewRedisCLient,
	NewRedisCache,
	NewMysqlDB,
	NewTransaction,
)

// Data .
type Data struct {
	rdb *goredis.Client
	db  *gorm.DB
}

type contextTxKey struct{}

// NewData .
func NewData(logger log.Logger, rdb *goredis.Client, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{rdb: rdb, db: db}, cleanup, nil
}

// InTx runs the given function in a database transaction.
// It starts a new transaction, calls the function with the transaction context,
// and commits/rollbacks based on the error returned by the function.
func (d *Data) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

// DB returns the database connection to use for the given context.
// If the context contains a transaction (tx), it returns the transaction.
// Otherwise, it returns the default database connection.
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
	redisConf := c.ServiceSet.Aiart.Data.Redis
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
	return mysql.NewDB(c.ServiceSet.Aiart.Data.Database.Source)
}

func NewRedisCache(client *goredis.Client) cache.Cache {
	return cache.NewRedisCache(client)
}
