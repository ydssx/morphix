// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/sms/v1/sms.proto

package smsv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/ydssx/morphix/api/sms/v1"
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
	// SMSServiceName is the fully-qualified name of the SMSService service.
	SMSServiceName = "smsv1.SMSService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SMSServiceSendSMSProcedure is the fully-qualified name of the SMSService's SendSMS RPC.
	SMSServiceSendSMSProcedure = "/smsv1.SMSService/SendSMS"
	// SMSServiceCheckSMSStatusProcedure is the fully-qualified name of the SMSService's CheckSMSStatus
	// RPC.
	SMSServiceCheckSMSStatusProcedure = "/smsv1.SMSService/CheckSMSStatus"
	// SMSServiceManageSMSTemplateProcedure is the fully-qualified name of the SMSService's
	// ManageSMSTemplate RPC.
	SMSServiceManageSMSTemplateProcedure = "/smsv1.SMSService/ManageSMSTemplate"
	// SMSServiceManageSMSSignatureProcedure is the fully-qualified name of the SMSService's
	// ManageSMSSignature RPC.
	SMSServiceManageSMSSignatureProcedure = "/smsv1.SMSService/ManageSMSSignature"
)

// SMSServiceClient is a client for the smsv1.SMSService service.
type SMSServiceClient interface {
	// 发送短信
	SendSMS(context.Context, *connect_go.Request[v1.SendSMSRequest]) (*connect_go.Response[v1.SendSMSResponse], error)
	// 查询短信状态
	CheckSMSStatus(context.Context, *connect_go.Request[v1.QuerySMSStatusRequest]) (*connect_go.Response[v1.QuerySMSStatusResponse], error)
	// 管理短信模板
	ManageSMSTemplate(context.Context, *connect_go.Request[v1.TemplateManagementRequest]) (*connect_go.Response[v1.TemplateManagementResponse], error)
	// 管理短信签名
	ManageSMSSignature(context.Context, *connect_go.Request[v1.SignatureManagementRequest]) (*connect_go.Response[v1.SignatureManagementResponse], error)
}

// NewSMSServiceClient constructs a client for the smsv1.SMSService service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSMSServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) SMSServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &sMSServiceClient{
		sendSMS: connect_go.NewClient[v1.SendSMSRequest, v1.SendSMSResponse](
			httpClient,
			baseURL+SMSServiceSendSMSProcedure,
			opts...,
		),
		checkSMSStatus: connect_go.NewClient[v1.QuerySMSStatusRequest, v1.QuerySMSStatusResponse](
			httpClient,
			baseURL+SMSServiceCheckSMSStatusProcedure,
			opts...,
		),
		manageSMSTemplate: connect_go.NewClient[v1.TemplateManagementRequest, v1.TemplateManagementResponse](
			httpClient,
			baseURL+SMSServiceManageSMSTemplateProcedure,
			opts...,
		),
		manageSMSSignature: connect_go.NewClient[v1.SignatureManagementRequest, v1.SignatureManagementResponse](
			httpClient,
			baseURL+SMSServiceManageSMSSignatureProcedure,
			opts...,
		),
	}
}

// sMSServiceClient implements SMSServiceClient.
type sMSServiceClient struct {
	sendSMS            *connect_go.Client[v1.SendSMSRequest, v1.SendSMSResponse]
	checkSMSStatus     *connect_go.Client[v1.QuerySMSStatusRequest, v1.QuerySMSStatusResponse]
	manageSMSTemplate  *connect_go.Client[v1.TemplateManagementRequest, v1.TemplateManagementResponse]
	manageSMSSignature *connect_go.Client[v1.SignatureManagementRequest, v1.SignatureManagementResponse]
}

// SendSMS calls smsv1.SMSService.SendSMS.
func (c *sMSServiceClient) SendSMS(ctx context.Context, req *connect_go.Request[v1.SendSMSRequest]) (*connect_go.Response[v1.SendSMSResponse], error) {
	return c.sendSMS.CallUnary(ctx, req)
}

// CheckSMSStatus calls smsv1.SMSService.CheckSMSStatus.
func (c *sMSServiceClient) CheckSMSStatus(ctx context.Context, req *connect_go.Request[v1.QuerySMSStatusRequest]) (*connect_go.Response[v1.QuerySMSStatusResponse], error) {
	return c.checkSMSStatus.CallUnary(ctx, req)
}

// ManageSMSTemplate calls smsv1.SMSService.ManageSMSTemplate.
func (c *sMSServiceClient) ManageSMSTemplate(ctx context.Context, req *connect_go.Request[v1.TemplateManagementRequest]) (*connect_go.Response[v1.TemplateManagementResponse], error) {
	return c.manageSMSTemplate.CallUnary(ctx, req)
}

// ManageSMSSignature calls smsv1.SMSService.ManageSMSSignature.
func (c *sMSServiceClient) ManageSMSSignature(ctx context.Context, req *connect_go.Request[v1.SignatureManagementRequest]) (*connect_go.Response[v1.SignatureManagementResponse], error) {
	return c.manageSMSSignature.CallUnary(ctx, req)
}

// SMSServiceHandler is an implementation of the smsv1.SMSService service.
type SMSServiceHandler interface {
	// 发送短信
	SendSMS(context.Context, *connect_go.Request[v1.SendSMSRequest]) (*connect_go.Response[v1.SendSMSResponse], error)
	// 查询短信状态
	CheckSMSStatus(context.Context, *connect_go.Request[v1.QuerySMSStatusRequest]) (*connect_go.Response[v1.QuerySMSStatusResponse], error)
	// 管理短信模板
	ManageSMSTemplate(context.Context, *connect_go.Request[v1.TemplateManagementRequest]) (*connect_go.Response[v1.TemplateManagementResponse], error)
	// 管理短信签名
	ManageSMSSignature(context.Context, *connect_go.Request[v1.SignatureManagementRequest]) (*connect_go.Response[v1.SignatureManagementResponse], error)
}

// NewSMSServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSMSServiceHandler(svc SMSServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	sMSServiceSendSMSHandler := connect_go.NewUnaryHandler(
		SMSServiceSendSMSProcedure,
		svc.SendSMS,
		opts...,
	)
	sMSServiceCheckSMSStatusHandler := connect_go.NewUnaryHandler(
		SMSServiceCheckSMSStatusProcedure,
		svc.CheckSMSStatus,
		opts...,
	)
	sMSServiceManageSMSTemplateHandler := connect_go.NewUnaryHandler(
		SMSServiceManageSMSTemplateProcedure,
		svc.ManageSMSTemplate,
		opts...,
	)
	sMSServiceManageSMSSignatureHandler := connect_go.NewUnaryHandler(
		SMSServiceManageSMSSignatureProcedure,
		svc.ManageSMSSignature,
		opts...,
	)
	return "/smsv1.SMSService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SMSServiceSendSMSProcedure:
			sMSServiceSendSMSHandler.ServeHTTP(w, r)
		case SMSServiceCheckSMSStatusProcedure:
			sMSServiceCheckSMSStatusHandler.ServeHTTP(w, r)
		case SMSServiceManageSMSTemplateProcedure:
			sMSServiceManageSMSTemplateHandler.ServeHTTP(w, r)
		case SMSServiceManageSMSSignatureProcedure:
			sMSServiceManageSMSSignatureHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSMSServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSMSServiceHandler struct{}

func (UnimplementedSMSServiceHandler) SendSMS(context.Context, *connect_go.Request[v1.SendSMSRequest]) (*connect_go.Response[v1.SendSMSResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("smsv1.SMSService.SendSMS is not implemented"))
}

func (UnimplementedSMSServiceHandler) CheckSMSStatus(context.Context, *connect_go.Request[v1.QuerySMSStatusRequest]) (*connect_go.Response[v1.QuerySMSStatusResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("smsv1.SMSService.CheckSMSStatus is not implemented"))
}

func (UnimplementedSMSServiceHandler) ManageSMSTemplate(context.Context, *connect_go.Request[v1.TemplateManagementRequest]) (*connect_go.Response[v1.TemplateManagementResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("smsv1.SMSService.ManageSMSTemplate is not implemented"))
}

func (UnimplementedSMSServiceHandler) ManageSMSSignature(context.Context, *connect_go.Request[v1.SignatureManagementRequest]) (*connect_go.Response[v1.SignatureManagementResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("smsv1.SMSService.ManageSMSSignature is not implemented"))
}
