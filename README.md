# 模块十作业

## 更新代码并打包上传

* `http_server.go`增加了延时和metrics处理API
```
docker build -t simdiak/httpserver .
docker push simdiak/httpserver:latest
```

## 更新了部署配置

* m8.yaml

## 安装prometheus

```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install kube-prometheus-stack prometheus-community/kube-prometheus-stack --create-namespace -n prometheus-stack
```

## 配置服务发现

```
kubectl create secret generic additional-configs --from-file=prometheus-additional.yaml -n prometheus-stack
```

## 配置RBAC

```
kubectl create -f prometheus-rbac.yaml
```
