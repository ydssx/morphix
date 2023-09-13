package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitMongoDB(url string) (*mongo.Client, func()) {
	ctx := context.Background()
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
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
