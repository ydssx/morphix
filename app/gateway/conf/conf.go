package conf

import (
	"log"

	"github.com/ydssx/morphix/common"
)

type Config struct {
	common.CommonConfig `mapstructure:"common,omitempty" json:"common"`
	Name                string `mapstructure:"name,omitempty" json:"name"`
	Addr                string `mapstructure:"addr,omitempty" json:"addr,omitempty"`
}

func MustLoad(v *Config, path ...string) {
	common.MustLoad(&v, path...)
	log.Print(v)
}
