package redis

import (
	"context"
	"sync"
	"time"

	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
	"github.com/ydssx/morphix/pkg/errors"
)

type Locker interface {
	// Lock 获取锁
	Lock(ctx context.Context, key string, opt ...LockerOption) error
	// Unlock 释放锁
	Unlock(ctx context.Context, key string) error
	// TryLock 尝试获取锁
	TryLock(ctx context.Context, key string, opt ...LockerOption) error
}

type lockOption struct {
	// 锁的有效期
	ttl time.Duration
	// 重试次数
	tries int
	// 重试间隔
	delay time.Duration
	// 超时时间
	timeout time.Duration
}

type LockerOption func(*lockOption)

// WithTTL 设置锁的有效期
func WithTTL(ttl time.Duration) LockerOption {
	return func(opt *lockOption) {
		opt.ttl = ttl
	}
}

// WithTries 设置重试次数
func WithTries(tries int) LockerOption {
	return func(opt *lockOption) {
		opt.tries = tries
	}
}

// WithDelay 设置重试间隔
func WithDelay(delay time.Duration) LockerOption {
	return func(opt *lockOption) {
		opt.delay = delay
	}
}

func WithTimeout(timeout time.Duration) LockerOption {
	return func(opt *lockOption) {
		opt.timeout = timeout
	}
}

type RedisLocker struct {
	*redislock.Client
	lockMap map[string]*redislock.Lock
	mu      sync.Mutex
}

var _ Locker = (*RedisLocker)(nil)

func NewLocker(cli *redis.Client) *RedisLocker {
	return &RedisLocker{Client: redislock.New(cli), lockMap: make(map[string]*redislock.Lock)}
}

func (r *RedisLocker) Lock(ctx context.Context, key string, opt ...LockerOption) error {
	lo := &lockOption{
		ttl:   10 * time.Second,
		tries: 3,
		delay: 500 * time.Millisecond,
	}
	for _, o := range opt {
		o(lo)
	}
	lock, err := r.Client.Obtain(ctx, key, lo.ttl, &redislock.Options{
		RetryStrategy: redislock.LinearBackoff(lo.delay),
	})
	if err != nil {
		return errors.Errorf("redislock: lock %s failed: %v", key, err)
	}
	r.mu.Lock()
	r.lockMap[key] = lock
	r.mu.Unlock()
	return nil
}

func (r *RedisLocker) Unlock(ctx context.Context, key string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.lockMap[key] == nil {
		return errors.New("redislock: not obtain lock")
	}
	if err := r.lockMap[key].Release(ctx); err != nil && !errors.Is(err, redislock.ErrLockNotHeld) {
		return errors.Errorf("redislock: unlock %s failed: %v", r.lockMap[key].Key(), err)
	}

	delete(r.lockMap, key)
	return nil
}

func (r *RedisLocker) TryLock(ctx context.Context, key string, opt ...LockerOption) error {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel()
	return r.Lock(ctx, key, opt...)
}
