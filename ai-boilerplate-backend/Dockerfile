##################################
# 第一阶段：构建GO可执行文件
##################################

# 使用官方的 Go 基础镜像作为构建环境
FROM golang:1.24-alpine AS builder
# 构建参数
ARG APP_VERSION=unknown
ARG APP_ENV=production
# 设置环境变量
ENV GOPROXY="https://goproxy.cn,direct"
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
# 设置工作目录
WORKDIR /src
# 复制项目源代码到工作目录
COPY . /src
# 构建可执行文件
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    mkdir -p bin/ && \
    go mod download && \
    go build -ldflags "-w -s -extldflags \"-static\" -X main.Version=${APP_VERSION}" -tags musl -o ./bin/server ./cmd/server

##################################
# 第二阶段：创建最终的运行时镜像
##################################

# 使用 Alpine 作为基础镜像，因为它非常轻量级
FROM alpine:3.20

# 重新声明构建参数，使其在当前阶段可用（不设默认值，使用构建时传入的值）
ARG APP_ENV

# 安装必要的证书（如果应用程序需要进行 HTTPS 请求）
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /app

# 从第一阶段的构建结果中复制可执行文件到当前工作目录
COPY --from=builder /src/bin/server /app/server

# 拷贝配置文件
COPY --from=builder /src/configs/ /app/configs

# 暴露服务端口
EXPOSE 8000 9000

# 设置环境变量
ENV APP_ENV=${APP_ENV}

# 设置容器启动时执行的命令
CMD ["/app/server", "-conf", "/app/configs"]