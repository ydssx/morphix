package eventhandler

import (
	"context"
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/ydssx/morphix/common/event"
	"github.com/ydssx/morphix/pkg/logger"
)

func UpdateOrderStatus(ctx context.Context, e cloudevents.Event) error { 
	fmt.Printf("Got Event Context: %+v\n", e.Context)
	data := &event.PayloadPaymentCompleted{}
	if err := e.DataAs(data); err != nil {
		logger.Infof(ctx, "Got Data Error: %s\n", err.Error())
	}
	logger.Infof(ctx, "Got Data: %+v\n", data)

	return nil

}
