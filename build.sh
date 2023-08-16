#!/bin/bash

set -eo pipefail


# 阿里云容器镜像服务相关信息
REGISTRY_URL="registry.cn-shenzhen.aliyuncs.com"
REGISTRY_NAMESPACE="ydssx"
registry_username="$REGISTRY_USERNAME"
registry_password="$REGISTRY_PASSWORD"

# 定义版本号
version="1.0.3"

# 登录镜像仓库
echo "$registry_password" | docker login --username "$registry_username" --password-stdin $REGISTRY_URL

# 定义要构建的微服务列表
services=("gateway" "user" "sms" "order" "payment")

# 构建和推送指定的微服务镜像
build_and_push() {
    service="$1"
    image_name="$REGISTRY_URL/$REGISTRY_NAMESPACE/$service:$version"

    echo "构建镜像：$image_name"

    # 构建镜像
    TAG="$version" docker-compose build "$service"

    # 推送镜像到阿里云容器镜像服务
    docker push "$image_name"

    echo "镜像构建和推送完成：$image_name"
    echo
}

# 最大并发数
max_concurrency=3
current_jobs=0

# 如果传入了服务名参数，则只构建和推送指定的服务
if [ $# -gt 0 ]; then
    for service in "$@"; do
        build_and_push "$service" &
        ((current_jobs++))
        
        if [ $current_jobs -ge $max_concurrency ]; then
            wait
            current_jobs=0
        fi
    done
else
    # 否则，构建和推送所有微服务镜像
    services=("gateway" "user" "sms" "order" "payment")
    for service in "${services[@]}"; do
        build_and_push "$service" &
        ((current_jobs++))
        
        if [ $current_jobs -ge $max_concurrency ]; then
            wait
            current_jobs=0
        fi
    done
fi

# 确保最后一批任务完成
wait