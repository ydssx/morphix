package biz

import (
	"context"

	orderv1 "github.com/ydssx/morphix/api/order/v1"
	paymentv1 "github.com/ydssx/morphix/api/payment/v1"
	productv1 "github.com/ydssx/morphix/api/product/v1"
	"github.com/ydssx/morphix/app/order/internal/model"
	"github.com/ydssx/morphix/pkg/errors"
	"github.com/ydssx/morphix/pkg/interceptors"
	"github.com/ydssx/morphix/pkg/util"
)

type Transaction interface {
	InTx(ctx context.Context, f func(ctx context.Context) error) error
}

type OrderRepo interface {
	CreateOrder(ctx context.Context, order model.Order) (orderID int64, err error)
	GetOrder(ctx context.Context, orderID int64) (order *model.Order, err error)
	UpdateOrderStatus(ctx context.Context, orderID int64, status int32) (err error)
	DeleteOrder(ctx context.Context, orderID int64) (err error)
	DeleteOrders(ctx context.Context, orderIDs []int64) (err error)
	ListOrders(ctx context.Context, cond *ListOrderCond) (orders []*model.Order, total int64, err error)

	// 创建订单项
	CreateOrderItem(ctx context.Context, orderItem ...model.OrderItem) (err error)
}

type ListOrderCond struct {
	PageNum  int64
	PageSize int64
	Status   string
	UserID   int64
}

type OrderUseCase struct {
	repo          OrderRepo
	tx            Transaction
	productClient productv1.ProductServiceClient
	paymentClient paymentv1.PaymentServiceClient
}

func NewOrderUseCase(tx Transaction, repo OrderRepo, productClient productv1.ProductServiceClient, paymentClient paymentv1.PaymentServiceClient) *OrderUseCase {
	return &OrderUseCase{repo: repo, tx: tx, productClient: productClient, paymentClient: paymentClient}
}

// 创建订单
func (b *OrderUseCase) CreateOrder(ctx context.Context, req *orderv1.CreateOrderRequest) (res *orderv1.CreateOrderResponse, err error) {
	res = new(orderv1.CreateOrderResponse)
	claim, _ := interceptors.AuthFromContext(ctx)

	var productIds []int64
	for _, item := range req.Items {
		productIds = append(productIds, int64(item.ProductId))
	}

	// 查询商品价格
	productResp, err := b.productClient.GetProducts(ctx, &productv1.GetProductsRequest{
		ProductIds: productIds,
	})
	if err != nil {
		return nil, errors.Wrap(err, "查询商品价格失败")
	}
	if len(productResp.Products) != len(req.Items) {
		return nil, errors.New("商品价格数量与请求数量不一致")
	}
	productPriceMap := make(map[int64]float64)
	for _, product := range productResp.Products {
		productPriceMap[product.Id] = float64(product.Price)
	}

	// 计算订单总价
	var totalPrice float64
	for _, item := range req.Items {
		totalPrice += productPriceMap[int64(item.ProductId)] * float64(item.Quantity)
	}

	err = b.tx.InTx(ctx, func(ctx context.Context) error {
		orderNum, err := util.GenerateOrderNumber()
		if err != nil {
			return errors.Wrap(err, "生成订单号失败")
		}
		order := model.Order{OrderNumber: orderNum, UserId: int(claim.Uid), Amount: totalPrice}
		orderID, err := b.repo.CreateOrder(ctx, order)
		if err != nil {
			return err
		}

		// 创建订单项
		orderItems := make([]model.OrderItem, 0)
		for _, item := range req.Items {
			orderItem := model.OrderItem{
				OrderId:   int(orderID),
				ProductId: int(item.ProductId),
				Quantity:  int(item.Quantity),
				Price:     productPriceMap[int64(item.ProductId)],
			}
			orderItems = append(orderItems, orderItem)
		}
		err = b.repo.CreateOrderItem(ctx, orderItems...)
		if err != nil {
			return errors.Wrap(err, "创建订单项失败")
		}

		return nil
	})
	return
}

// 查询订单
func (b *OrderUseCase) GetOrder(ctx context.Context, req *orderv1.GetOrderRequest) (res *orderv1.GetOrderResponse, err error) {
	res = new(orderv1.GetOrderResponse)
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

	return
}

// 查询订单列表
func (b *OrderUseCase) ListOrders(ctx context.Context, req *orderv1.ListOrdersRequest) (res *orderv1.ListOrdersResponse, err error) {
	res = new(orderv1.ListOrdersResponse)

	return
}
