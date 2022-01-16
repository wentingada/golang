#在最新的golang docker 镜像中build
FROM golang:latest AS builder
#在容器内部创建一个目录来存储http应用，并设为工作目录
RUN mkdir -p /go/src/httpserver
WORKDIR /go/src/httpserver
ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
    GO111MODULE=on \
    CGO_ENABLED=0
#复制程序到容器
COPY . /go/src/httpserver
#编译
RUN GOPS=linux go build -installsuffix cgo -o httpserver main.go

#在新的镜像里面执行部署
FROM busybox AS executor
COPY --from=builder /go/src/httpserver/httpserver /httpserver

WORKDIR /httpserver
#暴露8080给主机
EXPOSE 3030

#告诉Docker启动容器运行的命令
ENTRYPOINT ["./httpserver"]
