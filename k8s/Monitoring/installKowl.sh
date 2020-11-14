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
helm repo add cloudhut https://raw.githubusercontent.com/cloudhut/charts/master/archives
helm repo update
echo "================================"
echo "🚢 Apply K8s objects for Kowl"
echo "================================"
#kubectl apply -f kowl-port.yml
echo ""
echo "========================================================="
echo "📦 Install/Update Kowl"
echo "========================================================="
helm upgrade -i kowl  cloudhut/kowl -n monitoring -f kowlvalues.yaml 

output=$(helm list -n monitoring | grep "kowl"  | awk -F '\t' '{print $7}')
echo "$output"
echo ""
