package eventlistener

import (
	"context"
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/ydssx/morphix/common/event"
	"github.com/ydssx/morphix/pkg/logger"
	"github.com/ydssx/morphix/pkg/mq"
)

var subjectHandlerMap = map[event.Subject]mq.EventHandler{
	event.Subject_PaymentCompleted: updateOrderStatus,
}

func Init() {
	for subject, handler := range subjectHandlerMap {
		err := mq.AddEventListenerAsync(event.Subject_name[int32(subject)], handler)
		if err != nil {
			panic(err)
		}
	}
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
