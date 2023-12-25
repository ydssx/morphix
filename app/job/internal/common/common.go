package common

import (
	"context"
	"sync"

	"github.com/hibiken/asynq"
	smsv1 "github.com/ydssx/morphix/api/sms/v1"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
)

var (
	rdbClientOpt asynq.RedisClientOpt
	once         sync.Once
)

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

func NewAsynqInspector(c *conf.Bootstrap) *asynq.Inspector {
	return asynq.NewInspector(InitRedisOpt(c))
}

type ServiceClientSet struct {
	smsv1.SMSServiceClient
	userv1.UserServiceClient
}

// NewServiceClientSet creates a new ServiceClientSet instance.
// It takes a Bootstrap config and uses it to initialize the SMS and User service clients.
// Returns a pointer to the initialized ServiceClientSet.
func NewServiceClientSet(c *conf.Bootstrap) *ServiceClientSet {
	return &ServiceClientSet{
		common.NewSMSClient(c),
		common.NewUserClient(c),
	}
}

type clientSetKey struct{}

func NewContextWithServiceClientSet(ctx context.Context, clientSet *ServiceClientSet) context.Context {
	return context.WithValue(ctx, clientSetKey{}, clientSet)
}

func ClientSetFromContext(ctx context.Context) *ServiceClientSet {
	return ctx.Value(clientSetKey{}).(*ServiceClientSet)
}
