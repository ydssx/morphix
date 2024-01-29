package service

import (
	"context"

	paymentv1 "github.com/ydssx/morphix/api/payment/v1"
	"github.com/ydssx/morphix/common/event"
)

var _ paymentv1.PaymentServiceServer = (*PaymentService)(nil)

type PaymentService struct {
	paymentv1.UnimplementedPaymentServiceServer

	eventSink PaymentEventSinker
}

func NewPaymentService(eventSink PaymentEventSinker) *PaymentService {
	return &PaymentService{eventSink: eventSink}
}

// CancelPayment cancels an existing payment by order ID.
// It publishes a cancel payment event and returns a response with a completed status.
func (p *PaymentService) CancelPayment(ctx context.Context, req *paymentv1.CancelPaymentRequest) (*paymentv1.CancelPaymentResponse, error) {
	payload := event.PayloadCancelPayment{OrderId: req.OrderId}
	err := p.eventSink.OnCancelPayment(ctx, &payload)
	if err != nil {
		return nil, err
	}

	return &paymentv1.CancelPaymentResponse{Status: "COMPLETED"}, nil
}

// GetPayment implements paymentv1.PaymentServiceServer.
func (*PaymentService) GetPayment(context.Context, *paymentv1.GetPaymentRequest) (*paymentv1.GetPaymentResponse, error) {
	panic("unimplemented")
}

// MakePayment implements paymentv1.PaymentServiceServer.
func (p *PaymentService) MakePayment(ctx context.Context, req *paymentv1.MakePaymentRequest) (*paymentv1.PaymentResponse, error) {
	payload := event.PayloadPaymentCompleted{
		UserId:  1,
		Amount:  float32(req.Amount),
		// OrderId: "",
	}
	err := p.eventSink.OnMakePayment(ctx, &payload)
	if err != nil {
		return nil, err
	}

	return &paymentv1.PaymentResponse{OrderId: req.OrderId, Status: "COMPLETED"}, nil
}

// Refund implements paymentv1.PaymentServiceServer.
func (*PaymentService) Refund(context.Context, *paymentv1.RefundRequest) (*paymentv1.RefundResponse, error) {
	panic("unimplemented")
}
