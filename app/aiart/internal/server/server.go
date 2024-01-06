package server

import (
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
)

// ProviderSet is server providers.
var (
	ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer, NewServer)
)

func NewServer(httpServer *http.Server, grpcServer *grpc.Server) []transport.Server {
	return []transport.Server{
		httpServer,
		// grpcServer,
	}
}
