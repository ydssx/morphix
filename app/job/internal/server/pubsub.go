package server

import (
	"context"

	"github.com/ydssx/morphix/app/job/internal/common"
	"github.com/ydssx/morphix/app/job/internal/handler"
	"github.com/ydssx/morphix/pkg/pubsub"
)

type ListenerServer struct {
	sub       pubsub.Subscriber
	clientSet *common.ServiceClientSet
}

func NewListenerServer(sub pubsub.Subscriber, clientSet *common.ServiceClientSet) *ListenerServer {
	return &ListenerServer{sub: sub, clientSet: clientSet}
}

func (l *ListenerServer) Start(ctx context.Context) error {
	ctx = common.NewContextWithServiceClientSet(ctx, l.clientSet)
	for subject, handler := range handler.PubsubHandlerMap {
		err := l.sub.SubscribeAsync(ctx, subject.String(), handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func (*ListenerServer) Stop(context.Context) error {
	return nil
}
