package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) (*RedisCache, error) {
	return &RedisCache{client: client}, nil
}

func (c *RedisCache) Get(key string) (interface{}, error) {
	val, err := c.client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("key %s not found", key)
	}
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (c *RedisCache) Set(key string, value interface{}, expire time.Duration) error {
	err := c.client.Set(context.Background(), key, value, expire).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *RedisCache) Delete(key string) error {
	err := c.client.Del(context.Background(), key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *RedisCache) Clear() error {
	err := c.client.FlushDB(context.Background()).Err()
	if err != nil {
		return err
	}
	return nil
}
