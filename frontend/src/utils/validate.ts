/**
 * 验证手机号
 */
export function isPhone(value: string): boolean {
  return /^1[3-9]\d{9}$/.test(value)
}

/**
 * 验证邮箱
 */
export function isEmail(value: string): boolean {
  return /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/.test(value)
}

/**
 * 验证 URL
 */
export function isUrl(value: string): boolean {
  try {
    new URL(value)
    return true
  } catch {
    return false
  }
}

/**
 * 验证身份证号
 */
export function isIdCard(value: string): boolean {
  return /^[1-9]\d{5}(18|19|20)\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$/.test(value)
}

/**
 * 验证是否为空
 */
export function isEmpty(value: unknown): boolean {
  if (value === null || value === undefined) return true
  if (typeof value === 'string') return value.trim() === ''
  if (Array.isArray(value)) return value.length === 0
  if (typeof value === 'object') return Object.keys(value).length === 0
  return false
}

/**
 * 验证是否为数字
 */
export function isNumber(value: unknown): boolean {
  return typeof value === 'number' && !isNaN(value)
}

/**
 * 验证密码强度（至少8位，包含字母和数字）
 */
export function isStrongPassword(value: string): boolean {
  return /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d@$!%*#?&]{8,}$/.test(value)
}

/**
 * 验证用户名（字母开头，允许字母数字下划线，4-16位）
 */
export function isUsername(value: string): boolean {
  return /^[a-zA-Z][a-zA-Z0-9_]{3,15}$/.test(value)
}

/**
 * 创建表单验证规则
 */
export const rules = {
  required: (message = '此项为必填项') => ({
    required: true,
    message,
    trigger: 'blur'
  }),
  
  phone: (message = '请输入正确的手机号') => ({
    validator: (_: unknown, value: string) => isPhone(value),
    message,
    trigger: 'blur'
  }),
  
  email: (message = '请输入正确的邮箱地址') => ({
    validator: (_: unknown, value: string) => isEmail(value),
    message,
    trigger: 'blur'
  }),
  
  min: (min: number, message?: string) => ({
    min,
    message: message || `最少${min}个字符`,
    trigger: 'blur'
  }),
  
  max: (max: number, message?: string) => ({
    max,
    message: message || `最多${max}个字符`,
    trigger: 'blur'
  })
}
