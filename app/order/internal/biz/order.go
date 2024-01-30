package biz

import (
	"context"
	"fmt"
	"time"

	jobv1 "github.com/ydssx/morphix/api/job/v1"
	orderv1 "github.com/ydssx/morphix/api/order/v1"
	paymentv1 "github.com/ydssx/morphix/api/payment/v1"
	productv1 "github.com/ydssx/morphix/api/product/v1"
	quotev1 "github.com/ydssx/morphix/api/quote/v1"
	"github.com/ydssx/morphix/app/order/internal/model"
	"github.com/ydssx/morphix/pkg/errors"
	"github.com/ydssx/morphix/pkg/interceptors"
	"github.com/ydssx/morphix/pkg/redis"
	"github.com/ydssx/morphix/pkg/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Transaction interface {
	InTx(ctx context.Context, f func(ctx context.Context) error) error
}

type OrderRepo interface {
	CreateOrder(ctx context.Context, order model.Order) (orderID int64, err error)
	GetOrder(ctx context.Context, orderID int64) (order *model.Order, err error)
	UpdateOrderStatus(ctx context.Context, orderID int64, status string) (err error)
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
	locker        redis.Locker
	productClient productv1.ProductServiceClient
	paymentClient paymentv1.PaymentServiceClient
	quoteClient   quotev1.QuoteServiceClient
	jobClient     jobv1.JobServiceClient
}

func NewOrderUseCase(
	tx Transaction,
	repo OrderRepo,
	locker redis.Locker,
	productClient productv1.ProductServiceClient,
	paymentClient paymentv1.PaymentServiceClient,
	quoteClient quotev1.QuoteServiceClient,
	jobClient jobv1.JobServiceClient,
) *OrderUseCase {
	return &OrderUseCase{
		repo:          repo,
		tx:            tx,
		locker:        locker,
		productClient: productClient,
		paymentClient: paymentClient,
		quoteClient:   quoteClient,
		jobClient:     jobClient,
	}
}

// 创建订单
func (b *OrderUseCase) CreateOrder(ctx context.Context, req *orderv1.CreateOrderRequest) (res *orderv1.CreateOrderResponse, err error) {
	res = new(orderv1.CreateOrderResponse)
	claim, _ := interceptors.AuthFromContext(ctx)

	var productIds []int64
	for _, item := range req.Items {
		productIds = append(productIds, int64(item.ProductId))
	}

	// 查询商品库存
	productStockResp, err := b.productClient.GetProductsStock(ctx, &productv1.GetProductsStockRequest{Ids: productIds})
	if err != nil {
		return nil, errors.Wrap(err, "查询商品库存失败")
	}
	for _, item := range req.Items {
		if productStockResp.Stocks[int64(item.ProductId)] < int32(item.Quantity) {
			return nil, errors.Errorf("商品库存不足 [%d]", item.ProductId)
		}
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

	// 查询报价
	quoteResp, err := b.quoteClient.GetQuotes(ctx, &quotev1.GetQuotesRequest{
		ProductIds: productIds,
	})
	if err != nil {
		return nil, errors.Wrap(err, "查询报价失败")
	}
	quoteMap := make(map[int64]float64)
	for _, quote := range quoteResp.Quotes {
		quoteMap[quote.ProductId] = float64(quote.FinalPrice)
	}

	productPriceMap := make(map[int64]float64)
	for _, product := range productResp.Products {
		productPriceMap[product.Id] = float64(product.Price)
		if _, ok := quoteMap[product.Id]; !ok {
			return nil, errors.New("报价不存在")
		}
		productPriceMap[product.Id] = quoteMap[product.Id]
	}

	// 计算订单总价
	var totalPrice float64
	for _, item := range req.Items {
		totalPrice += productPriceMap[int64(item.ProductId)] * float64(item.Quantity)
	}

	orderNum, err := util.GenerateOrderNumber()
	if err != nil {
		return nil, errors.Wrap(err, "生成订单号失败")
	}
	err = b.tx.InTx(ctx, func(ctx context.Context) error {
		order := model.Order{OrderNumber: orderNum, UserId: int(claim.Uid), Amount: totalPrice}
		orderID, err := b.repo.CreateOrder(ctx, order)
		if err != nil {
			return errors.Wrap(err, "创建订单失败")
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
	if err != nil {
		return nil, err
	}

	// 订单超时自动取消
	_, err = b.jobClient.Enqueue(ctx, &jobv1.EnqueueRequest{
		JobType:   jobv1.JobType_ORDER_TIMEOUT,
		Payload:   []byte(orderNum),
		ProcessAt: timestamppb.New(time.Now().Add(time.Minute * 10)),
	})
	if err != nil {
		return nil, errors.Wrap(err, "创建订单超时任务失败")
	}

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

	lockKey := fmt.Sprintf("order:%d", req.OrderId)
	err = b.locker.Lock(ctx, lockKey, redis.WithTTL(time.Second*10))
	if err != nil {
		return nil, errors.Wrap(err, "获取订单锁失败")
	}
	defer b.locker.Unlock(ctx, lockKey)

	order, err := b.repo.GetOrder(ctx, int64(req.OrderId))
	if err != nil {
		return nil, errors.Wrap(err, "查询订单失败")
	}
	// 检查状态转换是否有效
	if !isStatusTransitionValid(orderv1.OrderStatus(orderv1.OrderStatus_value[order.Status]), req.Status) {
		return nil, errors.New("订单状态转换不合法")
	}

	err = b.repo.UpdateOrderStatus(ctx, int64(req.OrderId), req.Status.String())
	if err != nil {
		return nil, errors.Wrap(err, "更新订单状态失败")
	}

	return
}

// isStatusTransitionValid 检查订单状态转换是否有效
func isStatusTransitionValid(currentStatus orderv1.OrderStatus, newStatus orderv1.OrderStatus) bool {
	// 定义状态转换规则，例如只允许从较低的状态转换到较高的状态
	// 这里的实现应该根据业务逻辑来定义
	// 例如，如果订单状态有一个明确的生命周期，可以按照生命周期的顺序来进行比较
	// 如果状态转换更复杂，可能需要一个状态转换图来表示允许的转换
	return orderv1.OrderStatus_value[currentStatus.String()] < orderv1.OrderStatus_value[newStatus.String()]
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

// 支付订单
func (uc *OrderUseCase) PayOrder(ctx context.Context, req *orderv1.PayOrderRequest) (res *orderv1.PayOrderResponse, err error) {
	res = new(orderv1.PayOrderResponse)

	order, err := uc.repo.GetOrder(ctx, int64(req.OrderId))
	if err != nil {
		return nil, errors.Wrap(err, "查询订单失败")
	}
	if order.Status != orderv1.OrderStatus_PENDING.String() {
		return nil, errors.New("订单状态不正确")
	}

	// 调用支付系统接口，支付订单
	paymentResp, err := uc.paymentClient.MakePayment(ctx, &paymentv1.MakePaymentRequest{
		OrderId: int64(req.OrderId),
		Amount:  order.Amount,
	})
	if err != nil {
		return nil, errors.Wrap(err, "调用支付系统接口，支付订单失败")
	}
	res.PaymentUrl = paymentResp.PaymentUrl

	return
}

// 取消订单
func (uc *OrderUseCase) CancelOrder(ctx context.Context, req *orderv1.CancelOrderRequest) (res *orderv1.CancelOrderResponse, err error) {
	res = new(orderv1.CancelOrderResponse)

	// TODO:ADD logic here and delete this line.

	return
}
