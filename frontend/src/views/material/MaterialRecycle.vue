<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 28 })
const selectedItems = ref<string[]>([])

const recycleItems = ref([
  { id: 'R001', name: '促销视频A.mp4', type: 'video', size: '45MB', deletedAt: '2025-11-25', expiresIn: 5 },
  { id: 'R002', name: '产品图片-红色.jpg', type: 'image', size: '2.5MB', deletedAt: '2025-11-26', expiresIn: 6 },
  { id: 'R003', name: '品牌Banner.png', type: 'image', size: '1.8MB', deletedAt: '2025-11-27', expiresIn: 7 },
  { id: 'R004', name: '测试素材.mp4', type: 'video', size: '120MB', deletedAt: '2025-11-20', expiresIn: 1 }
])

const toggleSelect = (id: string) => {
  const idx = selectedItems.value.indexOf(id)
  if (idx > -1) selectedItems.value.splice(idx, 1)
  else selectedItems.value.push(id)
}

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '素材库' }, { name: '回收站' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">回收站</h1>
          <p class="mt-2 text-gray-600">已删除的素材将在30天后永久删除</p>
        </div>
        <div class="flex gap-3">
          <button v-if="selectedItems.length > 0" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
            恢复选中 ({{ selectedItems.length }})
          </button>
          <button class="px-4 py-2 border border-red-300 text-red-600 rounded-lg hover:bg-red-50">
            清空回收站
          </button>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">回收站文件</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">占用空间</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">169MB</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">即将过期</p>
        <p class="text-2xl font-bold text-red-600 mt-1">3</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">
              <input type="checkbox" class="rounded text-blue-600">
            </th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">文件名</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">类型</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">大小</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">删除时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">剩余天数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="item in recycleItems" :key="item.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <input type="checkbox" :checked="selectedItems.includes(item.id)"
                     @change="toggleSelect(item.id)" class="rounded text-blue-600">
            </td>
            <td class="px-6 py-4">
              <div class="flex items-center gap-3">
                <span class="text-xl">{{ item.type === 'video' ? '🎬' : '🖼️' }}</span>
                <span class="text-sm font-medium text-gray-900">{{ item.name }}</span>
              </div>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded text-xs',
                     item.type === 'video' ? 'bg-purple-100 text-purple-700' : 'bg-blue-100 text-blue-700']">
                {{ item.type === 'video' ? '视频' : '图片' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ item.size }}</td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ item.deletedAt }}</td>
            <td class="px-6 py-4">
              <span :class="['text-sm font-medium', item.expiresIn <= 3 ? 'text-red-600' : 'text-gray-600']">
                {{ item.expiresIn }} 天
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">恢复</button>
              <button class="text-red-600 hover:text-red-800">永久删除</button>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
