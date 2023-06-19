package common

import (
	"log"

	"github.com/spf13/viper"
)

type CommonConfig struct {
	Etcd          EtcdConf   `mapstructure:"etcd,omitempty" json:"etcd"`
	Jeager        JeagerConf `mapstructure:"jeager,omitempty" json:"jeager"`
	UserRpcClient RpcConf    `mapstructure:"userRpcClient,omitempty" json:"userRpcClient,omitempty"`
	SmsRpcClient  RpcConf    `mapstructure:"smsRpcClient,omitempty" json:"smsRpcClient,omitempty"`
}

type RpcConf struct {
	Network string `yaml:"network,omitempty" mapstructure:"network,omitempty" json:"network"`
	Addr    string `yaml:"addr,omitempty" mapstructure:"addr,omitempty" json:"addr"`
	Timeout int    `yaml:"timeout" mapstructure:"timeout,omitempty" json:"timeout"`
}

type EtcdConf struct {
	Endpoints []string `yaml:"endpoints,omitempty" mapstructure:"endpoints,omitempty" json:"endpoints"`
	Timeout   int      `yaml:"timeout,omitempty" mapstructure:"timeout,omitempty" json:"timeout"`
}

type JeagerConf struct {
	Addr string `yaml:"addr,omitempty" mapstructure:"addr,omitempty" json:"addr"`
}

func MustLoad(out interface{}, filePath ...string) {
	for _, v := range filePath {
		cfg := mergeConfig(v)
		viper.MergeConfigMap(cfg)
	}
	log.Print(viper.AllSettings())
	if err := viper.Unmarshal(&out); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

}

func mergeConfig(fielpath string) map[string]interface{} {
	v := viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath(fielpath)
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
	// if err := v.MergeInConfig(); err != nil {
	// 	log.Fatalf("Failed to merge config file: %v", err)
	// }
	return v.AllSettings()
}
