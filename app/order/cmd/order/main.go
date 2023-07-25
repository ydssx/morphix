package main

import (
	"context"
	"log"

	"github.com/ydssx/morphix/app/order/internal/eventlistener"
	"github.com/ydssx/morphix/pkg/mq"
)

func main() {
	close, err := mq.InitNats("http://localhost:4222")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	defer close(ctx)

	eventlistener.Init()

	log.Print("handler register success.")
	<-ctx.Done()
}
