package interceptors

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	"google.golang.org/grpc"
)

func ValidatorServer(opts ...validator.Option) grpc.UnaryServerInterceptor {
	return validator.UnaryServerInterceptor(opts...)
}

func ValidatorStreamServer(opts ...validator.Option) grpc.StreamServerInterceptor {
	return validator.StreamServerInterceptor(opts...)
}

func ValidatorClient(opts ...validator.Option) grpc.UnaryClientInterceptor {
	return validator.UnaryClientInterceptor(opts...)
}
