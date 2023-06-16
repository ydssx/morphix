package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/golang/glog"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ydssx/morphix/app/gateway/conf"
	"github.com/ydssx/morphix/pkg/provider"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

// Run starts a HTTP server and blocks while running if successful.
// The server will be shutdown when "ctx" is canceled.
func Run(ctx context.Context, c conf.Config) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tp, err := provider.InitTraceProvider("http://localhost:14268/api/traces", "gateway")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	server := gin.New()

	server.Use(gin.Logger(), gin.Recovery())

	server.Any("/metrics", gin.WrapH(promhttp.Handler()))
	// mux.HandleFunc("/openapiv2/", openAPIServer(opts.OpenAPIDir))
	// mux.HandleFunc("/healthz", healthzServer(conn))
	opts := []gwruntime.ServeMuxOption{}

	registerRpcServer(c)

	gw, err := newGateway(ctx, opts)
	if err != nil {
		return err
	}
	server.Any("/api/*any", gin.WrapH(gw))

	httpSrv := khttp.NewServer(khttp.Address(c.Addr))

	openAPIhandler := openapiv2.NewHandler()
	httpSrv.HandlePrefix("/q/", openAPIhandler)
	
	httpSrv.HandlePrefix("/", server)
	app := kratos.New(
		kratos.Name("gateway"),
		kratos.Server(
			httpSrv,
		),
	)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
	return nil
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// preflightHandler adds the necessary headers in order to serve
// CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
// We insist, don't do this without consideration in production systems.
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	glog.Infof("preflight request for %s", r.URL.Path)
}

// healthzServer returns a simple health handler which returns ok.
func healthzServer(conn *grpc.ClientConn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		if s := conn.GetState(); s != connectivity.Ready {
			http.Error(w, fmt.Sprintf("grpc server is %s", s), http.StatusBadGateway)
			return
		}
		fmt.Fprintln(w, "ok")
	}
}

var configFile = flag.String("f", "../configs/config.yaml", "the config file")

func main() {
	var config conf.Config
	conf.MustLoad(*configFile, &config)

	if err := Run(context.Background(), config); err != nil {
		panic(err)
	}
}
