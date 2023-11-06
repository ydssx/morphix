#! /bin/bash

helm upgrade --install etcd bitnami/etcd -f deploy/kubernetes/helm/etcd.yaml
helm upgrade --install mysql bitnami/mysql -f deploy/kubernetes/helm/mysql.yaml
helm upgrade --install nats nats/nats -f deploy/kubernetes/helm/nats.yaml
helm upgrade --install redis bitnami/redis -f deploy/kubernetes/helm/redis.yaml
helm upgrade --install phpmyadmin bitnami/phpmyadmin -f deploy/kubernetes/helm/phpmyadmin.yaml
helm upgrade --install prometheus prometheus-community/prometheus -f deploy/kubernetes/helm/prometheus.yaml
helm upgrade --install loki --namespace=loki-stack grafana/loki-stack -f deploy/kubernetes/helm/loki.yaml
helm upgrade --install otelcol opentelemetry-helm/opentelemetry-collector -f deploy/kubernetes/helm/otelcol.yaml
helm upgrade --install thanos bitnami/thanos -f deploy/kubernetes/helm/thanos.yaml
helm upgrade --install jaeger jaegertracing/jaeger -f deploy/kubernetes/helm/jaeger.yaml
helm upgrade --install argocd argo/argo-cd -f deploy/kubernetes/helm/argocd.yaml --namespace=argocd --create-namespace
helm upgrade --install sealed-secrets -n kube-system --set-string fullnameOverride=sealed-secrets-controller sealed-secrets/sealed-secrets