package dapr

import (
	"context"

	"github.com/dapr/go-sdk/client"
)

var (
	cli        client.Client
	pubSubName = "pubsub"
)

func Init() func(context.Context) error {
	var err error
	cli, err = client.NewClient()
	if err != nil {
		panic(err)
	}

	return func(context.Context) error {
		cli.Close()
		return nil
	}
}

func PublishEvent(ctx context.Context, topic string, payload interface{}) error {
	return cli.PublishEvent(ctx, pubSubName, topic, payload)
}
