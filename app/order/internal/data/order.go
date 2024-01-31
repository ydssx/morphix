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
func (o *OrderRepo) CreateOrder(ctx context.Context, order model.Order) (orderID int64, err error) {
	orderID, err = model.NewOrderModel(o.data.DB(ctx)).Create(order)
	return
}

// CreateOrderItem implements biz.OrderRepo.
func (o *OrderRepo) CreateOrderItem(ctx context.Context, orderItem ...model.OrderItem) (err error) {
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
func (*OrderRepo) GetOrder(ctx context.Context, orderNumber string) (order *model.Order, err error) {
	panic("unimplemented")
}

// ListOrders implements biz.OrderRepo.
func (r *OrderRepo) ListOrders(ctx context.Context, cond *biz.ListOrderCond) (orders []*model.Order, total int64, err error) {
	return
}

// UpdateOrderStatus implements biz.OrderRepo.
func (r *OrderRepo) UpdateOrderStatus(ctx context.Context, orderNumber string, status string) (err error) {
	err = model.NewOrderModel(r.data.DB(ctx)).SetOrderNumber(orderNumber).Updates(model.Order{Status: status})
	return
}

func NewOrderRepo(data *Data) biz.OrderRepo {
	return &OrderRepo{data: data}
}
