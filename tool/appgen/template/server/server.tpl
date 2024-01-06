package server

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer, NewServer)

func NewServer(httpServer *http.Server, grpcServer *grpc.Server) []transport.Server {
	return []transport.Server{
		httpServer,
		grpcServer,
	}
}
