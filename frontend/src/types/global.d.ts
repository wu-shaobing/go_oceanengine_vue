/**
 * 全局类型声明
 */

declare global {
  /**
   * 通用对象类型
   */
  type Recordable<T = unknown> = Record<string, T>
  
  /**
   * 可空类型
   */
  type Nullable<T> = T | null
  
  /**
   * 可选类型
   */
  type Optional<T> = T | undefined
  
  /**
   * 深度可选
   */
  type DeepPartial<T> = {
    [P in keyof T]?: T[P] extends object ? DeepPartial<T[P]> : T[P]
  }
  
  /**
   * 深度只读
   */
  type DeepReadonly<T> = {
    readonly [P in keyof T]: T[P] extends object ? DeepReadonly<T[P]> : T[P]
  }
  
  /**
   * 函数类型
   */
  type Fn<T = void> = (...args: unknown[]) => T
  
  /**
   * 异步函数类型
   */
  type AsyncFn<T = void> = (...args: unknown[]) => Promise<T>
  
  /**
   * 组件实例类型
   */
  type ComponentRef<T extends abstract new (...args: unknown[]) => unknown> = InstanceType<T>
  
  /**
   * Element 引用类型
   */
  type ElRef<T extends HTMLElement = HTMLElement> = Nullable<T>
  
  /**
   * 定时器类型
   */
  type TimeoutHandle = ReturnType<typeof setTimeout>
  type IntervalHandle = ReturnType<typeof setInterval>
  
  /**
   * 选项类型
   */
  interface SelectOption {
    label: string
    value: string | number
    disabled?: boolean
    children?: SelectOption[]
  }
  
  /**
   * 树节点类型
   */
  interface TreeNode {
    id: string | number
    label: string
    children?: TreeNode[]
    [key: string]: unknown
  }
  
  /**
   * 表格列类型
   */
  interface TableColumn {
    key: string
    title: string
    width?: number | string
    align?: 'left' | 'center' | 'right'
    sortable?: boolean
    fixed?: 'left' | 'right'
    slot?: string
    formatter?: (value: unknown, row: Recordable) => string
  }
  
  /**
   * 表单项类型
   */
  interface FormItem {
    field: string
    label: string
    type: 'input' | 'select' | 'date' | 'textarea' | 'switch' | 'checkbox' | 'radio'
    placeholder?: string
    rules?: FormRule[]
    options?: SelectOption[]
    props?: Recordable
  }
  
  /**
   * 表单规则类型
   */
  interface FormRule {
    required?: boolean
    message?: string
    pattern?: RegExp
    min?: number
    max?: number
    trigger?: 'blur' | 'change'
    validator?: (value: unknown) => boolean | string | Promise<boolean | string>
  }
}

export {}
