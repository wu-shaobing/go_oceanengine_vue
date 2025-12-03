# 状态管理

## 概述

本项目使用 Pinia 作为状态管理工具，它是 Vue 3 官方推荐的状态管理库。

## Pinia 配置

### 安装与初始化

```typescript
// src/stores/index.ts
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)

export default pinia
```

```typescript
// src/main.ts
import { createApp } from 'vue'
import App from './App.vue'
import pinia from './stores'

const app = createApp(App)
app.use(pinia)
app.mount('#app')
```

---

## Store 定义

### 1. Auth Store - 认证状态

```typescript
// src/stores/auth.ts
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { User, LoginRequest, LoginResponse } from '@/types/auth'
import { authApi } from '@/api/auth'
import { storage } from '@/utils/storage'

export const useAuthStore = defineStore('auth', () => {
  // State
  const token = ref<string>(storage.get('token') || '')
  const refreshToken = ref<string>(storage.get('refreshToken') || '')
  const user = ref<User | null>(null)
  const permissions = ref<string[]>([])

  // Getters
  const isLoggedIn = computed(() => !!token.value)
  const username = computed(() => user.value?.username || '')
  const avatar = computed(() => user.value?.avatar || '/default-avatar.png')

  // Actions
  const login = async (params: LoginRequest): Promise<void> => {
    const res = await authApi.login(params)
    setToken(res.access_token, res.refresh_token)
    await fetchUserInfo()
  }

  const logout = async (): Promise<void> => {
    try {
      await authApi.logout()
    } finally {
      clearAuth()
    }
  }

  const setToken = (accessToken: string, refresh: string): void => {
    token.value = accessToken
    refreshToken.value = refresh
    storage.set('token', accessToken)
    storage.set('refreshToken', refresh)
  }

  const fetchUserInfo = async (): Promise<void> => {
    const res = await authApi.getUserInfo()
    user.value = res.user
    permissions.value = res.permissions
  }

  const refreshAccessToken = async (): Promise<string> => {
    const res = await authApi.refreshToken(refreshToken.value)
    setToken(res.access_token, res.refresh_token)
    return res.access_token
  }

  const clearAuth = (): void => {
    token.value = ''
    refreshToken.value = ''
    user.value = null
    permissions.value = []
    storage.remove('token')
    storage.remove('refreshToken')
  }

  const hasPermission = (permission: string): boolean => {
    return permissions.value.includes(permission) || permissions.value.includes('*')
  }

  return {
    // State
    token,
    refreshToken,
    user,
    permissions,
    // Getters
    isLoggedIn,
    username,
    avatar,
    // Actions
    login,
    logout,
    setToken,
    fetchUserInfo,
    refreshAccessToken,
    clearAuth,
    hasPermission
  }
}, {
  persist: {
    key: 'auth',
    paths: ['token', 'refreshToken']
  }
})
```

### 2. App Store - 应用状态

```typescript
// src/stores/app.ts
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useAppStore = defineStore('app', () => {
  // State
  const sidebarCollapsed = ref(false)
  const theme = ref<'light' | 'dark'>('light')
  const locale = ref('zh-CN')
  const loading = ref(false)

  // Getters
  const isDark = computed(() => theme.value === 'dark')

  // Actions
  const toggleSidebar = (): void => {
    sidebarCollapsed.value = !sidebarCollapsed.value
  }

  const setTheme = (newTheme: 'light' | 'dark'): void => {
    theme.value = newTheme
    document.documentElement.classList.toggle('dark', newTheme === 'dark')
  }

  const setLocale = (newLocale: string): void => {
    locale.value = newLocale
  }

  const setLoading = (value: boolean): void => {
    loading.value = value
  }

  return {
    sidebarCollapsed,
    theme,
    locale,
    loading,
    isDark,
    toggleSidebar,
    setTheme,
    setLocale,
    setLoading
  }
}, {
  persist: {
    key: 'app',
    paths: ['sidebarCollapsed', 'theme', 'locale']
  }
})
```

### 3. Advertiser Store - 广告主状态

```typescript
// src/stores/advertiser.ts
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { Advertiser, AdvertiserListParams } from '@/types/advertiser'
import { advertiserApi } from '@/api/advertiser'

export const useAdvertiserStore = defineStore('advertiser', () => {
  // State
  const advertisers = ref<Advertiser[]>([])
  const currentAdvertiser = ref<Advertiser | null>(null)
  const total = ref(0)
  const loading = ref(false)
  const params = ref<AdvertiserListParams>({
    page: 1,
    page_size: 10,
    keyword: '',
    status: undefined
  })

  // Getters
  const advertiserOptions = computed(() =>
    advertisers.value.map(a => ({
      label: a.name,
      value: a.id
    }))
  )

  const hasMore = computed(() =>
    advertisers.value.length < total.value
  )

  // Actions
  const fetchList = async (): Promise<void> => {
    loading.value = true
    try {
      const res = await advertiserApi.getList(params.value)
      advertisers.value = res.list
      total.value = res.total
    } finally {
      loading.value = false
    }
  }

  const fetchDetail = async (id: number): Promise<Advertiser> => {
    loading.value = true
    try {
      const res = await advertiserApi.getDetail(id)
      currentAdvertiser.value = res
      return res
    } finally {
      loading.value = false
    }
  }

  const setParams = (newParams: Partial<AdvertiserListParams>): void => {
    params.value = { ...params.value, ...newParams }
  }

  const resetParams = (): void => {
    params.value = {
      page: 1,
      page_size: 10,
      keyword: '',
      status: undefined
    }
  }

  const syncAdvertiser = async (id: number): Promise<void> => {
    await advertiserApi.sync(id)
    await fetchDetail(id)
  }

  return {
    advertisers,
    currentAdvertiser,
    total,
    loading,
    params,
    advertiserOptions,
    hasMore,
    fetchList,
    fetchDetail,
    setParams,
    resetParams,
    syncAdvertiser
  }
})
```

---

## 使用示例

### 在组件中使用

```vue
<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { useAuthStore } from '@/stores/auth'
import { useAdvertiserStore } from '@/stores/advertiser'

// 使用 storeToRefs 保持响应式
const authStore = useAuthStore()
const { isLoggedIn, username } = storeToRefs(authStore)

const advertiserStore = useAdvertiserStore()
const { advertisers, loading, total } = storeToRefs(advertiserStore)

// 直接解构 actions（不需要 storeToRefs）
const { logout } = authStore
const { fetchList, setParams } = advertiserStore

// 页面加载时获取数据
onMounted(() => {
  fetchList()
})

// 搜索
const handleSearch = (keyword: string) => {
  setParams({ keyword, page: 1 })
  fetchList()
}
</script>

<template>
  <div>
    <p v-if="isLoggedIn">欢迎, {{ username }}</p>
    
    <div v-loading="loading">
      <div v-for="adv in advertisers" :key="adv.id">
        {{ adv.name }}
      </div>
    </div>
    
    <p>共 {{ total }} 条记录</p>
  </div>
</template>
```

### 组合式函数封装

```typescript
// src/composables/useAuth.ts
import { storeToRefs } from 'pinia'
import { useAuthStore } from '@/stores/auth'
import { useRouter } from 'vue-router'

export function useAuth() {
  const authStore = useAuthStore()
  const router = useRouter()
  
  const { isLoggedIn, user, permissions } = storeToRefs(authStore)
  
  const login = async (username: string, password: string) => {
    await authStore.login({ username, password })
    router.push('/dashboard')
  }
  
  const logout = async () => {
    await authStore.logout()
    router.push('/login')
  }
  
  const checkPermission = (permission: string) => {
    return authStore.hasPermission(permission)
  }
  
  return {
    isLoggedIn,
    user,
    permissions,
    login,
    logout,
    checkPermission
  }
}
```

---

## Store 之间的交互

```typescript
// src/stores/campaign.ts
import { defineStore, storeToRefs } from 'pinia'
import { useAdvertiserStore } from './advertiser'

export const useCampaignStore = defineStore('campaign', () => {
  const advertiserStore = useAdvertiserStore()
  const { currentAdvertiser } = storeToRefs(advertiserStore)

  const fetchCampaigns = async () => {
    if (!currentAdvertiser.value) {
      throw new Error('请先选择广告主')
    }
    
    const res = await campaignApi.getList({
      advertiser_id: currentAdvertiser.value.id
    })
    // ...
  }

  return { fetchCampaigns }
})
```

---

## 持久化配置

```typescript
// 完整持久化配置示例
export const useSettingsStore = defineStore('settings', () => {
  const fontSize = ref(14)
  const compactMode = ref(false)
  
  return { fontSize, compactMode }
}, {
  persist: {
    key: 'oceanengine-settings',     // 存储键名
    storage: localStorage,            // 存储方式
    paths: ['fontSize', 'compactMode'], // 持久化字段
    beforeRestore: (ctx) => {         // 恢复前钩子
      console.log('restoring', ctx.store.$id)
    },
    afterRestore: (ctx) => {          // 恢复后钩子
      console.log('restored', ctx.store.$id)
    }
  }
})
```

---

## 最佳实践

### 1. Store 职责单一
每个 Store 只管理一个领域的状态

### 2. 使用 storeToRefs
在组件中解构响应式状态时使用 `storeToRefs`

### 3. 避免直接修改状态
通过 actions 修改状态，便于追踪和调试

### 4. 合理使用持久化
只持久化必要的状态，避免存储敏感信息

### 5. 类型安全
为所有状态和参数定义完整的 TypeScript 类型
