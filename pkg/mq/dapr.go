package mq

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
)

// Subscription to tell the dapr what topic to subscribe.
//   - PubsubName: is the name of the component configured in the metadata of pubsub.yaml.
//   - Topic: is the name of the topic to subscribe.
//   - Route: tell dapr where to request the API to publish the message to the subscriber when get a message from topic.
//   - Match: (Optional) The CEL expression to match on the CloudEvent to select this route.
//   - Priority: (Optional) The priority order of the route when Match is specificed.
//     If not specified, the matches are evaluated in the order in which they are added.
var defaultSubscription = &common.Subscription{
	PubsubName: "messages",
	Topic:      "neworder",
	Route:      "/orders",
}

var importantSubscription = &common.Subscription{
	PubsubName: "messages",
	Topic:      "neworder",
	Route:      "/important",
	Match:      `event.type == "important"`,
	Priority:   1,
}

func main() {
	s := daprd.NewService(":8080")

	if err := s.AddTopicEventHandler(defaultSubscription, eventHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
	
	if err := s.AddTopicEventHandler(importantSubscription, importantEventHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
	
	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}
}

func eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s", e.PubsubName, e.Topic, e.ID, e.Data)
	return false, nil
}

func importantEventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("important event - PubsubName: %s, Topic: %s, ID: %s, Data: %s", e.PubsubName, e.Topic, e.ID, e.Data)
	return false, nil
}

var (
	// set the environment as instructions.
	pubsubName = os.Getenv("DAPR_PUBSUB_NAME")
	topicName  = "neworder"
)

func main1() {
	ctx := context.Background()
	publishEventData := []byte("ping")
	publishEventsData := []interface{}{"multi-ping", "multi-pong"}

	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	
	defer client.Close()
	client.SaveState(context.Background(), "", "", []byte{}, make(map[string]string), dapr.WithConsistency(dapr.StateConsistencyEventual))
	
	if err := client.PublishEvent(ctx, pubsubName, topicName, publishEventData); err != nil {
		panic(err)
	}

	// Publish multiple events
	if res := client.PublishEvents(ctx, pubsubName, topicName, publishEventsData); res.Error != nil {
		panic(err)
	}

	fmt.Println("data published")

	fmt.Println("Done (CTRL+C to Exit)")
}
