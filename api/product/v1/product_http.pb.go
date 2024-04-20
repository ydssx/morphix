// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             (unknown)
// source: api/product/v1/product.proto

package productv1

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

const OperationProductServiceCreateProduct = "/productv1.ProductService/CreateProduct"
const OperationProductServiceDeleteProduct = "/productv1.ProductService/DeleteProduct"
const OperationProductServiceGetProduct = "/productv1.ProductService/GetProduct"
const OperationProductServiceGetProductStock = "/productv1.ProductService/GetProductStock"
const OperationProductServiceGetProducts = "/productv1.ProductService/GetProducts"
const OperationProductServiceGetProductsStock = "/productv1.ProductService/GetProductsStock"
const OperationProductServiceUpdateProduct = "/productv1.ProductService/UpdateProduct"
const OperationProductServiceUpdateProductStock = "/productv1.ProductService/UpdateProductStock"

type ProductServiceHTTPServer interface {
	// CreateProduct 创建产品
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error)
	// DeleteProduct 删除产品
	DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error)
	// GetProduct 获取单个产品
	GetProduct(context.Context, *GetProductRequest) (*Product, error)
	GetProductStock(context.Context, *GetProductStockRequest) (*GetProductStockResponse, error)
	// GetProducts 获取产品列表
	GetProducts(context.Context, *GetProductsRequest) (*GetProductsResponse, error)
	// GetProductsStock 获取产品库存
	GetProductsStock(context.Context, *GetProductsStockRequest) (*GetProductsStockResponse, error)
	// UpdateProduct 更新产品信息
	UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductResponse, error)
	UpdateProductStock(context.Context, *UpdateProductStockRequest) (*UpdateProductStockResponse, error)
}

func RegisterProductServiceHTTPServer(s *http.Server, srv ProductServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/products", _ProductService_CreateProduct0_HTTP_Handler(srv))
	r.GET("/api/v1/products", _ProductService_GetProducts0_HTTP_Handler(srv))
	r.GET("/api/v1/products/{id}", _ProductService_GetProduct0_HTTP_Handler(srv))
	r.PUT("/api/v1/products/{id}", _ProductService_UpdateProduct0_HTTP_Handler(srv))
	r.DELETE("/api/v1/products/{id}", _ProductService_DeleteProduct0_HTTP_Handler(srv))
	r.GET("/api/v1/products/{id}/stock", _ProductService_GetProductStock0_HTTP_Handler(srv))
	r.PUT("/api/v1/products/{id}/stock", _ProductService_UpdateProductStock0_HTTP_Handler(srv))
	r.GET("/api/v1/products/stock", _ProductService_GetProductsStock0_HTTP_Handler(srv))
}

func _ProductService_CreateProduct0_HTTP_Handler(srv ProductServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateProductRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProductServiceCreateProduct)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateProduct(ctx, req.(*CreateProductRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateProductResponse)
		return ctx.Result(200, reply)
	}
}

func _ProductService_GetProducts0_HTTP_Handler(srv ProductServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetProductsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProductServiceGetProducts)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProducts(ctx, req.(*GetProductsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetProductsResponse)
		return ctx.Result(200, reply)
	}
}

func _ProductService_GetProduct0_HTTP_Handler(srv ProductServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetProductRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProductServiceGetProduct)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProduct(ctx, req.(*GetProductRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Product)
		return ctx.Result(200, reply)
	}
}

func _ProductService_UpdateProduct0_HTTP_Handler(srv ProductServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateProductRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProductServiceUpdateProduct)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateProduct(ctx, req.(*UpdateProductRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateProductResponse)
		return ctx.Result(200, reply)
	}
}

func _ProductService_DeleteProduct0_HTTP_Handler(srv ProductServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteProductRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProductServiceDeleteProduct)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteProduct(ctx, req.(*DeleteProductRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteProductResponse)
		return ctx.Result(200, reply)
	}
}

func _ProductService_GetProductStock0_HTTP_Handler(srv ProductServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetProductStockRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProductServiceGetProductStock)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProductStock(ctx, req.(*GetProductStockRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetProductStockResponse)
		return ctx.Result(200, reply)
	}
}

func _ProductService_UpdateProductStock0_HTTP_Handler(srv ProductServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateProductStockRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProductServiceUpdateProductStock)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateProductStock(ctx, req.(*UpdateProductStockRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateProductStockResponse)
		return ctx.Result(200, reply)
	}
}

func _ProductService_GetProductsStock0_HTTP_Handler(srv ProductServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetProductsStockRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProductServiceGetProductsStock)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProductsStock(ctx, req.(*GetProductsStockRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetProductsStockResponse)
		return ctx.Result(200, reply)
	}
}

type ProductServiceHTTPClient interface {
	CreateProduct(ctx context.Context, req *CreateProductRequest, opts ...http.CallOption) (rsp *CreateProductResponse, err error)
	DeleteProduct(ctx context.Context, req *DeleteProductRequest, opts ...http.CallOption) (rsp *DeleteProductResponse, err error)
	GetProduct(ctx context.Context, req *GetProductRequest, opts ...http.CallOption) (rsp *Product, err error)
	GetProductStock(ctx context.Context, req *GetProductStockRequest, opts ...http.CallOption) (rsp *GetProductStockResponse, err error)
	GetProducts(ctx context.Context, req *GetProductsRequest, opts ...http.CallOption) (rsp *GetProductsResponse, err error)
	GetProductsStock(ctx context.Context, req *GetProductsStockRequest, opts ...http.CallOption) (rsp *GetProductsStockResponse, err error)
	UpdateProduct(ctx context.Context, req *UpdateProductRequest, opts ...http.CallOption) (rsp *UpdateProductResponse, err error)
	UpdateProductStock(ctx context.Context, req *UpdateProductStockRequest, opts ...http.CallOption) (rsp *UpdateProductStockResponse, err error)
}

type ProductServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewProductServiceHTTPClient(client *http.Client) ProductServiceHTTPClient {
	return &ProductServiceHTTPClientImpl{client}
}

func (c *ProductServiceHTTPClientImpl) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...http.CallOption) (*CreateProductResponse, error) {
	var out CreateProductResponse
	pattern := "/api/v1/products"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationProductServiceCreateProduct))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ProductServiceHTTPClientImpl) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...http.CallOption) (*DeleteProductResponse, error) {
	var out DeleteProductResponse
	pattern := "/api/v1/products/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationProductServiceDeleteProduct))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ProductServiceHTTPClientImpl) GetProduct(ctx context.Context, in *GetProductRequest, opts ...http.CallOption) (*Product, error) {
	var out Product
	pattern := "/api/v1/products/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationProductServiceGetProduct))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ProductServiceHTTPClientImpl) GetProductStock(ctx context.Context, in *GetProductStockRequest, opts ...http.CallOption) (*GetProductStockResponse, error) {
	var out GetProductStockResponse
	pattern := "/api/v1/products/{id}/stock"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationProductServiceGetProductStock))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ProductServiceHTTPClientImpl) GetProducts(ctx context.Context, in *GetProductsRequest, opts ...http.CallOption) (*GetProductsResponse, error) {
	var out GetProductsResponse
	pattern := "/api/v1/products"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationProductServiceGetProducts))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ProductServiceHTTPClientImpl) GetProductsStock(ctx context.Context, in *GetProductsStockRequest, opts ...http.CallOption) (*GetProductsStockResponse, error) {
	var out GetProductsStockResponse
	pattern := "/api/v1/products/stock"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationProductServiceGetProductsStock))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ProductServiceHTTPClientImpl) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...http.CallOption) (*UpdateProductResponse, error) {
	var out UpdateProductResponse
	pattern := "/api/v1/products/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationProductServiceUpdateProduct))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ProductServiceHTTPClientImpl) UpdateProductStock(ctx context.Context, in *UpdateProductStockRequest, opts ...http.CallOption) (*UpdateProductStockResponse, error) {
	var out UpdateProductStockResponse
	pattern := "/api/v1/products/{id}/stock"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationProductServiceUpdateProductStock))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
