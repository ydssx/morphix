package biz

import (
	"context"

	orderv1 "github.com/ydssx/morphix/api/order/v1"
)

type OrderUseCase struct{}

func NewOrderUseCase() *OrderUseCase {
	return &OrderUseCase{}
}

// 创建订单
func (b *OrderUseCase) CreateOrder(ctx context.Context, req *orderv1.CreateOrderRequest) (res *orderv1.CreateOrderResponse, err error) {
	res = new(orderv1.CreateOrderResponse)

	// TODO:ADD logic here and delete this line.

	return
}

// 查询订单
func (b *OrderUseCase) GetOrder(ctx context.Context, req *orderv1.GetOrderRequest) (res *orderv1.GetOrderResponse, err error) {
	res = new(orderv1.GetOrderResponse)

	// TODO:ADD logic here and delete this line.

	return
}

// 更新订单状态
func (b *OrderUseCase) UpdateOrderStatus(ctx context.Context, req *orderv1.UpdateOrderStatusRequest) (res *orderv1.UpdateOrderStatusResponse, err error) {
	res = new(orderv1.UpdateOrderStatusResponse)

	// TODO:ADD logic here and delete this line.

	return
}

// 删除订单
func (b *OrderUseCase) DeleteOrder(ctx context.Context, req *orderv1.DeleteOrderRequest) (res *orderv1.DeleteOrderResponse, err error) {
	res = new(orderv1.DeleteOrderResponse)

	// TODO:ADD logic here and delete this line.

	return
}

// 查询订单列表
func (b *OrderUseCase) ListOrders(ctx context.Context, req *orderv1.ListOrdersRequest) (res *orderv1.ListOrdersResponse, err error) {
	res = new(orderv1.ListOrdersResponse)

	// TODO:ADD logic here and delete this line.

	return
}
