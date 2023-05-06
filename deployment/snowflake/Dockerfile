# 启动编译环境
FROM golang:1.20-alpine AS builder

# 配置编译环境
RUN go env -w GO111MODULE=auto
RUN go env -w GOPROXY=https://goproxy.cn,direct

# 拷贝源码到镜像中
COPY . /app

# 编译
# 禁用CGO缩小镜像体积并提高程序可移植性，设置GOOS跨平台编译
WORKDIR /app/server/snowflake/
RUN CGO_ENABLED=0 GOOS=linux go build -o snowflake .

FROM alpine:3.13
COPY --from=builder /app/server/snowflake/snowflake /snowflake

# 暴露容器端口, 并不真正暴露, 还需要运行时使用 -p 指定
# 帮助其他人了解容器内部的应用程序需要监听哪些端口
EXPOSE 7010

# 设置服务入口
ENTRYPOINT ["/snowflake"]