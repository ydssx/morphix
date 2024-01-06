package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/app/user/internal/service"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
)

// NewGRPCServer creates a new gRPC server for the user service.
// It registers the UserServiceServer, enables Dapr integration,
// and initializes the server based on the provided config.
// The returned server can then be started.
func NewGRPCServer(c *conf.Bootstrap, userSvc *service.UserService) *grpc.Server {
	srv := common.NewGRPCServer(c.ServiceSet.User.Server)

	userv1.RegisterUserServiceServer(srv, userSvc)

	return srv
}
