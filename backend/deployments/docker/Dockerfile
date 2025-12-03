# 构建阶段
FROM golang:1.21-alpine AS builder

# 设置环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 设置工作目录
WORKDIR /build

# 安装基础依赖
RUN apk add --no-cache git ca-certificates tzdata

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN go build -ldflags="-s -w" -o server ./cmd/server/main.go
RUN go build -ldflags="-s -w" -o migrate ./cmd/migrate/main.go
RUN go build -ldflags="-s -w" -o task ./cmd/task/main.go

# 运行阶段
FROM alpine:latest

# 安装依赖
RUN apk --no-cache add ca-certificates tzdata

# 设置时区
ENV TZ=Asia/Shanghai

# 创建工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /build/server /app/
COPY --from=builder /build/migrate /app/
COPY --from=builder /build/task /app/

# 复制配置文件
COPY config/ /app/config/

# 创建日志目录
RUN mkdir -p /app/logs

# 暴露端口
EXPOSE 8080

# 启动命令
CMD ["./server", "-config", "config/settings.yml"]
