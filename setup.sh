#! /bin/bash

helm upgrade --install etcd bitnami/etcd -f deploy/kubernetes/helm/etcd.yaml
helm upgrade --install mysql bitnami/mysql -f deploy/kubernetes/helm/mysql.yaml
helm upgrade --install nats nats/nats -f deploy/kubernetes/helm/nats.yaml
helm upgrade --install redis bitnami/redis -f deploy/kubernetes/helm/redis.yaml
helm upgrade --install prometheus prometheus-community/prometheus -f deploy/kubernetes/helm/prometheus.yaml