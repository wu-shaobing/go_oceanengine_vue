import axios, { type AxiosInstance, type AxiosRequestConfig, type InternalAxiosRequestConfig } from 'axios'
import { mockLogin, mockUserInfo } from '@/mock/auth'

// 是否启用 Mock
const ENABLE_MOCK = import.meta.env.VITE_ENABLE_MOCK === 'true'

// 响应数据结构
export interface ApiResponse<T = unknown> {
  code: number
  message: string
  data: T
  request_id?: string
}

// 分页响应
export interface PageResponse<T> {
  list: T[]
  total: number
  page: number
  page_size: number
}

// Mock 请求处理
const handleMockRequest = (config: InternalAxiosRequestConfig): ApiResponse | null => {
  if (!ENABLE_MOCK) return null
  
  const { url, method, data } = config
  const body = typeof data === 'string' ? JSON.parse(data) : data
  
  // Mock 登录
  if (url === '/auth/login' && method === 'post') {
    return mockLogin(body.username, body.password)
  }
  
  // Mock 获取用户信息
  if (url === '/auth/info' && method === 'get') {
    const token = localStorage.getItem('access_token') || ''
    return mockUserInfo(token)
  }
  
  // Mock 登出
  if (url === '/auth/logout' && method === 'post') {
    return { code: 0, message: 'success', data: null }
  }
  
  return null
}

const instance: AxiosInstance = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '/api/v1',
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// Request interceptor
instance.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('access_token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// Response interceptor
instance.interceptors.response.use(
  (response) => {
    const res = response.data as ApiResponse
    if (res.code === 0) {
      return res.data as unknown as typeof response
    }
    return Promise.reject(new Error(res.message || '请求失败'))
  },
  (error) => {
    const message = error.response?.data?.message || error.message || '请求失败'
    console.error('API Error:', message)
    return Promise.reject(new Error(message))
  }
)

// 处理 mock 响应
const processMockResponse = <T>(mockResult: ApiResponse | null): Promise<T> | null => {
  if (!mockResult) return null
  
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      if (mockResult.code === 0) {
        resolve(mockResult.data as T)
      } else {
        reject(new Error(mockResult.message))
      }
    }, 200) // 模拟网络延迟
  })
}

const request = {
  get<T>(url: string, params?: object, config?: AxiosRequestConfig): Promise<T> {
    const mockResult = handleMockRequest({ url, method: 'get', data: params } as InternalAxiosRequestConfig)
    const mockResponse = processMockResponse<T>(mockResult)
    if (mockResponse) return mockResponse
    return instance.get(url, { params, ...config }) as Promise<T>
  },
  post<T>(url: string, data?: object, config?: AxiosRequestConfig): Promise<T> {
    const mockResult = handleMockRequest({ url, method: 'post', data } as InternalAxiosRequestConfig)
    const mockResponse = processMockResponse<T>(mockResult)
    if (mockResponse) return mockResponse
    return instance.post(url, data, config) as Promise<T>
  },
  put<T>(url: string, data?: object, config?: AxiosRequestConfig): Promise<T> {
    const mockResult = handleMockRequest({ url, method: 'put', data } as InternalAxiosRequestConfig)
    const mockResponse = processMockResponse<T>(mockResult)
    if (mockResponse) return mockResponse
    return instance.put(url, data, config) as Promise<T>
  },
  delete<T>(url: string, params?: object, config?: AxiosRequestConfig): Promise<T> {
    const mockResult = handleMockRequest({ url, method: 'delete', data: params } as InternalAxiosRequestConfig)
    const mockResponse = processMockResponse<T>(mockResult)
    if (mockResponse) return mockResponse
    return instance.delete(url, { params, ...config }) as Promise<T>
  }
}

export default request
