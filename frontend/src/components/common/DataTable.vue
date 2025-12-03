<script setup lang="ts">
import { ref } from 'vue'
import Pagination from './Pagination.vue'

export interface Column {
  key: string
  title: string
  width?: string | number
  align?: 'left' | 'center' | 'right'
}

const props = defineProps<{
  columns: Column[]
  data: any[]
  loading?: boolean
  selectable?: boolean
  pagination?: {
    current: number
    total: number
    pageSize: number
  }
}>()

const emit = defineEmits<{
  'page-change': [page: number]
  'selection-change': [selected: any[]]
}>()

const selectedRows = ref<Set<number>>(new Set())

const toggleAll = (checked: boolean) => {
  if (checked) {
    props.data.forEach((_, index) => selectedRows.value.add(index))
  } else {
    selectedRows.value.clear()
  }
  emitSelection()
}

const toggleRow = (index: number) => {
  if (selectedRows.value.has(index)) {
    selectedRows.value.delete(index)
  } else {
    selectedRows.value.add(index)
  }
  emitSelection()
}

const emitSelection = () => {
  const selected = Array.from(selectedRows.value).map(i => props.data[i])
  emit('selection-change', selected)
}

const isAllSelected = () => {
  return props.data.length > 0 && selectedRows.value.size === props.data.length
}
</script>

<template>
  <div class="bg-white rounded-lg border border-gray-200">
    <!-- Loading Overlay -->
    <div v-if="loading" class="p-8 text-center text-gray-500">
      <svg class="animate-spin h-8 w-8 mx-auto text-blue-600" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
      <p class="mt-2">加载中...</p>
    </div>

    <!-- Table -->
    <div v-else class="overflow-x-auto">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th v-if="selectable" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              <input
                type="checkbox"
                :checked="isAllSelected()"
                @change="toggleAll(($event.target as HTMLInputElement).checked)"
                class="rounded border-gray-300"
              >
            </th>
            <th
              v-for="col in columns"
              :key="col.key"
:style="col.width ? { width: typeof col.width === 'number' ? `${col.width}px` : col.width } : {}"
              :class="[
                'px-6 py-3 text-xs font-medium text-gray-500 uppercase tracking-wider',
                col.align === 'center' ? 'text-center' : col.align === 'right' ? 'text-right' : 'text-left'
              ]"
            >
              {{ col.title }}
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr
            v-for="(row, rowIndex) in data"
            :key="rowIndex"
            class="hover:bg-gray-50 transition-colors"
          >
            <td v-if="selectable" class="px-6 py-4 whitespace-nowrap">
              <input
                type="checkbox"
                :checked="selectedRows.has(rowIndex)"
                @change="toggleRow(rowIndex)"
                class="rounded border-gray-300"
              >
            </td>
            <td
              v-for="col in columns"
              :key="col.key"
              :class="[
                'px-6 py-4 whitespace-nowrap text-sm',
                col.align === 'center' ? 'text-center' : col.align === 'right' ? 'text-right' : 'text-left'
              ]"
            >
              <slot :name="col.key" :row="row" :value="row[col.key]">
                {{ row[col.key] }}
              </slot>
            </td>
          </tr>
          <!-- Empty State -->
          <tr v-if="data.length === 0">
            <td :colspan="columns.length + (selectable ? 1 : 0)" class="px-6 py-12 text-center text-gray-500">
              暂无数据
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Pagination -->
    <Pagination
      v-if="pagination && data.length > 0"
      :current="pagination.current"
      :total="pagination.total"
      :page-size="pagination.pageSize"
      @update:current="emit('page-change', $event)"
    />
  </div>
</template>
