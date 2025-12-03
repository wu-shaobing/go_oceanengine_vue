import { defineStore } from 'pinia'
import { authApi, type UserInfo } from '@/api/auth'
import router from '@/router'

interface AuthState {
  token: string | null
  refreshToken: string | null
  userInfo: UserInfo | null
  permissions: string[]
}

export const useAuthStore = defineStore('auth', {
  state: (): AuthState => ({
    token: localStorage.getItem('access_token'),
    refreshToken: localStorage.getItem('refresh_token'),
    userInfo: null,
    permissions: []
  }),

  getters: {
    isLoggedIn: (state) => !!state.token,
    username: (state) => state.userInfo?.username || '',
    nickname: (state) => state.userInfo?.nickname || state.userInfo?.username || ''
  },

  actions: {
    async login(username: string, password: string) {
      try {
        const res = await authApi.login({ username, password })
        this.setToken(res.access_token, res.refresh_token)
        await this.fetchUserInfo()
        return true
      } catch (error) {
        console.error('Login failed:', error)
        return false
      }
    },

    setToken(accessToken: string, refreshToken?: string) {
      this.token = accessToken
      localStorage.setItem('access_token', accessToken)
      
      if (refreshToken) {
        this.refreshToken = refreshToken
        localStorage.setItem('refresh_token', refreshToken)
      }
    },

    async fetchUserInfo() {
      try {
        const res = await authApi.getUserInfo()
        this.userInfo = res.user
        this.permissions = res.permissions
        localStorage.setItem('permissions', JSON.stringify(res.permissions))
      } catch (error) {
        console.error('Fetch user info failed:', error)
        this.clearAuth()
      }
    },

    async refreshAccessToken() {
      if (!this.refreshToken) {
        throw new Error('No refresh token')
      }
      
      try {
        const res = await authApi.refreshToken(this.refreshToken)
        this.setToken(res.access_token, res.refresh_token)
      } catch (error) {
        this.clearAuth()
        throw error
      }
    },

    hasPermission(permissions: string | string[]): boolean {
      const required = Array.isArray(permissions) ? permissions : [permissions]
      return required.some(p => this.permissions.includes(p))
    },

    async logout() {
      try {
        await authApi.logout()
      } catch {
        // ignore logout error
      }
      this.clearAuth()
      router.push('/login')
    },

    clearAuth() {
      this.token = null
      this.refreshToken = null
      this.userInfo = null
      this.permissions = []
      localStorage.removeItem('access_token')
      localStorage.removeItem('refresh_token')
      localStorage.removeItem('permissions')
    }
  }
})
