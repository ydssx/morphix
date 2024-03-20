### 项目简介

一个简单的微服务项目，用于演示如何使用Morphix，包括如何使用Morphix在Kubernetes上部署一个微服务，并使用Morphix在Kubernetes上部署一个API网关。

### 项目目录结构

```shell
├── README.md
├── deploy
│   └── kubernetes
│       ├── gateway.yaml
│       └── ingress
│           └── ingress.yaml
├── app
│   ├── aiart     # aiart服务
│   ├── chat      # chat服务
│   ├── gateway   # gateway服务
│   ├── job       # job服务
│   ├── media     # media服务
│   ├── order     # order服务
│   ├── payment   # payment服务
│   ├── sms       # sms服务
│   ├── user      # user服务
├── common
│   └── conf
│       └── conf.proto
├── proto
│   └── chat.proto
│   └── gateway.proto
│   └── job.proto
│   └── media.proto
│   └── order.proto
│   └── payment.proto
│   └── sms.proto
│   └── user.proto
├── api
│   ├── chat/v1/chat.proto
│   ├── gateway/v1
│   ├── job/v1/job.proto
│   ├── media/v1
│   ├── order/v1
│   ├── payment/v1
│   ├── sms/v1
│   └── user/v1
└── go.mod
└── go.sum
```
### 部署

1. 为了演示方便，我们使用了本地的Kubernetes集群，并把应用部署在本地的Kubernetes集群上，这样我们可以直接在本地运行应用。
2. 我们使用了Kustomize，这个工具可以将Kubernetes的配置文件和应用程序代码一起部署到Kubernetes集群中，这种方式可以让我们更轻松的部署和管理应用。
