package redis

import (
	"context"

	"github.com/bsm/redislock"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"github.com/ydssx/morphix/pkg/errors"
)

func NewRedis(opt *redis.Options) *redis.Client {
	cli := redis.NewClient(opt)
	_, err := cli.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return cli
}

func NewRedisLock(cli *redis.Client) *redislock.Client {
	return redislock.New(cli)
}

func NewRedisCluster(opt *redis.ClusterOptions) *redis.ClusterClient {
	cli := redis.NewClusterClient(opt)
	_, err := cli.Ping(context.Background()).Result()
	if err != nil {
		log.Error(err)
	}
	return cli
}

type RedisPubSub struct {
	cli  *redis.Client
	subs map[string]*redis.PubSub
}

func NewRedisPubSub(cli *redis.Client) *RedisPubSub {
	return &RedisPubSub{cli: cli, subs: make(map[string]*redis.PubSub)}
}

func (ps *RedisPubSub) PublishMessage(topic string, message interface{}) error {
	err := ps.cli.Publish(context.Background(), topic, message).Err()
	if err != nil {
		return errors.Wrap(err, "failed to publish message")
	}
	return nil
}

func (ps *RedisPubSub) SubscribeToTopic(topic string, handler func(message []byte)) {
	sub := ps.cli.Subscribe(context.Background(), topic)
	ps.subs[topic] = sub
	ch := sub.Channel()
	go func() {
		for msg := range ch {
			if msg != nil {
				handler([]byte(msg.Payload))
			}
		}
		log.Infof("stopping subscriber for topic [%s]", topic)
	}()
}

func (ps *RedisPubSub) Close() error {
	var errs []error
	for t, v := range ps.subs {
		err := v.Close()
		errs = append(errs, errors.Wrap(err, "failed to close subscription for topic "+t))
	}
	return errors.Join(errs...)
}
