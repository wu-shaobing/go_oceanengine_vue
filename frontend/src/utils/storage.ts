const PREFIX = 'oceanengine_'

/**
 * 本地存储工具
 */
export const storage = {
  get<T>(key: string, defaultValue?: T): T | null {
    try {
      const item = localStorage.getItem(PREFIX + key)
      if (item === null) return defaultValue ?? null
      return JSON.parse(item)
    } catch {
      return defaultValue ?? null
    }
  },

  set<T>(key: string, value: T): void {
    try {
      localStorage.setItem(PREFIX + key, JSON.stringify(value))
    } catch (e) {
      console.error('Storage set error:', e)
    }
  },

  remove(key: string): void {
    localStorage.removeItem(PREFIX + key)
  },

  clear(): void {
    const keys = Object.keys(localStorage)
    keys.forEach(key => {
      if (key.startsWith(PREFIX)) {
        localStorage.removeItem(key)
      }
    })
  }
}

/**
 * 会话存储工具
 */
export const sessionStorage = {
  get<T>(key: string, defaultValue?: T): T | null {
    try {
      const item = window.sessionStorage.getItem(PREFIX + key)
      if (item === null) return defaultValue ?? null
      return JSON.parse(item)
    } catch {
      return defaultValue ?? null
    }
  },

  set<T>(key: string, value: T): void {
    try {
      window.sessionStorage.setItem(PREFIX + key, JSON.stringify(value))
    } catch (e) {
      console.error('SessionStorage set error:', e)
    }
  },

  remove(key: string): void {
    window.sessionStorage.removeItem(PREFIX + key)
  },

  clear(): void {
    const keys = Object.keys(window.sessionStorage)
    keys.forEach(key => {
      if (key.startsWith(PREFIX)) {
        window.sessionStorage.removeItem(key)
      }
    })
  }
}
