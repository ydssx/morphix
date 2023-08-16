package mq

import (
	"context"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/google/uuid"
	"github.com/ydssx/morphix/pkg/interceptors"
	"go.opentelemetry.io/otel/trace"
)

type Event struct {
	contentType ContentType
	source      string
	eventType   string
	traceID     string
}

type ContentType string

const (
	ApplicationJSON ContentType = cloudevents.ApplicationJSON
	ApplicationXML  ContentType = cloudevents.ApplicationXML
	TextPlain       ContentType = cloudevents.TextPlain
)

func getTraceIDFromCtx(ctx context.Context) string {
	if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
		return span.TraceID().String()
	}
	return ""
}

func newDefaultEvent(ctx context.Context) *Event {
	return &Event{
		contentType: ApplicationJSON,
		source:      interceptors.EventSourceFromCtx(ctx),
		eventType:   interceptors.EventTypeFromCtx(ctx),
		traceID:     getTraceIDFromCtx(ctx),
	}
}

type Option func(*Event)

func WithContentType(t ContentType) Option {
	return func(e *Event) { e.contentType = t }
}

func WithSource(s string) Option {
	return func(e *Event) { e.source = s }
}

func WithEventType(t string) Option {
	return func(e *Event) { e.eventType = t }
}

func NewEvent(ctx context.Context, payload interface{}, opts ...Option) (event.Event, error) {
	ev := newDefaultEvent(ctx)
	for _, v := range opts {
		v(ev)
	}

	e := cloudevents.NewEvent()
	e.SetID(uuid.New().String())
	e.SetType(ev.eventType)
	e.SetTime(time.Now().Local())
	e.SetSource(ev.source)
	e.SetExtension("traceid", ev.traceID)
	err := e.SetData(string(ev.contentType), payload)

	return e, err
}
