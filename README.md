# 模块八作业

## 更新代码并打包上传

* `http_server.go`增加了信号处理和运行标记，用于优雅启动退出和探活
```
docker build -t simdiak/http_server .
docker push simdiak/http_server:latest
```

## 登录dockerhub

```
kubectl create secret docker-registry dockerhub \
  --docker-server=hub.docker.com \
  --docker-username=simdiak \
  --docker-password=5b73709a-****-****-****-84bbd77ec402 \
  --docker-email=s***k@163.com
```

## 建立环境变量

```
kubectl create cm httpserver-env \
  --from-literal=HOST=0.0.0.0 \
  --from-literal=PORT=8080
```

## 部署

```
kubectl create -f m8.yaml
```

## 查看输出

```
# kubectl get po -w
NAME                                READY   STATUS     RESTARTS   AGE
httpserver-cd4f4d978-fvjrl          0/1     Init:0/1   0          4s
httpserver-cd4f4d978-fvjrl          0/1     PodInitializing   0          11s
httpserver-cd4f4d978-fvjrl          0/1     Running           0          14s
httpserver-cd4f4d978-fvjrl          1/1     Running           0          20s
# kubectl get deploy
NAME               READY   UP-TO-DATE   AVAILABLE   AGE
httpserver         1/1     1            1           5m
# kubectl describe po httpserver-cd4f4d978-fvjrl
...
Events:
  Type     Reason     Age                     From               Message
  ----     ------     ----                    ----               -------
  Normal   Scheduled  5m47s                   default-scheduler  Successfully assigned default/httpserver-cd4f4d978-fvjrl to gdl-test23-30002.i.nease.net
  Normal   Pulling    5m46s                   kubelet            Pulling image "simdiak/http_server:latest"
  Normal   Pulled     5m36s                   kubelet            Successfully pulled image "simdiak/http_server:latest" in 9.396664248s
  Normal   Created    5m36s                   kubelet            Created container init-httpserver
  Normal   Started    5m36s                   kubelet            Started container init-httpserver
  Normal   Pulling    5m36s                   kubelet            Pulling image "simdiak/http_server:latest"
  Normal   Pulled     5m33s                   kubelet            Successfully pulled image "simdiak/http_server:latest" in 2.42900509s
  Normal   Created    5m33s                   kubelet            Created container httpserver
  Normal   Started    5m33s                   kubelet            Started container httpserver
```

