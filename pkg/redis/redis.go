package redis

import (
	"context"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redsync/redsync/v4"
	syncredis "github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	"github.com/ydssx/morphix/pkg/errors"
)

// NewRedis 连接Redis并返回Client对象
func NewRedis(opt *redis.Options) (*redis.Client, error) {
	cli := redis.NewClient(opt)
	_, err := cli.Ping(context.Background()).Result()
	if err != nil {
		return nil, errors.Wrap(err, "redis connect failed")
	}
	log.Info("init redis success")
	return cli, nil
}

// NewRedisCluster 连接Redis集群并返回ClusterClient对象
func NewRedisCluster(opt *redis.ClusterOptions) *redis.ClusterClient {
	cli := redis.NewClusterClient(opt)
	_, err := cli.Ping(context.Background()).Result()
	if err != nil {
		log.Error(err)
	}
	return cli
}

type RedisSync struct {
	Redsync  *redsync.Redsync
	mutexMap map[string]*redsync.Mutex
	mutex    sync.Mutex
}

func NewRedisSync(cli *redis.Client) *RedisSync {
	pool := syncredis.NewPool(cli)
	return &RedisSync{Redsync: redsync.New(pool), mutexMap: make(map[string]*redsync.Mutex)}
}

func (r *RedisSync) Lock(ctx context.Context, key string, opt ...LockerOption) error {
	m := r.newMutex(opt, key)
	err := m.LockContext(ctx)
	if err != nil {
		return errors.Wrap(err, "redisSync: lock failed")
	}
	r.mutexMap[key] = m
	return nil
}

func (r *RedisSync) Unlock(ctx context.Context, key string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.mutexMap[key] == nil {
		return errors.New("redisSync: not obtain lock")
	}
	ok, err := r.mutexMap[key].UnlockContext(ctx)
	if err != nil {
		if _, is := err.(*redsync.ErrTaken); !is {
			return errors.Errorf("redisSync: unlock %s failed: %v", r.mutexMap[key].Name(), err)
		}
	}
	if !ok {
		return errors.New("redisSync: unlock failed")
	}
	delete(r.mutexMap, key)
	return nil
}

// TryLock 尝试获取锁
func (r *RedisSync) TryLock(ctx context.Context, key string, opt ...LockerOption) error {
	m := r.newMutex(opt, key)
	err := m.TryLockContext(ctx)
	if err != nil {
		return errors.Wrap(err, "redisSync: try lock failed")
	}
	r.mutexMap[key] = m
	return nil
}

func (r *RedisSync) newMutex(opt []LockerOption, key string) *redsync.Mutex {
	var o lockOption
	for _, f := range opt {
		f(&o)
	}
	var opts []redsync.Option
	if o.ttl != 0 {
		opts = append(opts, redsync.WithExpiry(o.ttl))
	}
	if o.tries != 0 {
		opts = append(opts, redsync.WithTries(o.tries))
	}
	if o.delay != 0 {
		opts = append(opts, redsync.WithRetryDelay(o.delay))
	}
	m := r.Redsync.NewMutex(key, opts...)
	return m
}
