import { ref, computed } from 'vue'

export interface FormRule {
  required?: boolean
  message?: string
  pattern?: RegExp
  min?: number
  max?: number
  validator?: (value: unknown) => boolean | string | Promise<boolean | string>
}

export interface FormField {
  value: unknown
  rules?: FormRule[]
  error?: string
  touched?: boolean
}

export interface UseFormOptions {
  initialValues: Record<string, unknown>
  rules?: Record<string, FormRule[]>
  onSubmit?: (values: Record<string, unknown>) => void | Promise<void>
}

export function useForm(options: UseFormOptions) {
  const { initialValues, rules = {}, onSubmit } = options
  
  // 表单值
  const values = ref<Record<string, unknown>>({ ...initialValues })
  
  // 表单错误
  const errors = ref<Record<string, string>>({})
  
  // 表单是否被修改
  const touched = ref<Record<string, boolean>>({})
  
  // 提交状态
  const submitting = ref(false)
  
  // 表单是否有效
  const isValid = computed(() => {
    return Object.keys(errors.value).length === 0
  })
  
  // 表单是否被修改过
  const isDirty = computed(() => {
    return Object.values(touched.value).some(Boolean)
  })
  
  // 验证单个字段
  const validateField = async (field: string): Promise<boolean> => {
    const fieldRules = rules[field]
    if (!fieldRules) return true
    
    const value = values.value[field]
    
    for (const rule of fieldRules) {
      // 必填验证
      if (rule.required) {
        if (value === undefined || value === null || value === '') {
          errors.value[field] = rule.message || '此字段为必填项'
          return false
        }
      }
      
      // 正则验证
      if (rule.pattern && typeof value === 'string') {
        if (!rule.pattern.test(value)) {
          errors.value[field] = rule.message || '格式不正确'
          return false
        }
      }
      
      // 最小长度/值验证
      if (rule.min !== undefined) {
        if (typeof value === 'string' && value.length < rule.min) {
          errors.value[field] = rule.message || `最少 ${rule.min} 个字符`
          return false
        }
        if (typeof value === 'number' && value < rule.min) {
          errors.value[field] = rule.message || `最小值为 ${rule.min}`
          return false
        }
      }
      
      // 最大长度/值验证
      if (rule.max !== undefined) {
        if (typeof value === 'string' && value.length > rule.max) {
          errors.value[field] = rule.message || `最多 ${rule.max} 个字符`
          return false
        }
        if (typeof value === 'number' && value > rule.max) {
          errors.value[field] = rule.message || `最大值为 ${rule.max}`
          return false
        }
      }
      
      // 自定义验证器
      if (rule.validator) {
        const result = await rule.validator(value)
        if (result !== true) {
          errors.value[field] = typeof result === 'string' ? result : (rule.message || '验证失败')
          return false
        }
      }
    }
    
    // 清除错误
    delete errors.value[field]
    return true
  }
  
  // 验证所有字段
  const validate = async (): Promise<boolean> => {
    const fields = Object.keys(rules)
    const results = await Promise.all(fields.map(validateField))
    return results.every(Boolean)
  }
  
  // 设置字段值
  const setFieldValue = (field: string, value: unknown) => {
    values.value[field] = value
    touched.value[field] = true
  }
  
  // 设置字段错误
  const setFieldError = (field: string, error: string) => {
    errors.value[field] = error
  }
  
  // 清除字段错误
  const clearFieldError = (field: string) => {
    delete errors.value[field]
  }
  
  // 重置表单
  const reset = () => {
    values.value = { ...initialValues }
    errors.value = {}
    touched.value = {}
  }
  
  // 提交表单
  const submit = async () => {
    const valid = await validate()
    if (!valid) return false
    
    if (onSubmit) {
      submitting.value = true
      try {
        await onSubmit(values.value)
        return true
      } catch (error) {
        console.error('Form submit error:', error)
        return false
      } finally {
        submitting.value = false
      }
    }
    
    return true
  }
  
  // 处理字段变化
  const handleChange = (field: string) => (event: Event) => {
    const target = event.target as HTMLInputElement
    setFieldValue(field, target.value)
  }
  
  // 处理字段失焦
  const handleBlur = (field: string) => async () => {
    touched.value[field] = true
    await validateField(field)
  }
  
  return {
    values,
    errors,
    touched,
    submitting,
    isValid,
    isDirty,
    validateField,
    validate,
    setFieldValue,
    setFieldError,
    clearFieldError,
    reset,
    submit,
    handleChange,
    handleBlur
  }
}
