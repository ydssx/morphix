package common

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/interceptors"
)

// NewGRPCServer 创建一个新的 gRPC 服务器。它接受一个 server 配置,并使用拦截器、地址、超时等选项来初始化 gRPC 服务器。
// 返回初始化后的 gRPC 服务器实例。
func NewGRPCServer(server *conf.Server) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			interceptors.TraceServer(),
			interceptors.LoggingServer(),
			interceptors.ValidatorServer(),
			interceptors.EventServer(),
			interceptors.AuthServer(),
			interceptors.MetricServer(),
			interceptors.RecoveryServer(),
		),
		grpc.StreamInterceptor(
			interceptors.TraceStreamServer(),
			interceptors.LoggingStreamServer(),
			interceptors.ValidatorStreamServer(),
			interceptors.AuthStreamServer(),
			interceptors.EventStreamServer(),
			interceptors.RecoveryStreamServer(),
		),
	}
	if server.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(server.Grpc.Addr))
	}
	if server.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(server.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)

	return srv
}
