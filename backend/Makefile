.PHONY: all build run clean test lint fmt mod help

# 变量
APP_NAME := oceanengine-backend
MAIN_FILE := cmd/server/main.go
BUILD_DIR := build
GO := go

# 默认目标
all: build

# 构建
build:
	@echo "构建 $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
	@echo "构建完成: $(BUILD_DIR)/$(APP_NAME)"

# 运行
run:
	@echo "启动服务..."
	$(GO) run $(MAIN_FILE)

# 运行 (开发模式，使用 air 热重载)
dev:
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "请先安装 air: go install github.com/air-verse/air@latest"; \
		exit 1; \
	fi

# 清理
clean:
	@echo "清理构建文件..."
	@rm -rf $(BUILD_DIR)
	@echo "清理完成"

# 测试
test:
	@echo "运行测试..."
	$(GO) test -v ./...

# 代码检查
lint:
	@echo "代码检查..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run ./...; \
	else \
		echo "请先安装 golangci-lint"; \
		exit 1; \
	fi

# 格式化代码
fmt:
	@echo "格式化代码..."
	$(GO) fmt ./...
	@echo "格式化完成"

# 整理依赖
mod:
	@echo "整理依赖..."
	$(GO) mod tidy
	@echo "依赖整理完成"

# 下载依赖
mod-download:
	@echo "下载依赖..."
	$(GO) mod download
	@echo "依赖下载完成"

# Docker 构建
docker-build:
	@echo "构建 Docker 镜像..."
	docker build -t $(APP_NAME):latest .

# Docker 运行
docker-run:
	@echo "运行 Docker 容器..."
	docker run -d --name $(APP_NAME) -p 8080:8080 $(APP_NAME):latest

# Docker Compose 启动
compose-up:
	@echo "启动服务..."
	docker-compose up -d

# Docker Compose 停止
compose-down:
	@echo "停止服务..."
	docker-compose down

# 数据库迁移 (如果使用 migrate 工具)
migrate-up:
	@echo "执行数据库迁移..."
	@if command -v migrate > /dev/null; then \
		migrate -path ./migrations -database "mysql://root:password@tcp(localhost:3306)/oceanengine" up; \
	else \
		echo "请先安装 migrate: go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; \
	fi

migrate-down:
	@echo "回滚数据库迁移..."
	@if command -v migrate > /dev/null; then \
		migrate -path ./migrations -database "mysql://root:password@tcp(localhost:3306)/oceanengine" down 1; \
	else \
		echo "请先安装 migrate"; \
	fi

# 生成 Swagger 文档
swagger:
	@echo "生成 Swagger 文档..."
	@if command -v swag > /dev/null; then \
		swag init -g $(MAIN_FILE) -o ./docs/swagger; \
	else \
		echo "请先安装 swag: go install github.com/swaggo/swag/cmd/swag@latest"; \
	fi

# 帮助
help:
	@echo "可用命令:"
	@echo "  make build         - 构建应用"
	@echo "  make run           - 运行应用"
	@echo "  make dev           - 开发模式 (热重载)"
	@echo "  make clean         - 清理构建文件"
	@echo "  make test          - 运行测试"
	@echo "  make lint          - 代码检查"
	@echo "  make fmt           - 格式化代码"
	@echo "  make mod           - 整理依赖"
	@echo "  make mod-download  - 下载依赖"
	@echo "  make docker-build  - 构建 Docker 镜像"
	@echo "  make docker-run    - 运行 Docker 容器"
	@echo "  make compose-up    - Docker Compose 启动"
	@echo "  make compose-down  - Docker Compose 停止"
	@echo "  make migrate-up    - 执行数据库迁移"
	@echo "  make migrate-down  - 回滚数据库迁移"
	@echo "  make swagger       - 生成 Swagger 文档"
	@echo "  make help          - 显示帮助"
