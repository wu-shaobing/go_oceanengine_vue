# 样式指南

## 概述

本项目使用 TailwindCSS 3.4 作为主要的样式框架，辅以少量自定义 CSS。遵循原子化 CSS 理念，保持样式的一致性和可维护性。

## TailwindCSS 配置

### 基础配置

```javascript
// tailwind.config.js
/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './index.html',
    './src/**/*.{vue,js,ts,jsx,tsx}'
  ],
  theme: {
    extend: {
      // 品牌色
      colors: {
        primary: {
          50: '#eff6ff',
          100: '#dbeafe',
          200: '#bfdbfe',
          300: '#93c5fd',
          400: '#60a5fa',
          500: '#3b82f6',  // 主色
          600: '#2563eb',
          700: '#1d4ed8',
          800: '#1e40af',
          900: '#1e3a8a'
        },
        success: {
          light: '#dcfce7',
          DEFAULT: '#22c55e',
          dark: '#15803d'
        },
        warning: {
          light: '#fef3c7',
          DEFAULT: '#f59e0b',
          dark: '#b45309'
        },
        danger: {
          light: '#fee2e2',
          DEFAULT: '#ef4444',
          dark: '#b91c1c'
        },
        // 灰度
        gray: {
          50: '#f9fafb',
          100: '#f3f4f6',
          200: '#e5e7eb',
          300: '#d1d5db',
          400: '#9ca3af',
          500: '#6b7280',
          600: '#4b5563',
          700: '#374151',
          800: '#1f2937',
          900: '#111827'
        }
      },
      // 字体
      fontFamily: {
        sans: [
          'Inter',
          '-apple-system',
          'BlinkMacSystemFont',
          'Segoe UI',
          'Roboto',
          'Helvetica Neue',
          'Arial',
          'sans-serif'
        ],
        mono: ['JetBrains Mono', 'Menlo', 'Monaco', 'Consolas', 'monospace']
      },
      // 字号
      fontSize: {
        'xs': ['0.75rem', { lineHeight: '1rem' }],
        'sm': ['0.875rem', { lineHeight: '1.25rem' }],
        'base': ['1rem', { lineHeight: '1.5rem' }],
        'lg': ['1.125rem', { lineHeight: '1.75rem' }],
        'xl': ['1.25rem', { lineHeight: '1.75rem' }],
        '2xl': ['1.5rem', { lineHeight: '2rem' }],
        '3xl': ['1.875rem', { lineHeight: '2.25rem' }]
      },
      // 圆角
      borderRadius: {
        'sm': '0.25rem',
        'DEFAULT': '0.375rem',
        'md': '0.5rem',
        'lg': '0.75rem',
        'xl': '1rem'
      },
      // 阴影
      boxShadow: {
        'sm': '0 1px 2px 0 rgb(0 0 0 / 0.05)',
        'DEFAULT': '0 1px 3px 0 rgb(0 0 0 / 0.1), 0 1px 2px -1px rgb(0 0 0 / 0.1)',
        'md': '0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1)',
        'lg': '0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1)',
        'card': '0 2px 8px 0 rgb(0 0 0 / 0.08)'
      },
      // 动画
      animation: {
        'fade-in': 'fadeIn 0.2s ease-out',
        'slide-up': 'slideUp 0.3s ease-out',
        'slide-down': 'slideDown 0.3s ease-out',
        'spin-slow': 'spin 2s linear infinite'
      },
      keyframes: {
        fadeIn: {
          '0%': { opacity: '0' },
          '100%': { opacity: '1' }
        },
        slideUp: {
          '0%': { transform: 'translateY(10px)', opacity: '0' },
          '100%': { transform: 'translateY(0)', opacity: '1' }
        },
        slideDown: {
          '0%': { transform: 'translateY(-10px)', opacity: '0' },
          '100%': { transform: 'translateY(0)', opacity: '1' }
        }
      },
      // 间距
      spacing: {
        '18': '4.5rem',
        '22': '5.5rem'
      }
    }
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography')
  ]
}
```

---

## 全局样式

### 基础重置

```css
/* src/styles/base.css */
@tailwind base;
@tailwind components;
@tailwind utilities;

@layer base {
  /* 根元素 */
  :root {
    --sidebar-width: 240px;
    --sidebar-collapsed-width: 64px;
    --header-height: 60px;
    --page-padding: 24px;
  }

  /* HTML/Body */
  html {
    @apply antialiased;
    font-size: 14px;
  }

  body {
    @apply bg-gray-50 text-gray-900;
    min-height: 100vh;
  }

  /* 链接 */
  a {
    @apply text-primary-600 hover:text-primary-700 transition-colors;
  }

  /* 滚动条 */
  ::-webkit-scrollbar {
    @apply w-2 h-2;
  }

  ::-webkit-scrollbar-track {
    @apply bg-gray-100 rounded;
  }

  ::-webkit-scrollbar-thumb {
    @apply bg-gray-300 rounded hover:bg-gray-400;
  }

  /* 选中文本 */
  ::selection {
    @apply bg-primary-100 text-primary-900;
  }
}
```

### 组件样式

```css
/* src/styles/components.css */
@layer components {
  /* 卡片 */
  .card {
    @apply bg-white rounded-lg shadow-card p-6;
  }

  .card-header {
    @apply flex items-center justify-between mb-4 pb-4 border-b border-gray-100;
  }

  .card-title {
    @apply text-lg font-semibold text-gray-900;
  }

  /* 按钮 */
  .btn {
    @apply inline-flex items-center justify-center gap-2 px-4 py-2 
           text-sm font-medium rounded-md transition-all duration-200
           focus:outline-none focus:ring-2 focus:ring-offset-2
           disabled:opacity-50 disabled:cursor-not-allowed;
  }

  .btn-primary {
    @apply btn bg-primary-600 text-white 
           hover:bg-primary-700 focus:ring-primary-500;
  }

  .btn-secondary {
    @apply btn bg-white text-gray-700 border border-gray-300
           hover:bg-gray-50 focus:ring-primary-500;
  }

  .btn-danger {
    @apply btn bg-danger text-white
           hover:bg-danger-dark focus:ring-danger;
  }

  .btn-ghost {
    @apply btn bg-transparent text-gray-600
           hover:bg-gray-100 focus:ring-gray-500;
  }

  .btn-sm {
    @apply px-3 py-1.5 text-xs;
  }

  .btn-lg {
    @apply px-6 py-3 text-base;
  }

  /* 表单 */
  .form-label {
    @apply block text-sm font-medium text-gray-700 mb-1;
  }

  .form-input {
    @apply block w-full px-3 py-2 text-sm
           border border-gray-300 rounded-md
           placeholder-gray-400
           focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500
           disabled:bg-gray-100 disabled:cursor-not-allowed;
  }

  .form-select {
    @apply form-input pr-10 appearance-none
           bg-no-repeat bg-right
           bg-[url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 20 20'%3e%3cpath stroke='%236b7280' stroke-linecap='round' stroke-linejoin='round' stroke-width='1.5' d='M6 8l4 4 4-4'/%3e%3c/svg%3e")]
           bg-[length:1.5em_1.5em];
  }

  .form-error {
    @apply mt-1 text-xs text-danger;
  }

  .form-hint {
    @apply mt-1 text-xs text-gray-500;
  }

  /* 表格 */
  .table-wrapper {
    @apply overflow-x-auto rounded-lg border border-gray-200;
  }

  .table {
    @apply min-w-full divide-y divide-gray-200;
  }

  .table th {
    @apply px-4 py-3 text-left text-xs font-semibold text-gray-600 
           uppercase tracking-wider bg-gray-50;
  }

  .table td {
    @apply px-4 py-3 text-sm text-gray-900 whitespace-nowrap;
  }

  .table tbody tr {
    @apply hover:bg-gray-50 transition-colors;
  }

  /* 徽章 */
  .badge {
    @apply inline-flex items-center px-2 py-0.5 
           text-xs font-medium rounded-full;
  }

  .badge-success {
    @apply badge bg-success-light text-success-dark;
  }

  .badge-warning {
    @apply badge bg-warning-light text-warning-dark;
  }

  .badge-danger {
    @apply badge bg-danger-light text-danger-dark;
  }

  .badge-gray {
    @apply badge bg-gray-100 text-gray-700;
  }

  /* 加载状态 */
  .skeleton {
    @apply animate-pulse bg-gray-200 rounded;
  }

  .spinner {
    @apply animate-spin h-5 w-5 border-2 border-gray-300 
           border-t-primary-600 rounded-full;
  }
}
```

### 工具类

```css
/* src/styles/utilities.css */
@layer utilities {
  /* 文本截断 */
  .text-ellipsis-2 {
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
  }

  .text-ellipsis-3 {
    display: -webkit-box;
    -webkit-line-clamp: 3;
    -webkit-box-orient: vertical;
    overflow: hidden;
  }

  /* 隐藏滚动条 */
  .scrollbar-hide {
    -ms-overflow-style: none;
    scrollbar-width: none;
  }

  .scrollbar-hide::-webkit-scrollbar {
    display: none;
  }

  /* 渐变背景 */
  .bg-gradient-primary {
    @apply bg-gradient-to-r from-primary-500 to-primary-700;
  }

  /* 玻璃效果 */
  .glass {
    @apply bg-white/80 backdrop-blur-sm;
  }

  /* 安全区域 */
  .safe-area-inset-bottom {
    padding-bottom: env(safe-area-inset-bottom);
  }
}
```

---

## 布局规范

### 页面结构

```vue
<template>
  <!-- 页面容器 -->
  <div class="page-container">
    <!-- 页面头部 -->
    <header class="page-header">
      <h1 class="page-title">页面标题</h1>
      <div class="page-actions">
        <button class="btn-primary">操作按钮</button>
      </div>
    </header>

    <!-- 页面内容 -->
    <main class="page-content">
      <!-- 筛选区域 -->
      <div class="filter-bar">
        <!-- 筛选表单 -->
      </div>

      <!-- 数据区域 -->
      <div class="card">
        <!-- 表格或其他内容 -->
      </div>
    </main>
  </div>
</template>

<style scoped>
.page-container {
  @apply min-h-full;
}

.page-header {
  @apply flex items-center justify-between mb-6;
}

.page-title {
  @apply text-2xl font-bold text-gray-900;
}

.page-actions {
  @apply flex items-center gap-3;
}

.filter-bar {
  @apply bg-white rounded-lg shadow-card p-4 mb-4;
}

.page-content {
  @apply space-y-4;
}
</style>
```

### 栅格系统

```vue
<template>
  <!-- 响应式栅格 -->
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
    <div class="card">卡片 1</div>
    <div class="card">卡片 2</div>
    <div class="card">卡片 3</div>
    <div class="card">卡片 4</div>
  </div>

  <!-- 固定比例布局 -->
  <div class="grid grid-cols-12 gap-4">
    <div class="col-span-8">主内容区</div>
    <div class="col-span-4">侧边栏</div>
  </div>
</template>
```

### 间距规范

| 场景 | 间距值 | Tailwind 类 |
|------|--------|-------------|
| 元素内边距（小） | 8px | `p-2` |
| 元素内边距（中） | 16px | `p-4` |
| 元素内边距（大） | 24px | `p-6` |
| 元素间距（小） | 8px | `gap-2` |
| 元素间距（中） | 16px | `gap-4` |
| 模块间距 | 24px | `mb-6` |
| 页面边距 | 24px | `p-6` |

---

## 颜色使用

### 语义化颜色

```vue
<template>
  <!-- 主色：品牌色、主要操作 -->
  <button class="bg-primary-600 text-white">主要按钮</button>
  <a class="text-primary-600">链接文字</a>

  <!-- 成功：正向状态、成功提示 -->
  <span class="badge-success">投放中</span>
  <p class="text-success">保存成功</p>

  <!-- 警告：需要注意、警示状态 -->
  <span class="badge-warning">余额不足</span>
  <p class="text-warning">请注意配置</p>

  <!-- 危险：错误、删除操作 -->
  <span class="badge-danger">已停用</span>
  <button class="btn-danger">删除</button>

  <!-- 灰度：次要文字、边框、背景 -->
  <p class="text-gray-500">次要说明文字</p>
  <div class="border-gray-200 bg-gray-50">背景区域</div>
</template>
```

### 文字颜色层级

| 层级 | 用途 | 类名 |
|------|------|------|
| 一级 | 标题、重要内容 | `text-gray-900` |
| 二级 | 正文内容 | `text-gray-700` |
| 三级 | 次要信息 | `text-gray-500` |
| 四级 | 占位符、禁用 | `text-gray-400` |

---

## 响应式设计

### 断点定义

| 断点 | 宽度 | 场景 |
|------|------|------|
| `sm` | 640px | 手机横屏 |
| `md` | 768px | 平板 |
| `lg` | 1024px | 小屏笔记本 |
| `xl` | 1280px | 桌面显示器 |
| `2xl` | 1536px | 大屏显示器 |

### 响应式编写

```vue
<template>
  <!-- 移动优先 -->
  <div class="
    p-4 md:p-6 lg:p-8
    grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3
    text-sm md:text-base
  ">
    内容
  </div>

  <!-- 隐藏/显示 -->
  <div class="hidden md:block">仅桌面端显示</div>
  <div class="block md:hidden">仅移动端显示</div>
</template>
```

---

## 暗色模式

### 配置支持

```javascript
// tailwind.config.js
export default {
  darkMode: 'class', // 或 'media'
  // ...
}
```

### 样式编写

```vue
<template>
  <div class="
    bg-white dark:bg-gray-800
    text-gray-900 dark:text-gray-100
    border-gray-200 dark:border-gray-700
  ">
    支持暗色模式的内容
  </div>
</template>
```

### 主题切换

```typescript
// src/composables/useTheme.ts
import { ref, watchEffect } from 'vue'

type Theme = 'light' | 'dark' | 'system'

export function useTheme() {
  const theme = ref<Theme>(
    (localStorage.getItem('theme') as Theme) || 'system'
  )

  const applyTheme = (t: Theme) => {
    const isDark =
      t === 'dark' ||
      (t === 'system' && window.matchMedia('(prefers-color-scheme: dark)').matches)

    document.documentElement.classList.toggle('dark', isDark)
  }

  const setTheme = (t: Theme) => {
    theme.value = t
    localStorage.setItem('theme', t)
    applyTheme(t)
  }

  watchEffect(() => applyTheme(theme.value))

  return { theme, setTheme }
}
```

---

## 动画与过渡

### 过渡效果

```vue
<template>
  <!-- 基础过渡 -->
  <button class="transition-colors duration-200 hover:bg-gray-100">
    悬停变色
  </button>

  <!-- 多属性过渡 -->
  <div class="transition-all duration-300 ease-out hover:scale-105 hover:shadow-lg">
    悬停放大
  </div>

  <!-- Vue 过渡 -->
  <Transition name="fade">
    <div v-if="show">淡入淡出内容</div>
  </Transition>
</template>

<style>
.fade-enter-active,
.fade-leave-active {
  @apply transition-opacity duration-200;
}

.fade-enter-from,
.fade-leave-to {
  @apply opacity-0;
}
</style>
```

### 预设动画

```vue
<template>
  <!-- 加载旋转 -->
  <div class="animate-spin">⟳</div>

  <!-- 脉冲 -->
  <div class="animate-pulse">加载中...</div>

  <!-- 自定义动画 -->
  <div class="animate-fade-in">淡入内容</div>
  <div class="animate-slide-up">上滑内容</div>
</template>
```

---

## 编码规范

### 类名顺序

遵循以下顺序组织 Tailwind 类名：

1. 布局（`flex`, `grid`, `block`）
2. 定位（`relative`, `absolute`）
3. 尺寸（`w-`, `h-`, `max-w-`）
4. 间距（`m-`, `p-`, `gap-`）
5. 排版（`text-`, `font-`）
6. 颜色（`bg-`, `text-`, `border-`）
7. 边框（`border`, `rounded`）
8. 效果（`shadow`, `opacity`）
9. 过渡（`transition`, `duration`）
10. 交互（`hover:`, `focus:`）

```vue
<!-- 推荐 -->
<div class="flex items-center justify-between w-full p-4 text-sm text-gray-700 bg-white border rounded-lg shadow-sm transition-colors hover:bg-gray-50">

<!-- 不推荐：随意排列 -->
<div class="hover:bg-gray-50 text-sm p-4 flex shadow-sm border bg-white rounded-lg items-center justify-between w-full text-gray-700 transition-colors">
```

### 避免过长类名

```vue
<!-- 类名过长时，使用 @apply 提取 -->
<template>
  <button class="action-btn">按钮</button>
</template>

<style scoped>
.action-btn {
  @apply inline-flex items-center justify-center gap-2
         px-4 py-2 text-sm font-medium
         bg-primary-600 text-white rounded-md
         transition-colors duration-200
         hover:bg-primary-700
         focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500
         disabled:opacity-50 disabled:cursor-not-allowed;
}
</style>
```

### 组件样式隔离

```vue
<template>
  <div class="my-component">
    <span class="label">标签</span>
  </div>
</template>

<style scoped>
/* 使用 scoped 避免样式污染 */
.my-component {
  @apply relative;
}

.label {
  @apply text-sm text-gray-600;
}

/* 深度选择器（影响子组件） */
:deep(.child-class) {
  @apply text-primary-600;
}
</style>
```
