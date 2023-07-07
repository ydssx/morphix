package common

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc/resolver/discovery"
	smsv1 "github.com/ydssx/morphix/api/sms/v1"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/pkg/interceptors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, rpcCliConf.Addr, opts...)
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
