package redis

import (
	"encoding/json"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestRedisPubSub_PublishMessage(t *testing.T) {
	cli, _ := NewRedis(&redis.Options{Addr: "localhost:6379"})
	pubsub := NewRedisPubSub(cli)
	pubsub.SubscribeToTopic("test", func(message []byte) {
		time.Sleep(time.Second * 1)
		log.Println("sub1:", string(message))
	})
	pubsub.SubscribeToTopic("test1", func(message []byte) { log.Println("sub2:", string(message)) })
	// go func() {
	// 	pubsub.SubscribeToTopic("test", func(message []byte) { log.Print("11111",string(message)) })
	// }()
	type Msg struct {
		Id int
	}
	go func() {
		for i := 0; i < 10; i++ {
			m, _ := json.Marshal(Msg{Id: i})
			pubsub.PublishMessage("test", m)
			time.Sleep(time.Second)
		}
		time.Sleep(time.Second * 2)
	}()
	go func() {
		for i := 0; i < 10; i++ {
			pubsub.PublishMessage("test", "published msg:"+strconv.Itoa(i))
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 10)
	err := pubsub.Close()
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second * 2)
}
