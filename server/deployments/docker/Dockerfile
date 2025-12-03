# 构建阶段
FROM golang:1.21-alpine AS builder

# 安装必要工具
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /build

# 复制整个项目 (包含SDK和server)
COPY . .

# 进入server目录构建
WORKDIR /build/server

# 下载依赖
RUN go mod download

# 构建二进制文件
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o /build/api-server \
    ./cmd/api

# 运行阶段
FROM alpine:3.18

# 安装ca证书和时区数据
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /build/api-server .

# 设置时区
ENV TZ=Asia/Shanghai

# 暴露端口
EXPOSE 8080

# 健康检查
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# 运行
CMD ["./api-server"]
