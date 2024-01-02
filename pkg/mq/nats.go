package mq

import (
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
	js nats.JetStreamContext
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
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	return &NATSMessageQueueService{
		nc: nc,
		ec: ec,
		js: js,
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

// QueueSubscribeToTopic 订阅指定主题的指定队列,并使用提供的处理程序处理接收到的消息。
// topic 是要订阅的主题,queue 是订阅的队列名称,handler 是接收到消息时调用的处理程序。
// 它会返回订阅信息和一个错误(如果有)。
func (mq *NATSMessageQueueService) QueueSubscribeToTopic(topic, queue string, handler Handler) error {
	_, err := mq.ec.QueueSubscribe(topic, queue, handler)
	return err
}

func (mq *NATSMessageQueueService) AddStream(cfg *nats.StreamConfig) error {
	_, err := mq.js.AddStream(cfg)
	return err
}

func (mq *NATSMessageQueueService) StreamSubscribe(subj string, cb nats.MsgHandler) {
	mq.js.Subscribe(subj, cb)
}

// Close 关闭与 NATS 服务器的连接
func (mq *NATSMessageQueueService) Close() {
	mq.nc.Drain()
}
