package nats

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/nats-io/nats.go"
)

var natsConn *nats.Conn

// InitNats connects to a NATS server and returns the connection.
// The connection is drained on cleanup, which should be called
// before the application exits.
func InitNats(url string) (conn *nats.Conn, cleanup func(), err error) {
	// Connect to NATS
	natsConn, err = nats.Connect(url, nats.Timeout(time.Second*5))
	if err != nil {
		panic("failed to connect to NATS: " + err.Error())
	}

	// Return a cleanup function to drain the connection on exit
	cleanup = func() {
		err = natsConn.Drain()
	}

	// Log success
	log.Info("init nats success")

	// Return the connection and cleanup function
	return natsConn, cleanup, err
}
