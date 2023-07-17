package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

// MessageQueue 是一个消息队列服务的接口
type MessageQueue interface {
	// PublishMessage 发布消息到指定的主题
	PublishMessage(topic string, message interface{}) error

	// SubscribeToTopic 订阅指定主题并处理接收到的消息
	SubscribeToTopic(topic string, handler func(message []byte)) error

	// Close 关闭与消息队列服务的连接
	Close()
}

// NATSMessageQueueService 是基于 NATS 的消息队列服务实现
type NATSMessageQueueService struct {
	nc *nats.Conn
	ec *nats.EncodedConn
}

type Handler interface{}

// NewNATSMessageQueueService 创建一个新的 NATSMessageQueueService 实例并连接到 NATS 服务器
func NewNATSMessageQueueService() (*NATSMessageQueueService, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}
	return &NATSMessageQueueService{
		nc: nc,
		ec: ec,
	}, nil
}

// PublishMessage 将消息发布到指定的主题
func (mq *NATSMessageQueueService) PublishMessage(topic string, message interface{}) error {
	err := mq.ec.Publish(topic, message)
	return err
}

// SubscribeToTopic 订阅指定主题并处理接收到的消息
func (mq *NATSMessageQueueService) SubscribeToTopic(topic string, handler Handler) error {
	_, err := mq.ec.Subscribe(topic, handler)
	return err
}

func (mq *NATSMessageQueueService) QueueSubscribeToTopic(topic, queue string, handler Handler) error {
	_, err := mq.ec.QueueSubscribe(topic, queue, handler)
	return err
}

// Close 关闭与 NATS 服务器的连接
func (mq *NATSMessageQueueService) Close() {
	mq.nc.Close()
}

func main() {
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
