package main

import (
	"context"
	"net/http"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc/resolver/discovery"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/app/gateway/conf"
	"github.com/ydssx/morphix/pkg/interceptors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type registerFn func(ctx context.Context, mux *gwruntime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

var handlers = make(map[string]registerFn)

func registerRpcServer(c conf.Config) {
	handlers[c.CommonConfig.UserRpcClient.Addr] = userv1.RegisterUserServiceHandlerFromEndpoint
}

func newGateway(ctx context.Context, opts []gwruntime.ServeMuxOption, r *etcd.Registry) (http.Handler, error) {

	mux := gwruntime.NewServeMux(opts...)

	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			interceptors.TraceClientInterceptor(),
			interceptors.MetricClientInterceptor()),
		grpc.WithResolvers(
			discovery.NewBuilder(
				r,
				discovery.WithInsecure(true),
				discovery.WithSubset(25),
				discovery.PrintDebugLog(true),
			)),
	}

	for endpoint, f := range handlers {
		if err := f(ctx, mux, endpoint, dialOpts); err != nil {
			return nil, err
		}
	}

	return mux, nil
}
