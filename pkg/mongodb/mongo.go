package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// InitMongoDB 初始化一个 MongoDB 连接。
// url 是 MongoDB 的连接字符串。
// 返回一个 MongoDB 客户端实例和一个清理函数用于断开连接。
// 如果连接失败会 panic。
func InitMongoDB(url string) (*mongo.Client, func()) {
	ctx := context.Background()
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(url), options.Client().SetMaxPoolSize(100), options.Client().SetMaxConnIdleTime(10*time.Second))
	if err != nil {
		panic("failed to connect to MongoDB: " + err.Error())
	}

	cleanup := func() { cli.Disconnect(ctx) }

	err = cli.Ping(ctx, readpref.Primary())
	if err != nil {
		cleanup()
		panic("failed to ping MongoDB: " + err.Error())
	}

	return cli, cleanup
}

type Mongo struct {
	db      *mongo.Database
	cli     *mongo.Client
	cleanup func()
}

func NewMongo(url string, dbName string) *Mongo {
	cli, cleanup := InitMongoDB(url)
	db := cli.Database(dbName)
	return &Mongo{db: db, cli: cli, cleanup: cleanup}
}

func (m *Mongo) Close() {
	m.cleanup()
}

func (m *Mongo) Database() *mongo.Database {
	return m.db
}

func (m *Mongo) Collection(collectionName string) *mongo.Collection {
	return m.db.Collection(collectionName)
}

func InitMongoCollection(db *mongo.Database, collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}

func InitMongoDatabase(cli *mongo.Client, dbName string) *mongo.Database {
	return cli.Database(dbName)
}
