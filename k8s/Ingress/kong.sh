#!/bin/sh 
helm repo add kong https://charts.konghq.com
helm repo update
helm upgrade -i kong  kong/kong --set ingressController.installCRDs=false --create-namespace  --namespace ingress
echo ""