package redis

import (
	"context"

	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
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
