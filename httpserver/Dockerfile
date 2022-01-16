#在最新的golang docker 镜像中build
FROM golang:latest AS builder
#在容器内部创建一个目录来存储http应用，并设为工作目录
WORKDIR /BUILD
ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
    GO111MODULE=off \
    CGO_ENABLED=0
#复制程序到容器
COPY . .
#编译
RUN go build -o httpserver 

#在新的镜像里面执行部署
FROM busybox AS executor
COPY --from=builder /BUILD/httpserver /

WORKDIR /
#暴露8000给主机
EXPOSE 8000

#告诉Docker启动容器运行的命令
ENTRYPOINT ["./httpserver"]
