# 组件设计规范

## 组件分类

### 1. 通用组件 (Common Components)

可在任何项目中复用的基础组件。

#### Button 按钮

```vue
<!-- components/common/Button.vue -->
<script setup lang="ts">
interface Props {
  type?: 'primary' | 'secondary' | 'danger' | 'ghost'
  size?: 'sm' | 'md' | 'lg'
  loading?: boolean
  disabled?: boolean
  block?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  type: 'primary',
  size: 'md',
  loading: false,
  disabled: false,
  block: false
})

const emit = defineEmits<{
  (e: 'click', event: MouseEvent): void
}>()

const classes = computed(() => {
  const base = 'inline-flex items-center justify-center font-medium rounded-lg transition-colors'
  
  const types = {
    primary: 'bg-blue-600 text-white hover:bg-blue-700 disabled:bg-blue-300',
    secondary: 'bg-gray-100 text-gray-700 hover:bg-gray-200 disabled:bg-gray-50',
    danger: 'bg-red-600 text-white hover:bg-red-700 disabled:bg-red-300',
    ghost: 'bg-transparent text-gray-700 hover:bg-gray-100'
  }
  
  const sizes = {
    sm: 'px-3 py-1.5 text-sm',
    md: 'px-4 py-2 text-sm',
    lg: 'px-6 py-3 text-base'
  }
  
  return [
    base,
    types[props.type],
    sizes[props.size],
    props.block ? 'w-full' : '',
    props.disabled || props.loading ? 'cursor-not-allowed opacity-60' : ''
  ].join(' ')
})
</script>

<template>
  <button
    :class="classes"
    :disabled="disabled || loading"
    @click="emit('click', $event)"
  >
    <svg
      v-if="loading"
      class="animate-spin -ml-1 mr-2 h-4 w-4"
      fill="none"
      viewBox="0 0 24 24"
    >
      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z" />
    </svg>
    <slot />
  </button>
</template>
```

#### Input 输入框

```vue
<!-- components/common/Input.vue -->
<script setup lang="ts">
interface Props {
  modelValue: string | number
  type?: 'text' | 'password' | 'email' | 'number'
  placeholder?: string
  disabled?: boolean
  error?: string
  prefix?: string
  suffix?: string
}

const props = withDefaults(defineProps<Props>(), {
  type: 'text',
  placeholder: '',
  disabled: false
})

const emit = defineEmits<{
  (e: 'update:modelValue', value: string | number): void
  (e: 'focus', event: FocusEvent): void
  (e: 'blur', event: FocusEvent): void
}>()

const inputClasses = computed(() => {
  const base = 'block w-full rounded-lg border px-3 py-2 text-sm transition-colors'
  const normal = 'border-gray-300 focus:border-blue-500 focus:ring-1 focus:ring-blue-500'
  const errorStyle = 'border-red-500 focus:border-red-500 focus:ring-1 focus:ring-red-500'
  const disabledStyle = 'bg-gray-50 cursor-not-allowed'
  
  return [
    base,
    props.error ? errorStyle : normal,
    props.disabled ? disabledStyle : ''
  ].join(' ')
})
</script>

<template>
  <div class="w-full">
    <div class="relative">
      <span v-if="prefix" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500">
        {{ prefix }}
      </span>
      <input
        :type="type"
        :value="modelValue"
        :placeholder="placeholder"
        :disabled="disabled"
        :class="[inputClasses, prefix ? 'pl-8' : '', suffix ? 'pr-8' : '']"
        @input="emit('update:modelValue', ($event.target as HTMLInputElement).value)"
        @focus="emit('focus', $event)"
        @blur="emit('blur', $event)"
      />
      <span v-if="suffix" class="absolute right-3 top-1/2 -translate-y-1/2 text-gray-500">
        {{ suffix }}
      </span>
    </div>
    <p v-if="error" class="mt-1 text-sm text-red-500">{{ error }}</p>
  </div>
</template>
```

#### Modal 弹窗

```vue
<!-- components/common/Modal.vue -->
<script setup lang="ts">
import { watch } from 'vue'

interface Props {
  modelValue: boolean
  title?: string
  width?: string
  closable?: boolean
  maskClosable?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  title: '',
  width: '500px',
  closable: true,
  maskClosable: true
})

const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
  (e: 'close'): void
  (e: 'confirm'): void
}>()

const close = () => {
  emit('update:modelValue', false)
  emit('close')
}

const handleMaskClick = () => {
  if (props.maskClosable) {
    close()
  }
}

// 禁止背景滚动
watch(() => props.modelValue, (value) => {
  document.body.style.overflow = value ? 'hidden' : ''
})
</script>

<template>
  <Teleport to="body">
    <Transition name="modal">
      <div
        v-if="modelValue"
        class="fixed inset-0 z-50 flex items-center justify-center"
      >
        <!-- 遮罩 -->
        <div
          class="absolute inset-0 bg-black/50"
          @click="handleMaskClick"
        />
        
        <!-- 弹窗内容 -->
        <div
          class="relative bg-white rounded-lg shadow-xl"
          :style="{ width }"
        >
          <!-- 头部 -->
          <div class="flex items-center justify-between px-6 py-4 border-b">
            <h3 class="text-lg font-medium text-gray-900">{{ title }}</h3>
            <button
              v-if="closable"
              class="text-gray-400 hover:text-gray-500"
              @click="close"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
          
          <!-- 内容 -->
          <div class="px-6 py-4">
            <slot />
          </div>
          
          <!-- 底部 -->
          <div v-if="$slots.footer" class="px-6 py-4 border-t bg-gray-50 rounded-b-lg">
            <slot name="footer" />
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<style scoped>
.modal-enter-active,
.modal-leave-active {
  transition: opacity 0.2s ease;
}

.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}
</style>
```

#### Table 表格

```vue
<!-- components/common/Table.vue -->
<script setup lang="ts" generic="T extends Record<string, any>">
interface Column<T> {
  key: keyof T | string
  title: string
  width?: string
  align?: 'left' | 'center' | 'right'
  sortable?: boolean
  render?: (row: T, index: number) => any
}

interface Props {
  data: T[]
  columns: Column<T>[]
  loading?: boolean
  rowKey?: keyof T
  stripe?: boolean
  border?: boolean
  emptyText?: string
}

const props = withDefaults(defineProps<Props>(), {
  loading: false,
  rowKey: 'id' as keyof T,
  stripe: false,
  border: false,
  emptyText: '暂无数据'
})

const emit = defineEmits<{
  (e: 'row-click', row: T, index: number): void
  (e: 'sort-change', key: string, order: 'asc' | 'desc'): void
}>()

const getCellValue = (row: T, key: string) => {
  const keys = key.split('.')
  let value: any = row
  for (const k of keys) {
    value = value?.[k]
  }
  return value
}
</script>

<template>
  <div class="overflow-x-auto">
    <table class="min-w-full divide-y divide-gray-200">
      <thead class="bg-gray-50">
        <tr>
          <th
            v-for="column in columns"
            :key="String(column.key)"
            :style="{ width: column.width }"
            :class="[
              'px-4 py-3 text-xs font-medium text-gray-500 uppercase tracking-wider',
              column.align === 'center' ? 'text-center' : '',
              column.align === 'right' ? 'text-right' : 'text-left'
            ]"
          >
            {{ column.title }}
          </th>
        </tr>
      </thead>
      <tbody class="bg-white divide-y divide-gray-200">
        <!-- Loading -->
        <tr v-if="loading">
          <td :colspan="columns.length" class="px-4 py-8 text-center">
            <div class="flex items-center justify-center">
              <svg class="animate-spin h-6 w-6 text-blue-500" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z" />
              </svg>
              <span class="ml-2 text-gray-500">加载中...</span>
            </div>
          </td>
        </tr>
        
        <!-- Empty -->
        <tr v-else-if="data.length === 0">
          <td :colspan="columns.length" class="px-4 py-8 text-center text-gray-500">
            {{ emptyText }}
          </td>
        </tr>
        
        <!-- Data -->
        <tr
          v-else
          v-for="(row, index) in data"
          :key="String(row[rowKey])"
          :class="[
            'hover:bg-gray-50 cursor-pointer',
            stripe && index % 2 === 1 ? 'bg-gray-50' : ''
          ]"
          @click="emit('row-click', row, index)"
        >
          <td
            v-for="column in columns"
            :key="String(column.key)"
            :class="[
              'px-4 py-3 text-sm text-gray-900',
              column.align === 'center' ? 'text-center' : '',
              column.align === 'right' ? 'text-right' : ''
            ]"
          >
            <slot :name="column.key" :row="row" :index="index">
              <template v-if="column.render">
                {{ column.render(row, index) }}
              </template>
              <template v-else>
                {{ getCellValue(row, String(column.key)) }}
              </template>
            </slot>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
```

---

### 2. 布局组件 (Layout Components)

#### AppLayout 应用布局

```vue
<!-- components/layout/AppLayout.vue -->
<script setup lang="ts">
import { ref } from 'vue'
import AppHeader from './AppHeader.vue'
import AppSidebar from './AppSidebar.vue'

const sidebarCollapsed = ref(false)

const toggleSidebar = () => {
  sidebarCollapsed.value = !sidebarCollapsed.value
}
</script>

<template>
  <div class="min-h-screen bg-gray-100">
    <AppHeader @toggle-sidebar="toggleSidebar" />
    
    <div class="flex">
      <AppSidebar :collapsed="sidebarCollapsed" />
      
      <main
        :class="[
          'flex-1 p-6 transition-all duration-300',
          sidebarCollapsed ? 'ml-16' : 'ml-64'
        ]"
      >
        <slot />
      </main>
    </div>
  </div>
</template>
```

---

### 3. 业务组件 (Business Components)

#### StatsCard 数据卡片

```vue
<!-- components/business/StatsCard.vue -->
<script setup lang="ts">
interface Props {
  title: string
  value: string | number
  change?: number
  changeLabel?: string
  icon?: string
  color?: 'blue' | 'green' | 'red' | 'yellow'
}

const props = withDefaults(defineProps<Props>(), {
  color: 'blue'
})

const colorClasses = computed(() => {
  const colors = {
    blue: 'bg-blue-50 text-blue-600',
    green: 'bg-green-50 text-green-600',
    red: 'bg-red-50 text-red-600',
    yellow: 'bg-yellow-50 text-yellow-600'
  }
  return colors[props.color]
})
</script>

<template>
  <div class="bg-white rounded-lg shadow p-6">
    <div class="flex items-center justify-between">
      <div>
        <p class="text-sm font-medium text-gray-500">{{ title }}</p>
        <p class="mt-1 text-2xl font-semibold text-gray-900">{{ value }}</p>
        
        <div v-if="change !== undefined" class="mt-2 flex items-center text-sm">
          <span
            :class="[
              'font-medium',
              change >= 0 ? 'text-green-600' : 'text-red-600'
            ]"
          >
            {{ change >= 0 ? '+' : '' }}{{ change }}%
          </span>
          <span v-if="changeLabel" class="ml-1 text-gray-500">
            {{ changeLabel }}
          </span>
        </div>
      </div>
      
      <div
        v-if="icon"
        :class="['p-3 rounded-lg', colorClasses]"
      >
        <component :is="icon" class="w-6 h-6" />
      </div>
    </div>
  </div>
</template>
```

#### StatusBadge 状态标签

```vue
<!-- components/business/StatusBadge.vue -->
<script setup lang="ts">
type Status = 'active' | 'inactive' | 'pending' | 'error' | 'success'

interface Props {
  status: Status
  text?: string
}

const props = defineProps<Props>()

const statusConfig: Record<Status, { bg: string; text: string; dot: string; label: string }> = {
  active: { bg: 'bg-green-100', text: 'text-green-800', dot: 'bg-green-400', label: '运行中' },
  inactive: { bg: 'bg-gray-100', text: 'text-gray-800', dot: 'bg-gray-400', label: '已暂停' },
  pending: { bg: 'bg-yellow-100', text: 'text-yellow-800', dot: 'bg-yellow-400', label: '审核中' },
  error: { bg: 'bg-red-100', text: 'text-red-800', dot: 'bg-red-400', label: '异常' },
  success: { bg: 'bg-green-100', text: 'text-green-800', dot: 'bg-green-400', label: '成功' }
}

const config = computed(() => statusConfig[props.status])
</script>

<template>
  <span
    :class="[
      'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
      config.bg,
      config.text
    ]"
  >
    <span :class="['w-1.5 h-1.5 rounded-full mr-1.5', config.dot]" />
    {{ text || config.label }}
  </span>
</template>
```

---

## 组件通信模式

### 1. Props + Events

```vue
<!-- 父组件 -->
<template>
  <ChildComponent
    :data="parentData"
    @update="handleUpdate"
  />
</template>

<!-- 子组件 -->
<script setup lang="ts">
const props = defineProps<{ data: string }>()
const emit = defineEmits<{ (e: 'update', value: string): void }>()
</script>
```

### 2. v-model 双向绑定

```vue
<!-- 父组件 -->
<template>
  <CustomInput v-model="inputValue" />
  <CustomSelect v-model:selected="selectedItem" />
</template>

<!-- 子组件 -->
<script setup lang="ts">
const props = defineProps<{ modelValue: string }>()
const emit = defineEmits<{ (e: 'update:modelValue', value: string): void }>()
</script>
```

### 3. Provide / Inject

```vue
<!-- 祖先组件 -->
<script setup lang="ts">
import { provide, ref } from 'vue'

const theme = ref('light')
provide('theme', theme)
</script>

<!-- 后代组件 -->
<script setup lang="ts">
import { inject } from 'vue'

const theme = inject('theme', ref('light'))
</script>
```

---

## 组件最佳实践

### 1. 使用组合式函数抽取逻辑

```typescript
// composables/useTable.ts
export function useTable<T>(fetchFn: () => Promise<T[]>) {
  const data = ref<T[]>([])
  const loading = ref(false)
  const error = ref<Error | null>(null)
  
  const fetch = async () => {
    loading.value = true
    try {
      data.value = await fetchFn()
    } catch (e) {
      error.value = e as Error
    } finally {
      loading.value = false
    }
  }
  
  onMounted(fetch)
  
  return { data, loading, error, refresh: fetch }
}
```

### 2. 组件懒加载

```typescript
// router/routes.ts
const routes = [
  {
    path: '/dashboard',
    component: () => import('@/views/dashboard/DashboardView.vue')
  }
]
```

### 3. 异步组件

```vue
<script setup lang="ts">
import { defineAsyncComponent } from 'vue'

const HeavyChart = defineAsyncComponent({
  loader: () => import('./HeavyChart.vue'),
  loadingComponent: LoadingSpinner,
  delay: 200
})
</script>
```
