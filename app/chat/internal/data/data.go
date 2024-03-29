package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	goredis "github.com/redis/go-redis/v9"
	"github.com/ydssx/morphix/app/chat/internal/biz"
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

// InTx 在一个数据库事务中执行给定的函数 fn。
// fn 接收一个上下文作为参数,该上下文中包含了一个事务数据库连接。
// 如果 fn 执行成功则提交事务,如果返回错误则回滚事务。
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

// IsInTx 判断当前上下文是否在事务中
func (d *Data) IsInTx(ctx context.Context) bool {
	_, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	return ok
}

func NewTransaction(d *Data) biz.Transaction {
	return d
}

func NewRedisCLient(c *conf.Bootstrap) (*goredis.Client, error) {
	redisConf := c.ServiceSet.Chat.Data.Redis
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
	return mysql.NewDB(c.ServiceSet.Chat.Data.Database.Source)
}

func NewRedisCache(client *goredis.Client) cache.Cache {
	return cache.NewRedisCache(client)
}
