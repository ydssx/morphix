package mq

import (
	"encoding/json"

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
}

// NewNATSMessageQueueService 创建一个新的 NATSMessageQueueService 实例并连接到 NATS 服务器
func NewNATSMessageQueueService() (*NATSMessageQueueService, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}

	return &NATSMessageQueueService{
		nc: nc,
	}, nil
}

// PublishMessage 将消息发布到指定的主题
func (mq *NATSMessageQueueService) PublishMessage(topic string, message interface{}) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = mq.nc.Publish(topic, data)
	if err != nil {
		return err
	}

	return nil
}

// SubscribeToTopic 订阅指定主题并处理接收到的消息
func (mq *NATSMessageQueueService) SubscribeToTopic(topic string, handler func(message []byte)) error {
	_, err := mq.nc.Subscribe(topic, func(msg *nats.Msg) {
		handler(msg.Data)
	})
	if err != nil {
		return err
	}
	return nil
}

// Close 关闭与 NATS 服务器的连接
func (mq *NATSMessageQueueService) Close() {
	mq.nc.Close()
}
