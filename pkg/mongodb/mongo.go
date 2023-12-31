package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitMongoDB(url string) (*mongo.Client, func()) {
	ctx := context.Background()
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(url), options.Client().SetMaxPoolSize(100), options.Client().SetMaxConnIdleTime(10*time.Second))
	if err != nil {
		panic(err)
	}

	cleanup := func() { cli.Disconnect(ctx) }

	err = cli.Ping(ctx, readpref.Primary())
	if err != nil {
		cleanup()
		panic(err)
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
