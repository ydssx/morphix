package service

import (
	"context"

	"github.com/ydssx/morphix/common/dapr"
	"github.com/ydssx/morphix/common/event"
	"github.com/ydssx/morphix/pkg/mq"
)

type PaymentEvents interface {
	OnMakePayment(ctx context.Context, payload *event.PayloadPaymentCompleted) error
	OnCancelPayment(ctx context.Context, payload *event.PayloadCancelPayment) error
}

var _ PaymentEvents = (*eventSender)(nil)

type eventSender struct {
	*dapr.DaprClient
	ce *mq.CloudEvent
}

func NewEventSender(daprClient *dapr.DaprClient, ce *mq.CloudEvent) PaymentEvents {
	return &eventSender{DaprClient: daprClient}
}

// OnMakePayment implements PaymentEvents.
func (e *eventSender) OnMakePayment(ctx context.Context, payload *event.PayloadPaymentCompleted) error {
	return e.PublishEvent(ctx, event.Subject_name[int32(event.Subject_PaymentCompleted)], payload)
}

// OnCancelPayment implements PaymentEvents.
func (e *eventSender) OnCancelPayment(ctx context.Context, payload *event.PayloadCancelPayment) error {
	return e.ce.PublishEvent(ctx, event.Subject_name[int32(event.Subject_CancelPayment)], payload)
}
