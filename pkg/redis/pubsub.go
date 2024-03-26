package redis

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"github.com/ydssx/morphix/pkg/errors"
)

type RedisPubSub struct {
	cli  *redis.Client
	subs map[string]*redis.PubSub
}

// NewRedisPubSub 创建RedisPubSub对象
func NewRedisPubSub(cli *redis.Client) *RedisPubSub {
	return &RedisPubSub{cli: cli, subs: make(map[string]*redis.PubSub)}
}

// PublishMessage publishes a message to the given topic.
// It returns an error if the publish failed.
func (ps *RedisPubSub) PublishMessage(topic string, message interface{}) error {
	err := ps.cli.Publish(context.Background(), topic, message).Err()
	if err != nil {
		return errors.Wrap(err, "发布消息失败")
	}
	return nil
}

// SubscribeToTopic subscribes to the given topic and calls the handler
// function whenever a new message is received on that topic.
func (ps *RedisPubSub) SubscribeToTopic(topic string, handler func(message []byte)) {
	sub := ps.cli.Subscribe(context.Background(), topic)
	ps.subs[topic] = sub

	ch := sub.Channel()
	go func() {
		for msg := range ch {
			if msg != nil {
				msg := msg
				go handler([]byte(msg.Payload))
			}
		}
		log.Infof("Stopped subscribing to messages on topic [%s]", topic)
	}()
}

// Close 关闭RedisPubSub对象
func (ps *RedisPubSub) Close() error {
	var errs []error
	for t, v := range ps.subs {
		err := v.Close()
		errs = append(errs, errors.Wrap(err, "关闭主题["+t+"]的订阅失败"))
	}
	return errors.Join(errs...)
}
