# OceanEngine 项目文件位置索引

> 快速定位项目中的关键文件

---

## 项目根目录

```
/Users/wushaobing911/Desktop/oceanengine/
```

---

## 一、后端 (Backend)

**根目录**: `backend/`

### 入口文件
| 文件 | 路径 | 说明 |
|------|------|------|
| 主程序 | `backend/cmd/server/main.go` | 应用启动入口 |
| 迁移工具 | `backend/cmd/migrate/main.go` | 数据库迁移 |
| 定时任务 | `backend/cmd/task/main.go` | 后台任务 |

### 配置文件
| 文件 | 路径 | 说明 |
|------|------|------|
| 主配置 | `backend/config/settings.yml` | 应用配置 |
| Go模块 | `backend/go.mod` | 依赖定义 |
| Docker | `backend/Dockerfile` | 镜像构建 |
| Makefile | `backend/Makefile` | 构建脚本 |

### 业务模块 (internal/app/)
| 模块 | 路径 | 说明 |
|------|------|------|
| 广告组 | `backend/internal/app/ad/` | ad api/dto/model/repository/service |
| 后台管理 | `backend/internal/app/admin/` | 用户/角色/菜单/日志 |
| 广告主 | `backend/internal/app/advertiser/` | advertiser 管理 |
| 人群定向 | `backend/internal/app/audience/` | audience 管理 |
| 广告系列 | `backend/internal/app/campaign/` | campaign 管理 |
| 创意管理 | `backend/internal/app/creative/` | creative 管理 |
| 素材管理 | `backend/internal/app/media/` | media 管理 |
| 数据报表 | `backend/internal/app/report/` | report 报表 |

### 中间件
| 文件 | 路径 | 说明 |
|------|------|------|
| CORS | `backend/internal/middleware/cors.go` | 跨域处理 |
| JWT | `backend/internal/middleware/jwt.go` | 认证中间件 |
| Logger | `backend/internal/middleware/logger.go` | 日志中间件 |
| RateLimit | `backend/internal/middleware/ratelimit.go` | 限流中间件 |
| Recovery | `backend/internal/middleware/recovery.go` | 异常恢复 |
| RequestID | `backend/internal/middleware/request_id.go` | 请求ID |

### 路由
| 文件 | 路径 | 说明 |
|------|------|------|
| 主路由 | `backend/internal/router/router.go` | 所有API路由定义 |

### 公共包 (pkg/)
| 包 | 路径 | 说明 |
|------|------|------|
| auth | `backend/pkg/auth/` | jwt.go, password.go |
| cache | `backend/pkg/cache/` | cache.go |
| database | `backend/pkg/database/` | mysql.go, redis.go |
| errcode | `backend/pkg/errcode/` | error.go, codes.go |
| logger | `backend/pkg/logger/` | logger.go |
| oceanengine | `backend/pkg/oceanengine/` | SDK封装 |
| response | `backend/pkg/response/` | response.go |
| utils | `backend/pkg/utils/` | pagination.go |

---

## 二、前端 (Frontend)

**根目录**: `frontend/`

### 入口文件
| 文件 | 路径 | 说明 |
|------|------|------|
| HTML入口 | `frontend/index.html` | SPA入口 |
| JS入口 | `frontend/src/main.ts` | Vue应用入口 |
| 根组件 | `frontend/src/App.vue` | 根组件 |

### 配置文件
| 文件 | 路径 | 说明 |
|------|------|------|
| Vite配置 | `frontend/vite.config.ts` | 构建配置 |
| TS配置 | `frontend/tsconfig.json` | TypeScript配置 |
| TailwindCSS | `frontend/tailwind.config.js` | 样式框架配置 |
| 依赖 | `frontend/package.json` | 包依赖 |
| 环境变量 | `frontend/.env` | 环境变量 |
| 开发环境 | `frontend/.env.development` | 开发环境变量 |
| 生产环境 | `frontend/.env.production` | 生产环境变量 |

### API接口 (src/api/)
| 文件 | 路径 | 说明 |
|------|------|------|
| 导出入口 | `frontend/src/api/index.ts` | API模块导出 |
| HTTP客户端 | `frontend/src/api/request.ts` | Axios封装 |
| 认证API | `frontend/src/api/auth.ts` | 登录/登出 |
| 广告主API | `frontend/src/api/advertiser.ts` | 广告主接口 |
| 广告系列API | `frontend/src/api/campaign.ts` | 广告系列接口 |
| 报表API | `frontend/src/api/report.ts` | 数据报表接口 |

### 状态管理 (src/stores/)
| 文件 | 路径 | 说明 |
|------|------|------|
| 入口 | `frontend/src/stores/index.ts` | Store导出 |
| 应用状态 | `frontend/src/stores/app.ts` | 应用全局状态 |
| 认证状态 | `frontend/src/stores/auth.ts` | 用户认证状态 |
| 广告主状态 | `frontend/src/stores/advertiser.ts` | 广告主状态 |

### 路由 (src/router/)
| 文件 | 路径 | 说明 |
|------|------|------|
| 入口 | `frontend/src/router/index.ts` | 路由配置入口 |
| 路由定义 | `frontend/src/router/routes.ts` | 路由表 (26KB) |
| 路由守卫 | `frontend/src/router/guards.ts` | 导航守卫 |

### 视图 (src/views/) - 28个模块
| 模块 | 路径 | 文件数 |
|------|------|--------|
| ad | `frontend/src/views/ad/` | 广告组 |
| advertiser | `frontend/src/views/advertiser/` | 广告主 |
| agent | `frontend/src/views/agent/` | 代理商 |
| audience | `frontend/src/views/audience/` | 人群 |
| auth | `frontend/src/views/auth/` | 认证 |
| campaign | `frontend/src/views/campaign/` | 广告系列 |
| clue | `frontend/src/views/clue/` | 线索 |
| creative | `frontend/src/views/creative/` | 创意 |
| dashboard | `frontend/src/views/dashboard/` | 看板 |
| douplus | `frontend/src/views/douplus/` | Dou+ |
| dpa | `frontend/src/views/dpa/` | DPA |
| enterprise | `frontend/src/views/enterprise/` | 企业 |
| error | `frontend/src/views/error/` | 错误页 |
| keyword | `frontend/src/views/keyword/` | 关键词 |
| local | `frontend/src/views/local/` | 本地推广 |
| material | `frontend/src/views/material/` | 素材 |
| media | `frontend/src/views/media/` | 媒体 |
| other | `frontend/src/views/other/` | 其他 |
| qianchuan | `frontend/src/views/qianchuan/` | 千川 |
| report | `frontend/src/views/report/` | 报表 |
| servemarket | `frontend/src/views/servemarket/` | 服务市场 |
| site | `frontend/src/views/site/` | 落地页 |
| star | `frontend/src/views/star/` | 星图 |
| system | `frontend/src/views/system/` | 系统 |
| tools | `frontend/src/views/tools/` | 工具 |
| track | `frontend/src/views/track/` | 追踪 |
| wallet | `frontend/src/views/wallet/` | 钱包 |
| workspace | `frontend/src/views/workspace/` | 工作台 |

### 组件 (src/components/)
| 目录 | 路径 | 说明 |
|------|------|------|
| 布局组件 | `frontend/src/components/layout/` | 布局 |
| 业务组件 | `frontend/src/components/business/` | 业务 |
| 通用组件 | `frontend/src/components/common/` | 通用 |

---

## 三、SDK

**根目录**: `sdk/`

### 核心文件
| 文件 | 路径 | 说明 |
|------|------|------|
| Go模块 | `sdk/go.mod` | 依赖定义 |
| 说明文档 | `sdk/README.md` | SDK说明 |
| 许可证 | `sdk/LICENSE` | 开源协议 |

### 主要目录
| 目录 | 路径 | 说明 |
|------|------|------|
| core | `sdk/marketing-api/core/` | 核心客户端 |
| enum | `sdk/marketing-api/enum/` | 枚举定义 |
| model | `sdk/marketing-api/model/` | 数据模型 |
| util | `sdk/marketing-api/util/` | 工具函数 |

---

## 四、Juliang API文档

**根目录**: `juliang/`

### 核心文件
| 文件 | 路径 | 说明 |
|------|------|------|
| SDK文档 | `juliang/README.md` | 完整API文档 |
| 页面生成器 | `juliang/generate-pages.js` | HTML生成脚本 |
| 页面配置 | `juliang/pages-config.json` | 页面配置 |
| 入口页面 | `juliang/index.html` | 文档入口 |
| 快速开始 | `juliang/QUICK_START.html` | 快速开始 |

---

## 五、项目文档

**根目录**: `docs/`

### 后端文档
| 序号 | 路径 | 说明 |
|------|------|------|
| 00 | `docs/backend/00-overview.md` | 总览 |
| 01 | `docs/backend/01-directory-structure.md` | 目录结构 |
| 02 | `docs/backend/02-database-design.md` | 数据库设计 |
| 03 | `docs/backend/03-api-design.md` | API设计 |
| 04 | `docs/backend/04-auth-design.md` | 认证设计 |
| 05 | `docs/backend/05-sdk-integration.md` | SDK集成 |
| 06 | `docs/backend/06-middleware-design.md` | 中间件 |
| 07 | `docs/backend/07-error-handling.md` | 错误处理 |
| 08 | `docs/backend/08-cache-design.md` | 缓存设计 |
| 09 | `docs/backend/09-logging-design.md` | 日志设计 |
| 10 | `docs/backend/10-deployment.md` | 部署方案 |

### 前端文档
| 序号 | 路径 | 说明 |
|------|------|------|
| 00 | `docs/frontend/00-overview.md` | 总览 |
| 01 | `docs/frontend/01-project-structure.md` | 项目结构 |
| 02 | `docs/frontend/02-component-design.md` | 组件设计 |
| 03 | `docs/frontend/03-state-management.md` | 状态管理 |
| 04 | `docs/frontend/04-api-integration.md` | API集成 |
| 05 | `docs/frontend/05-routing.md` | 路由设计 |
| 06 | `docs/frontend/06-style-guide.md` | 样式规范 |

### 通用文档
| 文件 | 路径 | 说明 |
|------|------|------|
| README | `docs/README.md` | 项目说明 |
| 快速开始 | `docs/getting-started.md` | 环境搭建 |
| 开发指南 | `docs/development-guide.md` | 开发规范 |
| API参考 | `docs/api-reference.md` | API参考 |

---

## 六、配置目录

### AI Agents
**位置**: `.warp/agents/` (91个Agent定义)

常用Agent:
- `backend-architect.md` - 后端架构
- `frontend-developer.md` - 前端开发
- `code-reviewer.md` - 代码审查
- `devops.md` - 运维部署
- `go-pro.md` - Go开发
- `vue-pro.md` - Vue开发

### GitHub配置
**位置**: `.github/`

### Claude配置
**位置**: `.claude/`

---

## 七、构建产物

| 类型 | 路径 | 说明 |
|------|------|------|
| 前端dist | `frontend/dist/` | 前端构建产物 |
| 后端binary | `backend/build/oceanengine-backend` | 后端二进制 |

---

## 八、部署配置

| 类型 | 路径 | 说明 |
|------|------|------|
| Docker | `backend/deployments/docker/` | Docker配置 |
| Kubernetes | `backend/deployments/kubernetes/` | K8s配置 |
| 脚本 | `backend/scripts/` | 部署脚本 |
