package redis

import (
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestRedisPubSub_PublishMessage(t *testing.T) {
	cli, _ := NewRedis(&redis.Options{Addr: "localhost:6379"})
	pubsub := NewRedisPubSub(cli)
	pubsub.SubscribeToTopic("test", func(message []byte) { log.Println("sub1:", string(message)) })
	pubsub.SubscribeToTopic("test", func(message []byte) { log.Println("sub2:", string(message)) })
	// go func() {
	// 	pubsub.SubscribeToTopic("test", func(message []byte) { log.Print("11111",string(message)) })
	// }()
	go func() {
		for i := 0; i < 10; i++ {
			pubsub.PublishMessage("test", "publish msg:"+strconv.Itoa(i))
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 5)
	err := pubsub.Close()
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second * 2)
}
