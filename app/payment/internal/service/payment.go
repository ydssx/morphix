package service

import (
	"context"

	paymentv1 "github.com/ydssx/morphix/api/payment/v1"
	"github.com/ydssx/morphix/constants"
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
	err := mq.Send(ctx, string(constants.TopicUserCharge), mq.NewEvent(nil))
	if err != nil {
		return nil, err
	}
	return &paymentv1.PaymentResponse{}, nil
}

// Refund implements paymentv1.PaymentServiceServer.
func (*PaymentService) Refund(context.Context, *paymentv1.RefundRequest) (*paymentv1.RefundResponse, error) {
	panic("unimplemented")
}
