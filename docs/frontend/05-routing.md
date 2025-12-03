# 路由设计

## 概述

本项目使用 Vue Router 4 实现路由管理，采用基于角色的访问控制，支持动态路由和懒加载。

## 路由配置

### 基础配置

```typescript
// src/router/index.ts
import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

// 布局组件
const MainLayout = () => import('@/components/layout/MainLayout.vue')
const AuthLayout = () => import('@/components/layout/AuthLayout.vue')

// 静态路由（无需认证）
export const constantRoutes: RouteRecordRaw[] = [
  {
    path: '/login',
    component: AuthLayout,
    children: [
      {
        path: '',
        name: 'Login',
        component: () => import('@/views/auth/LoginView.vue'),
        meta: { title: '登录' }
      }
    ]
  },
  {
    path: '/403',
    name: 'Forbidden',
    component: () => import('@/views/error/403.vue'),
    meta: { title: '无权限' }
  },
  {
    path: '/404',
    name: 'NotFound',
    component: () => import('@/views/error/404.vue'),
    meta: { title: '页面不存在' }
  }
]

// 动态路由（需要认证）
export const asyncRoutes: RouteRecordRaw[] = [
  {
    path: '/',
    component: MainLayout,
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/dashboard/DashboardView.vue'),
        meta: {
          title: '数据概览',
          icon: 'dashboard',
          permissions: ['dashboard:view']
        }
      }
    ]
  },
  {
    path: '/advertiser',
    component: MainLayout,
    redirect: '/advertiser/list',
    meta: {
      title: '广告主管理',
      icon: 'user-group',
      permissions: ['advertiser:view']
    },
    children: [
      {
        path: 'list',
        name: 'AdvertiserList',
        component: () => import('@/views/advertiser/AdvertiserList.vue'),
        meta: { title: '广告主列表' }
      },
      {
        path: ':id',
        name: 'AdvertiserDetail',
        component: () => import('@/views/advertiser/AdvertiserDetail.vue'),
        meta: { title: '广告主详情', hidden: true }
      }
    ]
  },
  {
    path: '/campaign',
    component: MainLayout,
    redirect: '/campaign/list',
    meta: {
      title: '广告系列',
      icon: 'folder',
      permissions: ['campaign:view']
    },
    children: [
      {
        path: 'list',
        name: 'CampaignList',
        component: () => import('@/views/campaign/CampaignList.vue'),
        meta: { title: '系列列表' }
      },
      {
        path: 'create',
        name: 'CampaignCreate',
        component: () => import('@/views/campaign/CampaignEdit.vue'),
        meta: { title: '创建系列', permissions: ['campaign:create'] }
      },
      {
        path: ':id/edit',
        name: 'CampaignEdit',
        component: () => import('@/views/campaign/CampaignEdit.vue'),
        meta: { title: '编辑系列', hidden: true, permissions: ['campaign:update'] }
      }
    ]
  },
  {
    path: '/report',
    component: MainLayout,
    redirect: '/report/dashboard',
    meta: {
      title: '数据报表',
      icon: 'chart-bar',
      permissions: ['report:view']
    },
    children: [
      {
        path: 'dashboard',
        name: 'ReportDashboard',
        component: () => import('@/views/report/ReportDashboard.vue'),
        meta: { title: '报表概览' }
      },
      {
        path: 'detail',
        name: 'ReportDetail',
        component: () => import('@/views/report/ReportDetail.vue'),
        meta: { title: '详细报表' }
      }
    ]
  },
  {
    path: '/audience',
    component: MainLayout,
    redirect: '/audience/list',
    meta: {
      title: '人群管理',
      icon: 'users',
      permissions: ['audience:view']
    },
    children: [
      {
        path: 'list',
        name: 'AudienceList',
        component: () => import('@/views/audience/AudienceList.vue'),
        meta: { title: '人群列表' }
      },
      {
        path: 'create',
        name: 'AudienceCreate',
        component: () => import('@/views/audience/AudienceEdit.vue'),
        meta: { title: '创建人群', permissions: ['audience:create'] }
      }
    ]
  },
  {
    path: '/system',
    component: MainLayout,
    redirect: '/system/user',
    meta: {
      title: '系统管理',
      icon: 'cog',
      permissions: ['system:view']
    },
    children: [
      {
        path: 'user',
        name: 'UserManage',
        component: () => import('@/views/system/UserManage.vue'),
        meta: { title: '用户管理', permissions: ['system:user'] }
      },
      {
        path: 'role',
        name: 'RoleManage',
        component: () => import('@/views/system/RoleManage.vue'),
        meta: { title: '角色管理', permissions: ['system:role'] }
      },
      {
        path: 'menu',
        name: 'MenuManage',
        component: () => import('@/views/system/MenuManage.vue'),
        meta: { title: '菜单管理', permissions: ['system:menu'] }
      }
    ]
  },
  // 404 兜底路由
  {
    path: '/:pathMatch(.*)*',
    redirect: '/404'
  }
]

// 创建路由实例
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: constantRoutes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    }
    return { top: 0 }
  }
})

export default router
```

---

## 路由守卫

### 全局守卫

```typescript
// src/router/guards.ts
import { Router } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useAppStore } from '@/stores/app'
import NProgress from 'nprogress'

// 白名单路由
const whiteList = ['/login', '/403', '/404']

export function setupRouterGuards(router: Router): void {
  // 前置守卫
  router.beforeEach(async (to, from, next) => {
    NProgress.start()
    
    const authStore = useAuthStore()
    const appStore = useAppStore()
    
    // 设置页面标题
    document.title = to.meta.title
      ? `${to.meta.title} - 巨量引擎管理系统`
      : '巨量引擎管理系统'
    
    // 白名单直接放行
    if (whiteList.includes(to.path)) {
      next()
      return
    }
    
    // 未登录跳转登录页
    if (!authStore.token) {
      next({ path: '/login', query: { redirect: to.fullPath } })
      return
    }
    
    // 已登录但未获取用户信息
    if (!authStore.userInfo) {
      try {
        await authStore.fetchUserInfo()
        
        // 根据权限生成动态路由
        const accessRoutes = await authStore.generateRoutes()
        accessRoutes.forEach(route => {
          router.addRoute(route)
        })
        
        // 重新导航到目标页面
        next({ ...to, replace: true })
      } catch (error) {
        authStore.clearAuth()
        next({ path: '/login', query: { redirect: to.fullPath } })
      }
      return
    }
    
    // 检查路由权限
    if (to.meta.permissions) {
      const hasPermission = authStore.hasPermission(to.meta.permissions as string[])
      if (!hasPermission) {
        next('/403')
        return
      }
    }
    
    next()
  })
  
  // 后置守卫
  router.afterEach((to) => {
    NProgress.done()
    
    // 记录访问历史（可用于标签页）
    const appStore = useAppStore()
    if (!to.meta.hidden) {
      appStore.addVisitedView(to)
    }
  })
  
  // 错误处理
  router.onError((error) => {
    console.error('路由错误:', error)
    NProgress.done()
  })
}
```

---

## 权限控制

### 动态路由生成

```typescript
// src/stores/auth.ts (部分)
import { asyncRoutes } from '@/router'
import type { RouteRecordRaw } from 'vue-router'

// 过滤有权限的路由
function filterAsyncRoutes(
  routes: RouteRecordRaw[],
  permissions: string[]
): RouteRecordRaw[] {
  const result: RouteRecordRaw[] = []
  
  routes.forEach(route => {
    const tmp = { ...route }
    
    // 检查路由权限
    if (hasRoutePermission(tmp, permissions)) {
      // 递归处理子路由
      if (tmp.children) {
        tmp.children = filterAsyncRoutes(tmp.children, permissions)
      }
      result.push(tmp)
    }
  })
  
  return result
}

function hasRoutePermission(
  route: RouteRecordRaw,
  permissions: string[]
): boolean {
  if (route.meta?.permissions) {
    const routePermissions = route.meta.permissions as string[]
    return routePermissions.some(p => permissions.includes(p))
  }
  return true
}

// Store action
async generateRoutes(): Promise<RouteRecordRaw[]> {
  const accessRoutes = filterAsyncRoutes(asyncRoutes, this.permissions)
  this.routes = accessRoutes
  return accessRoutes
}
```

### 权限指令

```typescript
// src/directives/permission.ts
import { Directive } from 'vue'
import { useAuthStore } from '@/stores/auth'

export const permissionDirective: Directive = {
  mounted(el, binding) {
    const authStore = useAuthStore()
    const { value } = binding
    
    if (value && value.length > 0) {
      const hasPermission = authStore.hasPermission(value)
      
      if (!hasPermission) {
        el.parentNode?.removeChild(el)
      }
    }
  }
}

// 注册指令
// main.ts
app.directive('permission', permissionDirective)
```

```vue
<!-- 使用示例 -->
<template>
  <button v-permission="['campaign:create']">
    创建广告系列
  </button>
</template>
```

---

## 路由元信息

### Meta 类型定义

```typescript
// src/types/router.d.ts
import 'vue-router'

declare module 'vue-router' {
  interface RouteMeta {
    // 页面标题
    title?: string
    // 菜单图标
    icon?: string
    // 所需权限
    permissions?: string[]
    // 是否隐藏（不在菜单显示）
    hidden?: boolean
    // 是否缓存
    keepAlive?: boolean
    // 是否固定标签
    affix?: boolean
    // 父级路由名称（用于面包屑）
    activeMenu?: string
    // 是否不显示面包屑
    breadcrumb?: boolean
  }
}
```

---

## 菜单生成

### 根据路由生成菜单

```typescript
// src/utils/menu.ts
import type { RouteRecordRaw } from 'vue-router'

export interface MenuItem {
  path: string
  title: string
  icon?: string
  children?: MenuItem[]
}

export function generateMenus(routes: RouteRecordRaw[]): MenuItem[] {
  const menus: MenuItem[] = []
  
  routes.forEach(route => {
    // 跳过隐藏路由
    if (route.meta?.hidden) return
    
    const menu: MenuItem = {
      path: route.path,
      title: route.meta?.title || '',
      icon: route.meta?.icon
    }
    
    // 处理子路由
    if (route.children && route.children.length > 0) {
      const visibleChildren = route.children.filter(
        child => !child.meta?.hidden
      )
      
      if (visibleChildren.length === 1) {
        // 只有一个子路由时，直接显示子路由
        const child = visibleChildren[0]
        menu.path = `${route.path}/${child.path}`.replace('//', '/')
        menu.title = child.meta?.title || menu.title
        menu.icon = child.meta?.icon || menu.icon
      } else if (visibleChildren.length > 1) {
        // 多个子路由时，递归生成子菜单
        menu.children = generateMenus(visibleChildren)
      }
    }
    
    menus.push(menu)
  })
  
  return menus
}
```

---

## 标签页导航

### 访问历史管理

```typescript
// src/stores/app.ts (部分)
import type { RouteLocationNormalized } from 'vue-router'

interface TagView {
  path: string
  name: string
  title: string
  affix?: boolean
}

export const useAppStore = defineStore('app', {
  state: () => ({
    visitedViews: [] as TagView[],
    cachedViews: [] as string[]
  }),
  
  actions: {
    addVisitedView(route: RouteLocationNormalized) {
      // 已存在则跳过
      if (this.visitedViews.some(v => v.path === route.path)) return
      
      this.visitedViews.push({
        path: route.path,
        name: route.name as string,
        title: route.meta.title as string || '未命名',
        affix: route.meta.affix
      })
      
      // 添加到缓存列表
      if (route.meta.keepAlive && route.name) {
        this.cachedViews.push(route.name as string)
      }
    },
    
    removeVisitedView(path: string) {
      const index = this.visitedViews.findIndex(v => v.path === path)
      if (index > -1) {
        const view = this.visitedViews[index]
        this.visitedViews.splice(index, 1)
        
        // 从缓存中移除
        const cachedIndex = this.cachedViews.indexOf(view.name)
        if (cachedIndex > -1) {
          this.cachedViews.splice(cachedIndex, 1)
        }
      }
    },
    
    closeOtherViews(path: string) {
      this.visitedViews = this.visitedViews.filter(
        v => v.path === path || v.affix
      )
    },
    
    closeAllViews() {
      this.visitedViews = this.visitedViews.filter(v => v.affix)
      this.cachedViews = []
    }
  }
})
```

### 页面缓存

```vue
<!-- MainLayout.vue -->
<template>
  <div class="main-layout">
    <AppSidebar />
    <div class="main-content">
      <AppHeader />
      <TagsView />
      <main class="page-content">
        <router-view v-slot="{ Component }">
          <keep-alive :include="cachedViews">
            <component :is="Component" />
          </keep-alive>
        </router-view>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useAppStore } from '@/stores/app'

const appStore = useAppStore()
const cachedViews = computed(() => appStore.cachedViews)
</script>
```

---

## 路由跳转

### 编程式导航

```typescript
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()

// 基本跳转
router.push('/dashboard')
router.push({ name: 'Dashboard' })

// 带参数
router.push({ name: 'AdvertiserDetail', params: { id: 123 } })
router.push({ path: '/campaign/list', query: { status: 'active' } })

// 替换当前页
router.replace('/login')

// 返回
router.back()
router.go(-1)

// 获取当前路由信息
console.log(route.path)
console.log(route.params.id)
console.log(route.query.keyword)
```

### 路由跳转工具

```typescript
// src/utils/navigation.ts
import router from '@/router'

export const navigation = {
  // 跳转到详情页
  toDetail(module: string, id: number | string) {
    router.push(`/${module}/${id}`)
  },
  
  // 跳转到编辑页
  toEdit(module: string, id: number | string) {
    router.push(`/${module}/${id}/edit`)
  },
  
  // 跳转到创建页
  toCreate(module: string) {
    router.push(`/${module}/create`)
  },
  
  // 带确认的返回
  goBackWithConfirm(hasChanges: boolean) {
    if (hasChanges) {
      if (confirm('有未保存的更改，确定离开吗？')) {
        router.back()
      }
    } else {
      router.back()
    }
  }
}
```
