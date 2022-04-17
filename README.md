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
  Normal   Scheduled  5m47s                   default-scheduler  Successfully assigned default/httpserver-cd4f4d978-fvjrl to xdl-test23-30002.i.xtest.net
  Normal   Pulling    5m46s                   kubelet            Pulling image "simdiak/http_server:latest"
  Normal   Pulled     5m36s                   kubelet            Successfully pulled image "simdiak/http_server:latest" in 9.396664248s
  Normal   Created    5m36s                   kubelet            Created container init-httpserver
  Normal   Started    5m36s                   kubelet            Started container init-httpserver
  Normal   Pulling    5m36s                   kubelet            Pulling image "simdiak/http_server:latest"
  Normal   Pulled     5m33s                   kubelet            Successfully pulled image "simdiak/http_server:latest" in 2.42900509s
  Normal   Created    5m33s                   kubelet            Created container httpserver
  Normal   Started    5m33s                   kubelet            Started container httpserver
```

## 部署service

```
# kubectl expose deploy httpserver --port=8080 --target-port=8080 --type=LoadBalancer
```

## 增加externalIP使内网可访问

```
# kubectl edit svc httpserver
...
  externalIPs:
  - 10.246.46.103
...

# kubectl get svc httpserver -oyaml
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: "2022-04-16T15:45:19Z"
  name: httpserver
  namespace: default
  resourceVersion: "3660297"
  uid: 02c46a68-202a-46cb-967b-9ab4fc7619b8
spec:
  allocateLoadBalancerNodePorts: true
  clusterIP: 10.100.25.24
  clusterIPs:
  - 10.100.25.24
  externalIPs:
  - 10.246.46.103
  externalTrafficPolicy: Cluster
  internalTrafficPolicy: Cluster
  ipFamilies:
  - IPv4
  ipFamilyPolicy: SingleStack
  ports:
  - nodePort: 31527
    port: 8080
    protocol: TCP
    targetPort: 8080
  selector:
    app: httpserver
  sessionAffinity: None
  type: LoadBalancer
status:
  loadBalancer: {}
```

## 查看结果

```
# kubectl get svc
NAME               TYPE           CLUSTER-IP     EXTERNAL-IP     PORT(S)           AGE
httpserver         LoadBalancer   10.100.25.24   10.246.46.103   8080:31527/TCP    15h
kubernetes         ClusterIP      10.96.0.1      <none>          443/TCP           18d
```
* 测试访问
```
$ curl http://10.246.46.103:8080
Hello, world!
```

## 安装ingress

```
# helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
"ingress-nginx" has been added to your repositories
# helm install ingress-nginx ingress-nginx/ingress-nginx --create-namespace --namespace ingress
NAME: ingress-nginx
LAST DEPLOYED: Sun Apr 17 13:56:36 2022
NAMESPACE: ingress
STATUS: deployed
REVISION: 1
TEST SUITE: None
...
```

## 导入tls证书

* 因为手上有一个可用的外网证书，就直接拿来用了
```
# kubectl create secret tls x-ntes-com --cert ssl/x.ntes.com.crt --key ssl/x.ntes.com.key
secret/x-ntes-com created
# kubectl get secret -owide
NAME                  TYPE                                  DATA   AGE
default-token-xrrq4   kubernetes.io/service-account-token   3      18d
dockerhub             kubernetes.io/dockerconfigjson        1      6d22h
x-ntes-com            kubernetes.io/tls                     2      2m36s
```

## 配置ingress

```
# kubectl create -f m8-ingress.yaml
ingress.networking.k8s.io/httpserver-tls created
# kubectl get ing
NAME             CLASS   HOSTS             ADDRESS   PORTS     AGE
httpserver-tls   nginx   z.x.ntes.com                80, 443   6s
```

## 这样在子网另一台机就可以访问它了

```
$ curl https://z.x.ntes.com --resolve z.x.ntes.com:443:10.246.46.103
Hello, world!
```

## 调整deploy使其高可用

* 调整为2个副本
```
-  replicas: 1
+  replicas: 2
```
* 增加节点反亲和性
```
+      affinity:
+        podAntiAffinity:
+          requiredDuringSchedulingIgnoredDuringExecution:
+          - labelSelector:
+              matchExpressions:
+              - key: app
+                operator: In
+                values:
+                - httpserver
+            topologyKey: "kubernetes.io/hostname"
```
* 更新它
```
# kubectl apply -f m8.yaml
Warning: resource deployments/httpserver is missing the kubectl.kubernetes.io/last-applied-configuration annotation which is required by kubectl apply. kubectl apply should only be used on resources created declaratively by either kubectl create --save-config or kubectl apply. The missing annotation will be patched automatically.
deployment.apps/httpserver configured
```

## 查看结果

```
# kubectl get deploy -owide
NAME               READY   UP-TO-DATE   AVAILABLE   AGE   CONTAINERS   IMAGES                       SELECTOR
httpserver         2/2     2            2           7d    httpserver   simdiak/http_server:latest   app=httpserver
# kubectl get po -owide
NAME                                READY   STATUS    RESTARTS   AGE    IP               NODE                           NOMINATED NODE   READINESS GATES
httpserver-6b88d6bf87-58qd2         1/1     Running   0          16s    10.111.194.91    xdl-test23-30002.i.xtest.net   <none>           <none>
httpserver-6b88d6bf87-729df         1/1     Running   0          55s    10.111.186.211   xdl-test24-30002.i.xtest.net   <none>           <none>
```
