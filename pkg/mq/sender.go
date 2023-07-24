package mq

import (
	"context"
	"fmt"

	cenats "github.com/cloudevents/sdk-go/protocol/nats/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/ydssx/morphix/pkg/logger"
)

func Send(ctx context.Context, subject string, payload interface{}, opts ...Option) error {
	p, err := cenats.NewSenderFromConn(natsServer, subject)
	if err != nil {
		return fmt.Errorf("Failed to create nats protocol, %s", err.Error())
	}
	defer p.Close(ctx)

	c, err := cloudevents.NewClient(p)
	if err != nil {
		return fmt.Errorf("Failed to create client, %s", err.Error())
	}

	event, err := NewEvent(ctx, payload, opts...)
	if err != nil {
		return err
	}

	if result := c.Send(ctx, event); cloudevents.IsUndelivered(result) {
		logger.Errorf(ctx, "failed to send: %v", result)
		return result
	} else {
		logger.Infof(ctx, "subject: %s, sent: %s, accepted: %t", p.Subject, event.ID(), cloudevents.IsACK(result))
	}
	return nil
}
