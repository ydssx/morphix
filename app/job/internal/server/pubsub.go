package server

import (
	"context"

	"github.com/ydssx/morphix/app/job/internal/common"
	"github.com/ydssx/morphix/app/job/internal/handler"
	"github.com/ydssx/morphix/pkg/mq"
)

type ListenerServer struct {
	ce        *mq.CloudEvent
	clientSet *common.ServiceClientSet
}

func NewListenerServer(ce *mq.CloudEvent, clientSet *common.ServiceClientSet) *ListenerServer {
	return &ListenerServer{ce: ce, clientSet: clientSet}
}

func (l *ListenerServer) Start(ctx context.Context) error {
	ctx = common.NewContextWithServiceClientSet(ctx, l.clientSet)
	for subject, handler := range handler.PubsubHandlerMap {
		err := l.ce.SubscribeToTopicAsync(ctx, subject.String(), handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func (*ListenerServer) Stop(context.Context) error {
	return nil
}
