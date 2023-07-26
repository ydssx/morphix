package mq

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/nats-io/nats.go"
)

var natsServer *nats.Conn

func InitNats(url string) (fn func(context.Context) error, err error) {
	log.Info("nats url:", url)
	natsServer, err = nats.Connect(url,nats.Timeout(time.Second*5))
	if err != nil {
		panic(err)
	}

	return func(context.Context) error { return natsServer.Drain() }, nil
}
