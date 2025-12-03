# 快速开始

本文档帮助你快速搭建 OceanEngine 广告管理平台的开发环境。

## 环境要求

### 必需
- Go 1.21+
- Node.js 18+
- MySQL 8.0+
- Redis 7.0+
- Git

### 推荐
- Docker & Docker Compose
- pnpm (Node 包管理器)
- VS Code / GoLand

## 获取源码

```bash
git clone https://github.com/your-org/oceanengine.git
cd oceanengine
```

## 方式一：Docker Compose 启动（推荐）

### 1. 启动所有服务

```bash
# 启动开发环境
docker-compose up -d

# 查看日志
docker-compose logs -f api
```

### 2. 初始化数据库

```bash
# 执行数据库迁移
docker-compose exec api ./server migrate

# 初始化基础数据
docker-compose exec api ./server seed
```

### 3. 访问服务

- 前端页面: http://localhost:3000
- 后端 API: http://localhost:8080
- API 文档: http://localhost:8080/swagger/index.html

## 方式二：本地启动

### 1. 后端服务

#### 安装依赖

```bash
cd backend
go mod download
```

#### 配置文件

```bash
# 复制配置模板
cp config/config.example.yaml config/config.yaml

# 编辑配置
vim config/config.yaml
```

配置示例：

```yaml
server:
  mode: debug
  port: 8080

database:
  host: localhost
  port: 3306
  database: oceanengine
  username: root
  password: your_password

redis:
  addr: localhost:6379
  password: ""
  db: 0

jwt:
  secret: your-jwt-secret-key
  access_expire: 2h
  refresh_expire: 168h

oceanengine:
  app_id: your-app-id
  secret: your-secret
```

#### 创建数据库

```bash
# 登录 MySQL
mysql -u root -p

# 创建数据库
CREATE DATABASE oceanengine DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

#### 运行迁移

```bash
# 运行数据库迁移
go run cmd/server/main.go migrate

# 初始化基础数据（可选）
go run cmd/server/main.go seed
```

#### 启动服务

```bash
# 开发模式运行
go run cmd/server/main.go

# 或编译后运行
go build -o server cmd/server/main.go
./server
```

### 2. 前端服务

#### 安装依赖

```bash
cd frontend

# 使用 pnpm（推荐）
pnpm install

# 或使用 npm
npm install
```

#### 配置环境变量

```bash
# 复制环境变量模板
cp .env.example .env.development

# 编辑配置
vim .env.development
```

环境变量示例：

```env
VITE_API_BASE_URL=http://localhost:8080/api/v1
VITE_APP_TITLE=OceanEngine 广告管理平台
```

#### 启动开发服务器

```bash
# 启动开发服务器
pnpm dev

# 或
npm run dev
```

访问 http://localhost:5173 查看前端页面。

## 验证安装

### 1. 检查后端服务

```bash
# 健康检查
curl http://localhost:8080/health

# 预期响应
{"status":"ok","timestamp":"2024-01-15T10:30:00Z"}
```

### 2. 检查数据库连接

```bash
# 通过 API 检查
curl http://localhost:8080/api/v1/ping

# 预期响应
{"code":0,"message":"pong"}
```

### 3. 测试登录

```bash
# 获取验证码
curl http://localhost:8080/api/v1/auth/captcha

# 登录（使用初始账号 admin/admin123）
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin123",
    "captcha_id": "xxx",
    "captcha_code": "xxx"
  }'
```

## 初始账号

系统初始化后会创建以下账号：

| 用户名 | 密码 | 角色 | 说明 |
|--------|------|------|------|
| admin | admin123 | 超级管理员 | 拥有所有权限 |
| test | test123 | 普通用户 | 测试账号 |

**注意**: 首次登录后请立即修改默认密码！

## 项目结构

```
oceanengine/
├── backend/                 # 后端代码
│   ├── cmd/                 # 入口程序
│   ├── config/              # 配置文件
│   ├── internal/            # 内部代码
│   └── pkg/                 # 公共包
│
├── frontend/                # 前端代码
│   ├── src/
│   │   ├── api/             # API 接口
│   │   ├── components/      # 组件
│   │   ├── views/           # 页面
│   │   ├── stores/          # 状态管理
│   │   └── router/          # 路由
│   └── public/              # 静态资源
│
├── docs/                    # 文档
│   ├── backend/             # 后端文档
│   └── frontend/            # 前端文档
│
└── deploy/                  # 部署配置
    ├── docker/              # Docker 配置
    └── k8s/                 # Kubernetes 配置
```

## 常用命令

### 后端

```bash
# 运行测试
go test ./...

# 代码检查
golangci-lint run

# 生成 Swagger 文档
swag init -g cmd/server/main.go

# 构建生产版本
go build -ldflags="-w -s" -o server cmd/server/main.go
```

### 前端

```bash
# 开发服务器
pnpm dev

# 构建生产版本
pnpm build

# 预览构建结果
pnpm preview

# 代码检查
pnpm lint

# 类型检查
pnpm typecheck
```

## 开发工具配置

### VS Code 推荐插件

```json
{
  "recommendations": [
    "golang.go",
    "vue.volar",
    "bradlc.vscode-tailwindcss",
    "dbaeumer.vscode-eslint",
    "esbenp.prettier-vscode"
  ]
}
```

### GoLand 配置

1. 设置 GOROOT 和 GOPATH
2. 启用 Go Modules
3. 配置 File Watchers 自动格式化

## 常见问题

### 1. 数据库连接失败

检查 MySQL 服务是否启动，配置文件中的连接信息是否正确。

```bash
# 检查 MySQL 服务
systemctl status mysql

# 测试连接
mysql -h localhost -u root -p
```

### 2. Redis 连接失败

```bash
# 检查 Redis 服务
redis-cli ping
```

### 3. 前端代理问题

确保 `vite.config.ts` 中的代理配置正确：

```typescript
export default defineConfig({
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
})
```

### 4. 端口冲突

修改配置文件中的端口，或停止占用端口的服务：

```bash
# 查找占用端口的进程
lsof -i :8080

# 杀死进程
kill -9 <PID>
```

## 下一步

- 阅读 [开发指南](./development-guide.md) 了解开发规范
- 查看 [API 文档](./api-reference.md) 了解接口详情
- 参考 [后端文档](./backend/00-overview.md) 了解后端架构
- 参考 [前端文档](./frontend/00-overview.md) 了解前端架构
