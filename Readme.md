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

    Usefull command recap :
    kubectl create ns loki-stack
    helm upgrade --install loki --namespace=loki-stack loki/loki-stack --set grafana.enabled=true
    kubectl get secret --namespace loki-stack loki-grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
    kubectl port-forward --namespace loki-stack service/loki-grafana 3000:80

    Click on the link to have further explanations

## 6. Add 2 new Api to Query K8S

    /inside/k8s ==> Query k8s from service inside the cluster
    /outside/k8s ==> Query from service outside the cluster
