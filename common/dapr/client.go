package dapr

import (
	"context"

	"github.com/dapr/go-sdk/client"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	cli        client.Client
	pubSubName = "pubsub"
)

func Init() func(ctx context.Context) error {
	var err error
	cli, err = client.NewClient()
	if err != nil {
		log.Error(err)
	}

	return func(context.Context) error {
		cli.Close()
		return nil
	}
}

func PublishEvent(ctx context.Context, topic string, payload interface{}) error {
	return cli.PublishEvent(ctx, pubSubName, topic, payload)
}
