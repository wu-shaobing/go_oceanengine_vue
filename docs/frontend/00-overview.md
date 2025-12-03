# 前端开发总览

## 项目概述

本项目是 OceanEngine 广告管理平台的前端部分，基于 Vue 3 + TypeScript + Vite 构建，提供广告投放管理的用户界面。

## 技术栈

### 核心框架
- **Vue 3.4+** - 渐进式 JavaScript 框架
- **TypeScript 5.x** - 类型安全
- **Vite 5.x** - 下一代前端构建工具

### UI 与样式
- **TailwindCSS 3.4** - 原子化 CSS 框架
- **Heroicons** - SVG 图标库
- **自定义组件** - 业务组件库

### 状态管理
- **Pinia** - Vue 3 官方推荐状态管理

### 路由
- **Vue Router 4** - 官方路由管理器

### 数据可视化
- **Chart.js 4** - 图表库
- **vue-chartjs** - Vue 封装

### HTTP 请求
- **Axios** - HTTP 客户端

### 工具库
- **VueUse** - Vue 组合式 API 工具集
- **dayjs** - 日期处理
- **lodash-es** - 工具函数

## 项目结构

```
frontend/
├── public/                    # 静态资源
│   ├── favicon.ico
│   └── logo.svg
│
├── src/
│   ├── api/                   # API 接口层
│   │   ├── index.ts           # API 入口
│   │   ├── request.ts         # Axios 封装
│   │   ├── auth.ts            # 认证接口
│   │   ├── advertiser.ts      # 广告主接口
│   │   ├── campaign.ts        # 广告系列接口
│   │   ├── creative.ts        # 创意接口
│   │   └── report.ts          # 报表接口
│   │
│   ├── assets/                # 资源文件
│   │   ├── images/
│   │   └── styles/
│   │       ├── main.css       # 全局样式
│   │       └── tailwind.css   # Tailwind 入口
│   │
│   ├── components/            # 公共组件
│   │   ├── common/            # 通用组件
│   │   │   ├── Button.vue
│   │   │   ├── Input.vue
│   │   │   ├── Modal.vue
│   │   │   ├── Table.vue
│   │   │   └── Pagination.vue
│   │   │
│   │   ├── layout/            # 布局组件
│   │   │   ├── AppLayout.vue
│   │   │   ├── AppHeader.vue
│   │   │   ├── AppSidebar.vue
│   │   │   └── Breadcrumb.vue
│   │   │
│   │   └── business/          # 业务组件
│   │       ├── StatsCard.vue
│   │       ├── StatusBadge.vue
│   │       └── DataTable.vue
│   │
│   ├── composables/           # 组合式函数
│   │   ├── useAuth.ts
│   │   ├── usePagination.ts
│   │   ├── useTable.ts
│   │   └── useForm.ts
│   │
│   ├── router/                # 路由配置
│   │   ├── index.ts
│   │   ├── routes.ts
│   │   └── guards.ts
│   │
│   ├── stores/                # 状态管理
│   │   ├── index.ts
│   │   ├── auth.ts
│   │   ├── app.ts
│   │   └── advertiser.ts
│   │
│   ├── types/                 # 类型定义
│   │   ├── api.ts
│   │   ├── advertiser.ts
│   │   ├── campaign.ts
│   │   └── common.ts
│   │
│   ├── utils/                 # 工具函数
│   │   ├── format.ts
│   │   ├── validate.ts
│   │   ├── storage.ts
│   │   └── helpers.ts
│   │
│   ├── views/                 # 页面视图
│   │   ├── dashboard/
│   │   ├── advertiser/
│   │   ├── campaign/
│   │   ├── creative/
│   │   ├── report/
│   │   ├── media/
│   │   └── system/
│   │
│   ├── App.vue                # 根组件
│   ├── main.ts                # 入口文件
│   └── env.d.ts               # 类型声明
│
├── .env                       # 环境变量
├── .env.development           # 开发环境变量
├── .env.production            # 生产环境变量
├── .eslintrc.cjs              # ESLint 配置
├── .prettierrc                # Prettier 配置
├── index.html                 # HTML 模板
├── package.json               # 项目配置
├── tailwind.config.js         # Tailwind 配置
├── tsconfig.json              # TypeScript 配置
└── vite.config.ts             # Vite 配置
```

## 设计原则

### 1. 组件化
- 单一职责：每个组件只做一件事
- 可复用性：通用组件与业务组件分离
- Props Down, Events Up：单向数据流

### 2. 类型安全
- 所有组件使用 TypeScript
- API 响应类型完整定义
- 避免使用 `any` 类型

### 3. 响应式设计
- 移动端优先
- 使用 Tailwind 响应式断点
- 弹性布局

### 4. 性能优化
- 路由懒加载
- 组件按需加载
- 图片懒加载
- 虚拟滚动（大列表）

## 开发规范

### 命名规范

```typescript
// 组件名：PascalCase
// 文件名：PascalCase.vue 或 kebab-case.vue
AdvertiserList.vue
advertiser-list.vue

// 组合式函数：use 前缀
useAuth.ts
usePagination.ts

// 类型：PascalCase，接口用 I 前缀（可选）
interface Advertiser { }
type AdvertiserStatus = 'active' | 'inactive'

// 常量：UPPER_SNAKE_CASE
const API_BASE_URL = '/api/v1'
const MAX_PAGE_SIZE = 100
```

### 组件规范

```vue
<script setup lang="ts">
// 1. 类型导入
import type { Advertiser } from '@/types/advertiser'

// 2. 组件导入
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'

// 3. 组合式函数
import { useAuth } from '@/composables/useAuth'

// 4. Props 定义
interface Props {
  id: number
  title?: string
}
const props = withDefaults(defineProps<Props>(), {
  title: '默认标题'
})

// 5. Emits 定义
const emit = defineEmits<{
  (e: 'update', value: string): void
  (e: 'delete', id: number): void
}>()

// 6. 响应式数据
const loading = ref(false)
const data = ref<Advertiser | null>(null)

// 7. 计算属性
const displayName = computed(() => data.value?.name ?? '未知')

// 8. 方法
const fetchData = async () => {
  // ...
}

// 9. 生命周期
onMounted(() => {
  fetchData()
})
</script>

<template>
  <!-- 模板内容 -->
</template>

<style scoped>
/* 组件样式 */
</style>
```

### 目录规范

```
views/
├── advertiser/
│   ├── AdvertiserList.vue      # 列表页
│   ├── AdvertiserDetail.vue    # 详情页
│   ├── AdvertiserCreate.vue    # 创建页
│   ├── AdvertiserEdit.vue      # 编辑页
│   └── components/             # 页面私有组件
│       ├── AdvertiserForm.vue
│       └── AdvertiserCard.vue
```

## 核心功能模块

### 1. 认证模块
- 登录/登出
- Token 管理
- 权限控制

### 2. 广告主管理
- 广告主列表
- 广告主详情
- 授权管理

### 3. 广告系列管理
- 系列列表
- 创建/编辑系列
- 状态管理

### 4. 数据报表
- 数据看板
- 趋势图表
- 数据导出

### 5. 素材管理
- 图片上传
- 视频上传
- 素材库

## 文档目录

| 文档 | 说明 |
|------|------|
| 01-project-structure.md | 项目结构详解 |
| 02-component-design.md | 组件设计规范 |
| 03-state-management.md | 状态管理 |
| 04-api-integration.md | API 集成 |
| 05-routing.md | 路由设计 |
| 06-style-guide.md | 样式规范 |
