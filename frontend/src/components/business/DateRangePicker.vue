<script setup lang="ts">
import { ref, computed } from 'vue'

interface DateRange {
  start: string
  end: string
}

interface Props {
  modelValue?: DateRange
  presets?: { label: string; value: string }[]
}

const props = withDefaults(defineProps<Props>(), {
  presets: () => [
    { label: '今日', value: 'today' },
    { label: '昨日', value: 'yesterday' },
    { label: '近7天', value: '7d' },
    { label: '近30天', value: '30d' },
    { label: '本月', value: 'month' }
  ]
})

const emit = defineEmits<{
  (e: 'update:modelValue', value: DateRange): void
  (e: 'change', value: DateRange): void
}>()

const selectedPreset = ref('7d')
const showPicker = ref(false)

const formatDate = (date: Date): string => {
  return date.toISOString().split('T')[0]
}

const getDateRange = (preset: string): DateRange => {
  const today = new Date()
  const start = new Date()
  
  switch (preset) {
    case 'today':
      return { start: formatDate(today), end: formatDate(today) }
    case 'yesterday':
      start.setDate(today.getDate() - 1)
      return { start: formatDate(start), end: formatDate(start) }
    case '7d':
      start.setDate(today.getDate() - 6)
      return { start: formatDate(start), end: formatDate(today) }
    case '30d':
      start.setDate(today.getDate() - 29)
      return { start: formatDate(start), end: formatDate(today) }
    case 'month':
      start.setDate(1)
      return { start: formatDate(start), end: formatDate(today) }
    default:
      start.setDate(today.getDate() - 6)
      return { start: formatDate(start), end: formatDate(today) }
  }
}

const handlePresetClick = (preset: string) => {
  selectedPreset.value = preset
  const range = getDateRange(preset)
  emit('update:modelValue', range)
  emit('change', range)
}

const displayText = computed(() => {
  if (props.modelValue) {
    return `${props.modelValue.start} 至 ${props.modelValue.end}`
  }
  const range = getDateRange(selectedPreset.value)
  return `${range.start} 至 ${range.end}`
})
</script>

<template>
  <div class="relative">
    <div class="flex items-center gap-2">
      <!-- 快捷选项 -->
      <div class="flex bg-gray-100 rounded-lg p-1">
        <button
          v-for="preset in presets"
          :key="preset.value"
          @click="handlePresetClick(preset.value)"
          class="px-3 py-1.5 text-sm rounded-md transition-colors"
          :class="selectedPreset === preset.value
            ? 'bg-white shadow text-blue-600'
            : 'text-gray-600 hover:text-gray-900'"
        >
          {{ preset.label }}
        </button>
      </div>
      
      <!-- 日期显示 -->
      <button
        @click="showPicker = !showPicker"
        class="px-4 py-2 border border-gray-300 rounded-lg text-sm text-gray-700 hover:bg-gray-50 flex items-center gap-2"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
            d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"/>
        </svg>
        {{ displayText }}
      </button>
    </div>
  </div>
</template>
