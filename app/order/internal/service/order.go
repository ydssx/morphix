package service

import (
	"context"

	orderv1 "github.com/ydssx/morphix/api/order/v1"
)

var _ orderv1.OrderServiceServer = (*OrderService)(nil)

type OrderService struct {
	orderv1.UnimplementedOrderServiceServer
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

// CreateOrder implements orderv1.OrderServiceServer.
func (*OrderService) CreateOrder(context.Context, *orderv1.CreateOrderRequest) (*orderv1.CreateOrderResponse, error) {
	panic("unimplemented")
}

// DeleteOrder implements orderv1.OrderServiceServer.
func (*OrderService) DeleteOrder(context.Context, *orderv1.DeleteOrderRequest) (*orderv1.DeleteOrderResponse, error) {
	panic("unimplemented")
}

// GetOrder implements orderv1.OrderServiceServer.
func (*OrderService) GetOrder(context.Context, *orderv1.GetOrderRequest) (*orderv1.GetOrderResponse, error) {
	panic("unimplemented")
}

// ListOrders implements orderv1.OrderServiceServer.
func (*OrderService) ListOrders(context.Context, *orderv1.ListOrdersRequest) (*orderv1.ListOrdersResponse, error) {
	panic("unimplemented")
}

// UpdateOrderStatus implements orderv1.OrderServiceServer.
func (*OrderService) UpdateOrderStatus(context.Context, *orderv1.UpdateOrderStatusRequest) (*orderv1.UpdateOrderStatusResponse, error) {
	panic("unimplemented")
}
