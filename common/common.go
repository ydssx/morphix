package common

import (
	"log"

	"github.com/spf13/viper"
)

type CommonConfig struct {
	Etcd   EtcdConf   `mapstructure:"etcd,omitempty"`
	Jeager JeagerConf `mapstructure:"jeager,omitempty"`
}

type RpcConf struct {
	Network string `yaml:"network,omitempty" mapstructure:"network,omitempty"`
	Addr    string `yaml:"addr,omitempty" mapstructure:"addr,omitempty"`
	Timeout string `yaml:"timeout" mapstructure:"timeout,omitempty"`
}

type EtcdConf struct {
	Endpoints []string `yaml:"endpoints,omitempty" mapstructure:"endpoints,omitempty"`
	Timeout   int      `yaml:"timeout,omitempty" mapstructure:"timeout,omitempty"`
}

type JeagerConf struct {
	Addr string `yaml:"addr,omitempty" mapstructure:"addr,omitempty"`
}

func MustLoad(baseFile string, out interface{}, file ...string) {
	config := viper.New()
	config.SetConfigType("yaml")
	config.AddConfigPath(baseFile)
	for _, v := range file {
		config.AddConfigPath(v)
	}

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
	if err := config.MergeInConfig(); err != nil {
		log.Fatalf("Failed to merge config file: %v", err)
	}
	if err := config.Unmarshal(&out); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}
}
