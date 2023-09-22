package limit

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

type LimiterType int

const (
	LimiterTypeConnection    LimiterType = iota // 连接数和ip地址限流
	LimiterTypeSlidingWindow                    // 滑动窗口限流
	LimiterTypeTokenBucket                      // 令牌桶限流
)

type option struct {
	ratePerSecond int
	burst         int
}

func defaultConfig() *option {
	return &option{ratePerSecond: 10, burst: 20}
}

type Option func(*option)

func WithRatePerSecond(ratePerSecond int) Option {
	return func(o *option) { o.ratePerSecond = ratePerSecond }
}

func WithBurst(burst int) Option {
	return func(o *option) { o.burst = burst }
}

type Limiter interface {
	Allow() bool
}

type RedisLimiter struct {
	*redis_rate.Limiter
}

func NewRedisLimiter(rdb *redis.Client) *RedisLimiter {
	return &RedisLimiter{redis_rate.NewLimiter(rdb)}
}

func (l *RedisLimiter) Limit(ctx context.Context) error {
	key := LimitKeyFromCtx(ctx).(string)
	if l.Allow(key) {
		return nil
	}
	return errors.New("rate limited.")
}

func (l *RedisLimiter) Allow(key string, opts ...Option) bool {
	opt := defaultConfig()
	for _, v := range opts {
		v(opt)
	}

	r, err := l.Limiter.Allow(context.Background(), key, perSecond(opt.ratePerSecond, opt.burst))
	if err != nil {
		panic(err)
	}

	return r.Allowed > 0
}

func (l *RedisLimiter) Reset(key string) {
	_ = l.Limiter.Reset(context.Background(), key)
}

func perSecond(rate, burst int) redis_rate.Limit {
	burst = int(math.Max(float64(rate), float64(burst)))
	return redis_rate.Limit{
		Rate:   rate,
		Period: time.Second,
		Burst:  burst,
	}
}

type limitKey struct{}

func LimitKeyFromCtx(ctx context.Context) any {
	return ctx.Value(limitKey{})
}

func CtxWithLimitKey(ctx context.Context, value any) context.Context {
	return context.WithValue(ctx, limitKey{}, value)
}
