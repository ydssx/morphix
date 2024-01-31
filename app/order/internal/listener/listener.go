package listener

import (
	"context"
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	orderv1 "github.com/ydssx/morphix/api/order/v1"
	"github.com/ydssx/morphix/common/event"
	"github.com/ydssx/morphix/pkg/logger"
	"github.com/ydssx/morphix/pkg/mq"
)

var subjectHandlerMap = map[event.Subject]mq.EventHandler{
	event.Subject_PaymentCompleted: updateOrderStatus,
	event.Subject_CancelPayment:    updateOrderStatus,
}

func updateOrderStatus(ctx context.Context, e cloudevents.Event) error {
	fmt.Printf("Got Event Context: %+v\n", e.Context)
	data := &event.PayloadPaymentCompleted{}
	if err := e.DataAs(data); err != nil {
		logger.Infof(ctx, "Got Data Error: %s\n", err.Error())
	}
	logger.Infof(ctx, "Got Data: %+v\n", data)

	if _, err := GetUcFromContext(ctx).UpdateOrderStatus(ctx, &orderv1.UpdateOrderStatusRequest{OrderNumber: "", Status: orderv1.OrderStatus_COMPLETED}); err != nil {
		logger.Errorf(ctx, "UpdateOrderStatus Error: %s\n", err.Error())
		return err
	}

	return nil
}
