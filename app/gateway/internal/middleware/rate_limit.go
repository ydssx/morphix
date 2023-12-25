package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/ydssx/morphix/pkg/limit"
)

// RateLimit is a middleware function that uses a RedisLimiter to limit request rates.
// It gets a redis client and returns a gin HandlerFunc.
// The returned handler function extracts a limit key from the request URL path and method.
// It checks if the request is allowed by the redis limiter using the limit key.
// If allowed, it calls ctx.Next() to process the next handler.
// If rate limited, it aborts the request with 429 Too Many Requests status and a JSON body.
func RateLimit(rdb *redis.Client) gin.HandlerFunc {
	limiter := limit.NewRedisLimiter(rdb)
	return func(ctx *gin.Context) {
		limitKey := ctx.Request.URL.Path + ":" + ctx.Request.Method
		if !limiter.Allow(limitKey, limit.WithRatePerSecond(10), limit.WithBurst(20)) {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, map[string]interface{}{"code": -1, "msg": "操作频繁,请稍后重试."})
			return
		}
		ctx.Next()
	}
}
