package common

import (
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/interceptors"
	"github.com/ydssx/morphix/pkg/logger"
	"github.com/ydssx/morphix/pkg/middleware/kratos"
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

// NewHTTPServer 创建一个新的 HTTP 服务器。它接受一个 server 配置,并使用中间件、地址、超时等选项来初始化 HTTP 服务器。
// 返回初始化后的 HTTP 服务器实例。
func NewHTTPServer(server *conf.Server) *http.Server {
	opts := []http.ServerOption{
		http.Middleware(
			logging.Server(logger.DefaultLogger),
			validate.Validator(),
			kratos.AuthServer(),
			kratos.MetricServer(),
			recovery.Recovery(),
			ratelimit.Server(),
		),
	}
	if server.Http.Addr != "" {
		opts = append(opts, http.Address(server.Http.Addr))
	}
	if server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(server.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	return srv
}
