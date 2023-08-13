package dapr

import (
	"context"
	"errors"

	"github.com/dapr/go-sdk/client"
)

var (
	cli        client.Client
	pubSubName = "pubsub"
	storeName  = "statestore"
)

type DaprClient struct {
	client.Client
}

func NewDaprClient() (*DaprClient, func(), error) {
	cli, err := client.NewClient()
	if err != nil {
		return nil, nil, err
	}
	return &DaprClient{cli}, cli.Close, nil
}

func (d *DaprClient) PublishEvent(ctx context.Context, topic string, payload interface{}) error {
	return d.Client.PublishEvent(ctx, pubSubName, topic, payload)
}

func (d *DaprClient) TryLock(ctx context.Context, key string) error {
	resp, err := d.TryLockAlpha1(ctx, storeName, &client.LockRequest{})
	if err != nil {
		return err
	}
	if !resp.Success {
		return errors.New("fail to obtain lock.")
	}
	return nil
}
