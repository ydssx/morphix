package common

import (
	"sync"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/ydssx/morphix/common/conf"
	etcdclient "go.etcd.io/etcd/client/v3"
)

var (
	registry *etcd.Registry
	once     sync.Once
)

// NewEtcdRegistry 创建一个新的 etcd 注册中心客户端。
// 如果配置了用户名和密码,会进行认证。
// 它实现了 kratos/contrib/registry 接口。
func NewEtcdRegistry(c *conf.Etcd) *etcd.Registry {
	once.Do(func() {
		client, err := etcdclient.New(etcdclient.Config{
			Endpoints:   c.Endpoints,
			DialTimeout: c.Timeout.AsDuration(),
			Username:    c.Username,
			Password:    c.Password,
		})
		if err != nil {
			panic("failed to connect to etcd: " + err.Error())
		}
		registry = etcd.New(client)
	})

	return registry
}
