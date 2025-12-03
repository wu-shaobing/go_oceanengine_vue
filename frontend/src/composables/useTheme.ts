import { ref, watch, computed } from 'vue'
import { storage } from '@/utils/storage'

export type ThemeMode = 'light' | 'dark' | 'system'

const THEME_STORAGE_KEY = 'theme-mode'

// 全局状态
const themeMode = ref<ThemeMode>('system')
const systemDark = ref(false)

// 检测系统主题
const detectSystemTheme = () => {
  if (typeof window !== 'undefined' && window.matchMedia) {
    systemDark.value = window.matchMedia('(prefers-color-scheme: dark)').matches
  }
}

// 初始化系统主题监听
if (typeof window !== 'undefined' && window.matchMedia) {
  const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)')
  mediaQuery.addEventListener('change', (e) => {
    systemDark.value = e.matches
  })
  detectSystemTheme()
}

// 从存储加载主题
const savedTheme = storage.get<ThemeMode>(THEME_STORAGE_KEY)
if (savedTheme && ['light', 'dark', 'system'].includes(savedTheme)) {
  themeMode.value = savedTheme
}

export function useTheme() {
  // 实际应用的主题
  const isDark = computed(() => {
    if (themeMode.value === 'system') {
      return systemDark.value
    }
    return themeMode.value === 'dark'
  })
  
  // 主题名称
  const themeName = computed(() => isDark.value ? 'dark' : 'light')
  
  // 设置主题模式
  const setThemeMode = (mode: ThemeMode) => {
    themeMode.value = mode
    storage.set(THEME_STORAGE_KEY, mode)
  }
  
  // 切换主题
  const toggleTheme = () => {
    if (themeMode.value === 'light') {
      setThemeMode('dark')
    } else if (themeMode.value === 'dark') {
      setThemeMode('system')
    } else {
      setThemeMode('light')
    }
  }
  
  // 应用主题到 DOM
  const applyTheme = () => {
    if (typeof document !== 'undefined') {
      const root = document.documentElement
      if (isDark.value) {
        root.classList.add('dark')
        root.style.colorScheme = 'dark'
      } else {
        root.classList.remove('dark')
        root.style.colorScheme = 'light'
      }
    }
  }
  
  // 监听主题变化并应用
  watch(isDark, applyTheme, { immediate: true })
  
  // 主题颜色配置
  const colors = computed(() => {
    if (isDark.value) {
      return {
        background: '#1f2937',
        surface: '#374151',
        primary: '#3b82f6',
        secondary: '#6b7280',
        text: '#f9fafb',
        textSecondary: '#9ca3af',
        border: '#4b5563',
        error: '#ef4444',
        success: '#22c55e',
        warning: '#f59e0b'
      }
    }
    return {
      background: '#ffffff',
      surface: '#f9fafb',
      primary: '#3b82f6',
      secondary: '#6b7280',
      text: '#111827',
      textSecondary: '#6b7280',
      border: '#e5e7eb',
      error: '#ef4444',
      success: '#22c55e',
      warning: '#f59e0b'
    }
  })
  
  // 获取主题图标
  const themeIcon = computed(() => {
    switch (themeMode.value) {
      case 'light':
        return 'sun'
      case 'dark':
        return 'moon'
      default:
        return 'computer'
    }
  })
  
  // 获取主题标签
  const themeLabel = computed(() => {
    switch (themeMode.value) {
      case 'light':
        return '浅色模式'
      case 'dark':
        return '深色模式'
      default:
        return '跟随系统'
    }
  })
  
  return {
    themeMode,
    isDark,
    themeName,
    colors,
    themeIcon,
    themeLabel,
    setThemeMode,
    toggleTheme,
    applyTheme
  }
}
