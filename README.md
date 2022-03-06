# 模块三作业

## 构建

docker build -t simdiak/http_server .

docker push simdiak/http_server:latest

## 下载

docker pull simdiak/http_server:latest

## 运行

docker run simdiak/http_server:latest

## 查看IP

```
$ docker inspect 961a8753d70b|grep Pid
            "Pid": 9223,
            "PidMode": "",
            "PidsLimit": null,
$ sudo nsenter -t 9223 -n ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
8290: eth0@if8291: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1400 qdisc noqueue state UP group default
    link/ether 02:42:ac:11:00:04 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.4/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```

