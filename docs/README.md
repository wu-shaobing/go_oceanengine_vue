# OceanEngine 广告管理平台文档

## 项目简介

OceanEngine 广告管理平台是基于巨量引擎 Marketing API 开发的广告投放管理系统，提供广告主管理、广告创建、数据报表等一站式广告管理功能。

## 系统架构

```
┌─────────────────────────────────────────────────────────────────────────┐
│                              用户浏览器                                   │
└─────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                           Nginx (反向代理)                                │
│                    - 静态资源 (前端)                                      │
│                    - API 代理 (/api -> 后端)                              │
└─────────────────────────────────────────────────────────────────────────┘
                                    │
            ┌───────────────────────┴───────────────────────┐
            │                                               │
            ▼                                               ▼
┌─────────────────────────┐                    ┌─────────────────────────┐
│       Frontend          │                    │        Backend          │
│    (Vue 3 + Vite)       │                    │      (Go + Gin)         │
│                         │                    │                         │
│  - 广告主管理            │                    │  - RESTful API          │
│  - 广告系列管理          │                    │  - 业务逻辑处理          │
│  - 数据报表展示          │                    │  - SDK 调用封装          │
│  - 素材管理              │                    │  - 数据持久化            │
└─────────────────────────┘                    └─────────────────────────┘
                                                            │
                    ┌───────────────────────────────────────┴───────────┐
                    │                       │                           │
                    ▼                       ▼                           ▼
          ┌─────────────────┐     ┌─────────────────┐       ┌─────────────────┐
          │     MySQL       │     │     Redis       │       │ Ocean Engine    │
          │   (数据存储)     │     │  (缓存/队列)    │       │     API         │
          └─────────────────┘     └─────────────────┘       └─────────────────┘
```

## 技术栈

### 前端
- **框架**: Vue 3 + TypeScript
- **构建**: Vite 5.x
- **UI**: TailwindCSS + 自定义组件
- **状态**: Pinia
- **路由**: Vue Router 4
- **图表**: Chart.js + vue-chartjs
- **HTTP**: Axios

### 后端
- **语言**: Go 1.21+
- **框架**: Gin 1.9
- **ORM**: GORM 1.25
- **数据库**: MySQL 8.0
- **缓存**: Redis 7.0
- **认证**: JWT + Casbin
- **日志**: Zap
- **配置**: Viper

## 文档目录

### 后端文档
- [00-overview.md](./backend/00-overview.md) - 后端开发总览
- [01-directory-structure.md](./backend/01-directory-structure.md) - 项目目录结构
- [02-database-design.md](./backend/02-database-design.md) - 数据库设计
- [03-api-design.md](./backend/03-api-design.md) - API 接口设计
- [04-auth-design.md](./backend/04-auth-design.md) - 认证授权设计
- [05-sdk-integration.md](./backend/05-sdk-integration.md) - SDK 集成方案
- [06-middleware-design.md](./backend/06-middleware-design.md) - 中间件设计
- [07-error-handling.md](./backend/07-error-handling.md) - 错误处理方案
- [08-cache-design.md](./backend/08-cache-design.md) - 缓存设计
- [09-logging-design.md](./backend/09-logging-design.md) - 日志设计
- [10-deployment.md](./backend/10-deployment.md) - 部署方案

### 前端文档
- [00-overview.md](./frontend/00-overview.md) - 前端开发总览
- [01-project-structure.md](./frontend/01-project-structure.md) - 项目结构说明
- [02-component-design.md](./frontend/02-component-design.md) - 组件设计规范
- [03-state-management.md](./frontend/03-state-management.md) - 状态管理
- [04-api-integration.md](./frontend/04-api-integration.md) - API 集成
- [05-routing.md](./frontend/05-routing.md) - 路由设计
- [06-style-guide.md](./frontend/06-style-guide.md) - 样式规范

### 其他文档
- [getting-started.md](./getting-started.md) - 快速开始
- [development-guide.md](./development-guide.md) - 开发指南
- [api-reference.md](./api-reference.md) - API 参考

## 核心功能模块

### 1. 广告主管理
- 广告主账户列表与详情
- 广告主授权与解绑
- 账户余额与资金流水
- 资质信息管理

### 2. 广告系列管理
- 广告系列 CRUD
- 预算与出价设置
- 投放状态控制
- 批量操作

### 3. 广告组管理
- 广告组创建与编辑
- 定向设置（地域、人群、兴趣等）
- 出价策略配置
- 投放时段设置

### 4. 创意管理
- 创意模板选择
- 素材上传与管理
- 创意预览
- 创意审核状态

### 5. 素材库
- 图片素材管理
- 视频素材管理
- 素材分类与标签
- 素材复用

### 6. 数据报表
- 实时数据看板
- 多维度报表（广告主/系列/创意）
- 数据趋势分析
- 报表导出

### 7. 系统管理
- 用户管理
- 角色权限
- 操作日志
- 系统配置

## 开发环境要求

### 后端
- Go 1.21+
- MySQL 8.0+
- Redis 7.0+
- Docker (可选)

### 前端
- Node.js 18+
- pnpm 8+ (推荐) 或 npm

## 快速开始

详见 [getting-started.md](./getting-started.md)

## 贡献指南

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交变更 (`git commit -m 'feat: add amazing feature'`)
4. 推送分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

## 许可证

MIT License
