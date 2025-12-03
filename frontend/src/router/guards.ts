import type { Router } from 'vue-router'

// 白名单路由（无需认证）
const whiteList = ['/login', '/403', '/404']

// 简单的进度条实现
const createProgressBar = () => {
  let bar: HTMLElement | null = null
  
  return {
    start: () => {
      if (typeof document === 'undefined') return
      bar = document.createElement('div')
      bar.style.cssText = `
        position: fixed;
        top: 0;
        left: 0;
        height: 3px;
        background: linear-gradient(to right, #3b82f6, #1d4ed8);
        z-index: 9999;
        animation: progress 0.8s ease-in-out infinite;
        width: 0%;
      `
      document.body.appendChild(bar)
      requestAnimationFrame(() => {
        if (bar) bar.style.width = '80%'
      })
    },
    done: () => {
      if (bar) {
        bar.style.width = '100%'
        setTimeout(() => {
          if (bar) {
            bar.remove()
            bar = null
          }
        }, 200)
      }
    }
  }
}

const progress = createProgressBar()

/**
 * 设置路由守卫
 */
export function setupRouterGuards(router: Router) {
  // 前置守卫
  router.beforeEach(async (to, _from, next) => {
    // 开始进度条
    progress.start()
    
    // 设置页面标题
    const title = to.meta.title as string
    document.title = title ? `${title} - 巨量引擎管理平台` : '巨量引擎管理平台'
    
    // 获取 token
    const token = localStorage.getItem('access_token')
    
    // 白名单路由直接放行
    if (whiteList.includes(to.path)) {
      // 已登录用户访问登录页，重定向到首页
      if (to.path === '/login' && token) {
        next('/')
        return
      }
      next()
      return
    }
    
    // 未登录跳转登录页
    if (!token) {
      next({
        path: '/login',
        query: { redirect: to.fullPath }
      })
      return
    }
    
    // 权限检查（如有需要）
    const requiredPermissions = to.meta.permissions as string[] | undefined
    if (requiredPermissions && requiredPermissions.length > 0) {
      // TODO: 从 store 获取用户权限进行校验
      const hasPermission = true // 临时放行
      if (!hasPermission) {
        next('/403')
        return
      }
    }
    
    next()
  })
  
  // 后置守卫
  router.afterEach(() => {
    // 结束进度条
    progress.done()
  })
  
  // 错误处理
  router.onError((error) => {
    console.error('Router error:', error)
    progress.done()
  })
}

/**
 * 检查是否有权限
 */
export function hasPermission(permissions: string[], requiredPermissions: string[]): boolean {
  if (!requiredPermissions || requiredPermissions.length === 0) {
    return true
  }
  return requiredPermissions.some(permission => permissions.includes(permission))
}

/**
 * 检查是否有角色
 */
export function hasRole(roles: string[], requiredRoles: string[]): boolean {
  if (!requiredRoles || requiredRoles.length === 0) {
    return true
  }
  return requiredRoles.some(role => roles.includes(role))
}
