package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	aiartv1 "github.com/ydssx/morphix/api/aiart/v1"
	"github.com/ydssx/morphix/app/aiart/internal/service"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
)

func NewGRPCServer(c *conf.Bootstrap, svc *service.ArtService) *grpc.Server {
	srv := common.NewGRPCServer(c.ServiceSet.Aiart.Server)

	aiartv1.RegisterArtServiceServer(srv, svc)

	return srv
}
