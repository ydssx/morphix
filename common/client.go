package common

import (
	"context"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	goredis "github.com/redis/go-redis/v9"
	smsv1 "github.com/ydssx/morphix/api/sms/v1"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/interceptors"
	"github.com/ydssx/morphix/pkg/logger"
	"github.com/ydssx/morphix/pkg/redis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func NewSMSClient(c *conf.Bootstrap) smsv1.SMSServiceClient {
	conn := createConn(c.Etcd, c.SmsRpcClient)

	return smsv1.NewSMSServiceClient(conn)
}

func NewUserClient(c *conf.Bootstrap) userv1.UserServiceClient {
	conn := createConn(c.Etcd, c.UserRpcClient)

	return userv1.NewUserServiceClient(conn)
}

func createConn(etcdConf *conf.Etcd, rpcCliConf *conf.ClientConf) *grpc.ClientConn {
	r := NewEtcdRegistry(etcdConf)

	return CreateClientConn(rpcCliConf, r)
}

func CreateClientConn(rpcCliConf *conf.ClientConf, r *etcd.Registry) *grpc.ClientConn {
	ctx := context.Background()
	conn, err := kgrpc.DialInsecure(ctx,
		kgrpc.WithEndpoint(rpcCliConf.Addr),
		kgrpc.WithTimeout(rpcCliConf.Timeout.AsDuration()),
		kgrpc.WithDiscovery(r),
		kgrpc.WithUnaryInterceptor(
			interceptors.TraceClient(),
			interceptors.LoggingClient(logger.DefaultLogger),
			interceptors.MetricClient(),
		),
	)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", rpcCliConf.Addr, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", rpcCliConf.Addr, cerr)
			}
		}()
	}()

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
