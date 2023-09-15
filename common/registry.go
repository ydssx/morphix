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

func NewEtcdRegistry(c *conf.Etcd) *etcd.Registry {
	once.Do(func() {
		client, err := etcdclient.New(etcdclient.Config{
			Endpoints:   c.Endpoints,
			DialTimeout: c.Timeout.AsDuration(),
			Username:    c.Username,
			Password:    c.Password,
		})
		if err != nil {
			panic(err)
		}
		registry = etcd.New(client)
	})

	return registry
}
