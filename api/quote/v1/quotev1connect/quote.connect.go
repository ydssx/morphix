// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/quote/v1/quote.proto

package quotev1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/ydssx/morphix/api/quote/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// QuoteServiceName is the fully-qualified name of the QuoteService service.
	QuoteServiceName = "quote.QuoteService"
	// CouponServiceName is the fully-qualified name of the CouponService service.
	CouponServiceName = "quote.CouponService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// QuoteServiceCreateQuoteProcedure is the fully-qualified name of the QuoteService's CreateQuote
	// RPC.
	QuoteServiceCreateQuoteProcedure = "/quote.QuoteService/CreateQuote"
	// QuoteServiceGetQuotesProcedure is the fully-qualified name of the QuoteService's GetQuotes RPC.
	QuoteServiceGetQuotesProcedure = "/quote.QuoteService/GetQuotes"
	// QuoteServiceGetQuoteProcedure is the fully-qualified name of the QuoteService's GetQuote RPC.
	QuoteServiceGetQuoteProcedure = "/quote.QuoteService/GetQuote"
	// CouponServiceGetUserCouponsProcedure is the fully-qualified name of the CouponService's
	// GetUserCoupons RPC.
	CouponServiceGetUserCouponsProcedure = "/quote.CouponService/GetUserCoupons"
	// CouponServiceUseCouponProcedure is the fully-qualified name of the CouponService's UseCoupon RPC.
	CouponServiceUseCouponProcedure = "/quote.CouponService/UseCoupon"
)

// QuoteServiceClient is a client for the quote.QuoteService service.
type QuoteServiceClient interface {
	// 创建报价
	CreateQuote(context.Context, *connect_go.Request[v1.CreateQuoteRequest]) (*connect_go.Response[v1.CreateQuoteResponse], error)
	// 获取报价列表
	GetQuotes(context.Context, *connect_go.Request[v1.GetQuotesRequest]) (*connect_go.Response[v1.GetQuotesResponse], error)
	// 获取单个报价
	GetQuote(context.Context, *connect_go.Request[v1.GetQuoteRequest]) (*connect_go.Response[v1.Quote], error)
}

// NewQuoteServiceClient constructs a client for the quote.QuoteService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewQuoteServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) QuoteServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &quoteServiceClient{
		createQuote: connect_go.NewClient[v1.CreateQuoteRequest, v1.CreateQuoteResponse](
			httpClient,
			baseURL+QuoteServiceCreateQuoteProcedure,
			opts...,
		),
		getQuotes: connect_go.NewClient[v1.GetQuotesRequest, v1.GetQuotesResponse](
			httpClient,
			baseURL+QuoteServiceGetQuotesProcedure,
			opts...,
		),
		getQuote: connect_go.NewClient[v1.GetQuoteRequest, v1.Quote](
			httpClient,
			baseURL+QuoteServiceGetQuoteProcedure,
			opts...,
		),
	}
}

// quoteServiceClient implements QuoteServiceClient.
type quoteServiceClient struct {
	createQuote *connect_go.Client[v1.CreateQuoteRequest, v1.CreateQuoteResponse]
	getQuotes   *connect_go.Client[v1.GetQuotesRequest, v1.GetQuotesResponse]
	getQuote    *connect_go.Client[v1.GetQuoteRequest, v1.Quote]
}

// CreateQuote calls quote.QuoteService.CreateQuote.
func (c *quoteServiceClient) CreateQuote(ctx context.Context, req *connect_go.Request[v1.CreateQuoteRequest]) (*connect_go.Response[v1.CreateQuoteResponse], error) {
	return c.createQuote.CallUnary(ctx, req)
}

// GetQuotes calls quote.QuoteService.GetQuotes.
func (c *quoteServiceClient) GetQuotes(ctx context.Context, req *connect_go.Request[v1.GetQuotesRequest]) (*connect_go.Response[v1.GetQuotesResponse], error) {
	return c.getQuotes.CallUnary(ctx, req)
}

// GetQuote calls quote.QuoteService.GetQuote.
func (c *quoteServiceClient) GetQuote(ctx context.Context, req *connect_go.Request[v1.GetQuoteRequest]) (*connect_go.Response[v1.Quote], error) {
	return c.getQuote.CallUnary(ctx, req)
}

// QuoteServiceHandler is an implementation of the quote.QuoteService service.
type QuoteServiceHandler interface {
	// 创建报价
	CreateQuote(context.Context, *connect_go.Request[v1.CreateQuoteRequest]) (*connect_go.Response[v1.CreateQuoteResponse], error)
	// 获取报价列表
	GetQuotes(context.Context, *connect_go.Request[v1.GetQuotesRequest]) (*connect_go.Response[v1.GetQuotesResponse], error)
	// 获取单个报价
	GetQuote(context.Context, *connect_go.Request[v1.GetQuoteRequest]) (*connect_go.Response[v1.Quote], error)
}

// NewQuoteServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewQuoteServiceHandler(svc QuoteServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	quoteServiceCreateQuoteHandler := connect_go.NewUnaryHandler(
		QuoteServiceCreateQuoteProcedure,
		svc.CreateQuote,
		opts...,
	)
	quoteServiceGetQuotesHandler := connect_go.NewUnaryHandler(
		QuoteServiceGetQuotesProcedure,
		svc.GetQuotes,
		opts...,
	)
	quoteServiceGetQuoteHandler := connect_go.NewUnaryHandler(
		QuoteServiceGetQuoteProcedure,
		svc.GetQuote,
		opts...,
	)
	return "/quote.QuoteService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case QuoteServiceCreateQuoteProcedure:
			quoteServiceCreateQuoteHandler.ServeHTTP(w, r)
		case QuoteServiceGetQuotesProcedure:
			quoteServiceGetQuotesHandler.ServeHTTP(w, r)
		case QuoteServiceGetQuoteProcedure:
			quoteServiceGetQuoteHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedQuoteServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedQuoteServiceHandler struct{}

func (UnimplementedQuoteServiceHandler) CreateQuote(context.Context, *connect_go.Request[v1.CreateQuoteRequest]) (*connect_go.Response[v1.CreateQuoteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("quote.QuoteService.CreateQuote is not implemented"))
}

func (UnimplementedQuoteServiceHandler) GetQuotes(context.Context, *connect_go.Request[v1.GetQuotesRequest]) (*connect_go.Response[v1.GetQuotesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("quote.QuoteService.GetQuotes is not implemented"))
}

func (UnimplementedQuoteServiceHandler) GetQuote(context.Context, *connect_go.Request[v1.GetQuoteRequest]) (*connect_go.Response[v1.Quote], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("quote.QuoteService.GetQuote is not implemented"))
}

// CouponServiceClient is a client for the quote.CouponService service.
type CouponServiceClient interface {
	// 获取用户拥有的优惠券列表
	GetUserCoupons(context.Context, *connect_go.Request[v1.GetUserCouponsRequest]) (*connect_go.Response[v1.GetUserCouponsResponse], error)
	// 使用优惠券
	UseCoupon(context.Context, *connect_go.Request[v1.UseCouponRequest]) (*connect_go.Response[v1.UseCouponResponse], error)
}

// NewCouponServiceClient constructs a client for the quote.CouponService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCouponServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) CouponServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &couponServiceClient{
		getUserCoupons: connect_go.NewClient[v1.GetUserCouponsRequest, v1.GetUserCouponsResponse](
			httpClient,
			baseURL+CouponServiceGetUserCouponsProcedure,
			opts...,
		),
		useCoupon: connect_go.NewClient[v1.UseCouponRequest, v1.UseCouponResponse](
			httpClient,
			baseURL+CouponServiceUseCouponProcedure,
			opts...,
		),
	}
}

// couponServiceClient implements CouponServiceClient.
type couponServiceClient struct {
	getUserCoupons *connect_go.Client[v1.GetUserCouponsRequest, v1.GetUserCouponsResponse]
	useCoupon      *connect_go.Client[v1.UseCouponRequest, v1.UseCouponResponse]
}

// GetUserCoupons calls quote.CouponService.GetUserCoupons.
func (c *couponServiceClient) GetUserCoupons(ctx context.Context, req *connect_go.Request[v1.GetUserCouponsRequest]) (*connect_go.Response[v1.GetUserCouponsResponse], error) {
	return c.getUserCoupons.CallUnary(ctx, req)
}

// UseCoupon calls quote.CouponService.UseCoupon.
func (c *couponServiceClient) UseCoupon(ctx context.Context, req *connect_go.Request[v1.UseCouponRequest]) (*connect_go.Response[v1.UseCouponResponse], error) {
	return c.useCoupon.CallUnary(ctx, req)
}

// CouponServiceHandler is an implementation of the quote.CouponService service.
type CouponServiceHandler interface {
	// 获取用户拥有的优惠券列表
	GetUserCoupons(context.Context, *connect_go.Request[v1.GetUserCouponsRequest]) (*connect_go.Response[v1.GetUserCouponsResponse], error)
	// 使用优惠券
	UseCoupon(context.Context, *connect_go.Request[v1.UseCouponRequest]) (*connect_go.Response[v1.UseCouponResponse], error)
}

// NewCouponServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCouponServiceHandler(svc CouponServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	couponServiceGetUserCouponsHandler := connect_go.NewUnaryHandler(
		CouponServiceGetUserCouponsProcedure,
		svc.GetUserCoupons,
		opts...,
	)
	couponServiceUseCouponHandler := connect_go.NewUnaryHandler(
		CouponServiceUseCouponProcedure,
		svc.UseCoupon,
		opts...,
	)
	return "/quote.CouponService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CouponServiceGetUserCouponsProcedure:
			couponServiceGetUserCouponsHandler.ServeHTTP(w, r)
		case CouponServiceUseCouponProcedure:
			couponServiceUseCouponHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCouponServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCouponServiceHandler struct{}

func (UnimplementedCouponServiceHandler) GetUserCoupons(context.Context, *connect_go.Request[v1.GetUserCouponsRequest]) (*connect_go.Response[v1.GetUserCouponsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("quote.CouponService.GetUserCoupons is not implemented"))
}

func (UnimplementedCouponServiceHandler) UseCoupon(context.Context, *connect_go.Request[v1.UseCouponRequest]) (*connect_go.Response[v1.UseCouponResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("quote.CouponService.UseCoupon is not implemented"))
}
