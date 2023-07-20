package mq

import (
	"context"

	"github.com/nats-io/nats.go"
)

var natsServer *nats.Conn

func InitNats(url string) (fn func(context.Context) error, err error) {
	natsServer, err = nats.Connect(url)
	if err != nil {
		panic(err)
	}

	return func(context.Context) error { return natsServer.Drain() }, nil
}
