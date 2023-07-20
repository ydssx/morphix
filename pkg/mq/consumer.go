package mq

import (
	"context"
	"fmt"
	"log"

	cenats "github.com/cloudevents/sdk-go/protocol/nats/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type EventHandler func(ctx context.Context, event cloudevents.Event) error

func NewConsumer(subject string, handler EventHandler) {
	p, err := cenats.NewConsumerFromConn(natsServer, subject)
	if err != nil {
		log.Fatalf("failed to create nats protocol, %s", err.Error())
	}
	ctx := context.Background()
	defer p.Close(ctx)

	c, err := cloudevents.NewClient(p)
	if err != nil {
		log.Fatalf("failed to create client, %s", err.Error())
	}

	if err := c.StartReceiver(ctx, handler); err != nil {
		log.Printf("failed to start nats receiver, %s", err.Error())
	}
}

func receive(ctx context.Context, event cloudevents.Event) error {
	fmt.Printf("Got Event Context: %+v\n", event.Context)

	data := &Example{}
	if err := event.DataAs(data); err != nil {
		fmt.Printf("Got Data Error: %s\n", err.Error())
	}
	fmt.Printf("Got Data: %+v\n", data)

	fmt.Printf("----------------------------\n")
	return nil
}