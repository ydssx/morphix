#!/bin/bash

set -eo pipefail

# 阿里云容器镜像服务相关信息
REGISTRY_URL="registry.cn-shenzhen.aliyuncs.com"
REGISTRY_NAMESPACE="ydssx"
registry_username="$REGISTRY_USERNAME"
registry_password="$REGISTRY_PASSWORD"

# 定义版本号
version="1.0.1"

# 登录镜像仓库
docker login --username "$registry_username" --password "$registry_password" $REGISTRY_URL

# 查找微服务目录下的 Dockerfile 文件，并迭代构建和推送镜像
services=("gateway" "user" "sms")
for service in "${services[@]}"; do
    image_name="$REGISTRY_URL/$REGISTRY_NAMESPACE/$service:$version"
    
    echo "构建镜像：$image_name"
    
    # 构建镜像
    TAG="$version" docker-compose build "$service"
        
    # 推送镜像到阿里云容器镜像服务
    docker push "$image_name"
    
    echo "镜像构建和推送完成：$image_name"
    echo
done
