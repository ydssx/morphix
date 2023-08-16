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

func (c *CloudEvent) AddEventListener(ctx context.Context, subject string, handler EventHandler, opts ...cenats.ConsumerOption) (err error) {
	p, err := cenats.NewConsumerFromConn(c.natsConn, subject, opts...)
	if err != nil {
		logger.Fatalf(ctx, "failed to create nats protocol, %s", err.Error())
	}
	defer p.Close(ctx)

	ce, err := cloudevents.NewClient(p, client.WithObservabilityService(NewObserver()))
	if err != nil {
		logger.Fatalf(ctx, "failed to create client, %s", err.Error())
	}
	if err := ce.StartReceiver(ctx, handler); err != nil {
		logger.Fatalf(ctx, "failed to start nats receiver, %s", err.Error())
	}
	return
}

func (c *CloudEvent) AddEventListenerAsync(subject string, handler EventHandler, opts ...cenats.ConsumerOption) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	errChan := make(chan error, 1)

	go func() {
		errChan <- c.AddEventListener(context.Background(), subject, handler, opts...)
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
