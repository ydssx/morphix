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
	*dapr.DaprClient
}

func NewEventSender(daprClient *dapr.DaprClient) PaymentEvents {
	return &eventSender{daprClient}
}

// OnMakePayment implements PaymentEvents.
func (e *eventSender) OnMakePayment(ctx context.Context, payload *event.PayloadPaymentCompleted) error {
	return e.PublishEvent(ctx, event.Subject_name[int32(event.Subject_PaymentCompleted)], payload)
}
