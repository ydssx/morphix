// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/quote/v1/quote.proto

package quotev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	QuoteService_CreateQuote_FullMethodName    = "/quote.QuoteService/CreateQuote"
	QuoteService_GetQuotes_FullMethodName      = "/quote.QuoteService/GetQuotes"
	QuoteService_GetQuote_FullMethodName       = "/quote.QuoteService/GetQuote"
	QuoteService_GetUserCoupons_FullMethodName = "/quote.QuoteService/GetUserCoupons"
	QuoteService_UseCoupon_FullMethodName      = "/quote.QuoteService/UseCoupon"
)

// QuoteServiceClient is the client API for QuoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuoteServiceClient interface {
	// 创建报价
	CreateQuote(ctx context.Context, in *CreateQuoteRequest, opts ...grpc.CallOption) (*CreateQuoteResponse, error)
	// 获取报价列表
	GetQuotes(ctx context.Context, in *GetQuotesRequest, opts ...grpc.CallOption) (*GetQuotesResponse, error)
	// 获取单个报价
	GetQuote(ctx context.Context, in *GetQuoteRequest, opts ...grpc.CallOption) (*Quote, error)
	// 获取用户拥有的优惠券列表
	GetUserCoupons(ctx context.Context, in *GetUserCouponsRequest, opts ...grpc.CallOption) (*GetUserCouponsResponse, error)
	// 使用优惠券
	UseCoupon(ctx context.Context, in *UseCouponRequest, opts ...grpc.CallOption) (*UseCouponResponse, error)
}

type quoteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuoteServiceClient(cc grpc.ClientConnInterface) QuoteServiceClient {
	return &quoteServiceClient{cc}
}

func (c *quoteServiceClient) CreateQuote(ctx context.Context, in *CreateQuoteRequest, opts ...grpc.CallOption) (*CreateQuoteResponse, error) {
	out := new(CreateQuoteResponse)
	err := c.cc.Invoke(ctx, QuoteService_CreateQuote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quoteServiceClient) GetQuotes(ctx context.Context, in *GetQuotesRequest, opts ...grpc.CallOption) (*GetQuotesResponse, error) {
	out := new(GetQuotesResponse)
	err := c.cc.Invoke(ctx, QuoteService_GetQuotes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quoteServiceClient) GetQuote(ctx context.Context, in *GetQuoteRequest, opts ...grpc.CallOption) (*Quote, error) {
	out := new(Quote)
	err := c.cc.Invoke(ctx, QuoteService_GetQuote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quoteServiceClient) GetUserCoupons(ctx context.Context, in *GetUserCouponsRequest, opts ...grpc.CallOption) (*GetUserCouponsResponse, error) {
	out := new(GetUserCouponsResponse)
	err := c.cc.Invoke(ctx, QuoteService_GetUserCoupons_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quoteServiceClient) UseCoupon(ctx context.Context, in *UseCouponRequest, opts ...grpc.CallOption) (*UseCouponResponse, error) {
	out := new(UseCouponResponse)
	err := c.cc.Invoke(ctx, QuoteService_UseCoupon_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuoteServiceServer is the server API for QuoteService service.
// All implementations should embed UnimplementedQuoteServiceServer
// for forward compatibility
type QuoteServiceServer interface {
	// 创建报价
	CreateQuote(context.Context, *CreateQuoteRequest) (*CreateQuoteResponse, error)
	// 获取报价列表
	GetQuotes(context.Context, *GetQuotesRequest) (*GetQuotesResponse, error)
	// 获取单个报价
	GetQuote(context.Context, *GetQuoteRequest) (*Quote, error)
	// 获取用户拥有的优惠券列表
	GetUserCoupons(context.Context, *GetUserCouponsRequest) (*GetUserCouponsResponse, error)
	// 使用优惠券
	UseCoupon(context.Context, *UseCouponRequest) (*UseCouponResponse, error)
}

// UnimplementedQuoteServiceServer should be embedded to have forward compatible implementations.
type UnimplementedQuoteServiceServer struct {
}

func (UnimplementedQuoteServiceServer) CreateQuote(context.Context, *CreateQuoteRequest) (*CreateQuoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuote not implemented")
}
func (UnimplementedQuoteServiceServer) GetQuotes(context.Context, *GetQuotesRequest) (*GetQuotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuotes not implemented")
}
func (UnimplementedQuoteServiceServer) GetQuote(context.Context, *GetQuoteRequest) (*Quote, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuote not implemented")
}
func (UnimplementedQuoteServiceServer) GetUserCoupons(context.Context, *GetUserCouponsRequest) (*GetUserCouponsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCoupons not implemented")
}
func (UnimplementedQuoteServiceServer) UseCoupon(context.Context, *UseCouponRequest) (*UseCouponResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UseCoupon not implemented")
}

// UnsafeQuoteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuoteServiceServer will
// result in compilation errors.
type UnsafeQuoteServiceServer interface {
	mustEmbedUnimplementedQuoteServiceServer()
}

func RegisterQuoteServiceServer(s grpc.ServiceRegistrar, srv QuoteServiceServer) {
	s.RegisterService(&QuoteService_ServiceDesc, srv)
}

func _QuoteService_CreateQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuoteServiceServer).CreateQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuoteService_CreateQuote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuoteServiceServer).CreateQuote(ctx, req.(*CreateQuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuoteService_GetQuotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuoteServiceServer).GetQuotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuoteService_GetQuotes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuoteServiceServer).GetQuotes(ctx, req.(*GetQuotesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuoteService_GetQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuoteServiceServer).GetQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuoteService_GetQuote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuoteServiceServer).GetQuote(ctx, req.(*GetQuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuoteService_GetUserCoupons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserCouponsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuoteServiceServer).GetUserCoupons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuoteService_GetUserCoupons_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuoteServiceServer).GetUserCoupons(ctx, req.(*GetUserCouponsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuoteService_UseCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UseCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuoteServiceServer).UseCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuoteService_UseCoupon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuoteServiceServer).UseCoupon(ctx, req.(*UseCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuoteService_ServiceDesc is the grpc.ServiceDesc for QuoteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuoteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "quote.QuoteService",
	HandlerType: (*QuoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQuote",
			Handler:    _QuoteService_CreateQuote_Handler,
		},
		{
			MethodName: "GetQuotes",
			Handler:    _QuoteService_GetQuotes_Handler,
		},
		{
			MethodName: "GetQuote",
			Handler:    _QuoteService_GetQuote_Handler,
		},
		{
			MethodName: "GetUserCoupons",
			Handler:    _QuoteService_GetUserCoupons_Handler,
		},
		{
			MethodName: "UseCoupon",
			Handler:    _QuoteService_UseCoupon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/quote/v1/quote.proto",
}
