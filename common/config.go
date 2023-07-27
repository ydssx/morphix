package common

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config
type Config struct {
	Gateway          Gateway   `yaml:"gateway"`
	User             User      `yaml:"user"`
	Sms              Sms       `yaml:"sms"`
	Order            Order     `yaml:"order"`
	Payment          Payment   `yaml:"payment"`
	Etcd             Etcd      `yaml:"etcd"`
	Jaeger           Jaeger    `yaml:"jaeger"`
	Otelcol          Otelcol   `yaml:"otelcol"`
	Nats             Nats      `yaml:"nats"`
	UserRpcClient    RpcClient `yaml:"userRpcClient"`
	SmsRpcClient     RpcClient `yaml:"smsRpcClient"`
	PaymentRpcClient RpcClient `yaml:"paymentRpcClient"`
	OrderRpcClient   RpcClient `yaml:"orderRpcClient"`
}

// Server
type Server struct {
	Grpc Grpc `yaml:"grpc"`
}

// Database
type Database struct {
	Driver string `yaml:"driver"`
	Source string `yaml:"source"`
}

// SmsData
type SmsData struct {
	Database Database `yaml:"database"`
	Redis    Redis    `yaml:"redis"`
}

// RpcClient
type RpcClient struct {
	Network string `yaml:"network"`
	Addr    string `yaml:"addr"`
	Timeout int    `yaml:"timeout"`
}

// SmsServerGrpc
type SmsServerGrpc struct {
	Addr    string `yaml:"addr"`
	Timeout int    `yaml:"timeout"`
}

// Gateway
type Gateway struct {
	Name string `yaml:"name"`
	Addr string `yaml:"addr"`
}

// Sms
type Sms struct {
	Name   string `yaml:"name"`
	Server Server `yaml:"server"`
	Data   Data   `yaml:"data"`
}

// Grpc
type Grpc struct {
	Addr    string `yaml:"addr"`
	Timeout int    `yaml:"timeout"`
}

// SmsDataRedis
type Redis struct {
	Addr         string `yaml:"addr"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
	DialTimeout  int    `yaml:"dial_timeout"`
}

// Etcd
type Etcd struct {
	Endpoints []string `yaml:"endpoints"`
	Timeout   int      `yaml:"timeout"`
	Username  string   `yaml:"username"`
	Password  string   `yaml:"password"`
}

// Jaeger
type Jaeger struct {
	Addr string `yaml:"addr"`
}
type Otelcol struct {
	Addr string `yaml:"addr"`
}

type Nats struct {
	Addr    string `yaml:"addr"`
	Timeout int    `yaml:"timeout"`
}

// User
type User struct {
	Server Server `yaml:"server"`
	Data   Data   `yaml:"data"`
	Name   string `yaml:"name"`
}
type Payment struct {
	Server Server `yaml:"server"`
	Data   Data   `yaml:"data"`
	Name   string `yaml:"name"`
}
type Order struct {
	Server Server `yaml:"server"`
	Data   Data   `yaml:"data"`
	Name   string `yaml:"name"`
}

// Data
type Data struct {
	Database Database `yaml:"database"`
	Redis    Redis    `yaml:"redis"`
}

func MustLoad(out *Config, path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(content, &out)
	if err != nil {
		panic(err)
	}
}
