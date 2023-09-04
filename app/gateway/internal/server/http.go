package server

import (
	"context"
	"net/http"

	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	orderv1 "github.com/ydssx/morphix/api/order/v1"
	paymentv1 "github.com/ydssx/morphix/api/payment/v1"
	smsv1 "github.com/ydssx/morphix/api/sms/v1"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/app/gateway/internal/middleware"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type registerFn func(ctx context.Context, mux *gwruntime.ServeMux, conn *grpc.ClientConn) (err error)

var handlers = make(map[*conf.ClientConf]registerFn)

func registerRpcHandler(c *conf.Bootstrap) {
	handlers[c.UserRpcClient] = userv1.RegisterUserServiceHandler
	handlers[c.SmsRpcClient] = smsv1.RegisterSMSServiceHandler
	handlers[c.PaymentRpcClient] = paymentv1.RegisterPaymentServiceHandler
	handlers[c.OrderRpcClient] = orderv1.RegisterOrderServiceHandler
}

func NewHTTPServer(c *conf.Bootstrap) *khttp.Server {
	httpSrv := khttp.NewServer(
		khttp.Address(c.Gateway.Server.Http.Addr),
	)

	openAPIhandler := openapiv2.NewHandler()
	httpSrv.HandlePrefix("/q/", openAPIhandler)

	ginHandler := newGinHandler(context.Background(), c)
	httpSrv.HandlePrefix("/", ginHandler)

	return httpSrv
}

func newGinHandler(ctx context.Context, c *conf.Bootstrap) *gin.Engine {
	server := gin.New()
	server.ContextWithFallback = true
	rdb := common.NewRedisClient(c)
	server.Use(gin.Logger(), ginprom.PromMiddleware(nil), middleware.RateLimit(rdb), gin.Recovery())
	
	server.GET("/metrics", gin.WrapH(promhttp.Handler()))
	server.GET("/healthz", func(c *gin.Context) { c.String(http.StatusOK, "%s", "ok") })
	server.Any("/api/*any", gin.WrapH(newGateway(ctx, c)))

	server.Use(middleware.Auth())
	server.GET("/auth", func(ctx *gin.Context) {
		auth := middleware.AuthFromGinContext(ctx)
		util.OKWithData(ctx, auth)
	})

	return server
}

func newGateway(ctx context.Context, c *conf.Bootstrap) http.Handler {
	registerRpcHandler(c)

	withMeta := gwruntime.WithMetadata(func(ctx context.Context, r *http.Request) metadata.MD {
		return metadata.New(map[string]string{"external-request": "true"})
	})
	// withTrace := gwruntime.WithOutgoingHeaderMatcher(func(s string) (string, bool) {
	// 	if s == "trace-id" {
	// 		return s, true
	// 	}
	// 	return s, true
	// })
	opts := []gwruntime.ServeMuxOption{withMeta}

	r := common.NewEtcdRegistry(c.Etcd)

	mux := gwruntime.NewServeMux(opts...)
	for cliConf, f := range handlers {
		conn := common.CreateClientConn(cliConf, r)
		if err := f(ctx, mux, conn); err != nil {
			panic(err)
		}
	}

	return mux
}
