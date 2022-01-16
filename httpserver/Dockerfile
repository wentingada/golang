#得到最新的golang docker 镜像
FROM golang:latest
#在容器内部创建一个目录来存储http应用，并设为工作目录
RUN mkdir -p /go/src/httpserver
WORKDIR /go/src/httpserver
ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
    GO111MODULE=on \
    CGO_ENABLED=1

#复制程序到容器
COPY . /go/src/httpserver
#下载安装第三方依赖库到容器中
#gin：
#wrapper：
RUN go get github.com/codegangsta/gin
RUN go-wrapper download
RUN go-wrapper install
#设置PORT环境变量
ENV PORT 8080
#暴露8080给主机
EXPOSE 3030
#启动容器时运行Gin
#CMD gin run
#告诉Docker启动容器运行的命令
CMD ["go-wrapper","run"]
