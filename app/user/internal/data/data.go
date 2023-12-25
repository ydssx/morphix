package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	goredis "github.com/redis/go-redis/v9"
	"github.com/ydssx/morphix/app/user/internal/biz"
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/cache"
	"github.com/ydssx/morphix/pkg/mongodb"
	"github.com/ydssx/morphix/pkg/mysql"
	"github.com/ydssx/morphix/pkg/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewUserRepoCacheDecorator,
	NewUserRepo,
	NewRedisCLient,
	NewRedisCache,
	NewMysqlDB,
	NewTransaction,
	NewCollection,
)

// Data .
type Data struct {
	rdb        *goredis.Client
	db         *gorm.DB
	collection *mongo.Collection
}

type contextTxKey struct{}

// NewData .
func NewData(logger log.Logger, rdb *goredis.Client, db *gorm.DB, collection *mongo.Collection) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		err := collection.Database().Client().Disconnect(context.Background())
		if err!= nil {
			log.NewHelper(logger).Error("close mongodb client failed", err)
		}
	}
	return &Data{rdb: rdb, db: db, collection: collection}, cleanup, nil
}

// InTx 在一个数据库事务中执行函数 fn。
// 使用给定的上下文 ctx 创建一个事务,在事务中执行 fn,如果成功提交事务,否则回滚。
// fn 函数在事务上下文中执行,可以通过 ctx 访问事务数据库连接。
func (d *Data) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

// DB 返回当前上下文绑定的数据库连接。
// 如果上下文中存在事务(tx),则返回事务连接,否则返回默认数据库连接。
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

func NewCollection(c *conf.Bootstrap) *mongo.Collection {
	mongoConf := c.ServiceSet.User.Data.Mongo
	db := mongodb.NewMongo(mongoConf.Addr, mongoConf.Database)
	return db.Collection(mongoConf.Collection)
}
