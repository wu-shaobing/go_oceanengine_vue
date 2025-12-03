// Mock 认证数据
export const mockUsers = [
  {
    id: 1,
    username: 'admin',
    password: 'admin123',
    nickname: '管理员',
    avatar: '',
    email: 'admin@oceanengine.com',
    phone: '13800138000',
    role: {
      id: 1,
      name: '超级管理员',
      key: 'admin'
    }
  }
]

export const mockLogin = (username: string, password: string) => {
  const user = mockUsers.find(u => u.username === username && u.password === password)
  if (user) {
    return {
      code: 0,
      message: 'success',
      data: {
        access_token: 'mock_access_token_' + Date.now(),
        refresh_token: 'mock_refresh_token_' + Date.now(),
        expires_in: 7200
      }
    }
  }
  return {
    code: 401,
    message: '用户名或密码错误',
    data: null
  }
}

export const mockUserInfo = (token: string) => {
  if (token && token.startsWith('mock_')) {
    const user = mockUsers[0]
    return {
      code: 0,
      message: 'success',
      data: {
        user: {
          id: user.id,
          username: user.username,
          nickname: user.nickname,
          avatar: user.avatar,
          email: user.email,
          phone: user.phone,
          role: user.role
        },
        permissions: ['*']
      }
    }
  }
  return {
    code: 401,
    message: '未登录或登录已过期',
    data: null
  }
}
