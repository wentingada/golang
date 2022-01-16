### go语言基础示例代码
copy from https://github.com/cncamp/golang
use own httpserver code
use own dockerfile to:
- 构建本地镜像
- 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化
  >在Dockerfile目录下：
  >docker build -t httpserver:0.0 .
- 将镜像推送至 docker 官方镜像仓库
  >使用阿里云镜像仓库
  - 在阿里云个人实例页面创建新的命名空间
  - 用阿里云账号、密码命令行登录阿里云仓库:docker login --username=[xxx] registry.cn-hangzhou.aliyuncs.com
  - 将我要推送的镜像重新命名，也叫打tag:docker tag [image-id] [image全名：registry.cn-hangzhou.aliyuncs.com/[namespace]/[httpserver:0.0]]
  - 推送到阿里云镜像仓库：docker push [image全名]
  - 从阿里云镜像仓库下载：docker pull [image全名]
  
- 通过 docker 命令本地启动 httpserver
  - 后台运行
  >docker run -d -p 8000:8000 --name server-v0.0 httpserver:0.0
  >>docker exec -it server-v0.0 /bin/sh
  - 交互运行
- 通过 nsenter 进入容器查看 IP 配置
    - 获取容器pid
    >docker inspect server-v0.0 -f '{{.State.Pid}}'
    >>23069
    - 也可以通过ContainerID
    >docker inspect [ContainerID] -f '{{.State.Pid}}'
    >>23069
    - 进入容器网络空间，执行nsenter命令
    - nsenter -n -t[pid]
    >ifconfig
     >>eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        >>inet 172.17.0.5  netmask 255.255.0.0  broadcast 0.0.0.0
        >>inet6 fe80::42:acff:fe11:5  prefixlen 64  scopeid 0x20<link>
        >>ether 02:42:ac:11:00:05  txqueuelen 0  (Ethernet)
        >>RX packets 3  bytes 222 (222.0 B)
        >>RX errors 0  dropped 0  overruns 0  frame 0
        >>TX packets 8  bytes 648 (648.0 B)
        >>TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
    
    >ip addr
>>1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN qlen 1
    >>link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    >>inet 127.0.0.1/8 scope host lo
       >>valid_lft forever preferred_lft forever
    >>inet6 ::1/128 scope host 
       >>valid_lft forever preferred_lft forever
>>16: eth0@if17: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP 
    >>link/ether 02:42:ac:11:00:05 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    >>inet 172.17.0.5/16 scope global eth0
       >>valid_lft forever preferred_lft forever
    >>inet6 fe80::42:acff:fe11:5/64 scope link 
       >>valid_lft forever preferred_lft forever

>ping www.baidu.com
 >>PING www.a.shifen.com (180.101.49.12) 56(84) bytes of data.
 >>64 bytes from 180.101.49.12 (180.101.49.12): icmp_seq=1 ttl=49 time=10.1 ms
 >>64 bytes from 180.101.49.12 (180.101.49.12): icmp_seq=2 ttl=49 time=10.1 ms
