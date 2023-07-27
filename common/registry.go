package common

import (
	"time"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	etcdclient "go.etcd.io/etcd/client/v3"
)

func NewEtcdRegistry(c Etcd) *etcd.Registry {
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints:   c.Endpoints,
		DialTimeout: time.Duration(c.Timeout) * time.Second,
		Username:    c.Username,
		Password:    c.Password,
	})
	if err != nil {
		panic(err)
	}
	r := etcd.New(client)
	return r
}
