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
  
- 通过 docker 命令本地启动 httpserver
  - 后台运行
  >docker run -d -p :3030 --name server-v0.0 httpserver:0.0
  >>docker exec -it server-v0.0 /bin/sh
  - 交互运行
- 通过 nsenter 进入容器查看 IP 配置
