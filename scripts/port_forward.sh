#! /bin/bash

kubectl port-forward pods/etcd-0 2379:2379 -n default &
kubectl port-forward pods/mysql-0 3306:3306 -n default &
kubectl port-forward pods/nats-0 4222:4222 -n default &
kubectl port-forward pods/redis-master-0 6379:6379 -n default &
kubectl port-forward svc/otelcol 4317:4317 -n morphix &