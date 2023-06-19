package conf

import (
	"log"

	"github.com/ydssx/morphix/common"
)

type Config struct {
	common.CommonConfig
	Name    string         `mapstructure:"name,omitempty"`
	Addr    string         `mapstructure:"addr,omitempty" json:"addr,omitempty"`
	UserRpc common.RpcConf `mapstructure:"userRpcClient,omitempty" json:"userRpcClient,omitempty"`
}

func MustLoad(path string, v *Config) {
	common.MustLoad("../../../configs", &v, path)
	log.Print(v)
}
