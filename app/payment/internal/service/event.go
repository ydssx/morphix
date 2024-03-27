package service

import (
	"context"

	"github.com/ydssx/morphix/common/dapr"
	"github.com/ydssx/morphix/common/event"
	"github.com/ydssx/morphix/pkg/pubsub"
)

type PaymentEventSinker interface {
	OnMakePayment(ctx context.Context, payload *event.PayloadPaymentCompleted) error
	OnCancelPayment(ctx context.Context, payload *event.PayloadCancelPayment) error
}

var _ PaymentEventSinker = (*eventSender)(nil)

type eventSender struct {
	*dapr.DaprClient
	ce *pubsub.CloudEvent
}

func NewEventSender(daprClient *dapr.DaprClient, ce *pubsub.CloudEvent) PaymentEventSinker {
	return &eventSender{DaprClient: daprClient, ce: ce}
}

// OnMakePayment implements PaymentEvents.
func (e *eventSender) OnMakePayment(ctx context.Context, payload *event.PayloadPaymentCompleted) error {
	return e.PublishEvent(ctx, event.Subject_PaymentCompleted.String(), payload)
}

// OnCancelPayment implements PaymentEvents.
func (e *eventSender) OnCancelPayment(ctx context.Context, payload *event.PayloadCancelPayment) error {
	return e.ce.PublishMessage(ctx, event.Subject_CancelPayment.String(), payload)
}
