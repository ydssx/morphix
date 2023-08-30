package interceptors

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
)

func RetryCLient() grpc.UnaryClientInterceptor {
	return retry.UnaryClientInterceptor(retry.WithMax(3))
}
