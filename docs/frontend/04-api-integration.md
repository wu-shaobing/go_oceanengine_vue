# API 集成

## 概述

本项目使用 Axios 进行 HTTP 请求，封装了统一的请求/响应拦截、错误处理和类型支持。

## Axios 封装

### 请求实例

```typescript
// src/api/request.ts
import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, AxiosError } from 'axios'
import { useAuthStore } from '@/stores/auth'
import { useAppStore } from '@/stores/app'
import router from '@/router'

// 响应数据结构
interface ApiResponse<T = any> {
  code: number
  message: string
  data: T
  request_id?: string
}

// 分页响应
interface PageResponse<T> {
  list: T[]
  total: number
  page: number
  page_size: number
}

// 创建实例
const instance: AxiosInstance = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '/api/v1',
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// 请求拦截器
instance.interceptors.request.use(
  (config) => {
    const authStore = useAuthStore()
    const appStore = useAppStore()
    
    // 添加 Token
    if (authStore.token) {
      config.headers.Authorization = `Bearer ${authStore.token}`
    }
    
    // 显示加载状态
    appStore.setLoading(true)
    
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
instance.interceptors.response.use(
  (response: AxiosResponse<ApiResponse>) => {
    const appStore = useAppStore()
    appStore.setLoading(false)
    
    const { code, message, data } = response.data
    
    // 业务成功
    if (code === 0) {
      return data
    }
    
    // 业务错误
    return Promise.reject(new Error(message || '请求失败'))
  },
  async (error: AxiosError<ApiResponse>) => {
    const appStore = useAppStore()
    appStore.setLoading(false)
    
    const authStore = useAuthStore()
    const status = error.response?.status
    const data = error.response?.data
    
    // Token 过期，尝试刷新
    if (status === 401 && authStore.refreshToken) {
      try {
        await authStore.refreshAccessToken()
        // 重试原请求
        return instance(error.config!)
      } catch {
        authStore.clearAuth()
        router.push('/login')
        return Promise.reject(new Error('登录已过期，请重新登录'))
      }
    }
    
    // 权限不足
    if (status === 403) {
      router.push('/403')
      return Promise.reject(new Error('没有访问权限'))
    }
    
    // 资源不存在
    if (status === 404) {
      return Promise.reject(new Error('请求的资源不存在'))
    }
    
    // 服务器错误
    if (status && status >= 500) {
      return Promise.reject(new Error('服务器错误，请稍后重试'))
    }
    
    // 网络错误
    if (error.message === 'Network Error') {
      return Promise.reject(new Error('网络连接失败，请检查网络'))
    }
    
    // 超时
    if (error.code === 'ECONNABORTED') {
      return Promise.reject(new Error('请求超时，请稍后重试'))
    }
    
    return Promise.reject(new Error(data?.message || '请求失败'))
  }
)

// 封装请求方法
export const request = {
  get<T>(url: string, params?: object, config?: AxiosRequestConfig): Promise<T> {
    return instance.get(url, { params, ...config })
  },
  
  post<T>(url: string, data?: object, config?: AxiosRequestConfig): Promise<T> {
    return instance.post(url, data, config)
  },
  
  put<T>(url: string, data?: object, config?: AxiosRequestConfig): Promise<T> {
    return instance.put(url, data, config)
  },
  
  delete<T>(url: string, params?: object, config?: AxiosRequestConfig): Promise<T> {
    return instance.delete(url, { params, ...config })
  },
  
  upload<T>(url: string, file: File, onProgress?: (percent: number) => void): Promise<T> {
    const formData = new FormData()
    formData.append('file', file)
    
    return instance.post(url, formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
      onUploadProgress: (e) => {
        if (e.total && onProgress) {
          onProgress(Math.round((e.loaded * 100) / e.total))
        }
      }
    })
  }
}

export default instance
export type { ApiResponse, PageResponse }
```

---

## API 模块定义

### 认证 API

```typescript
// src/api/auth.ts
import { request } from './request'

export interface LoginRequest {
  username: string
  password: string
  captcha_id?: string
  captcha_code?: string
}

export interface LoginResponse {
  access_token: string
  refresh_token: string
  expires_in: number
}

export interface UserInfo {
  id: number
  username: string
  nickname: string
  avatar: string
  email: string
  phone: string
  role: {
    id: number
    name: string
    key: string
  }
}

export const authApi = {
  // 登录
  login(data: LoginRequest) {
    return request.post<LoginResponse>('/auth/login', data)
  },
  
  // 登出
  logout() {
    return request.post<void>('/auth/logout')
  },
  
  // 刷新 Token
  refreshToken(refreshToken: string) {
    return request.post<LoginResponse>('/auth/refresh', { refresh_token: refreshToken })
  },
  
  // 获取用户信息
  getUserInfo() {
    return request.get<{ user: UserInfo; permissions: string[] }>('/auth/info')
  },
  
  // 获取验证码
  getCaptcha() {
    return request.get<{ captcha_id: string; captcha_image: string }>('/auth/captcha')
  },
  
  // 修改密码
  changePassword(data: { old_password: string; new_password: string }) {
    return request.post<void>('/auth/password', data)
  }
}
```

### 广告主 API

```typescript
// src/api/advertiser.ts
import { request, PageResponse } from './request'

export interface Advertiser {
  id: number
  advertiser_id: number
  name: string
  company: string
  status: string
  balance: number
  valid_balance: number
  created_at: string
  last_sync_at: string
}

export interface AdvertiserListParams {
  page: number
  page_size: number
  keyword?: string
  status?: string
}

export const advertiserApi = {
  // 获取列表
  getList(params: AdvertiserListParams) {
    return request.get<PageResponse<Advertiser>>('/advertiser', params)
  },
  
  // 获取详情
  getDetail(id: number) {
    return request.get<Advertiser>(`/advertiser/${id}`)
  },
  
  // 同步数据
  sync(id: number) {
    return request.post<void>(`/advertiser/${id}/sync`)
  },
  
  // 获取余额
  getBalance(id: number) {
    return request.get<{
      balance: number
      valid_balance: number
      cash_balance: number
    }>(`/advertiser/${id}/balance`)
  },
  
  // 获取资金流水
  getFundTransactions(id: number, params: { start_date: string; end_date: string; page: number; page_size: number }) {
    return request.get<PageResponse<{
      transaction_seq: string
      transaction_type: string
      amount: number
      transaction_time: string
    }>>(`/advertiser/${id}/funds`, params)
  }
}
```

### 广告系列 API

```typescript
// src/api/campaign.ts
import { request, PageResponse } from './request'

export interface Campaign {
  id: number
  campaign_id: number
  advertiser_id: number
  name: string
  budget_mode: string
  budget: number
  landing_type: string
  status: string
  opt_status: string
  created_at: string
}

export interface CampaignCreateRequest {
  advertiser_id: number
  name: string
  budget_mode: 'BUDGET_MODE_INFINITE' | 'BUDGET_MODE_DAY'
  budget?: number
  landing_type: string
  marketing_goal?: string
}

export const campaignApi = {
  // 获取列表
  getList(params: {
    advertiser_id: number
    page: number
    page_size: number
    status?: string
    name?: string
  }) {
    return request.get<PageResponse<Campaign>>('/campaign', params)
  },
  
  // 获取详情
  getDetail(id: number) {
    return request.get<Campaign>(`/campaign/${id}`)
  },
  
  // 创建
  create(data: CampaignCreateRequest) {
    return request.post<{ id: number }>('/campaign', data)
  },
  
  // 更新
  update(id: number, data: Partial<CampaignCreateRequest>) {
    return request.put<void>(`/campaign/${id}`, data)
  },
  
  // 更新状态
  updateStatus(id: number, status: 'enable' | 'disable' | 'delete') {
    return request.post<void>(`/campaign/${id}/status`, { status })
  },
  
  // 批量更新状态
  batchUpdateStatus(ids: number[], status: 'enable' | 'disable' | 'delete') {
    return request.post<void>('/campaign/batch/status', { ids, status })
  }
}
```

### 数据报表 API

```typescript
// src/api/report.ts
import { request, PageResponse } from './request'

export interface ReportData {
  stat_datetime: string
  cost: number
  show: number
  click: number
  ctr: number
  convert: number
  convert_cost: number
  convert_rate: number
}

export interface ReportParams {
  advertiser_id: number
  start_date: string
  end_date: string
  group_by?: string[]
  page?: number
  page_size?: number
}

export const reportApi = {
  // 广告主报表
  getAdvertiserReport(params: ReportParams) {
    return request.get<ReportData[]>('/report/advertiser', params)
  },
  
  // 广告系列报表
  getCampaignReport(params: ReportParams & { campaign_ids?: number[] }) {
    return request.get<PageResponse<ReportData & { campaign_id: number; campaign_name: string }>>('/report/campaign', params)
  },
  
  // 实时数据
  getRealtime(advertiser_id: number) {
    return request.get<{
      cost: number
      show: number
      click: number
      convert: number
      update_time: string
    }>('/report/realtime', { advertiser_id })
  },
  
  // 导出报表
  exportReport(params: ReportParams & { type: 'advertiser' | 'campaign' | 'ad' }) {
    return request.post<{ task_id: string }>('/report/export', params)
  },
  
  // 获取导出结果
  getExportResult(taskId: string) {
    return request.get<{ status: string; url?: string }>(`/report/export/${taskId}`)
  }
}
```

---

## 使用示例

### 在组件中使用

```vue
<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { advertiserApi, type Advertiser } from '@/api/advertiser'

const loading = ref(false)
const advertisers = ref<Advertiser[]>([])
const total = ref(0)

const fetchData = async () => {
  loading.value = true
  try {
    const res = await advertiserApi.getList({
      page: 1,
      page_size: 10
    })
    advertisers.value = res.list
    total.value = res.total
  } catch (error) {
    console.error('获取数据失败:', error)
  } finally {
    loading.value = false
  }
}

onMounted(fetchData)
</script>
```

### 使用组合式函数

```typescript
// src/composables/useRequest.ts
import { ref, UnwrapRef } from 'vue'

export function useRequest<T, P extends any[]>(
  fn: (...args: P) => Promise<T>
) {
  const data = ref<T | null>(null)
  const loading = ref(false)
  const error = ref<Error | null>(null)

  const execute = async (...args: P): Promise<T | null> => {
    loading.value = true
    error.value = null
    
    try {
      const result = await fn(...args)
      data.value = result as UnwrapRef<T>
      return result
    } catch (e) {
      error.value = e as Error
      return null
    } finally {
      loading.value = false
    }
  }

  return {
    data,
    loading,
    error,
    execute
  }
}

// 使用示例
const { data, loading, execute } = useRequest(advertiserApi.getList)

onMounted(() => {
  execute({ page: 1, page_size: 10 })
})
```

---

## 错误处理

### 全局错误处理

```typescript
// src/utils/error-handler.ts
import { showToast } from '@/utils/toast'

export function handleApiError(error: Error): void {
  const message = error.message || '操作失败'
  
  // 显示错误提示
  showToast({
    type: 'error',
    message
  })
  
  // 上报错误（可选）
  // reportError(error)
}
```

### 在组件中处理

```vue
<script setup lang="ts">
import { handleApiError } from '@/utils/error-handler'

const handleSubmit = async () => {
  try {
    await campaignApi.create(formData)
    showToast({ type: 'success', message: '创建成功' })
    router.push('/campaign')
  } catch (error) {
    handleApiError(error as Error)
  }
}
</script>
```

---

## 取消请求

```typescript
// 使用 AbortController 取消请求
const controller = new AbortController()

const fetchData = async () => {
  try {
    const data = await request.get('/api/data', {}, {
      signal: controller.signal
    })
  } catch (error) {
    if (axios.isCancel(error)) {
      console.log('请求已取消')
    }
  }
}

// 取消请求
controller.abort()
```

```typescript
// 组合式函数封装
export function useCancelableRequest<T>(
  fn: (signal: AbortSignal) => Promise<T>
) {
  const controller = ref<AbortController | null>(null)
  
  const execute = async () => {
    // 取消之前的请求
    controller.value?.abort()
    controller.value = new AbortController()
    
    return fn(controller.value.signal)
  }
  
  const cancel = () => {
    controller.value?.abort()
  }
  
  onUnmounted(cancel)
  
  return { execute, cancel }
}
```
