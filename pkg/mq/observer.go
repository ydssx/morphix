package mq

import (
	"context"

	"github.com/cloudevents/sdk-go/v2/binding"
	"github.com/cloudevents/sdk-go/v2/client"
	"github.com/cloudevents/sdk-go/v2/event"
)

var _ client.ObservabilityService = (*Observer)(nil)

type Observer struct{}

func NewObserver() *Observer {
	return &Observer{}
}

// InboundContextDecorators implements client.ObservabilityService.
func (*Observer) InboundContextDecorators() []func(context.Context, binding.Message) context.Context {
	return nil
}

// RecordCallingInvoker implements client.ObservabilityService.
func (*Observer) RecordCallingInvoker(ctx context.Context, event *event.Event) (context.Context, func(errOrResult error)) {
	return ctx, func(errOrResult error) {}
}

// RecordReceivedMalformedEvent implements client.ObservabilityService.
func (*Observer) RecordReceivedMalformedEvent(ctx context.Context, err error) {
}

// RecordRequestEvent implements client.ObservabilityService.
func (*Observer) RecordRequestEvent(ctx context.Context, e event.Event) (context.Context, func(errOrResult error, event *event.Event)) {
	return ctx, func(errOrResult error, event *event.Event) {}
}

// RecordSendingEvent implements client.ObservabilityService.
func (*Observer) RecordSendingEvent(ctx context.Context, event event.Event) (context.Context, func(errOrResult error)) {
	return ctx, func(errOrResult error) {}
}
