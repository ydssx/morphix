package mq

import (
	"context"
	"log"

	cenats "github.com/cloudevents/sdk-go/protocol/nats/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func Send(ctx context.Context, subject string, payload interface{}, opts ...Option) error {
	p, err := cenats.NewSenderFromConn(natsServer, subject)
	if err != nil {
		log.Fatalf("Failed to create nats protocol, %s", err.Error())
	}
	defer p.Close(ctx)

	c, err := cloudevents.NewClient(p)
	if err != nil {
		log.Fatalf("Failed to create client, %s", err.Error())
	}

	event, err := NewEvent(ctx, payload, opts...)
	if err != nil {
		return err
	}

	if result := c.Send(ctx, event); cloudevents.IsUndelivered(result) {
		log.Printf("failed to send: %v", result)
		return result
	} else {
		log.Printf("sent: %s, accepted: %t", event.ID(), cloudevents.IsACK(result))
	}
	return nil
}
