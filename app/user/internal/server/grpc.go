package server

import (
	"google.golang.org/grpc"

	user "github.com/ydssx/morphix/app/user/api"
	"github.com/ydssx/morphix/app/user/internal/service"
)

func NewGRPCServer() *grpc.Server {

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	s := grpc.NewServer()
	user.RegisterUserServiceServer(s, service.NewUserService())

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return s

}
