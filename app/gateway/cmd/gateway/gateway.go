package main

import (
	"context"
	"net/http"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	orderv1 "github.com/ydssx/morphix/api/order/v1"
	paymentv1 "github.com/ydssx/morphix/api/payment/v1"
	smsv1 "github.com/ydssx/morphix/api/sms/v1"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/common"
	"google.golang.org/grpc"
)

type registerFn func(ctx context.Context, mux *gwruntime.ServeMux, conn *grpc.ClientConn) (err error)

var handlers = make(map[common.RpcClient]registerFn)

func registerRpcHandler(c common.Config) {
	handlers[c.UserRpcClient] = userv1.RegisterUserServiceHandler
	handlers[c.SmsRpcClient] = smsv1.RegisterSMSServiceHandler
	handlers[c.PaymentRpcClient] = paymentv1.RegisterPaymentServiceHandler
	handlers[c.OrderRpcClient] = orderv1.RegisterOrderServiceHandler
}

func newGateway(ctx context.Context, r *etcd.Registry, opts ...gwruntime.ServeMuxOption) (http.Handler, error) {
	mux := gwruntime.NewServeMux(opts...)

	for cliConf, f := range handlers {
		conn := common.CreateClientConn(cliConf, r)
		if err := f(ctx, mux, conn); err != nil {
			return nil, err
		}
	}

	return mux, nil
}
