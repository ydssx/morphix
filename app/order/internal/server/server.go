package server

import (
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"github.com/ydssx/morphix/app/order/internal/listener"
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/errors"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer, NewServer)

func NewServer(httpServer *http.Server, grpcServer *grpc.Server, listenerServer *listener.ListenerServer, config *conf.Bootstrap) ([]transport.Server, error) {
	servers := []transport.Server{
		listenerServer,
	}
	if !config.ServiceSet.Order.Server.Http.Disabled {
		servers = append(servers, httpServer)
	}
	if !config.ServiceSet.Order.Server.Grpc.Disabled {
		servers = append(servers, grpcServer)
	}
	if len(servers) == 0 {
		return nil, errors.New("no server configured")
	}

	return servers, nil
}
