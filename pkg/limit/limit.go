package limit

import (
	"context"
	"math"
	"time"

	"github.com/go-redis/redis_rate/v10"
)

type LimiterType int

const (
	LimiterTypeConnection    LimiterType = iota // 连接数和ip地址限流
	LimiterTypeSlidingWindow                    // 滑动窗口限流
	LimiterTypeTokenBucket                      // 令牌桶限流
)

type Limiter interface {
	Allow() bool
}

var redisLimiter = redis_rate.NewLimiter(nil)

func Allow(ratePerSecond, burst int, key string) bool {
	burst = int(math.Max(float64(ratePerSecond), float64(burst)))

	r, err := redisLimiter.Allow(context.Background(), key, perSecond(ratePerSecond, burst))
	if err != nil {
		panic(err)
	}

	return r.Allowed > 0
}

func Reset(key string) {
	redisLimiter.Reset(context.Background(), key)
}

func perSecond(rate, burst int) redis_rate.Limit {
	return redis_rate.Limit{
		Rate:   rate,
		Period: time.Second,
		Burst:  burst,
	}
}
