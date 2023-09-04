package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/ydssx/morphix/pkg/limit"
)

func RateLimit(rdb *redis.Client) gin.HandlerFunc {
	limiter := limit.NewRedisLimiter(rdb)
	return func(ctx *gin.Context) {
		limitKey := ctx.Request.URL.Path + ":" + ctx.Request.Method
		if !limiter.Allow(10, 20, limitKey) {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, map[string]interface{}{"code": -1, "msg": "操作频繁,请稍后重试."})
			return
		}
		ctx.Next()
	}
}
