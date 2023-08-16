package mq

import (
	"time"

	"github.com/nats-io/nats.go"
)

var natsConn *nats.Conn

func InitNats(url string) (conn *nats.Conn, cleanup func(), err error) {
	natsConn, err = nats.Connect(url, nats.Timeout(time.Second*5))
	if err != nil {
		panic(err)
	}
	cleanup = func() {
		err = natsConn.Drain()
	}
	return natsConn, cleanup, err
}
