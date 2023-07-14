package common

import (
	"context"

	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	smsv1 "github.com/ydssx/morphix/api/sms/v1"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/pkg/interceptors"
	"github.com/ydssx/morphix/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func NewSMSClient(c *Config) smsv1.SMSServiceClient {
	conn := createConn(c.Etcd, c.SmsRpcClient)

	return smsv1.NewSMSServiceClient(conn)
}

func NewUserClient(c *Config) userv1.UserServiceClient {
	conn := createConn(c.Etcd, c.UserRpcClient)

	return userv1.NewUserServiceClient(conn)
}

func createConn(etcdConf Etcd, rpcCliConf RpcClient) *grpc.ClientConn {
	r := NewEtcdRegistry(etcdConf)

	ctx := context.Background()
	conn, err := kgrpc.DialInsecure(ctx,
		kgrpc.WithEndpoint(rpcCliConf.Addr),
		kgrpc.WithDiscovery(r),
		kgrpc.WithUnaryInterceptor(
			interceptors.TraceClientInterceptor(),
			interceptors.LoggingClientInterceptor(logger.DefaultLogger),
			interceptors.MetricClientInterceptor(),
		),
	)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", rpcCliConf.Addr, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", rpcCliConf.Addr, cerr)
			}
		}()
	}()

	return conn
}
