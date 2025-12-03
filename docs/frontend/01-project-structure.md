# 项目结构

## 目录结构

```
frontend/
├── public/                     # 静态资源
│   ├── favicon.ico            # 网站图标
│   └── logo.svg               # Logo 图片
│
├── src/                        # 源代码
│   ├── api/                   # API 接口
│   │   ├── request.ts         # Axios 封装
│   │   ├── auth.ts            # 认证相关 API
│   │   ├── advertiser.ts      # 广告主 API
│   │   ├── campaign.ts        # 广告系列 API
│   │   ├── report.ts          # 报表 API
│   │   └── index.ts           # API 统一导出
│   │
│   ├── assets/                # 静态资源
│   │   ├── images/            # 图片资源
│   │   └── icons/             # 图标资源
│   │
│   ├── components/            # 组件
│   │   ├── common/            # 通用基础组件
│   │   │   ├── Button.vue
│   │   │   ├── Input.vue
│   │   │   ├── Modal.vue
│   │   │   ├── Table.vue
│   │   │   └── index.ts
│   │   │
│   │   ├── layout/            # 布局组件
│   │   │   ├── MainLayout.vue
│   │   │   ├── AuthLayout.vue
│   │   │   ├── AppHeader.vue
│   │   │   ├── AppSidebar.vue
│   │   │   └── AppBreadcrumb.vue
│   │   │
│   │   └── business/          # 业务组件
│   │       ├── StatsCard.vue
│   │       ├── StatusBadge.vue
│   │       ├── DateRangePicker.vue
│   │       └── AdvertiserSelector.vue
│   │
│   ├── composables/           # 组合式函数
│   │   ├── useRequest.ts      # 请求 Hook
│   │   ├── useTable.ts        # 表格 Hook
│   │   ├── useForm.ts         # 表单 Hook
│   │   ├── useTheme.ts        # 主题 Hook
│   │   └── index.ts
│   │
│   ├── directives/            # 自定义指令
│   │   ├── permission.ts      # 权限指令
│   │   ├── loading.ts         # 加载指令
│   │   └── index.ts
│   │
│   ├── router/                # 路由配置
│   │   ├── index.ts           # 路由实例
│   │   ├── guards.ts          # 路由守卫
│   │   └── routes.ts          # 路由定义
│   │
│   ├── stores/                # Pinia 状态管理
│   │   ├── auth.ts            # 认证状态
│   │   ├── app.ts             # 应用状态
│   │   ├── advertiser.ts      # 广告主状态
│   │   └── index.ts
│   │
│   ├── styles/                # 全局样式
│   │   ├── base.css           # 基础样式
│   │   ├── components.css     # 组件样式
│   │   ├── utilities.css      # 工具类
│   │   └── main.css           # 入口文件
│   │
│   ├── types/                 # TypeScript 类型
│   │   ├── api.d.ts           # API 类型
│   │   ├── router.d.ts        # 路由类型
│   │   ├── env.d.ts           # 环境变量类型
│   │   └── global.d.ts        # 全局类型
│   │
│   ├── utils/                 # 工具函数
│   │   ├── format.ts          # 格式化工具
│   │   ├── validate.ts        # 校验工具
│   │   ├── storage.ts         # 存储工具
│   │   ├── error-handler.ts   # 错误处理
│   │   └── index.ts
│   │
│   ├── views/                 # 页面视图
│   │   ├── auth/              # 认证页面
│   │   │   └── LoginView.vue
│   │   │
│   │   ├── dashboard/         # 仪表盘
│   │   │   └── DashboardView.vue
│   │   │
│   │   ├── advertiser/        # 广告主管理
│   │   │   ├── AdvertiserList.vue
│   │   │   └── AdvertiserDetail.vue
│   │   │
│   │   ├── campaign/          # 广告系列
│   │   │   ├── CampaignList.vue
│   │   │   └── CampaignEdit.vue
│   │   │
│   │   ├── report/            # 数据报表
│   │   │   ├── ReportDashboard.vue
│   │   │   └── ReportDetail.vue
│   │   │
│   │   ├── audience/          # 人群管理
│   │   │   ├── AudienceList.vue
│   │   │   └── AudienceEdit.vue
│   │   │
│   │   ├── system/            # 系统管理
│   │   │   ├── UserManage.vue
│   │   │   ├── RoleManage.vue
│   │   │   └── MenuManage.vue
│   │   │
│   │   └── error/             # 错误页面
│   │       ├── 403.vue
│   │       └── 404.vue
│   │
│   ├── App.vue                # 根组件
│   └── main.ts                # 入口文件
│
├── .env                        # 环境变量
├── .env.development           # 开发环境
├── .env.production            # 生产环境
├── index.html                 # HTML 模板
├── package.json               # 依赖配置
├── tsconfig.json              # TypeScript 配置
├── vite.config.ts             # Vite 配置
├── tailwind.config.js         # Tailwind 配置
└── postcss.config.js          # PostCSS 配置
```

---

## 核心文件说明

### 入口文件

```typescript
// src/main.ts
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

import App from './App.vue'
import router from './router'
import { setupRouterGuards } from './router/guards'

// 样式
import './styles/main.css'

// 指令
import { permissionDirective } from './directives/permission'

// 创建应用
const app = createApp(App)

// Pinia
const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)
app.use(pinia)

// 路由
app.use(router)
setupRouterGuards(router)

// 注册指令
app.directive('permission', permissionDirective)

// 挂载
app.mount('#app')
```

### 根组件

```vue
<!-- src/App.vue -->
<template>
  <router-view />
</template>

<script setup lang="ts">
// 全局初始化逻辑（如有需要）
</script>
```

### Vite 配置

```typescript
// vite.config.ts
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': resolve(__dirname, 'src')
    }
  },
  server: {
    port: 3000,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true
      }
    }
  },
  build: {
    outDir: 'dist',
    sourcemap: false,
    minify: 'terser',
    rollupOptions: {
      output: {
        manualChunks: {
          vue: ['vue', 'vue-router', 'pinia'],
          charts: ['chart.js', 'vue-chartjs']
        }
      }
    }
  }
})
```

### TypeScript 配置

```json
// tsconfig.json
{
  "compilerOptions": {
    "target": "ES2020",
    "useDefineForClassFields": true,
    "module": "ESNext",
    "lib": ["ES2020", "DOM", "DOM.Iterable"],
    "skipLibCheck": true,
    "moduleResolution": "bundler",
    "allowImportingTsExtensions": true,
    "resolveJsonModule": true,
    "isolatedModules": true,
    "noEmit": true,
    "jsx": "preserve",
    "strict": true,
    "noUnusedLocals": true,
    "noUnusedParameters": true,
    "noFallthroughCasesInSwitch": true,
    "baseUrl": ".",
    "paths": {
      "@/*": ["src/*"]
    }
  },
  "include": ["src/**/*.ts", "src/**/*.d.ts", "src/**/*.tsx", "src/**/*.vue"],
  "references": [{ "path": "./tsconfig.node.json" }]
}
```

### 环境变量

```bash
# .env
VITE_APP_TITLE=巨量引擎管理系统

# .env.development
VITE_API_BASE_URL=/api/v1
VITE_MOCK_ENABLED=false

# .env.production
VITE_API_BASE_URL=https://api.example.com/v1
VITE_MOCK_ENABLED=false
```

```typescript
// src/types/env.d.ts
/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_APP_TITLE: string
  readonly VITE_API_BASE_URL: string
  readonly VITE_MOCK_ENABLED: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
```

---

## 模块约定

### API 模块

每个 API 模块包含：
- 接口类型定义
- 请求方法封装
- 统一的命名规范（`xxxApi`）

```typescript
// src/api/example.ts
import { request, PageResponse } from './request'

// 类型定义
export interface Example {
  id: number
  name: string
}

export interface ExampleCreateRequest {
  name: string
}

// API 对象
export const exampleApi = {
  getList(params: { page: number; page_size: number }) {
    return request.get<PageResponse<Example>>('/example', params)
  },
  
  getDetail(id: number) {
    return request.get<Example>(`/example/${id}`)
  },
  
  create(data: ExampleCreateRequest) {
    return request.post<{ id: number }>('/example', data)
  },
  
  update(id: number, data: Partial<ExampleCreateRequest>) {
    return request.put<void>(`/example/${id}`, data)
  },
  
  delete(id: number) {
    return request.delete<void>(`/example/${id}`)
  }
}
```

### 组件模块

组件分为三类：

1. **通用组件**（`components/common/`）：与业务无关的 UI 组件
2. **布局组件**（`components/layout/`）：页面布局相关组件
3. **业务组件**（`components/business/`）：与业务相关的复合组件

### 视图模块

视图按功能模块划分，每个模块包含：
- 列表页（`XxxList.vue`）
- 详情页（`XxxDetail.vue`）
- 编辑页（`XxxEdit.vue`，新增/编辑共用）

### Store 模块

每个 Store 包含：
- `state`：状态定义
- `getters`：计算属性
- `actions`：业务方法

```typescript
// src/stores/example.ts
import { defineStore } from 'pinia'

interface ExampleState {
  list: Example[]
  current: Example | null
}

export const useExampleStore = defineStore('example', {
  state: (): ExampleState => ({
    list: [],
    current: null
  }),
  
  getters: {
    activeItems: (state) => state.list.filter(item => item.active)
  },
  
  actions: {
    async fetchList() {
      // ...
    }
  }
})
```

---

## 命名规范

### 文件命名

| 类型 | 规范 | 示例 |
|------|------|------|
| 组件 | PascalCase | `Button.vue`, `UserList.vue` |
| 组合式函数 | camelCase + use 前缀 | `useRequest.ts`, `useTable.ts` |
| Store | camelCase | `auth.ts`, `advertiser.ts` |
| API | camelCase | `auth.ts`, `campaign.ts` |
| 工具函数 | camelCase | `format.ts`, `validate.ts` |
| 类型定义 | `.d.ts` 后缀 | `api.d.ts`, `router.d.ts` |

### 变量命名

| 类型 | 规范 | 示例 |
|------|------|------|
| 变量/函数 | camelCase | `userName`, `fetchData` |
| 常量 | UPPER_SNAKE_CASE | `API_BASE_URL`, `MAX_SIZE` |
| 组件名 | PascalCase | `MyButton`, `DataTable` |
| Props | camelCase | `modelValue`, `pageSize` |
| Emits | kebab-case | `update:model-value`, `page-change` |
| CSS 类 | kebab-case | `page-header`, `btn-primary` |

---

## 依赖说明

### 生产依赖

```json
{
  "dependencies": {
    "vue": "^3.4.0",
    "vue-router": "^4.2.0",
    "pinia": "^2.1.0",
    "pinia-plugin-persistedstate": "^3.2.0",
    "axios": "^1.6.0",
    "chart.js": "^4.4.0",
    "vue-chartjs": "^5.3.0",
    "dayjs": "^1.11.0",
    "nprogress": "^0.2.0"
  }
}
```

### 开发依赖

```json
{
  "devDependencies": {
    "@vitejs/plugin-vue": "^5.0.0",
    "vite": "^5.0.0",
    "typescript": "^5.3.0",
    "vue-tsc": "^1.8.0",
    "tailwindcss": "^3.4.0",
    "postcss": "^8.4.0",
    "autoprefixer": "^10.4.0",
    "@tailwindcss/forms": "^0.5.0",
    "@tailwindcss/typography": "^0.5.0",
    "@types/node": "^20.0.0",
    "@types/nprogress": "^0.2.0",
    "eslint": "^8.56.0",
    "eslint-plugin-vue": "^9.20.0",
    "@typescript-eslint/eslint-plugin": "^6.19.0",
    "@typescript-eslint/parser": "^6.19.0"
  }
}
```
