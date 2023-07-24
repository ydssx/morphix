package service

import (
	"context"

	paymentv1 "github.com/ydssx/morphix/api/payment/v1"
	"github.com/ydssx/morphix/common/event"
	"github.com/ydssx/morphix/pkg/mq"
)

var _ paymentv1.PaymentServiceServer = (*PaymentService)(nil)

type PaymentService struct {
	paymentv1.UnimplementedPaymentServiceServer
}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

// CancelPayment implements paymentv1.PaymentServiceServer.
func (*PaymentService) CancelPayment(context.Context, *paymentv1.CancelPaymentRequest) (*paymentv1.CancelPaymentResponse, error) {
	panic("unimplemented")
}

// GetPayment implements paymentv1.PaymentServiceServer.
func (*PaymentService) GetPayment(context.Context, *paymentv1.GetPaymentRequest) (*paymentv1.GetPaymentResponse, error) {
	panic("unimplemented")
}

// MakePayment implements paymentv1.PaymentServiceServer.
func (*PaymentService) MakePayment(ctx context.Context, req *paymentv1.MakePaymentRequest) (*paymentv1.PaymentResponse, error) {
	payload := event.PayloadUserCharge{
		UserId:  1,
		Amount:  float32(req.Amount),
		OrderId: req.OrderId,
	}
	err := mq.Send(ctx, string(event.Subject_name[int32(event.Subject_PaymentProcessed)]), &payload)
	if err != nil {
		return nil, err
	}
	return &paymentv1.PaymentResponse{}, nil
}

// Refund implements paymentv1.PaymentServiceServer.
func (*PaymentService) Refund(context.Context, *paymentv1.RefundRequest) (*paymentv1.RefundResponse, error) {
	panic("unimplemented")
}
