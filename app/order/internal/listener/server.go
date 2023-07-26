package listener

import (
	"context"

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
