package service

import (
	"context"

	productv1 "github.com/ydssx/morphix/api/product/v1"
	"github.com/ydssx/morphix/app/product/internal/biz"
)

var _ = context.Background

type ProductService struct {
	uc *biz.ProductUseCase

	productv1.UnimplementedProductServiceServer
}

func NewProductService(uc *biz.ProductUseCase) *ProductService {
	return &ProductService{uc: uc}
}

// 创建产品
func (s *ProductService) CreateProduct(ctx context.Context, req *productv1.CreateProductRequest) (res *productv1.CreateProductResponse, err error) {
	return s.uc.CreateProduct(ctx, req)
}

// 获取产品列表
func (s *ProductService) GetProducts(ctx context.Context, req *productv1.GetProductsRequest) (res *productv1.GetProductsResponse, err error) {
	return s.uc.GetProducts(ctx, req)
}

// 获取单个产品
func (s *ProductService) GetProduct(ctx context.Context, req *productv1.GetProductRequest) (res *productv1.Product, err error) {
	return s.uc.GetProduct(ctx, req)
}

// 更新产品信息
func (s *ProductService) UpdateProduct(ctx context.Context, req *productv1.UpdateProductRequest) (res *productv1.UpdateProductResponse, err error) {
	return s.uc.UpdateProduct(ctx, req)
}

// 删除产品
func (s *ProductService) DeleteProduct(ctx context.Context, req *productv1.DeleteProductRequest) (res *productv1.DeleteProductResponse, err error) {
	return s.uc.DeleteProduct(ctx, req)
}
