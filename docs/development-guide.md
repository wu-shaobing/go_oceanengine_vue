# 开发指南

## 开发环境

### 系统要求

- **操作系统**：macOS / Linux / Windows
- **Go**：1.21+
- **Node.js**：18.x+ (推荐 20.x)
- **MySQL**：8.0+
- **Redis**：7.0+
- **Docker**：24.0+（可选，用于容器化开发）

### 工具推荐

| 工具 | 用途 | 下载地址 |
|------|------|----------|
| VS Code | 代码编辑器 | https://code.visualstudio.com |
| GoLand | Go IDE | https://www.jetbrains.com/go |
| TablePlus | 数据库管理 | https://tableplus.com |
| Postman | API 测试 | https://www.postman.com |
| Another Redis Desktop Manager | Redis 管理 | https://goanother.com |

### VS Code 插件

**Go 开发：**
- Go (官方插件)
- Go Test Explorer

**前端开发：**
- Vue - Official
- TypeScript Vue Plugin (Volar)
- Tailwind CSS IntelliSense
- ESLint
- Prettier

---

## 项目初始化

### 1. 克隆项目

```bash
git clone https://github.com/your-org/oceanengine.git
cd oceanengine
```

### 2. 配置后端

```bash
# 复制配置文件
cp config.example.yaml config.yaml

# 编辑配置（修改数据库、Redis 等连接信息）
vim config.yaml

# 安装依赖
go mod download

# 初始化数据库
mysql -u root -p < scripts/init.sql

# 启动后端服务
go run main.go
```

### 3. 配置前端

```bash
cd frontend

# 安装依赖
npm install

# 复制环境变量
cp .env.example .env.development

# 启动开发服务器
npm run dev
```

### 4. 使用 Docker Compose（推荐）

```bash
# 一键启动所有服务
docker-compose up -d

# 查看日志
docker-compose logs -f

# 停止服务
docker-compose down
```

---

## 开发工作流

### Git 分支策略

```
main          # 生产分支，保持稳定
├── develop   # 开发分支，日常开发
├── feature/* # 功能分支
├── bugfix/*  # Bug 修复分支
├── hotfix/*  # 紧急修复分支
└── release/* # 发布分支
```

### 分支命名规范

```bash
# 功能开发
feature/add-audience-management
feature/campaign-optimization

# Bug 修复
bugfix/fix-login-error
bugfix/report-data-mismatch

# 紧急修复
hotfix/fix-token-refresh

# 发布
release/v1.2.0
```

### 提交规范

使用 [Conventional Commits](https://www.conventionalcommits.org/) 规范：

```bash
# 格式
<type>(<scope>): <subject>

# 类型
feat:     新功能
fix:      Bug 修复
docs:     文档更新
style:    代码格式（不影响功能）
refactor: 重构
perf:     性能优化
test:     测试相关
chore:    构建/工具相关

# 示例
feat(campaign): add batch update status API
fix(auth): fix token refresh logic
docs(api): update API documentation
refactor(advertiser): optimize database queries
```

### 开发流程

```bash
# 1. 从 develop 创建功能分支
git checkout develop
git pull origin develop
git checkout -b feature/my-feature

# 2. 开发并提交
git add .
git commit -m "feat(module): add new feature"

# 3. 推送并创建 PR
git push origin feature/my-feature
# 在 GitHub/GitLab 创建 Pull Request

# 4. Code Review 后合并
# 合并后删除功能分支
git branch -d feature/my-feature
```

---

## 后端开发

### 目录结构

```
├── api/           # API 接口层
│   └── v1/        # v1 版本接口
├── internal/      # 内部模块
│   ├── dao/       # 数据访问层
│   ├── dto/       # 数据传输对象
│   ├── model/     # 数据模型
│   └── service/   # 业务逻辑层
├── middleware/    # 中间件
├── pkg/           # 公共包
└── config/        # 配置
```

### 添加新 API

**1. 定义数据模型**

```go
// internal/model/example.go
type Example struct {
    ID        uint      `gorm:"primarykey"`
    Name      string    `gorm:"size:100;not null"`
    Status    int       `gorm:"default:1"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

**2. 定义 DTO**

```go
// internal/dto/example.go
type ExampleCreateReq struct {
    Name string `json:"name" binding:"required,max=100"`
}

type ExampleResp struct {
    ID     uint   `json:"id"`
    Name   string `json:"name"`
    Status int    `json:"status"`
}
```

**3. 实现 DAO**

```go
// internal/dao/example.go
type ExampleDao struct {
    db *gorm.DB
}

func (d *ExampleDao) Create(example *model.Example) error {
    return d.db.Create(example).Error
}

func (d *ExampleDao) GetByID(id uint) (*model.Example, error) {
    var example model.Example
    err := d.db.First(&example, id).Error
    return &example, err
}
```

**4. 实现 Service**

```go
// internal/service/example.go
type ExampleService struct {
    dao *dao.ExampleDao
}

func (s *ExampleService) Create(req *dto.ExampleCreateReq) (*dto.ExampleResp, error) {
    example := &model.Example{
        Name: req.Name,
    }
    if err := s.dao.Create(example); err != nil {
        return nil, err
    }
    return &dto.ExampleResp{
        ID:     example.ID,
        Name:   example.Name,
        Status: example.Status,
    }, nil
}
```

**5. 实现 API Handler**

```go
// api/v1/example.go
type ExampleApi struct {
    service *service.ExampleService
}

// CreateExample godoc
// @Summary 创建示例
// @Tags Example
// @Accept json
// @Produce json
// @Param data body dto.ExampleCreateReq true "请求参数"
// @Success 200 {object} response.Response{data=dto.ExampleResp}
// @Router /example [post]
func (a *ExampleApi) Create(c *gin.Context) {
    var req dto.ExampleCreateReq
    if err := c.ShouldBindJSON(&req); err != nil {
        response.FailWithMessage(c, err.Error())
        return
    }
    
    resp, err := a.service.Create(&req)
    if err != nil {
        response.FailWithMessage(c, err.Error())
        return
    }
    
    response.OkWithData(c, resp)
}
```

**6. 注册路由**

```go
// router/example.go
func InitExampleRouter(r *gin.RouterGroup) {
    api := &v1.ExampleApi{}
    group := r.Group("example")
    {
        group.POST("", api.Create)
        group.GET(":id", api.GetByID)
        group.PUT(":id", api.Update)
        group.DELETE(":id", api.Delete)
    }
}
```

### 常用命令

```bash
# 运行服务
go run main.go

# 运行测试
go test ./...

# 生成 Swagger 文档
swag init

# 代码格式化
go fmt ./...

# 代码检查
golangci-lint run
```

---

## 前端开发

### 添加新页面

**1. 创建视图组件**

```vue
<!-- src/views/example/ExampleList.vue -->
<template>
  <div class="page-container">
    <header class="page-header">
      <h1 class="page-title">示例列表</h1>
      <button class="btn-primary" @click="handleCreate">
        新建
      </button>
    </header>
    
    <main class="page-content">
      <DataTable
        :columns="columns"
        :data="list"
        :loading="loading"
        :total="total"
        @page-change="handlePageChange"
      />
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { exampleApi, type Example } from '@/api/example'
import DataTable from '@/components/common/Table.vue'

const loading = ref(false)
const list = ref<Example[]>([])
const total = ref(0)
const page = ref(1)

const columns = [
  { key: 'id', title: 'ID', width: 80 },
  { key: 'name', title: '名称' },
  { key: 'status', title: '状态' }
]

const fetchData = async () => {
  loading.value = true
  try {
    const res = await exampleApi.getList({ page: page.value, page_size: 10 })
    list.value = res.list
    total.value = res.total
  } finally {
    loading.value = false
  }
}

const handlePageChange = (newPage: number) => {
  page.value = newPage
  fetchData()
}

const handleCreate = () => {
  // 跳转创建页
}

onMounted(fetchData)
</script>
```

**2. 定义 API**

```typescript
// src/api/example.ts
import { request, PageResponse } from './request'

export interface Example {
  id: number
  name: string
  status: number
}

export const exampleApi = {
  getList(params: { page: number; page_size: number }) {
    return request.get<PageResponse<Example>>('/example', params)
  },
  
  create(data: { name: string }) {
    return request.post<{ id: number }>('/example', data)
  }
}
```

**3. 添加路由**

```typescript
// src/router/index.ts
{
  path: '/example',
  component: MainLayout,
  children: [
    {
      path: 'list',
      name: 'ExampleList',
      component: () => import('@/views/example/ExampleList.vue'),
      meta: { title: '示例列表' }
    }
  ]
}
```

### 常用命令

```bash
# 启动开发服务器
npm run dev

# 类型检查
npm run typecheck

# 代码检查
npm run lint

# 构建生产版本
npm run build

# 预览构建结果
npm run preview
```

---

## 调试技巧

### 后端调试

**1. 使用 Delve 调试器**

```bash
# 安装
go install github.com/go-delve/delve/cmd/dlv@latest

# 启动调试
dlv debug main.go
```

**2. VS Code 调试配置**

```json
// .vscode/launch.json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch Server",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}/main.go"
    }
  ]
}
```

### 前端调试

**1. Vue DevTools**

安装 [Vue DevTools](https://devtools.vuejs.org/) 浏览器插件。

**2. 源码映射**

开发模式下自动启用 Source Map，可在浏览器中直接调试 TypeScript。

**3. 网络请求**

使用浏览器 Network 面板或 Axios 拦截器查看请求详情。

---

## 测试

### 后端测试

```bash
# 运行所有测试
go test ./...

# 运行特定包测试
go test ./internal/service/...

# 查看覆盖率
go test -cover ./...

# 生成覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### 前端测试

```bash
# 运行单元测试
npm run test

# 运行 E2E 测试
npm run test:e2e
```

---

## 部署

### 手动部署

```bash
# 后端构建
go build -o server main.go

# 前端构建
cd frontend && npm run build

# 复制构建产物到服务器
scp server user@server:/app/
scp -r frontend/dist user@server:/app/static/
```

### Docker 部署

```bash
# 构建镜像
docker build -t oceanengine:latest .

# 运行容器
docker run -d -p 8080:8080 oceanengine:latest
```

### CI/CD

项目使用 GitHub Actions 进行持续集成和部署，配置文件位于 `.github/workflows/`。

---

## 常见问题

### Q: 后端启动失败，提示数据库连接错误

检查 `config.yaml` 中的数据库配置，确保：
- MySQL 服务已启动
- 用户名密码正确
- 数据库已创建

### Q: 前端请求 API 报 CORS 错误

开发模式下，确保 `vite.config.ts` 中配置了正确的代理：

```typescript
proxy: {
  '/api': {
    target: 'http://localhost:8080',
    changeOrigin: true
  }
}
```

### Q: TypeScript 类型错误

```bash
# 重新生成类型声明
npm run typecheck

# 清除缓存
rm -rf node_modules/.vite
npm run dev
```

### Q: Go 依赖下载慢

配置 Go 代理：

```bash
go env -w GOPROXY=https://goproxy.cn,direct
```

---

## 参考资源

- [Go 官方文档](https://go.dev/doc/)
- [Gin 框架文档](https://gin-gonic.com/docs/)
- [GORM 文档](https://gorm.io/docs/)
- [Vue 3 文档](https://vuejs.org/)
- [TypeScript 文档](https://www.typescriptlang.org/docs/)
- [TailwindCSS 文档](https://tailwindcss.com/docs)
- [巨量引擎 API 文档](https://open.oceanengine.com/doc/)
