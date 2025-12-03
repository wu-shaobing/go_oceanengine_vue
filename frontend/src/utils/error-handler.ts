/**
 * 错误处理工具
 */

export interface ErrorInfo {
  message: string
  code?: string | number
  stack?: string
  type: 'api' | 'runtime' | 'promise' | 'vue' | 'network'
  timestamp: number
  url?: string
  componentName?: string
}

// 错误日志队列
const errorQueue: ErrorInfo[] = []
const MAX_ERROR_QUEUE_SIZE = 50

/**
 * 添加错误到队列
 */
const addErrorToQueue = (error: ErrorInfo) => {
  errorQueue.unshift(error)
  if (errorQueue.length > MAX_ERROR_QUEUE_SIZE) {
    errorQueue.pop()
  }
}

/**
 * 获取错误队列
 */
export const getErrorQueue = (): ErrorInfo[] => {
  return [...errorQueue]
}

/**
 * 清空错误队列
 */
export const clearErrorQueue = () => {
  errorQueue.length = 0
}

/**
 * 处理 API 错误
 */
export const handleApiError = (error: unknown, url?: string): ErrorInfo => {
  const errorInfo: ErrorInfo = {
    message: '请求失败',
    type: 'api',
    timestamp: Date.now(),
    url
  }
  
  if (error instanceof Error) {
    errorInfo.message = error.message
    errorInfo.stack = error.stack
  } else if (typeof error === 'object' && error !== null) {
    const err = error as Record<string, unknown>
    errorInfo.message = (err.message as string) || (err.msg as string) || '请求失败'
    errorInfo.code = err.code as string | number
  }
  
  addErrorToQueue(errorInfo)
  return errorInfo
}

/**
 * 处理运行时错误
 */
export const handleRuntimeError = (error: Error): ErrorInfo => {
  const errorInfo: ErrorInfo = {
    message: error.message,
    stack: error.stack,
    type: 'runtime',
    timestamp: Date.now()
  }
  
  addErrorToQueue(errorInfo)
  return errorInfo
}

/**
 * 处理 Promise 拒绝
 */
export const handlePromiseError = (reason: unknown): ErrorInfo => {
  const errorInfo: ErrorInfo = {
    message: '未处理的 Promise 错误',
    type: 'promise',
    timestamp: Date.now()
  }
  
  if (reason instanceof Error) {
    errorInfo.message = reason.message
    errorInfo.stack = reason.stack
  } else if (typeof reason === 'string') {
    errorInfo.message = reason
  }
  
  addErrorToQueue(errorInfo)
  return errorInfo
}

/**
 * 处理 Vue 组件错误
 */
export const handleVueError = (error: Error, componentName?: string): ErrorInfo => {
  const errorInfo: ErrorInfo = {
    message: error.message,
    stack: error.stack,
    type: 'vue',
    timestamp: Date.now(),
    componentName
  }
  
  addErrorToQueue(errorInfo)
  return errorInfo
}

/**
 * 处理网络错误
 */
export const handleNetworkError = (status?: number): ErrorInfo => {
  let message = '网络错误'
  
  switch (status) {
    case 400:
      message = '请求参数错误'
      break
    case 401:
      message = '未授权，请重新登录'
      break
    case 403:
      message = '拒绝访问'
      break
    case 404:
      message = '请求资源不存在'
      break
    case 408:
      message = '请求超时'
      break
    case 500:
      message = '服务器内部错误'
      break
    case 501:
      message = '服务未实现'
      break
    case 502:
      message = '网关错误'
      break
    case 503:
      message = '服务不可用'
      break
    case 504:
      message = '网关超时'
      break
    default:
      message = `连接错误 ${status || ''}`
  }
  
  const errorInfo: ErrorInfo = {
    message,
    code: status,
    type: 'network',
    timestamp: Date.now()
  }
  
  addErrorToQueue(errorInfo)
  return errorInfo
}

/**
 * 显示错误消息
 */
export const showErrorMessage = (message: string, duration = 3000) => {
  // 可以替换为实际的 toast/notification 组件
  console.error('[Error]', message)
  
  // 简单的错误提示实现
  if (typeof document !== 'undefined') {
    const toast = document.createElement('div')
    toast.className = 'error-toast'
    toast.textContent = message
    toast.style.cssText = `
      position: fixed;
      top: 20px;
      left: 50%;
      transform: translateX(-50%);
      background-color: #ef4444;
      color: white;
      padding: 12px 24px;
      border-radius: 8px;
      font-size: 14px;
      z-index: 10000;
      box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
      animation: fadeIn 0.3s ease;
    `
    
    document.body.appendChild(toast)
    
    setTimeout(() => {
      toast.style.animation = 'fadeOut 0.3s ease'
      setTimeout(() => toast.remove(), 300)
    }, duration)
  }
}

/**
 * 初始化全局错误处理
 */
export const initErrorHandler = () => {
  // 处理未捕获的错误
  window.addEventListener('error', (event) => {
    handleRuntimeError(event.error || new Error(event.message))
  })
  
  // 处理未处理的 Promise 拒绝
  window.addEventListener('unhandledrejection', (event) => {
    handlePromiseError(event.reason)
  })
}

export default {
  handleApiError,
  handleRuntimeError,
  handlePromiseError,
  handleVueError,
  handleNetworkError,
  showErrorMessage,
  initErrorHandler,
  getErrorQueue,
  clearErrorQueue
}
