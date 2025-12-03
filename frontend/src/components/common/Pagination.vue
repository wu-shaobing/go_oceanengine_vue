<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  current: number
  total: number
  pageSize: number
}>()

const emit = defineEmits<{
  'update:current': [page: number]
}>()

const totalPages = computed(() => Math.ceil(props.total / props.pageSize))
const startItem = computed(() => (props.current - 1) * props.pageSize + 1)
const endItem = computed(() => Math.min(props.current * props.pageSize, props.total))

const visiblePages = computed(() => {
  const pages: number[] = []
  const start = Math.max(1, props.current - 2)
  const end = Math.min(totalPages.value, props.current + 2)
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  return pages
})

const goToPage = (page: number) => {
  if (page >= 1 && page <= totalPages.value) {
    emit('update:current', page)
  }
}
</script>

<template>
  <div class="px-6 py-4 border-t border-gray-200 flex items-center justify-between">
    <div class="text-sm text-gray-600">
      显示 <span class="font-medium">{{ startItem }}-{{ endItem }}</span> 条，共 <span class="font-medium">{{ total }}</span> 条
    </div>
    <div class="flex items-center gap-2">
      <button
        :disabled="current <= 1"
        @click="goToPage(current - 1)"
        class="px-3 py-1 border border-gray-300 rounded-lg text-sm hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
      >
        上一页
      </button>
      <button
        v-for="page in visiblePages"
        :key="page"
        @click="goToPage(page)"
        :class="[
          'px-3 py-1 rounded-lg text-sm',
          page === current
            ? 'bg-blue-600 text-white'
            : 'border border-gray-300 hover:bg-gray-50'
        ]"
      >
        {{ page }}
      </button>
      <button
        :disabled="current >= totalPages"
        @click="goToPage(current + 1)"
        class="px-3 py-1 border border-gray-300 rounded-lg text-sm hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
      >
        下一页
      </button>
    </div>
  </div>
</template>
