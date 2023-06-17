package conf

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Addr    string   `yaml:"addr,omitempty" json:"addr,omitempty"`
	UserRpc RpcConf  `yaml:"userRpc,omitempty"`
	Etcd    EtcdConf `yaml:"etcd,omitempty"`
}

type RpcConf struct {
	Network string `yaml:"network,omitempty"`
	Addr    string `yaml:"addr,omitempty"`
	Timeout string `yaml:"timeout"`
}

type EtcdConf struct {
	Endpoints []string `yaml:"endpoints,omitempty"`
	Timeout   int      `yaml:"timeout,omitempty"`
}

func MustLoad(path string, v *Config) {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(content, &v)
	if err != nil {
		panic(err)
	}
}
