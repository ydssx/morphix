package mq

import (
	"context"
	"log"

	cenats "github.com/cloudevents/sdk-go/protocol/nats/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
)

func Send(ctx context.Context, subject string, event event.Event) error {
	p, err := cenats.NewSenderFromConn(natsServer, subject)
	if err != nil {
		log.Fatalf("Failed to create nats protocol, %s", err.Error())
	}
	defer p.Close(ctx)

	c, err := cloudevents.NewClient(p)
	if err != nil {
		log.Fatalf("Failed to create client, %s", err.Error())
	}

	if result := c.Send(context.Background(), event); cloudevents.IsUndelivered(result) {
		log.Printf("failed to send: %v", result)
		return result
	} else {
		log.Printf("sent: %s, accepted: %t", event.ID(), cloudevents.IsACK(result))
	}
	return nil
}
