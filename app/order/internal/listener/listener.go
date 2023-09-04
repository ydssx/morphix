package listener

import (
	"context"
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/dapr/go-sdk/service/common"
	"github.com/ydssx/morphix/common/event"
	"github.com/ydssx/morphix/pkg/logger"
	"github.com/ydssx/morphix/pkg/mq"
)

var subjectHandlerMap = map[event.Subject]mq.EventHandler{
	event.Subject_PaymentCompleted: updateOrderStatus,
	event.Subject_CancelPayment:    updateOrderStatus,
}

var daprSubjectHandlerMap = map[event.Subject]common.TopicEventHandler{
	event.Subject_PaymentCompleted: updateOrder,
}

func updateOrderStatus(ctx context.Context, e cloudevents.Event) error {
	fmt.Printf("Got Event Context: %+v\n", e.Context)
	data := &event.PayloadPaymentCompleted{}
	if err := e.DataAs(data); err != nil {
		logger.Infof(ctx, "Got Data Error: %s\n", err.Error())
	}
	logger.Infof(ctx, "Got Data: %+v\n", data)

	return nil

}

func updateOrder(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	data := &event.PayloadPaymentCompleted{}
	if err := e.Struct(data); err != nil {
		logger.Infof(ctx, "Got Data Error: %s\n", err.Error())
	}
	logger.Infof(ctx, "Got Data: %+v\n", data)

	return false, nil
}
