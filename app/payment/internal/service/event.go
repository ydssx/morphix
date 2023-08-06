package service

import (
	"context"

	"github.com/ydssx/morphix/common/dapr"
	"github.com/ydssx/morphix/common/event"
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
	return dapr.PublishEvent(ctx, event.Subject_name[int32(event.Subject_PaymentCompleted)], payload)
}
