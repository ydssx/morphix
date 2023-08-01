package common

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/ydssx/morphix/common/conf"
	etcdclient "go.etcd.io/etcd/client/v3"
)

func NewEtcdRegistry(c *conf.Etcd) *etcd.Registry {
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints:   c.Endpoints,
		DialTimeout: c.Timeout.AsDuration(),
		Username:    c.Username,
		Password:    c.Password,
	})
	if err != nil {
		panic(err)
	}
	r := etcd.New(client)
	return r
}
