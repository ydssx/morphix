// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             (unknown)
// source: api/quote/v1/quote.proto

package quotev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationQuoteServiceGetQuotes = "/quote.QuoteService/GetQuotes"

type QuoteServiceHTTPServer interface {
	// GetQuotes 获取报价列表
	GetQuotes(context.Context, *GetQuotesRequest) (*GetQuotesResponse, error)
}

func RegisterQuoteServiceHTTPServer(s *http.Server, srv QuoteServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/api/v1/quotes", _QuoteService_GetQuotes0_HTTP_Handler(srv))
}

func _QuoteService_GetQuotes0_HTTP_Handler(srv QuoteServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetQuotesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQuoteServiceGetQuotes)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetQuotes(ctx, req.(*GetQuotesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetQuotesResponse)
		return ctx.Result(200, reply)
	}
}

type QuoteServiceHTTPClient interface {
	GetQuotes(ctx context.Context, req *GetQuotesRequest, opts ...http.CallOption) (rsp *GetQuotesResponse, err error)
}

type QuoteServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewQuoteServiceHTTPClient(client *http.Client) QuoteServiceHTTPClient {
	return &QuoteServiceHTTPClientImpl{client}
}

func (c *QuoteServiceHTTPClientImpl) GetQuotes(ctx context.Context, in *GetQuotesRequest, opts ...http.CallOption) (*GetQuotesResponse, error) {
	var out GetQuotesResponse
	pattern := "/api/v1/quotes"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationQuoteServiceGetQuotes))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
