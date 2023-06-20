package service

import (
	"context"

	"github.com/google/wire"
	smsv1 "github.com/ydssx/morphix/api/sms/v1"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/pkg/interceptors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService, NewSMSClient)

func NewSMSClient(c *common.Config) smsv1.SMSServiceClient {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptors.TraceClientInterceptor()),
	}
	conn, err := grpc.DialContext(context.Background(), c.SmsRpcClient.Addr, opts...)
	if err != nil {
		panic(err)
	}
	return smsv1.NewSMSServiceClient(conn)
}
