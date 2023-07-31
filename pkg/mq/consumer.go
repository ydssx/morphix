package mq

import (
	"context"
	"log"
	"time"

	cenats "github.com/cloudevents/sdk-go/protocol/nats/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/client"
)

type EventHandler func(ctx context.Context, event cloudevents.Event) error

func AddEventListener(ctx context.Context, subject string, handler EventHandler, opts ...cenats.ConsumerOption) (err error) {
	p, err := cenats.NewConsumerFromConn(natsServer, subject, opts...)
	if err != nil {
		log.Fatalf("failed to create nats protocol, %s", err.Error())
	}
	defer p.Close(ctx)

	c, err := cloudevents.NewClient(p, client.WithObservabilityService(NewObserver()))
	if err != nil {
		log.Fatalf("failed to create client, %s", err.Error())
	}
	if err := c.StartReceiver(ctx, handler); err != nil {
		log.Fatalf("failed to start nats receiver, %s", err.Error())
	}
	return
}

func AddEventListenerAsync(subject string, handler EventHandler, opts ...cenats.ConsumerOption) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	errChan := make(chan error, 1)

	go func() {
		errChan <- AddEventListener(context.Background(), subject, handler, opts...)
	}()

	select {
	case err = <-errChan:
		// 成功接收到 Goroutine 完成的错误值
	case <-ctx.Done():
		// err = ctx.Err()
	}
	close(errChan)
	
	return err
}
