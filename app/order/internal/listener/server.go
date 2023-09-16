package listener

import (
	"context"

	daprcommon "github.com/dapr/go-sdk/service/common"
	"github.com/google/wire"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/pkg/mq"
)

var ProviderSet = wire.NewSet(NewListenerServer, common.NewNatsConn, common.NewCloudEvent)

type ListenerServer struct {
	ce *mq.CloudEvent
}

func NewListenerServer(ce *mq.CloudEvent) *ListenerServer {
	return &ListenerServer{ce: ce}
}

func (l *ListenerServer) Start(ctx context.Context) error {
	for subject, handler := range subjectHandlerMap {
		err := l.ce.AddEventListenerAsync(ctx, subject.String(), handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func (*ListenerServer) Stop(context.Context) error {
	return nil
}

func RegisterListener(srv daprcommon.Service) {
	for subject, handler := range daprSubjectHandlerMap {
		sub := &daprcommon.Subscription{
			PubsubName: "pubsub",
			Topic:      subject.String(),
		}
		err := srv.AddTopicEventHandler(sub, handler)
		if err != nil {
			panic(err)
		}
	}
}
