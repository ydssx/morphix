package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ydssx/morphix/app/gateway/conf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

// Endpoint describes a gRPC endpoint
type Endpoint struct {
	Name, Network, Addr string
}

// Options is a set of options to be passed to Run
type Options struct {
	// Addr is the address to listen
	Addr string

	// GRPCServer defines an endpoint of a gRPC service
	GRPCServer []Endpoint

	// OpenAPIDir is a path to a directory from which the server
	// serves OpenAPI specs.
	OpenAPIDir string

	// Mux is a list of options to be passed to the gRPC-Gateway multiplexer
	Mux []gwruntime.ServeMuxOption
}

// Run starts a HTTP server and blocks while running if successful.
// The server will be shutdown when "ctx" is canceled.
func Run(ctx context.Context, c conf.Config) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	server := gin.New()
	server.Use(gin.Logger())
	// mux.HandleFunc("/openapiv2/", openAPIServer(opts.OpenAPIDir))
	// mux.HandleFunc("/healthz", healthzServer(conn))
	opts := []gwruntime.ServeMuxOption{}

	registerRpcServer(c)

	gw, err := newGateway(ctx, opts)
	if err != nil {
		return err
	}

	// server.GET("/test", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Ok")
	// })

	server.Any("*any", gin.WrapH(gw))

	err = server.Run(c.Addr)
	if err != nil {
		panic(err)
	}
	return nil
}

// openAPIServer returns OpenAPI specification files located under "/openapiv2/"
func openAPIServer(dir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
			glog.Errorf("Not Found: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}

		glog.Infof("Serving %s", r.URL.Path)
		p := strings.TrimPrefix(r.URL.Path, "docs/")
		p = path.Join(dir, p)
		http.ServeFile(w, r, p)
	}
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
