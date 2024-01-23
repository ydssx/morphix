package biz

import (
	"context"

	productv1 "github.com/ydssx/morphix/api/product/v1"
	"github.com/ydssx/morphix/app/product/internal/model"
	"github.com/ydssx/morphix/pkg/errors"
	"github.com/ydssx/morphix/pkg/interceptors"
)

type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
}

type ProductRepo interface {
	CreateProduct(ctx context.Context, product *model.Product) error
	GetProduct(ctx context.Context, id int64) (*model.Product, error)
	GetProducts(ctx context.Context, cond *ListProductsCond) ([]*model.Product, int64, error)
	UpdateProduct(ctx context.Context, product *model.Product) error
	DeleteProduct(ctx context.Context, id int64) error
}

type ListProductsCond struct {
	PageNum    int64
	PageSize   int64
	Status     string
	UserID     int64
	ProductIds []int64
	Name       string
}

type ProductUseCase struct {
	repo ProductRepo
}

func NewProductUseCase(repo ProductRepo) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

// 创建产品
func (uc *ProductUseCase) CreateProduct(ctx context.Context, req *productv1.CreateProductRequest) (res *productv1.CreateProductResponse, err error) {
	res = new(productv1.CreateProductResponse)

	// TODO:ADD logic here and delete this line.

	return
}

// 获取产品列表
func (uc *ProductUseCase) GetProducts(ctx context.Context, req *productv1.GetProductsRequest) (res *productv1.GetProductsResponse, err error) {
	res = new(productv1.GetProductsResponse)

	claim, _ := interceptors.AuthFromContext(ctx)

	products, total, err := uc.repo.GetProducts(ctx, &ListProductsCond{
		PageNum:  req.Page,
		PageSize: req.PageSize,
		// Status:     req.Status,
		UserID:     claim.Uid,
		ProductIds: req.ProductIds,
		Name:       req.Name,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get products")
	}
	
	res.Products = make([]*productv1.Product, 0)
	for _, product := range products {
		res.Products = append(res.Products, &productv1.Product{
			Id:          int64(product.ID),
			Name:        product.Name,
			Description: product.Description,
			Price:       float32(product.Price),
		})
	}
	res.Total = total

	return
}

// 获取单个产品
func (uc *ProductUseCase) GetProduct(ctx context.Context, req *productv1.GetProductRequest) (res *productv1.Product, err error) {
	res = new(productv1.Product)

	// TODO:ADD logic here and delete this line.

	return
}

// 更新产品信息
func (uc *ProductUseCase) UpdateProduct(ctx context.Context, req *productv1.UpdateProductRequest) (res *productv1.UpdateProductResponse, err error) {
	res = new(productv1.UpdateProductResponse)

	// TODO:ADD logic here and delete this line.

	return
}

// 删除产品
func (uc *ProductUseCase) DeleteProduct(ctx context.Context, req *productv1.DeleteProductRequest) (res *productv1.DeleteProductResponse, err error) {
	res = new(productv1.DeleteProductResponse)

	// TODO:ADD logic here and delete this line.

	return
}
