# 模块十二作业

## 安装istio

略

## 导入证书到istio-system

因为没有外网IP，所以手动加入证书
```
# kubectl create secret -n istio-system tls x-ntes-com --cert ssl/x.test.com.crt --key ssl/x.test.com.key
```

## 建立istio服务和gateway

```
# kubectl create -f istio-specs.yaml
```

