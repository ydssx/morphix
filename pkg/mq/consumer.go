package mq

import (
	"context"
	"log"
	"time"

	cenats "github.com/cloudevents/sdk-go/protocol/nats/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type EventHandler func(ctx context.Context, event cloudevents.Event) error

func AddEventHandler(ctx context.Context, subject string, handler EventHandler) (err error) {
	p, err := cenats.NewConsumerFromConn(natsServer, subject)
	if err != nil {
		log.Fatalf("failed to create nats protocol, %s", err.Error())
	}
	defer p.Close(ctx)

	c, err := cloudevents.NewClient(p)
	if err != nil {
		log.Fatalf("failed to create client, %s", err.Error())
	}
	if err := c.StartReceiver(ctx, handler); err != nil {
		log.Fatalf("failed to start nats receiver, %s", err.Error())
	}
	return
}

func AddEventHandlerAsync(subject string, handler EventHandler) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	errChan := make(chan error, 1)

	go func() {
		errChan <- AddEventHandler(context.Background(), subject, handler)
	}()

	select {
	case err = <-errChan:
		// 成功接收到 Goroutine 完成的错误值
	case <-ctx.Done():
		// err = ctx.Err()
	}

	return err
}
