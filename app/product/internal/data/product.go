package data

import (
	"context"

	"github.com/ydssx/morphix/app/product/internal/biz"
	"github.com/ydssx/morphix/app/product/internal/model"
)

type ProductRepo struct {
	data *Data
}

// GetProducts implements biz.ProductRepo.
func (*ProductRepo) GetProducts(ctx context.Context, cond *biz.ListProductsCond) ([]*model.Product, int64, error) {
	panic("unimplemented")
}

// CreateProduct implements biz.ProductRepo.
func (r *ProductRepo) CreateProduct(ctx context.Context, product *model.Product) error {
	return model.NewProductModel(r.data.DB(ctx)).Create(*product)
}

// DeleteProduct implements biz.ProductRepo.
func (*ProductRepo) DeleteProduct(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// GetProduct implements biz.ProductRepo.
func (*ProductRepo) GetProduct(ctx context.Context, id int64) (*model.Product, error) {
	panic("unimplemented")
}

// UpdateProduct implements biz.ProductRepo.
func (*ProductRepo) UpdateProduct(ctx context.Context, product *model.Product) error {
	panic("unimplemented")
}

func NewProductRepo(data *Data) biz.ProductRepo {
	return &ProductRepo{data: data}
}
