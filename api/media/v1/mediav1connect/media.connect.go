// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/media/v1/media.proto

package mediav1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/ydssx/morphix/api/media/v1"
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
	// MediaServiceName is the fully-qualified name of the MediaService service.
	MediaServiceName = "mediav1.MediaService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// MediaServiceUploadVideoProcedure is the fully-qualified name of the MediaService's UploadVideo
	// RPC.
	MediaServiceUploadVideoProcedure = "/mediav1.MediaService/UploadVideo"
	// MediaServiceTranscodeVideoProcedure is the fully-qualified name of the MediaService's
	// TranscodeVideo RPC.
	MediaServiceTranscodeVideoProcedure = "/mediav1.MediaService/TranscodeVideo"
	// MediaServicePlayVideoProcedure is the fully-qualified name of the MediaService's PlayVideo RPC.
	MediaServicePlayVideoProcedure = "/mediav1.MediaService/PlayVideo"
	// MediaServiceTakeScreenshotProcedure is the fully-qualified name of the MediaService's
	// TakeScreenshot RPC.
	MediaServiceTakeScreenshotProcedure = "/mediav1.MediaService/TakeScreenshot"
)

// MediaServiceClient is a client for the mediav1.MediaService service.
type MediaServiceClient interface {
	// 视频上传
	UploadVideo(context.Context, *connect_go.Request[v1.VideoUploadRequest]) (*connect_go.Response[v1.VideoUploadResponse], error)
	// 视频转码
	TranscodeVideo(context.Context, *connect_go.Request[v1.VideoTranscodeRequest]) (*connect_go.Response[v1.VideoTranscodeResponse], error)
	// 视频播放
	PlayVideo(context.Context, *connect_go.Request[v1.VideoPlaybackRequest]) (*connect_go.Response[v1.VideoPlaybackResponse], error)
	// 视频截图
	TakeScreenshot(context.Context, *connect_go.Request[v1.VideoScreenshotRequest]) (*connect_go.Response[v1.VideoScreenshotResponse], error)
}

// NewMediaServiceClient constructs a client for the mediav1.MediaService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMediaServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) MediaServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &mediaServiceClient{
		uploadVideo: connect_go.NewClient[v1.VideoUploadRequest, v1.VideoUploadResponse](
			httpClient,
			baseURL+MediaServiceUploadVideoProcedure,
			opts...,
		),
		transcodeVideo: connect_go.NewClient[v1.VideoTranscodeRequest, v1.VideoTranscodeResponse](
			httpClient,
			baseURL+MediaServiceTranscodeVideoProcedure,
			opts...,
		),
		playVideo: connect_go.NewClient[v1.VideoPlaybackRequest, v1.VideoPlaybackResponse](
			httpClient,
			baseURL+MediaServicePlayVideoProcedure,
			opts...,
		),
		takeScreenshot: connect_go.NewClient[v1.VideoScreenshotRequest, v1.VideoScreenshotResponse](
			httpClient,
			baseURL+MediaServiceTakeScreenshotProcedure,
			opts...,
		),
	}
}

// mediaServiceClient implements MediaServiceClient.
type mediaServiceClient struct {
	uploadVideo    *connect_go.Client[v1.VideoUploadRequest, v1.VideoUploadResponse]
	transcodeVideo *connect_go.Client[v1.VideoTranscodeRequest, v1.VideoTranscodeResponse]
	playVideo      *connect_go.Client[v1.VideoPlaybackRequest, v1.VideoPlaybackResponse]
	takeScreenshot *connect_go.Client[v1.VideoScreenshotRequest, v1.VideoScreenshotResponse]
}

// UploadVideo calls mediav1.MediaService.UploadVideo.
func (c *mediaServiceClient) UploadVideo(ctx context.Context, req *connect_go.Request[v1.VideoUploadRequest]) (*connect_go.Response[v1.VideoUploadResponse], error) {
	return c.uploadVideo.CallUnary(ctx, req)
}

// TranscodeVideo calls mediav1.MediaService.TranscodeVideo.
func (c *mediaServiceClient) TranscodeVideo(ctx context.Context, req *connect_go.Request[v1.VideoTranscodeRequest]) (*connect_go.Response[v1.VideoTranscodeResponse], error) {
	return c.transcodeVideo.CallUnary(ctx, req)
}

// PlayVideo calls mediav1.MediaService.PlayVideo.
func (c *mediaServiceClient) PlayVideo(ctx context.Context, req *connect_go.Request[v1.VideoPlaybackRequest]) (*connect_go.Response[v1.VideoPlaybackResponse], error) {
	return c.playVideo.CallUnary(ctx, req)
}

// TakeScreenshot calls mediav1.MediaService.TakeScreenshot.
func (c *mediaServiceClient) TakeScreenshot(ctx context.Context, req *connect_go.Request[v1.VideoScreenshotRequest]) (*connect_go.Response[v1.VideoScreenshotResponse], error) {
	return c.takeScreenshot.CallUnary(ctx, req)
}

// MediaServiceHandler is an implementation of the mediav1.MediaService service.
type MediaServiceHandler interface {
	// 视频上传
	UploadVideo(context.Context, *connect_go.Request[v1.VideoUploadRequest]) (*connect_go.Response[v1.VideoUploadResponse], error)
	// 视频转码
	TranscodeVideo(context.Context, *connect_go.Request[v1.VideoTranscodeRequest]) (*connect_go.Response[v1.VideoTranscodeResponse], error)
	// 视频播放
	PlayVideo(context.Context, *connect_go.Request[v1.VideoPlaybackRequest]) (*connect_go.Response[v1.VideoPlaybackResponse], error)
	// 视频截图
	TakeScreenshot(context.Context, *connect_go.Request[v1.VideoScreenshotRequest]) (*connect_go.Response[v1.VideoScreenshotResponse], error)
}

// NewMediaServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMediaServiceHandler(svc MediaServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mediaServiceUploadVideoHandler := connect_go.NewUnaryHandler(
		MediaServiceUploadVideoProcedure,
		svc.UploadVideo,
		opts...,
	)
	mediaServiceTranscodeVideoHandler := connect_go.NewUnaryHandler(
		MediaServiceTranscodeVideoProcedure,
		svc.TranscodeVideo,
		opts...,
	)
	mediaServicePlayVideoHandler := connect_go.NewUnaryHandler(
		MediaServicePlayVideoProcedure,
		svc.PlayVideo,
		opts...,
	)
	mediaServiceTakeScreenshotHandler := connect_go.NewUnaryHandler(
		MediaServiceTakeScreenshotProcedure,
		svc.TakeScreenshot,
		opts...,
	)
	return "/mediav1.MediaService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case MediaServiceUploadVideoProcedure:
			mediaServiceUploadVideoHandler.ServeHTTP(w, r)
		case MediaServiceTranscodeVideoProcedure:
			mediaServiceTranscodeVideoHandler.ServeHTTP(w, r)
		case MediaServicePlayVideoProcedure:
			mediaServicePlayVideoHandler.ServeHTTP(w, r)
		case MediaServiceTakeScreenshotProcedure:
			mediaServiceTakeScreenshotHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedMediaServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedMediaServiceHandler struct{}

func (UnimplementedMediaServiceHandler) UploadVideo(context.Context, *connect_go.Request[v1.VideoUploadRequest]) (*connect_go.Response[v1.VideoUploadResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("mediav1.MediaService.UploadVideo is not implemented"))
}

func (UnimplementedMediaServiceHandler) TranscodeVideo(context.Context, *connect_go.Request[v1.VideoTranscodeRequest]) (*connect_go.Response[v1.VideoTranscodeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("mediav1.MediaService.TranscodeVideo is not implemented"))
}

func (UnimplementedMediaServiceHandler) PlayVideo(context.Context, *connect_go.Request[v1.VideoPlaybackRequest]) (*connect_go.Response[v1.VideoPlaybackResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("mediav1.MediaService.PlayVideo is not implemented"))
}

func (UnimplementedMediaServiceHandler) TakeScreenshot(context.Context, *connect_go.Request[v1.VideoScreenshotRequest]) (*connect_go.Response[v1.VideoScreenshotResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("mediav1.MediaService.TakeScreenshot is not implemented"))
}
