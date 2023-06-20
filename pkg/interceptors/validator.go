package interceptors

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	"google.golang.org/grpc"
)

func ValidatorServerInterceptor(opts ...validator.Option) grpc.UnaryServerInterceptor {
	return validator.UnaryServerInterceptor(opts...)
}

func ValidatorClientInterceptor(opts ...validator.Option) grpc.UnaryClientInterceptor {
	return validator.UnaryClientInterceptor(opts...)
}
