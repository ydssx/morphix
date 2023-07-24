package main

import (
	"context"
	"log"

	"github.com/ydssx/morphix/app/order/internal/eventhandler"
	"github.com/ydssx/morphix/common/event"
	"github.com/ydssx/morphix/pkg/mq"
)

var subjectHandlerMap = map[event.Subject]mq.EventHandler{
	event.Subject_PaymentCompleted: eventhandler.UpdateOrderStatus,
}

func main() {
	close, err := mq.InitNats("http://localhost:4222")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	defer close(ctx)

	for subject, handler := range subjectHandlerMap {
		err = mq.AddEventHandlerAsync(event.Subject_name[int32(subject)], handler)
		if err != nil {
			panic(err)
		}
	}
	log.Print("handler register success.")
	<-ctx.Done()
}
