package interceptors

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/ratelimit"
	"google.golang.org/grpc"
)

func RateLimitServer(limiter ratelimit.Limiter) grpc.UnaryServerInterceptor {
	return ratelimit.UnaryServerInterceptor(limiter)
}
