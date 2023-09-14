package common

import (
	"sync"

	"github.com/hibiken/asynq"
	"github.com/ydssx/morphix/common/conf"
)

var rdbClientOpt asynq.RedisClientOpt
var once sync.Once

func InitRedisOpt(c *conf.Bootstrap) asynq.RedisClientOpt {
	redisConf := c.Redis
	once.Do(func() {
		rdbClientOpt = asynq.RedisClientOpt{
			Addr:     redisConf.Addr,
			Password: redisConf.Password,
			DB:       1,
		}
	})
	return rdbClientOpt
}

func NewAsynqClient(c *conf.Bootstrap) *asynq.Client {
	return asynq.NewClient(InitRedisOpt(c))
}
