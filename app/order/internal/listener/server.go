package listener

import (
	"context"

	"github.com/google/wire"
	"github.com/ydssx/morphix/app/order/internal/biz"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/pkg/mq"
)

var ProviderSet = wire.NewSet(NewListenerServer, common.NewNatsConn, common.NewCloudEvent)

type ListenerServer struct {
	ce *mq.CloudEvent
	uc *biz.OrderUseCase
}

type ucKey struct{}

func GetUcFromContext(ctx context.Context) *biz.OrderUseCase {
	return ctx.Value(ucKey{}).(*biz.OrderUseCase)
}

func NewListenerServer(ce *mq.CloudEvent, uc *biz.OrderUseCase) *ListenerServer {
	return &ListenerServer{ce: ce, uc: uc}
}

func (l *ListenerServer) Start(ctx context.Context) error {
	ctx = context.WithValue(ctx, ucKey{}, l.uc)
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
