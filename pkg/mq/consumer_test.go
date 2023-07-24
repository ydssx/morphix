package mq

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/ydssx/morphix/common/event"
	"github.com/ydssx/morphix/pkg/util"
)

func TestNewConsumer(t *testing.T) {
	url := "http://localhost:4222"
	sub := "test-pubsub"
	close, _ := InitNats(url)
	defer close(context.Background())
	go AddEventHandler(context.Background(), sub, receive)
	for i := 0; i < 10; i++ {
		payload := Example{Sequence: i, Message: "hello world"}
		Send(context.Background(), sub, &payload, WithSource("user login"))
	}
	time.Sleep(time.Second * 10)
}
func TestMq(t *testing.T) {
	url := "http://localhost:4222"
	close, _ := InitNats(url)
	defer close(context.Background())
	payload := event.PayloadPaymentCompleted{
		UserId:  1,
		Amount:  100,
		OrderId: util.GenerateCode(6),
	}
	x := event.Subject_name[int32(event.Subject_PaymentCompleted)]
	// go AddEventHandler(x, receive)
	Send(context.Background(), x, &payload)
	time.Sleep(time.Second * 5)

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

func receive(ctx context.Context, e cloudevents.Event) error {
	fmt.Printf("Got Event Context: %+v\n", e.Context)
	data := &event.PayloadPaymentCompleted{}
	if err := e.DataAs(data); err != nil {
		fmt.Printf("Got Data Error: %s\n", err.Error())
	}
	fmt.Printf("Got Data: %+v\n", data)

	fmt.Printf("----------------------------\n")
	return nil
}
