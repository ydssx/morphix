package server

import (
	daprd "github.com/dapr/go-sdk/service/grpc"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/app/user/internal/service"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
)

func NewGRPCServer(c *conf.Bootstrap, userSvc *service.UserService) *grpc.Server {
	srv := common.NewGRPCServer(c.ServiceSet.User.Server)

	userv1.RegisterUserServiceServer(srv, userSvc)

	daprd.NewServiceWithGrpcServer(nil, srv.Server)

	return srv
}
