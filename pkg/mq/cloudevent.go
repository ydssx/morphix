package mq

import (
	"context"
	"fmt"
	"time"

	cenats "github.com/cloudevents/sdk-go/protocol/nats/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/client"
	"github.com/nats-io/nats.go"
	"github.com/ydssx/morphix/pkg/logger"
)

type EventHandler func(ctx context.Context, event cloudevents.Event) error

type CloudEvent struct {
	natsConn *nats.Conn
}

func NewCloudEvent(conn *nats.Conn) *CloudEvent {
	return &CloudEvent{natsConn: conn}
}

func (c *CloudEvent) PublishEvent(ctx context.Context, subject string, payload interface{}, opts ...Option) error {
	p, err := cenats.NewSenderFromConn(c.natsConn, subject)
	if err != nil {
		return fmt.Errorf("Failed to create nats protocol, %s", err.Error())
	}
	defer p.Close(ctx)

	ce, err := cloudevents.NewClient(p)
	if err != nil {
		return fmt.Errorf("Failed to create client, %s", err.Error())
	}

	event, err := NewEvent(ctx, payload, opts...)
	if err != nil {
		return err
	}

	if result := ce.Send(ctx, event); cloudevents.IsUndelivered(result) {
		logger.Errorf(ctx, "failed to send: %v", result)
		return result
	} else {
		logger.Infof(ctx, "subject: %s, sent: %s, accepted: %t", p.Subject, event.ID(), cloudevents.IsACK(result))
	}
	return nil
}

func (c *CloudEvent) AddEventListener(ctx context.Context, subject string, handler EventHandler, opts ...cenats.ConsumerOption) error {
	consumer, err := cenats.NewConsumerFromConn(c.natsConn, subject, opts...)
	if err != nil {
		return fmt.Errorf("failed to create nats protocol: %s", err)
	}
	defer consumer.Close(ctx)

	client, err := cloudevents.NewClient(consumer, client.WithObservabilityService(NewObserver()))
	if err != nil {
		return fmt.Errorf("failed to create client: %s", err)
	}

	if err := client.StartReceiver(ctx, handler); err != nil {
		return fmt.Errorf("failed to start nats receiver: %s", err)
	}

	return nil
}

func (c *CloudEvent) AddEventListenerAsync(ctx context.Context, subject string, handler EventHandler, opts ...cenats.ConsumerOption) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	errChan := make(chan error, 1)

	go func() {
		defer close(errChan)
		errChan <- c.AddEventListener(ctx, subject, handler, opts...)
	}()

	select {
	case err = <-errChan:
		// 成功接收到 Goroutine 完成的错误值
	case <-ctx.Done():
		// err = ctx.Err()
	}

	return err
}
