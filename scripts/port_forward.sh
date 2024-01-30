#! /bin/bash

kubectl port-forward pods/etcd-0 2379:2379 -n default &
kubectl port-forward pods/mysql-0 3306:3306 -n default &
kubectl port-forward pods/nats-0 4222:4222 -n default &
kubectl port-forward pods/redis-master-0 6379:6379 -n default &
kubectl port-forward svc/otelcol 4317:4317 -n morphix &
kubectl port-forward --namespace kubeapps service/kubeapps 8080:80 &
kubectl port-forward service/loki-stack-grafana 3000:80 -n loki-stack &
kubectl -n kubernetes-dashboard port-forward service/kubernetes-dashboard 8443:443 &
kubectl port-forward service/jaeger-query 16686:16686 &