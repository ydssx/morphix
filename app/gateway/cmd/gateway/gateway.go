package main

import (
	"context"
	"net/http"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/pkg/interceptors"
	"github.com/ydssx/morphix/pkg/logger"
	"google.golang.org/grpc"
)

type registerFn func(ctx context.Context, mux *gwruntime.ServeMux, conn *grpc.ClientConn) (err error)

var handlers = make(map[string]registerFn)

func registerRpcHandler(c common.Config) {
	handlers[c.UserRpcClient.Addr] = userv1.RegisterUserServiceHandler
}

func newGateway(ctx context.Context, r *etcd.Registry, opts ...gwruntime.ServeMuxOption) (http.Handler, error) {
	mux := gwruntime.NewServeMux(opts...)

	for endpoint, f := range handlers {
		conn, err := kgrpc.DialInsecure(ctx,
			kgrpc.WithEndpoint(endpoint),
			kgrpc.WithDiscovery(r),
			kgrpc.WithUnaryInterceptor(
				interceptors.TraceClientInterceptor(),
				interceptors.LoggingClientInterceptor(logger.DefaultLogger),
				interceptors.MetricClientInterceptor(),
			),
		)
		if err != nil {
			return nil, err
		}

		if err := f(ctx, mux, conn); err != nil {
			return nil, err
		}
	}
	
	return mux, nil
}
