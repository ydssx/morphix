package data

import (
	"context"

	"github.com/ydssx/morphix/app/order/internal/biz"
	"github.com/ydssx/morphix/app/order/internal/model"
)

type OrderRepo struct {
	data *Data
}

// CreateOrder implements biz.OrderRepo.
func (*OrderRepo) CreateOrder(ctx context.Context, order model.Order) (orderID int64, err error) {
	panic("unimplemented")
}

// CreateOrderItem implements biz.OrderRepo.
func (*OrderRepo) CreateOrderItem(ctx context.Context, orderItem ...model.OrderItem) (err error) {
	panic("unimplemented")
}

// DeleteOrder implements biz.OrderRepo.
func (*OrderRepo) DeleteOrder(ctx context.Context, orderID int64) (err error) {
	panic("unimplemented")
}

// DeleteOrders implements biz.OrderRepo.
func (*OrderRepo) DeleteOrders(ctx context.Context, orderIDs []int64) (err error) {
	panic("unimplemented")
}

// GetOrder implements biz.OrderRepo.
func (*OrderRepo) GetOrder(ctx context.Context, orderID int64) (order *model.Order, err error) {
	panic("unimplemented")
}

// ListOrders implements biz.OrderRepo.
func (*OrderRepo) ListOrders(ctx context.Context, cond *biz.ListOrderCond) (orders []*model.Order, total int64, err error) {
	panic("unimplemented")
}

// UpdateOrderStatus implements biz.OrderRepo.
func (*OrderRepo) UpdateOrderStatus(ctx context.Context, orderID int64, status int32) (err error) {
	panic("unimplemented")
}

func NewOrderRepo(data *Data) biz.OrderRepo {
	return &OrderRepo{data: data}
}
