import 'vue-router'

declare module 'vue-router' {
  interface RouteMeta {
    /**
     * 页面标题
     */
    title?: string
    
    /**
     * 图标
     */
    icon?: string
    
    /**
     * 是否需要登录
     */
    requiresAuth?: boolean
    
    /**
     * 所需权限
     */
    permissions?: string[]
    
    /**
     * 所需角色
     */
    roles?: string[]
    
    /**
     * 是否在菜单中隐藏
     */
    hidden?: boolean
    
    /**
     * 是否缓存页面
     */
    keepAlive?: boolean
    
    /**
     * 是否固定在标签栏
     */
    affix?: boolean
    
    /**
     * 面包屑中是否显示
     */
    breadcrumb?: boolean
    
    /**
     * 当前激活的菜单路径
     */
    activeMenu?: string
    
    /**
     * 是否不显示在面包屑中
     */
    noBreadcrumb?: boolean
  }
}

export {}
