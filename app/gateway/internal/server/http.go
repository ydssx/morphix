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
	jobv1 "github.com/ydssx/morphix/api/job/v1"
	orderv1 "github.com/ydssx/morphix/api/order/v1"
	paymentv1 "github.com/ydssx/morphix/api/payment/v1"
	smsv1 "github.com/ydssx/morphix/api/sms/v1"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/app/gateway/internal/middleware"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/docs"
	"github.com/ydssx/morphix/pkg/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type registerFn func(ctx context.Context, mux *gwruntime.ServeMux, conn *grpc.ClientConn) (err error)

var handlers = make(map[*conf.ClientConf]registerFn)

// registerRPCHandler registers RPC handlers for all configured RPC clients.
// It maps each RPC client in the Bootstrap config to the appropriate
// register handler function for that API. This allows the gateway to set
// up RPC proxying for each configured downstream service.
func registerRPCHandler(c *conf.Bootstrap) {
	clientSet := c.ClientSet
	handlers[clientSet.UserRpcClient] = userv1.RegisterUserServiceHandler
	handlers[clientSet.SmsRpcClient] = smsv1.RegisterSMSServiceHandler
	handlers[clientSet.PaymentRpcClient] = paymentv1.RegisterPaymentServiceHandler
	handlers[clientSet.OrderRpcClient] = orderv1.RegisterOrderServiceHandler
	handlers[clientSet.JobRpcClient] = jobv1.RegisterJobServiceHandler
}

func NewHTTPServer(ctx context.Context, c *conf.Bootstrap) *khttp.Server {
	httpSrv := khttp.NewServer(
		khttp.Address(c.ServiceSet.Gateway.Server.Http.Addr),
	)

	openAPIhandler := openapiv2.NewHandler()
	httpSrv.HandlePrefix("/q/", openAPIhandler)

	ginHandler := newGinHandler(ctx, c)
	httpSrv.HandlePrefix("/", ginHandler)

	return httpSrv
}

func newGinHandler(ctx context.Context, conf *conf.Bootstrap) *gin.Engine {
	server := gin.New()
	server.ContextWithFallback = true
	redisClient := common.MustNewRedisClient(conf)
	server.Use(
		gin.Logger(),
		ginprom.PromMiddleware(nil),
		middleware.RateLimit(redisClient),
		gin.Recovery(),
	)
	server.GET("/metrics", gin.WrapH(promhttp.Handler()))
	server.GET("/healthz", healthHandler)
	server.GET("/docs", docsHandler)
	server.Any("/api/*any", gatewayHandler(ctx, conf))
	server.Use(middleware.Auth())
	server.GET("/auth", authHandler)
	return server
}

func healthHandler(c *gin.Context) {
	c.String(http.StatusOK, "%s", "ok")
}

func docsHandler(c *gin.Context) {
	c.Writer.Write(docs.ApiDocs)
}

func gatewayHandler(ctx context.Context, conf *conf.Bootstrap) gin.HandlerFunc {
	return gin.WrapH(newGateway(ctx, conf))
}

func authHandler(c *gin.Context) {
	auth := middleware.AuthFromGinContext(c)
	util.OKWithData(c, auth)
}

// newGateway returns a new instance of http.Handler that acts as a gateway for handling RPC requests.
//
// It takes a context.Context object and a *conf.Bootstrap object as parameters.
// The context.Context object represents the current execution context of the function.
// The *conf.Bootstrap object contains the configuration settings for the bootstrap process.
//
// The function registers the RPC handler using the provided bootstrap configuration.
// It creates a gwruntime.ServeMux with metadata and an Etcd registry.
// It then iterates through the handlers and creates a client connection for each handler.
// Finally, it returns the created gwruntime.ServeMux as an http.Handler.
func newGateway(ctx context.Context, c *conf.Bootstrap) http.Handler {
	registerRPCHandler(c)

	withMeta := gwruntime.WithMetadata(func(ctx context.Context, r *http.Request) metadata.MD {
		return metadata.New(map[string]string{"external-request": "true", "request-id": util.GetUUID()})
	})

	r := common.NewEtcdRegistry(c.Etcd)

	muxOpts := []gwruntime.ServeMuxOption{withMeta}
	mux := gwruntime.NewServeMux(muxOpts...)

	for clientConf, handlerFunc := range handlers {
		clientConn := common.CreateClientConn(ctx, clientConf, r)
		if err := handlerFunc(ctx, mux, clientConn); err != nil {
			panic(err)
		}
	}

	return mux
}
