package mq

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func TestNewConsumer(t *testing.T) {
	url := "http://localhost:4222"
	sub := "test-pubsub"
	close, _ := InitNats(url)
	defer close(context.Background())
	go NewConsumer(sub, receive)
	for i := 0; i < 10; i++ {
		payload := Example{Sequence: i, Message: "hello world"}
		Send(context.Background(), sub, payload, WithSource("user login"))
	}
	time.Sleep(time.Second * 10)
}

func TestMQ(t *testing.T) {
	type msg struct {
		Id int `json:"id,omitempty"`
	}
	conn, _ := NewNATSMessageQueueService()
	queue := "test-queue"
	topic := "test-topic"
	err := conn.QueueSubscribeToTopic(topic, queue, func(m *msg) {
		log.Printf("这是第一个订阅者,Id:%v", m.Id)
	})
	if err != nil {
		log.Print(err)
	}
	err = conn.QueueSubscribeToTopic(topic, queue, func(m *msg) {
		log.Printf("这是第二个订阅者,Id:%v", m.Id)
	})
	if err != nil {
		log.Print(err)
	}
	for i := 0; i < 10; i++ {
		err = conn.PublishMessage(topic, msg{Id: i})
		if err != nil {
			log.Print(err)
		}
	}
	time.Sleep(time.Second * 10)
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
