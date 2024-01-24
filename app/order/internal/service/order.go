package service

import (
	"context"

	orderv1 "github.com/ydssx/morphix/api/order/v1"
	"github.com/ydssx/morphix/app/order/internal/biz"
)

type OrderService struct {
	uc *biz.OrderUseCase

	orderv1.UnimplementedOrderServiceServer
}

func NewOrderService(uc *biz.OrderUseCase) *OrderService {
	return &OrderService{uc: uc}
}

// 创建订单
func (s *OrderService) CreateOrder(ctx context.Context, req *orderv1.CreateOrderRequest) (res *orderv1.CreateOrderResponse, err error) {
	return s.uc.CreateOrder(ctx, req)
}

// 查询订单
func (s *OrderService) GetOrder(ctx context.Context, req *orderv1.GetOrderRequest) (res *orderv1.GetOrderResponse, err error) {
	return s.uc.GetOrder(ctx, req)
}

// 更新订单状态
func (s *OrderService) UpdateOrderStatus(ctx context.Context, req *orderv1.UpdateOrderStatusRequest) (res *orderv1.UpdateOrderStatusResponse, err error) {
	return s.uc.UpdateOrderStatus(ctx, req)
}

// 删除订单
func (s *OrderService) DeleteOrder(ctx context.Context, req *orderv1.DeleteOrderRequest) (res *orderv1.DeleteOrderResponse, err error) {
	return s.uc.DeleteOrder(ctx, req)
}

// 查询订单列表
func (s *OrderService) ListOrders(ctx context.Context, req *orderv1.ListOrdersRequest) (res *orderv1.ListOrdersResponse, err error) {
	return s.uc.ListOrders(ctx, req)
}

// 支付订单
func (s *OrderService) PayOrder(ctx context.Context, req *orderv1.PayOrderRequest) (res *orderv1.PayOrderResponse, err error) {
	return s.uc.PayOrder(ctx, req)
}
