package mq

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/nats-io/nats.go"
)

var natsConn *nats.Conn

func InitNats(url string) (conn *nats.Conn, cleanup func(), err error) {
	natsConn, err = nats.Connect(url, nats.Timeout(time.Second*5))
	if err != nil {
		panic("failed to connect to NATS: " + err.Error())
	}
	cleanup = func() {
		err = natsConn.Drain()
	}
	log.Info("init nats success")
	return natsConn, cleanup, err
}
