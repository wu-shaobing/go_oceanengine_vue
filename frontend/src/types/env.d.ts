/// <reference types="vite/client" />

interface ImportMetaEnv {
  /**
   * 应用标题
   */
  readonly VITE_APP_TITLE: string
  
  /**
   * API 基础路径
   */
  readonly VITE_API_BASE_URL: string
  
  /**
   * API 超时时间（毫秒）
   */
  readonly VITE_API_TIMEOUT: string
  
  /**
   * 是否启用 Mock
   */
  readonly VITE_ENABLE_MOCK: string
  
  /**
   * 是否启用压缩
   */
  readonly VITE_BUILD_COMPRESS: string
  
  /**
   * 压缩类型
   */
  readonly VITE_BUILD_COMPRESS_TYPE: 'gzip' | 'brotli' | 'both'
  
  /**
   * 是否删除 console
   */
  readonly VITE_DROP_CONSOLE: string
  
  /**
   * 公共路径
   */
  readonly VITE_PUBLIC_PATH: string
  
  /**
   * 是否启用 HTTPS
   */
  readonly VITE_USE_HTTPS: string
  
  /**
   * 代理目标地址
   */
  readonly VITE_PROXY_TARGET: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
