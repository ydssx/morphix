package common

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config
type Config struct {
	Gateway       Gateway   `yaml:"gateway"`
	User          User      `yaml:"user"`
	Sms           Sms       `yaml:"sms"`
	Etcd          Etcd      `yaml:"etcd"`
	Jaeger        Jaeger    `yaml:"jaeger"`
	Otelcol       Otelcol   `yaml:"otelcol"`
	UserRpcClient RpcClient `yaml:"userRpcClient"`
	SmsRpcClient  RpcClient `yaml:"smsRpcClient"`
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
	Database SmsDataDatabase `yaml:"database"`
	Redis    SmsDataRedis    `yaml:"redis"`
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

// Redis
type Redis struct {
	Addr         string `yaml:"addr"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
}

// Sms
type Sms struct {
	Name   string    `yaml:"name"`
	Server SmsServer `yaml:"server"`
	Data   SmsData   `yaml:"data"`
}

// SmsServer
type SmsServer struct {
	Grpc SmsServerGrpc `yaml:"grpc"`
}

// Grpc
type Grpc struct {
	Addr    string `yaml:"addr"`
	Timeout int    `yaml:"timeout"`
}

// SmsDataRedis
type SmsDataRedis struct {
	Addr         string `yaml:"addr"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
}

// Etcd
type Etcd struct {
	Endpoints []string `yaml:"endpoints"`
	Timeout   int      `yaml:"timeout"`
}

// Jaeger
type Jaeger struct {
	Addr string `yaml:"addr"`
}
type Otelcol struct {
	Addr string `yaml:"addr"`
}

// User
type User struct {
	Server Server `yaml:"server"`
	Data   Data   `yaml:"data"`
	Name   string `yaml:"name"`
}

// Data
type Data struct {
	Database Database `yaml:"database"`
	Redis    Redis    `yaml:"redis"`
}

// SmsDataDatabase
type SmsDataDatabase struct {
	Driver string `yaml:"driver"`
	Source string `yaml:"source"`
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
