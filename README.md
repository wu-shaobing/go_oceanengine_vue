# OceanEngine 广告管理平台

基于巨量引擎 Marketing API 的广告投放管理系统，提供广告主管理、广告创建、数据报表等一站式服务。

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Vue Version](https://img.shields.io/badge/Vue-3.4+-4FC08D?style=flat&logo=vue.js)](https://vuejs.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## ✨ 核心特性

- 🎯 **广告主管理** - OAuth 授权、账户余额、资金流水
- 📊 **广告系列管理** - 创建/编辑/暂停广告系列、预算控制
- 🎨 **创意管理** - 素材上传、创意模板、审核状态
- 📈 **数据报表** - 实时数据看板、多维度报表、趋势分析
- 🛒 **千川电商** - 巨量千川全域推广、商品广告
- ⭐ **星图达人** - 星图任务管理、达人营销
- 🏪 **本地推** - 本地生活推广、门店管理
- 🔐 **权限管理** - 基于角色的访问控制（RBAC）

## 📦 技术栈

### 前端
- **框架**: Vue 3.4+ (Composition API)
- **构建**: Vite 7.2+
- **语言**: TypeScript 5.3+
- **状态**: Pinia
- **路由**: Vue Router 4
- **样式**: TailwindCSS 3.4+
- **图表**: Chart.js + vue-chartjs
- **HTTP**: Axios 1.6+

### 后端
- **语言**: Go 1.21+
- **框架**: Gin 1.9
- **ORM**: GORM 1.30+
- **数据库**: MySQL 8.0+
- **缓存**: Redis 7.0+
- **认证**: JWT (golang-jwt/jwt)
- **日志**: Zap + Lumberjack
- **配置**: Viper

## 🚀 快速开始

### 环境要求

- Go 1.21+
- Node.js 18+
- MySQL 8.0+
- Redis 7.0+
- Docker & Docker Compose (可选)

### 1. 克隆项目

```bash
git clone <repository-url>
cd oceanengine
```

### 2. 后端启动

#### 方式一：Docker Compose（推荐）

```bash
cd backend

# 创建环境变量文件
cp .env.example .env
# 编辑 .env 填入必要的密钥和配置

# 启动所有服务（MySQL + Redis + 后端）
docker compose up -d

# 查看日志
docker compose logs -f app
```

#### 方式二：本地运行

```bash
cd backend

# 安装依赖
go mod download

# 创建配置文件
cp config/settings.example.yml config/settings.yml
# 编辑 config/settings.yml 或设置环境变量

# 运行数据库迁移
go run cmd/migrate/main.go

# 启动服务
go run cmd/server/main.go
```

后端默认运行在 `http://localhost:8080`

### 3. 前端启动

```bash
cd frontend

# 安装依赖
npm install

# 开发模式
npm run dev

# 构建生产
npm run build
```

前端默认运行在 `http://localhost:3000`

## 📁 项目结构

```
oceanengine/
├── backend/              # Go 后端
│   ├── cmd/             # 命令行入口
│   │   ├── server/      # HTTP 服务器
│   │   ├── migrate/     # 数据库迁移
│   │   └── task/        # 定时任务
│   ├── config/          # 配置管理
│   ├── internal/        # 内部包
│   │   ├── app/        # 业务模块
│   │   ├── middleware/ # 中间件
│   │   └── router/     # 路由
│   ├── pkg/            # 可复用包
│   │   ├── auth/       # 认证
│   │   ├── cache/      # 缓存
│   │   ├── database/   # 数据库
│   │   └── oceanengine/ # SDK
│   ├── scripts/        # 脚本
│   └── deployments/    # 部署配置
│
├── frontend/           # Vue 前端
│   ├── src/
│   │   ├── api/       # API 调用
│   │   ├── components/ # 组件
│   │   ├── composables/ # 组合式函数
│   │   ├── router/    # 路由配置
│   │   ├── stores/    # Pinia 状态
│   │   ├── views/     # 页面视图
│   │   └── utils/     # 工具函数
│   └── public/        # 静态资源
│
├── docs/              # 项目文档
│   ├── backend/       # 后端文档
│   └── frontend/      # 前端文档
│
└── sdk/               # 原始 SDK（参考）
```

## 🔧 配置说明

### 后端环境变量

关键环境变量（在 `backend/.env` 中配置）：

```bash
# 数据库
DB_HOST=localhost
DB_PORT=3306
DB_USER=oceanengine
DB_PASSWORD=<your-password>
DB_NAME=oceanengine

# Redis
REDIS_PASSWORD=<your-redis-password>

# JWT
JWT_SECRET_KEY=<your-jwt-secret>

# 巨量引擎
OCEAN_APP_ID=<your-app-id>
OCEAN_SECRET=<your-app-secret>

# CORS（生产环境）
CORS_ALLOWED_ORIGINS=https://yourdomain.com
```

### 生成安全密钥

```bash
# JWT Secret（32 字节）
openssl rand -base64 32

# 数据库密码（24 字节）
openssl rand -base64 24
```

## 📚 文档

详细文档请查看 [`docs/`](./docs) 目录：

- [快速开始](./docs/getting-started.md)
- [开发指南](./docs/development-guide.md)
- [API 参考](./docs/api-reference.md)
- [后端文档](./docs/backend/)
- [前端文档](./docs/frontend/)

**上线检查**：
- [上线前检查清单](./PRE_LAUNCH_CHECKLIST_REPORT.md)
- [关键问题修复指南](./CRITICAL_FIXES.md)

## 🔐 安全特性

- ✅ JWT 身份认证
- ✅ bcrypt 密码加密
- ✅ CORS 跨域配置
- ✅ 安全响应头（X-Frame-Options, CSP 等）
- ✅ Rate Limiting 限流
- ✅ 环境变量敏感信息管理
- ✅ SQL 注入防护（GORM）
- ✅ XSS 防护

## 🧪 测试

```bash
# 后端测试
cd backend
go test ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# 前端测试
cd frontend
npm run test
npm run test:ui
```

## 📦 部署

### Docker 部署

```bash
cd backend
docker compose -f docker-compose.yml up -d
```

### Kubernetes 部署

```bash
cd backend/deployments/kubernetes
kubectl apply -f .
```

## 🤝 贡献

欢迎贡献！请遵循以下步骤：

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'feat: add amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

## 📄 许可证

本项目采用 [MIT](LICENSE) 许可证。

## 🙏 致谢

- [巨量引擎 Marketing API](https://open.oceanengine.com/)
- [Gin Web Framework](https://gin-gonic.com/)
- [Vue.js](https://vuejs.org/)
- [GORM](https://gorm.io/)

## 📞 支持

- 📧 Email: support@example.com
- 📖 文档: [docs/](./docs)
- 🐛 问题反馈: [Issues](../../issues)

---

**⚠️ 上线前必读**：请务必查看 [PRE_LAUNCH_CHECKLIST_REPORT.md](./PRE_LAUNCH_CHECKLIST_REPORT.md) 完成所有安全检查！
