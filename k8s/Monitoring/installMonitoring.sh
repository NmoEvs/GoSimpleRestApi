#!/bin/bash
LC_TIME=C
echo "================================"
echo "🏁 K8s Pulse checker"
echo "================================"
date
echo "================================"
echo "🔧 Create Namespace Monitoring"
echo "================================"
kubectl create ns monitoring
echo "================================"
echo "🔧 Add Helm Repositories"
echo "================================"
helm repo add prometheus-com https://prometheus-community.github.io/helm-charts
helm repo add loki https://grafana.github.io/loki/charts
helm repo add victoria-metrics  https://victoriametrics.github.io/helm-charts/
echo "================================"
echo "🚢 Apply K8s objects for Pulse"
echo "================================"
kubectl apply -f grafanaDatasourceSecret.yml -f pv_victoriaMetrics.yaml -f pv_loki.yaml
echo ""
echo "========================================================="
echo "📦 Install/Update Kube-Prometheus-Stack"
echo "========================================================="
helm upgrade -i kube-prometheus-stack  prometheus-com/kube-prometheus-stack --version 9.4.4 -n monitoring -f prometheusStackValues.yaml --create-namespace  --namespace monitoring
# Output error => 
#Bug https://github.com/prometheus-community/helm-charts/issues/155
output=$(helm list -n monitoring | grep "kube-prometheus-stack"  | awk -F '\t' '{print $7}')
echo "$output"
echo ""
echo "========================================================="
echo "📦 Get current Loki Stack Version installed"
echo "========================================================="
# Upgrade is buggy because Loki  is StatefullSet Need to uninstall first
helm upgrade -i loki loki/loki-stack -n monitoring -f lokiValues.yaml
output=$(helm list -n monitoring | grep "loki-stack"  | awk -F '\t' '{print $7}')
echo Loki-stack version installed : "$output"
echo ""
echo "========================================================="
echo "📦 Get current Victoria Metrics Stack Version installed"
echo "========================================================="
# Upgrade is buggy because Victoria is StatefullSet
helm upgrade -i  vmetrics victoria-metrics/victoria-metrics-single -n monitoring  -f victoriaMetricsValues.yaml
output=$(helm list -n monitoring | grep "vmetrics"  | awk -F '\t' '{print $7}')
echo "$output"

echo ""