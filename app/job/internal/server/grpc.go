package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jobv1 "github.com/ydssx/morphix/api/job/v1"
	"github.com/ydssx/morphix/app/job/internal/service"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
)

func NewGRPCServer(c *conf.Bootstrap, srv *service.JobService) *grpc.Server {

	s := common.NewGRPCServer(c.Job.Server)

	jobv1.RegisterJobServiceServer(s, srv)

	return s
}
