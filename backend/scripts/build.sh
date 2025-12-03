#!/bin/bash

# OceanEngine Backend 构建脚本

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 项目根目录
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$PROJECT_ROOT"

# 版本信息
VERSION="${VERSION:-$(git describe --tags --always --dirty 2>/dev/null || echo 'dev')}"
BUILD_TIME=$(date -u '+%Y-%m-%d_%H:%M:%S')
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo 'unknown')

# 输出目录
OUTPUT_DIR="${OUTPUT_DIR:-./bin}"

# 构建参数
LDFLAGS="-s -w -X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME} -X main.GitCommit=${GIT_COMMIT}"

echo -e "${GREEN}============================================${NC}"
echo -e "${GREEN}  OceanEngine Backend Build Script${NC}"
echo -e "${GREEN}============================================${NC}"
echo -e "${YELLOW}Version: ${VERSION}${NC}"
echo -e "${YELLOW}Build Time: ${BUILD_TIME}${NC}"
echo -e "${YELLOW}Git Commit: ${GIT_COMMIT}${NC}"
echo ""

# 创建输出目录
mkdir -p "$OUTPUT_DIR"

# 安装依赖
echo -e "${YELLOW}Installing dependencies...${NC}"
go mod download
go mod tidy

# 代码检查
echo -e "${YELLOW}Running go vet...${NC}"
go vet ./...

# 构建服务端
echo -e "${YELLOW}Building server...${NC}"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="$LDFLAGS" -o "$OUTPUT_DIR/server" ./cmd/server/main.go

# 构建迁移工具
echo -e "${YELLOW}Building migrate tool...${NC}"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="$LDFLAGS" -o "$OUTPUT_DIR/migrate" ./cmd/migrate/main.go

# 构建定时任务
echo -e "${YELLOW}Building task runner...${NC}"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="$LDFLAGS" -o "$OUTPUT_DIR/task" ./cmd/task/main.go

# 复制配置文件
echo -e "${YELLOW}Copying config files...${NC}"
cp -r config "$OUTPUT_DIR/"

echo ""
echo -e "${GREEN}============================================${NC}"
echo -e "${GREEN}  Build completed successfully!${NC}"
echo -e "${GREEN}============================================${NC}"
echo -e "Output directory: ${OUTPUT_DIR}"
ls -la "$OUTPUT_DIR"
