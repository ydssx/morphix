package mq

import (
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/google/uuid"
)

type Example struct {
	Sequence int    `json:"id"`
	Message  string `json:"message"`
}

type Event struct {
	contentType ContentType
	source      string
	eventType   string
}

type ContentType string

const (
	ApplicationJSON ContentType = cloudevents.ApplicationJSON
	ApplicationXML  ContentType = cloudevents.ApplicationXML
	TextPlain       ContentType = cloudevents.TextPlain
)

func newDefaultEvent() *Event {
	return &Event{
		contentType: ApplicationJSON,
		source:      "api",
		eventType:   "update",
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

func NewEvent(payload interface{}, opts ...Option) event.Event {
	ev := newDefaultEvent()
	for _, v := range opts {
		v(ev)
	}

	e := cloudevents.NewEvent()
	e.SetID(uuid.New().String())
	e.SetType(ev.eventType)
	e.SetTime(time.Now())
	e.SetSource(ev.source)
	_ = e.SetData(string(ev.contentType), payload)

	return e
}
