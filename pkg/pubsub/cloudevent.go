package pubsub

import (
	"context"
	"fmt"
	"time"

	cenats "github.com/cloudevents/sdk-go/protocol/nats/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/client"
	"github.com/nats-io/nats.go"
	"github.com/ydssx/morphix/pkg/logger"
)

type CloudEventer struct {
	natsConn *nats.Conn
}

func NewCloudEventer(conn *nats.Conn) *CloudEventer {
	return &CloudEventer{natsConn: conn}
}

// PublishMessage 发布一个云事件到指定的主题。
//
// 参数:
//   - ctx: 上下文
//   - subject: 主题名称
//   - payload: 事件数据
//   - opts: 可选的发布选项
//
// 返回错误信息:
//   - 如果创建 nats 协议失败
//   - 如果创建客户端失败
//   - 如果发送事件失败
//
// 发布成功则返回 nil。
func (c *CloudEventer) PublishMessage(ctx context.Context, subject string, payload interface{}, opts ...Option) error {
	p, err := cenats.NewSenderFromConn(c.natsConn, subject)
	if err != nil {
		return fmt.Errorf("Failed to create nats protocol, %s", err.Error())
	}
	defer p.Close(ctx)

	ce, err := cloudevents.NewClient(p)
	if err != nil {
		return fmt.Errorf("Failed to create client, %s", err.Error())
	}

	event, err := NewEvent(ctx, payload, opts...)
	if err != nil {
		return err
	}

	if result := ce.Send(ctx, event); cloudevents.IsUndelivered(result) {
		logger.Errorf(ctx, "failed to send: %v", result)
		return err
	} else {
		logger.Infof(ctx, "subject: %s, sent: %s, accepted: %t", p.Subject, event.ID(), cloudevents.IsACK(result))
	}
	return nil
}

// SubscribeToTopic adds an event listener to the CloudEvent.
//
// It takes the following parameters:
//   - ctx: the context.Context object for the function.
//   - subject: the string representing the subject of the event.
//   - handler: the EventHandler object to handle the event.
//   - opts: variadic cenats.ConsumerOption options for the consumer.
//
// It returns an error if any of the following operations fail:
//   - creating the nats protocol
//   - creating the client
//   - starting the nats receiver
//   - closing the consumer
func (c *CloudEventer) Subscribe(ctx context.Context, subject string, handler EventHandler, opts ...Subscription) error {
	cons := new(Consumer)
	for _, opt := range opts {
		opt(cons)
	}
	var opt []cenats.ConsumerOption
	if cons.Queue != "" {
		opt = append(opt, cenats.WithQueueSubscriber(cons.Queue))
	}

	consumer, err := cenats.NewConsumerFromConn(c.natsConn, subject, opt...)
	if err != nil {
		return fmt.Errorf("failed to create nats protocol: %s", err)
	}
	defer consumer.Close(ctx)

	client, err := cloudevents.NewClient(consumer, client.WithObservabilityService(NewObserver()))
	if err != nil {
		return fmt.Errorf("failed to create client: %s", err)
	}

	if err := client.StartReceiver(ctx, handler); err != nil {
		return fmt.Errorf("failed to start nats receiver: %s", err)
	}

	return nil
}

// SubscribeToTopicAsync adds an asynchronous event listener to the CloudEvent.
// It takes the same parameters as AddEventListener, but returns immediately
// and runs the listener in a new goroutine. The context is used to control
// cancelation and timeout.
func (c *CloudEventer) SubscribeAsync(ctx context.Context, subject string, handler EventHandler, opts ...Subscription) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	errChan := make(chan error, 1)

	go func() {
		defer close(errChan)
		errChan <- c.Subscribe(ctx, subject, handler, opts...)
	}()

	select {
	case err = <-errChan:
		// 成功接收到 Goroutine 完成的错误值
	case <-ctx.Done():
		// err = ctx.Err()
	}

	return err
}

func (c *CloudEventer) Close() error {
	return c.natsConn.Drain()
}