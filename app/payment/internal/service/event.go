package service

import (
	"context"

	"github.com/ydssx/morphix/common/event"
	"github.com/ydssx/morphix/pkg/mq"
)

type PaymentEvents interface {
	OnMakePayment(ctx context.Context, payload *event.PayloadPaymentCompleted) error
}

var _ PaymentEvents = (*eventSender)(nil)

type eventSender struct {
}

func NewEventSender() PaymentEvents {
	return &eventSender{}
}

// OnMakePayment implements PaymentEvents.
func (*eventSender) OnMakePayment(ctx context.Context, payload *event.PayloadPaymentCompleted) error {
	return mq.Send(ctx, event.Subject_name[int32(event.Subject_PaymentCompleted)], payload)
}
