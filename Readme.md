# Simple Rest API

## 1. Create Kind Cluster

    kind create cluster --name test --config clusterSimpleRestApi.yml

## 2. Create DockerFile Image localy

    docker build --tag k8s/simplerestapi:<version> .

## 3. Load images into Cluster nodes

    kind load docker-image k8s/simplerestapi:<version> --name test

## 4. Apply manifest to the cluster (*2)

     kubectl apply -f ./k8s 

## 5. [Install Loki Stack](https://hub.helm.sh/charts/loki/loki-stack)
