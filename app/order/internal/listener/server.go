package listener

import (
	"context"

	"github.com/dapr/go-sdk/service/common"
	"github.com/google/wire"
	"github.com/ydssx/morphix/common/event"
	"github.com/ydssx/morphix/pkg/mq"
)

var ProviderSet = wire.NewSet(NewListenerServer)

type ListenerServer struct {
}

func NewListenerServer() *ListenerServer {
	return &ListenerServer{}
}

func (*ListenerServer) Start(context.Context) error {
	for subject, handler := range subjectHandlerMap {
		err := mq.AddEventListenerAsync(event.Subject_name[int32(subject)], handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func (*ListenerServer) Stop(context.Context) error {
	return nil
}

func RegisterListener(srv common.Service) {
	for subject, handler := range daprSubjectHandlerMap {
		sub := &common.Subscription{
			PubsubName: "pubsub",
			Topic:      event.Subject_name[int32(subject)],
		}
		err := srv.AddTopicEventHandler(sub, handler)
		if err != nil {
			panic(err)
		}
	}
}
