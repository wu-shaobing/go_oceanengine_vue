<script setup lang="ts">
import { ref, watch, computed } from 'vue'

interface Advertiser {
  id: number
  name: string
  company: string
  status: string
}

interface Props {
  modelValue?: number | number[]
  multiple?: boolean
  placeholder?: string
  disabled?: boolean
  clearable?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  modelValue: undefined,
  multiple: false,
  placeholder: '请选择广告主',
  disabled: false,
  clearable: true
})

const emit = defineEmits<{
  (e: 'update:modelValue', value: number | number[] | undefined): void
  (e: 'change', value: number | number[] | undefined): void
}>()

const isOpen = ref(false)
const searchQuery = ref('')
const loading = ref(false)

// 模拟广告主数据
const advertisers = ref<Advertiser[]>([
  { id: 1, name: '广告主A', company: '北京科技有限公司', status: 'active' },
  { id: 2, name: '广告主B', company: '上海电商有限公司', status: 'active' },
  { id: 3, name: '广告主C', company: '深圳游戏有限公司', status: 'active' },
  { id: 4, name: '广告主D', company: '杭州教育有限公司', status: 'disabled' },
  { id: 5, name: '广告主E', company: '广州金融有限公司', status: 'active' },
  { id: 6, name: '广告主F', company: '成都旅游有限公司', status: 'active' }
])

const filteredAdvertisers = computed(() => {
  if (!searchQuery.value) return advertisers.value
  const query = searchQuery.value.toLowerCase()
  return advertisers.value.filter(
    a => a.name.toLowerCase().includes(query) || a.company.toLowerCase().includes(query)
  )
})

const selectedAdvertisers = computed(() => {
  if (!props.modelValue) return []
  const ids = Array.isArray(props.modelValue) ? props.modelValue : [props.modelValue]
  return advertisers.value.filter(a => ids.includes(a.id))
})

const displayText = computed(() => {
  if (selectedAdvertisers.value.length === 0) return ''
  if (props.multiple) {
    return `已选择 ${selectedAdvertisers.value.length} 个广告主`
  }
  return selectedAdvertisers.value[0]?.name || ''
})

const isSelected = (id: number) => {
  if (!props.modelValue) return false
  if (Array.isArray(props.modelValue)) {
    return props.modelValue.includes(id)
  }
  return props.modelValue === id
}

const toggleSelect = (advertiser: Advertiser) => {
  if (props.disabled || advertiser.status === 'disabled') return
  
  if (props.multiple) {
    const currentValue = Array.isArray(props.modelValue) ? [...props.modelValue] : []
    const index = currentValue.indexOf(advertiser.id)
    if (index >= 0) {
      currentValue.splice(index, 1)
    } else {
      currentValue.push(advertiser.id)
    }
    emit('update:modelValue', currentValue)
    emit('change', currentValue)
  } else {
    emit('update:modelValue', advertiser.id)
    emit('change', advertiser.id)
    isOpen.value = false
  }
}

const clear = () => {
  emit('update:modelValue', props.multiple ? [] : undefined)
  emit('change', props.multiple ? [] : undefined)
}

const removeItem = (id: number) => {
  if (!props.multiple || !Array.isArray(props.modelValue)) return
  const newValue = props.modelValue.filter(v => v !== id)
  emit('update:modelValue', newValue)
  emit('change', newValue)
}

watch(isOpen, (val) => {
  if (val) {
    searchQuery.value = ''
  }
})
</script>

<template>
  <div class="relative">
    <!-- 选择框 -->
    <div
      class="min-h-[42px] px-3 py-2 border rounded-lg cursor-pointer transition-colors flex items-center gap-2"
      :class="[
        disabled ? 'bg-gray-100 cursor-not-allowed' : 'bg-white hover:border-blue-400',
        isOpen ? 'border-blue-500 ring-2 ring-blue-100' : 'border-gray-300'
      ]"
      @click="!disabled && (isOpen = !isOpen)"
    >
      <!-- 单选模式 -->
      <template v-if="!multiple">
        <span v-if="displayText" class="text-gray-900">{{ displayText }}</span>
        <span v-else class="text-gray-400">{{ placeholder }}</span>
      </template>
      
      <!-- 多选模式 -->
      <template v-else>
        <div v-if="selectedAdvertisers.length > 0" class="flex flex-wrap gap-1">
          <span
            v-for="adv in selectedAdvertisers"
            :key="adv.id"
            class="inline-flex items-center gap-1 px-2 py-0.5 bg-blue-100 text-blue-800 text-sm rounded"
          >
            {{ adv.name }}
            <button
              class="hover:text-blue-600"
              @click.stop="removeItem(adv.id)"
            >
              <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
              </svg>
            </button>
          </span>
        </div>
        <span v-else class="text-gray-400">{{ placeholder }}</span>
      </template>
      
      <div class="ml-auto flex items-center gap-1">
        <!-- 清除按钮 -->
        <button
          v-if="clearable && selectedAdvertisers.length > 0 && !disabled"
          class="text-gray-400 hover:text-gray-600"
          @click.stop="clear"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
          </svg>
        </button>
        
        <!-- 下拉箭头 -->
        <svg
          class="w-4 h-4 text-gray-400 transition-transform"
          :class="{ 'rotate-180': isOpen }"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
        </svg>
      </div>
    </div>
    
    <!-- 下拉面板 -->
    <Transition
      enter-active-class="transition ease-out duration-100"
      enter-from-class="transform opacity-0 scale-95"
      enter-to-class="transform opacity-100 scale-100"
      leave-active-class="transition ease-in duration-75"
      leave-from-class="transform opacity-100 scale-100"
      leave-to-class="transform opacity-0 scale-95"
    >
      <div
        v-if="isOpen"
        class="absolute z-50 mt-1 w-full bg-white border border-gray-200 rounded-lg shadow-lg"
      >
        <!-- 搜索框 -->
        <div class="p-2 border-b border-gray-100">
          <input
            v-model="searchQuery"
            type="text"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-100 focus:border-blue-500"
            placeholder="搜索广告主..."
            @click.stop
          />
        </div>
        
        <!-- 广告主列表 -->
        <div class="max-h-60 overflow-y-auto">
          <div v-if="loading" class="p-4 text-center text-gray-500">
            加载中...
          </div>
          <div v-else-if="filteredAdvertisers.length === 0" class="p-4 text-center text-gray-500">
            未找到广告主
          </div>
          <div
            v-else
            v-for="advertiser in filteredAdvertisers"
            :key="advertiser.id"
            class="px-3 py-2 cursor-pointer transition-colors"
            :class="[
              advertiser.status === 'disabled' ? 'opacity-50 cursor-not-allowed' : 'hover:bg-gray-50',
              isSelected(advertiser.id) ? 'bg-blue-50' : ''
            ]"
            @click="toggleSelect(advertiser)"
          >
            <div class="flex items-center justify-between">
              <div>
                <p class="font-medium text-gray-900">{{ advertiser.name }}</p>
                <p class="text-sm text-gray-500">{{ advertiser.company }}</p>
              </div>
              <div class="flex items-center gap-2">
                <span
                  v-if="advertiser.status === 'disabled'"
                  class="text-xs text-gray-400"
                >
                  已禁用
                </span>
                <svg
                  v-if="isSelected(advertiser.id)"
                  class="w-5 h-5 text-blue-600"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
                </svg>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Transition>
    
    <!-- 点击外部关闭 -->
    <div
      v-if="isOpen"
      class="fixed inset-0 z-40"
      @click="isOpen = false"
    ></div>
  </div>
</template>
