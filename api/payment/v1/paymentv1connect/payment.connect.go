// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/payment/v1/payment.proto

package paymentv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/ydssx/morphix/api/payment/v1"
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
	// PaymentServiceName is the fully-qualified name of the PaymentService service.
	PaymentServiceName = "paymentv1.PaymentService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PaymentServiceMakePaymentProcedure is the fully-qualified name of the PaymentService's
	// MakePayment RPC.
	PaymentServiceMakePaymentProcedure = "/paymentv1.PaymentService/MakePayment"
	// PaymentServiceGetPaymentProcedure is the fully-qualified name of the PaymentService's GetPayment
	// RPC.
	PaymentServiceGetPaymentProcedure = "/paymentv1.PaymentService/GetPayment"
	// PaymentServiceCancelPaymentProcedure is the fully-qualified name of the PaymentService's
	// CancelPayment RPC.
	PaymentServiceCancelPaymentProcedure = "/paymentv1.PaymentService/CancelPayment"
	// PaymentServiceRefundProcedure is the fully-qualified name of the PaymentService's Refund RPC.
	PaymentServiceRefundProcedure = "/paymentv1.PaymentService/Refund"
)

// PaymentServiceClient is a client for the paymentv1.PaymentService service.
type PaymentServiceClient interface {
	// 发起支付
	MakePayment(context.Context, *connect_go.Request[v1.MakePaymentRequest]) (*connect_go.Response[v1.PaymentResponse], error)
	// 查询支付状态
	GetPayment(context.Context, *connect_go.Request[v1.GetPaymentRequest]) (*connect_go.Response[v1.GetPaymentResponse], error)
	// 取消支付
	CancelPayment(context.Context, *connect_go.Request[v1.CancelPaymentRequest]) (*connect_go.Response[v1.CancelPaymentResponse], error)
	// 退款
	Refund(context.Context, *connect_go.Request[v1.RefundRequest]) (*connect_go.Response[v1.RefundResponse], error)
}

// NewPaymentServiceClient constructs a client for the paymentv1.PaymentService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPaymentServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) PaymentServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &paymentServiceClient{
		makePayment: connect_go.NewClient[v1.MakePaymentRequest, v1.PaymentResponse](
			httpClient,
			baseURL+PaymentServiceMakePaymentProcedure,
			opts...,
		),
		getPayment: connect_go.NewClient[v1.GetPaymentRequest, v1.GetPaymentResponse](
			httpClient,
			baseURL+PaymentServiceGetPaymentProcedure,
			opts...,
		),
		cancelPayment: connect_go.NewClient[v1.CancelPaymentRequest, v1.CancelPaymentResponse](
			httpClient,
			baseURL+PaymentServiceCancelPaymentProcedure,
			opts...,
		),
		refund: connect_go.NewClient[v1.RefundRequest, v1.RefundResponse](
			httpClient,
			baseURL+PaymentServiceRefundProcedure,
			opts...,
		),
	}
}

// paymentServiceClient implements PaymentServiceClient.
type paymentServiceClient struct {
	makePayment   *connect_go.Client[v1.MakePaymentRequest, v1.PaymentResponse]
	getPayment    *connect_go.Client[v1.GetPaymentRequest, v1.GetPaymentResponse]
	cancelPayment *connect_go.Client[v1.CancelPaymentRequest, v1.CancelPaymentResponse]
	refund        *connect_go.Client[v1.RefundRequest, v1.RefundResponse]
}

// MakePayment calls paymentv1.PaymentService.MakePayment.
func (c *paymentServiceClient) MakePayment(ctx context.Context, req *connect_go.Request[v1.MakePaymentRequest]) (*connect_go.Response[v1.PaymentResponse], error) {
	return c.makePayment.CallUnary(ctx, req)
}

// GetPayment calls paymentv1.PaymentService.GetPayment.
func (c *paymentServiceClient) GetPayment(ctx context.Context, req *connect_go.Request[v1.GetPaymentRequest]) (*connect_go.Response[v1.GetPaymentResponse], error) {
	return c.getPayment.CallUnary(ctx, req)
}

// CancelPayment calls paymentv1.PaymentService.CancelPayment.
func (c *paymentServiceClient) CancelPayment(ctx context.Context, req *connect_go.Request[v1.CancelPaymentRequest]) (*connect_go.Response[v1.CancelPaymentResponse], error) {
	return c.cancelPayment.CallUnary(ctx, req)
}

// Refund calls paymentv1.PaymentService.Refund.
func (c *paymentServiceClient) Refund(ctx context.Context, req *connect_go.Request[v1.RefundRequest]) (*connect_go.Response[v1.RefundResponse], error) {
	return c.refund.CallUnary(ctx, req)
}

// PaymentServiceHandler is an implementation of the paymentv1.PaymentService service.
type PaymentServiceHandler interface {
	// 发起支付
	MakePayment(context.Context, *connect_go.Request[v1.MakePaymentRequest]) (*connect_go.Response[v1.PaymentResponse], error)
	// 查询支付状态
	GetPayment(context.Context, *connect_go.Request[v1.GetPaymentRequest]) (*connect_go.Response[v1.GetPaymentResponse], error)
	// 取消支付
	CancelPayment(context.Context, *connect_go.Request[v1.CancelPaymentRequest]) (*connect_go.Response[v1.CancelPaymentResponse], error)
	// 退款
	Refund(context.Context, *connect_go.Request[v1.RefundRequest]) (*connect_go.Response[v1.RefundResponse], error)
}

// NewPaymentServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPaymentServiceHandler(svc PaymentServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	paymentServiceMakePaymentHandler := connect_go.NewUnaryHandler(
		PaymentServiceMakePaymentProcedure,
		svc.MakePayment,
		opts...,
	)
	paymentServiceGetPaymentHandler := connect_go.NewUnaryHandler(
		PaymentServiceGetPaymentProcedure,
		svc.GetPayment,
		opts...,
	)
	paymentServiceCancelPaymentHandler := connect_go.NewUnaryHandler(
		PaymentServiceCancelPaymentProcedure,
		svc.CancelPayment,
		opts...,
	)
	paymentServiceRefundHandler := connect_go.NewUnaryHandler(
		PaymentServiceRefundProcedure,
		svc.Refund,
		opts...,
	)
	return "/paymentv1.PaymentService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PaymentServiceMakePaymentProcedure:
			paymentServiceMakePaymentHandler.ServeHTTP(w, r)
		case PaymentServiceGetPaymentProcedure:
			paymentServiceGetPaymentHandler.ServeHTTP(w, r)
		case PaymentServiceCancelPaymentProcedure:
			paymentServiceCancelPaymentHandler.ServeHTTP(w, r)
		case PaymentServiceRefundProcedure:
			paymentServiceRefundHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPaymentServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPaymentServiceHandler struct{}

func (UnimplementedPaymentServiceHandler) MakePayment(context.Context, *connect_go.Request[v1.MakePaymentRequest]) (*connect_go.Response[v1.PaymentResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("paymentv1.PaymentService.MakePayment is not implemented"))
}

func (UnimplementedPaymentServiceHandler) GetPayment(context.Context, *connect_go.Request[v1.GetPaymentRequest]) (*connect_go.Response[v1.GetPaymentResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("paymentv1.PaymentService.GetPayment is not implemented"))
}

func (UnimplementedPaymentServiceHandler) CancelPayment(context.Context, *connect_go.Request[v1.CancelPaymentRequest]) (*connect_go.Response[v1.CancelPaymentResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("paymentv1.PaymentService.CancelPayment is not implemented"))
}

func (UnimplementedPaymentServiceHandler) Refund(context.Context, *connect_go.Request[v1.RefundRequest]) (*connect_go.Response[v1.RefundResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("paymentv1.PaymentService.Refund is not implemented"))
}