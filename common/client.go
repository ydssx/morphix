package common

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc/resolver/discovery"
	smsv1 "github.com/ydssx/morphix/api/sms/v1"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/pkg/interceptors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptors.TraceClientInterceptor()),
		grpc.WithResolvers(
			discovery.NewBuilder(
				r,
				discovery.WithInsecure(true),
				discovery.WithSubset(25),
				discovery.PrintDebugLog(true),
			)),
	}
	conn, err := grpc.DialContext(context.Background(), rpcCliConf.Addr, opts...)
	if err != nil {
		panic(err)
	}
	return conn
}
