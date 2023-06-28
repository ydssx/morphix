package main

import (
	"context"
	"net/http"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc/resolver/discovery"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/pkg/interceptors"
	"github.com/ydssx/morphix/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type registerFn func(ctx context.Context, mux *gwruntime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

var handlers = make(map[string]registerFn)

func registerRpcHandler(c common.Config) {
	handlers[c.UserRpcClient.Addr] = userv1.RegisterUserServiceHandlerFromEndpoint
}

func newGateway(ctx context.Context, r *etcd.Registry, opts ...gwruntime.ServeMuxOption) (http.Handler, error) {
	mux := gwruntime.NewServeMux(opts...)

	builder := discovery.NewBuilder(r, discovery.WithSubset(25), discovery.PrintDebugLog(true))

	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			interceptors.TraceClientInterceptor(),
			interceptors.LoggingClientInterceptor(logger.DefaultLogger),
			interceptors.MetricClientInterceptor(),
		),
		grpc.WithResolvers(builder),
	}

	for endpoint, f := range handlers {
		if err := f(ctx, mux, endpoint, dialOpts); err != nil {
			return nil, err
		}
	}

	return mux, nil
}
