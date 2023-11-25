package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) Cache {
	return &RedisCache{client: client}
}

func (c *RedisCache) Get(key string, result interface{}) error {
	val, err := c.client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return fmt.Errorf("key %s not found", key)
	}
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(val), &result)
	return err
}

func (c *RedisCache) Set(key string, value interface{}, expire time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	randomExpire := expire + time.Duration(rand.Int63n(int64(time.Second)))

	err = c.client.Set(context.Background(), key, string(data), randomExpire).Err()
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
