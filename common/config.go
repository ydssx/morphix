package common

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config
type Config struct {
	Gateway       Gateway       `yaml:"gateway"`
	User          User          `yaml:"user"`
	Sms           Sms           `yaml:"sms"`
	Etcd          Etcd          `yaml:"etcd"`
	Jeager        Jeager        `yaml:"jeager"`
	UserRpcClient UserRpcClient `yaml:"userRpcClient"`
	SmsRpcClient  SmsRpcClient  `yaml:"smsRpcClient"`
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

// UserRpcClient
type UserRpcClient struct {
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

// SmsRpcClient
type SmsRpcClient struct {
	Addr    string `yaml:"addr"`
	Timeout int    `yaml:"timeout"`
	Network string `yaml:"network"`
}

// Etcd
type Etcd struct {
	Endpoints []string `yaml:"endpoints"`
	Timeout   int      `yaml:"timeout"`
}

// Jeager
type Jeager struct {
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