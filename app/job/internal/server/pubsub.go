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


// Start starts the pubsub listener server.
// It subscribes to the specified subjects with the corresponding handlers.
func (l *ListenerServer) Start(ctx context.Context) error {
	// Use the service client set in the context.
	ctx = common.NewContextWithServiceClientSet(ctx, l.clientSet)

	// Iterate over the topic-handler pairs in the handler map.
	for subject, handler := range handler.PubsubHandlerMap {
		// Subscribe to the topic with the handler asynchronously.
		if err := l.sub.SubscribeAsync(ctx, subject.String(), handler); err != nil {
			// Return the error if subscription fails.
			return err
		}
	}

	// Return nil if all subscriptions succeed.
	return nil
}

func (l *ListenerServer) Stop(ctx context.Context) error {
	return l.sub.UnsubscribeAll(ctx)
}

