### go语言基础示例代码
作业二：
构建本地镜像
编写 Dockerfile httpserver 容器化
将镜像推送至 docker 官方镜像仓库
通过 docker 命令本地启动 httpserver
通过 nsenter 进入容器查看 IP 配置

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
