package common

import (
	"context"

	_ "github.com/dtm-labs/driver-kratos"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/nats-io/nats.go"
	goredis "github.com/redis/go-redis/v9"
	jobv1 "github.com/ydssx/morphix/api/job/v1"
	smsv1 "github.com/ydssx/morphix/api/sms/v1"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/interceptors"
	"github.com/ydssx/morphix/pkg/mq"
	"github.com/ydssx/morphix/pkg/redis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// NewSMSClient 创建一个 SMS 服务的客户端连接。
// 接收配置参数 c,使用 etcd 注册中心和配置的 SmsRpcClient 来创建一个 gRPC 连接,
// 然后用这个连接生成一个 smsv1.SMSServiceClient 返回。
func NewSMSClient(c *conf.Bootstrap) smsv1.SMSServiceClient {
	conn := createConn(c.Etcd, c.ClientSet.SmsRpcClient)

	return smsv1.NewSMSServiceClient(conn)
}

// NewUserClient 创建一个 User 服务的客户端连接。
// 接收配置参数 c,使用 etcd 注册中心和配置的 UserRpcClient 来创建一个 gRPC 连接,
// 然后用这个连接生成一个 userv1.UserServiceClient 返回。
func NewUserClient(c *conf.Bootstrap) userv1.UserServiceClient {
	conn := createConn(c.Etcd, c.ClientSet.UserRpcClient)

	return userv1.NewUserServiceClient(conn)
}

// NewJobClient 创建一个 Job 服务的客户端连接。
// 接收配置参数 c,使用 etcd 注册中心和配置的 JobRpcClient 来创建一个 gRPC 连接,
// 然后用这个连接生成一个 jobv1.JobServiceClient 返回。
func NewJobClient(c *conf.Bootstrap) jobv1.JobServiceClient {
	conn := createConn(c.Etcd, c.ClientSet.JobRpcClient)

	return jobv1.NewJobServiceClient(conn)
}

func createConn(etcdConf *conf.Etcd, rpcCliConf *conf.ClientConf) *grpc.ClientConn {
	r := NewEtcdRegistry(etcdConf)

	return CreateClientConn(context.Background(), rpcCliConf, r)
}

// CreateClientConn 使用给定的配置创建一个 gRPC 客户端连接。
// 接收上下文,客户端配置和服务发现注册中心作为参数。
// 使用 kgrpc 封装的 gRPC dial 函数建立连接,配置 interceptor、超时、发现等信息。
// 返回建立的 gRPC 客户端连接。如果发生错误会 panic。
func CreateClientConn(ctx context.Context, rpcCliConf *conf.ClientConf, r *etcd.Registry) *grpc.ClientConn {
	conn, err := kgrpc.DialInsecure(ctx,
		kgrpc.WithEndpoint(rpcCliConf.Addr),
		kgrpc.WithTimeout(rpcCliConf.Timeout.AsDuration()),
		kgrpc.WithDiscovery(r),
		kgrpc.WithUnaryInterceptor(
			interceptors.TraceClient(),
			interceptors.LoggingClient(),
			interceptors.MetricClient(),
		),
		kgrpc.WithStreamInterceptor(
			interceptors.LoggingStreamClient(),
			interceptors.TraceStreamClient(),
		),
		kgrpc.WithOptions(grpc.WithKeepaliveParams(keepalive.ClientParameters{})),
	)
	if err != nil {
		panic(err)
	}

	return conn
}

func NewRedisCluster(c *conf.Bootstrap) *goredis.ClusterClient {
	clusterConf := c.RedisCluster
	return redis.NewRedisCluster(&goredis.ClusterOptions{
		Addrs:        clusterConf.Addr,
		Username:     clusterConf.Username,
		Password:     clusterConf.Password,
		ReadTimeout:  clusterConf.ReadTimeout.AsDuration(),
		DialTimeout:  clusterConf.DialTimeout.AsDuration(),
		WriteTimeout: clusterConf.WriteTimeout.AsDuration(),
	})
}

func NewRedisClient(c *conf.Bootstrap) *goredis.Client {
	clientConf := c.Redis
	return redis.NewRedis(&goredis.Options{
		Addr:         clientConf.Addr,
		Username:     clientConf.Username,
		Password:     clientConf.Password,
		ReadTimeout:  clientConf.ReadTimeout.AsDuration(),
		DialTimeout:  clientConf.DialTimeout.AsDuration(),
		WriteTimeout: clientConf.WriteTimeout.AsDuration(),
	})
}

func NewNatsConn(c *conf.Bootstrap) (conn *nats.Conn, cleanup func(), err error) {
	return mq.InitNats(c.Nats.Addr)
}

func NewCloudEvent(conn *nats.Conn) *mq.CloudEvent {
	return mq.NewCloudEvent(conn)
}
