// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/aiart/v1/aiart.proto

package aiartv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/ydssx/morphix/api/aiart/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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
	// ArtServiceName is the fully-qualified name of the ArtService service.
	ArtServiceName = "aiartv1.ArtService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ArtServiceGenerateImageProcedure is the fully-qualified name of the ArtService's GenerateImage
	// RPC.
	ArtServiceGenerateImageProcedure = "/aiartv1.ArtService/GenerateImage"
	// ArtServiceGetGenerateStatusProcedure is the fully-qualified name of the ArtService's
	// GetGenerateStatus RPC.
	ArtServiceGetGenerateStatusProcedure = "/aiartv1.ArtService/GetGenerateStatus"
	// ArtServiceGetGeneratedImageProcedure is the fully-qualified name of the ArtService's
	// GetGeneratedImage RPC.
	ArtServiceGetGeneratedImageProcedure = "/aiartv1.ArtService/GetGeneratedImage"
	// ArtServiceGetModelInfoProcedure is the fully-qualified name of the ArtService's GetModelInfo RPC.
	ArtServiceGetModelInfoProcedure = "/aiartv1.ArtService/GetModelInfo"
	// ArtServiceImageToImageProcedure is the fully-qualified name of the ArtService's ImageToImage RPC.
	ArtServiceImageToImageProcedure = "/aiartv1.ArtService/ImageToImage"
)

// ArtServiceClient is a client for the aiartv1.ArtService service.
type ArtServiceClient interface {
	// 生成图像
	GenerateImage(context.Context, *connect_go.Request[v1.GenerateImageRequest]) (*connect_go.Response[v1.GenerateImageResponse], error)
	// 获取生成任务状态
	GetGenerateStatus(context.Context, *connect_go.Request[v1.GetGenerateStatusRequest]) (*connect_go.Response[v1.GenerateStatusResponse], error)
	// 获取已生成的图像
	GetGeneratedImage(context.Context, *connect_go.Request[v1.GetGeneratedImageRequest]) (*connect_go.Response[v1.GetGeneratedImageResponse], error)
	// 获取模型信息
	GetModelInfo(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.GetModelInfoResponse], error)
	ImageToImage(context.Context, *connect_go.Request[v1.ImageToImageRequest]) (*connect_go.Response[v1.ImageToImageResponse], error)
}

// NewArtServiceClient constructs a client for the aiartv1.ArtService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewArtServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ArtServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &artServiceClient{
		generateImage: connect_go.NewClient[v1.GenerateImageRequest, v1.GenerateImageResponse](
			httpClient,
			baseURL+ArtServiceGenerateImageProcedure,
			opts...,
		),
		getGenerateStatus: connect_go.NewClient[v1.GetGenerateStatusRequest, v1.GenerateStatusResponse](
			httpClient,
			baseURL+ArtServiceGetGenerateStatusProcedure,
			opts...,
		),
		getGeneratedImage: connect_go.NewClient[v1.GetGeneratedImageRequest, v1.GetGeneratedImageResponse](
			httpClient,
			baseURL+ArtServiceGetGeneratedImageProcedure,
			opts...,
		),
		getModelInfo: connect_go.NewClient[emptypb.Empty, v1.GetModelInfoResponse](
			httpClient,
			baseURL+ArtServiceGetModelInfoProcedure,
			opts...,
		),
		imageToImage: connect_go.NewClient[v1.ImageToImageRequest, v1.ImageToImageResponse](
			httpClient,
			baseURL+ArtServiceImageToImageProcedure,
			opts...,
		),
	}
}

// artServiceClient implements ArtServiceClient.
type artServiceClient struct {
	generateImage     *connect_go.Client[v1.GenerateImageRequest, v1.GenerateImageResponse]
	getGenerateStatus *connect_go.Client[v1.GetGenerateStatusRequest, v1.GenerateStatusResponse]
	getGeneratedImage *connect_go.Client[v1.GetGeneratedImageRequest, v1.GetGeneratedImageResponse]
	getModelInfo      *connect_go.Client[emptypb.Empty, v1.GetModelInfoResponse]
	imageToImage      *connect_go.Client[v1.ImageToImageRequest, v1.ImageToImageResponse]
}

// GenerateImage calls aiartv1.ArtService.GenerateImage.
func (c *artServiceClient) GenerateImage(ctx context.Context, req *connect_go.Request[v1.GenerateImageRequest]) (*connect_go.Response[v1.GenerateImageResponse], error) {
	return c.generateImage.CallUnary(ctx, req)
}

// GetGenerateStatus calls aiartv1.ArtService.GetGenerateStatus.
func (c *artServiceClient) GetGenerateStatus(ctx context.Context, req *connect_go.Request[v1.GetGenerateStatusRequest]) (*connect_go.Response[v1.GenerateStatusResponse], error) {
	return c.getGenerateStatus.CallUnary(ctx, req)
}

// GetGeneratedImage calls aiartv1.ArtService.GetGeneratedImage.
func (c *artServiceClient) GetGeneratedImage(ctx context.Context, req *connect_go.Request[v1.GetGeneratedImageRequest]) (*connect_go.Response[v1.GetGeneratedImageResponse], error) {
	return c.getGeneratedImage.CallUnary(ctx, req)
}

// GetModelInfo calls aiartv1.ArtService.GetModelInfo.
func (c *artServiceClient) GetModelInfo(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.GetModelInfoResponse], error) {
	return c.getModelInfo.CallUnary(ctx, req)
}

// ImageToImage calls aiartv1.ArtService.ImageToImage.
func (c *artServiceClient) ImageToImage(ctx context.Context, req *connect_go.Request[v1.ImageToImageRequest]) (*connect_go.Response[v1.ImageToImageResponse], error) {
	return c.imageToImage.CallUnary(ctx, req)
}

// ArtServiceHandler is an implementation of the aiartv1.ArtService service.
type ArtServiceHandler interface {
	// 生成图像
	GenerateImage(context.Context, *connect_go.Request[v1.GenerateImageRequest]) (*connect_go.Response[v1.GenerateImageResponse], error)
	// 获取生成任务状态
	GetGenerateStatus(context.Context, *connect_go.Request[v1.GetGenerateStatusRequest]) (*connect_go.Response[v1.GenerateStatusResponse], error)
	// 获取已生成的图像
	GetGeneratedImage(context.Context, *connect_go.Request[v1.GetGeneratedImageRequest]) (*connect_go.Response[v1.GetGeneratedImageResponse], error)
	// 获取模型信息
	GetModelInfo(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.GetModelInfoResponse], error)
	ImageToImage(context.Context, *connect_go.Request[v1.ImageToImageRequest]) (*connect_go.Response[v1.ImageToImageResponse], error)
}

// NewArtServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewArtServiceHandler(svc ArtServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	artServiceGenerateImageHandler := connect_go.NewUnaryHandler(
		ArtServiceGenerateImageProcedure,
		svc.GenerateImage,
		opts...,
	)
	artServiceGetGenerateStatusHandler := connect_go.NewUnaryHandler(
		ArtServiceGetGenerateStatusProcedure,
		svc.GetGenerateStatus,
		opts...,
	)
	artServiceGetGeneratedImageHandler := connect_go.NewUnaryHandler(
		ArtServiceGetGeneratedImageProcedure,
		svc.GetGeneratedImage,
		opts...,
	)
	artServiceGetModelInfoHandler := connect_go.NewUnaryHandler(
		ArtServiceGetModelInfoProcedure,
		svc.GetModelInfo,
		opts...,
	)
	artServiceImageToImageHandler := connect_go.NewUnaryHandler(
		ArtServiceImageToImageProcedure,
		svc.ImageToImage,
		opts...,
	)
	return "/aiartv1.ArtService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ArtServiceGenerateImageProcedure:
			artServiceGenerateImageHandler.ServeHTTP(w, r)
		case ArtServiceGetGenerateStatusProcedure:
			artServiceGetGenerateStatusHandler.ServeHTTP(w, r)
		case ArtServiceGetGeneratedImageProcedure:
			artServiceGetGeneratedImageHandler.ServeHTTP(w, r)
		case ArtServiceGetModelInfoProcedure:
			artServiceGetModelInfoHandler.ServeHTTP(w, r)
		case ArtServiceImageToImageProcedure:
			artServiceImageToImageHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedArtServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedArtServiceHandler struct{}

func (UnimplementedArtServiceHandler) GenerateImage(context.Context, *connect_go.Request[v1.GenerateImageRequest]) (*connect_go.Response[v1.GenerateImageResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("aiartv1.ArtService.GenerateImage is not implemented"))
}

func (UnimplementedArtServiceHandler) GetGenerateStatus(context.Context, *connect_go.Request[v1.GetGenerateStatusRequest]) (*connect_go.Response[v1.GenerateStatusResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("aiartv1.ArtService.GetGenerateStatus is not implemented"))
}

func (UnimplementedArtServiceHandler) GetGeneratedImage(context.Context, *connect_go.Request[v1.GetGeneratedImageRequest]) (*connect_go.Response[v1.GetGeneratedImageResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("aiartv1.ArtService.GetGeneratedImage is not implemented"))
}

func (UnimplementedArtServiceHandler) GetModelInfo(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.GetModelInfoResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("aiartv1.ArtService.GetModelInfo is not implemented"))
}

func (UnimplementedArtServiceHandler) ImageToImage(context.Context, *connect_go.Request[v1.ImageToImageRequest]) (*connect_go.Response[v1.ImageToImageResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("aiartv1.ArtService.ImageToImage is not implemented"))
}
