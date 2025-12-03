import request from './request'

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
  login(data: LoginRequest) {
    return request.post<LoginResponse>('/auth/login', data)
  },

  logout() {
    return request.post<void>('/auth/logout')
  },

  refreshToken(refreshToken: string) {
    return request.post<LoginResponse>('/auth/refresh', { refresh_token: refreshToken })
  },

  getUserInfo() {
    return request.get<{ user: UserInfo; permissions: string[] }>('/auth/info')
  },

  getCaptcha() {
    return request.get<{ captcha_id: string; captcha_image: string }>('/auth/captcha')
  },

  changePassword(data: { old_password: string; new_password: string }) {
    return request.post<void>('/auth/password', data)
  }
}
