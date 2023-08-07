package redis

import (
	"context"

	"github.com/bsm/redislock"
	"github.com/go-kratos/kratos/v2/log"
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

func NewRedisCluster(opt *redis.ClusterOptions) *redis.ClusterClient {
	cli := redis.NewClusterClient(opt)
	_, err := cli.Ping(context.Background()).Result()
	if err != nil {
		log.Error(err)
	}
	return cli
}
