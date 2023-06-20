package service

import (
	"context"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc/resolver/discovery"
	"github.com/google/wire"
	smsv1 "github.com/ydssx/morphix/api/sms/v1"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/pkg/interceptors"
	etcdclient "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService, NewSMSClient)

func NewSMSClient(c *common.Config) smsv1.SMSServiceClient {
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: c.Etcd.Endpoints,
	})
	if err != nil {
		log.Fatal(err)
	}
	r := etcd.New(client)

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
	conn, err := grpc.DialContext(context.Background(), c.SmsRpcClient.Addr, opts...)
	if err != nil {
		panic(err)
	}
	return smsv1.NewSMSServiceClient(conn)
}
