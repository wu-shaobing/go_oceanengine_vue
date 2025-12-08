<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 24 })

const packages = ref([
  { id: 'BF001', name: '电商蓝海词包', industry: '电商', keywords: 1256, cpc: 2.8, status: 'active', createdAt: '2025-11-01' },
  { id: 'BF002', name: '教育行业蓝海', industry: '教育', keywords: 892, cpc: 3.2, status: 'active', createdAt: '2025-10-28' },
  { id: 'BF003', name: '游戏推广词包', industry: '游戏', keywords: 456, cpc: 4.5, status: 'paused', createdAt: '2025-10-25' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleViewKeywords = (pkg: typeof packages.value[0]) => {
  alert(`查看流量包关键词: ${pkg.name}`)
}

const handleSettings = (pkg: typeof packages.value[0]) => {
  alert(`设置流量包: ${pkg.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '关键词工具' }, { name: '蓝海流量包' }]" />
      <h1 class="text-3xl font-bold text-gray-900">蓝海流量包</h1>
      <p class="mt-2 text-gray-600">管理蓝海关键词流量包</p>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总流量包</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总关键词</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">2,604</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均CPC</p>
        <p class="text-2xl font-bold text-green-600 mt-1">¥3.5</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">流量包</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">行业</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">关键词数</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">平均CPC</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">创建时间</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="pkg in packages" :key="pkg.id" class="hover:bg-gray-50">
              <td class="px-6 py-4">
                <div class="text-sm font-medium text-gray-900">{{ pkg.name }}</div>
                <div class="text-xs text-gray-500">ID: {{ pkg.id }}</div>
              </td>
              <td class="px-6 py-4">
                <span class="px-2 py-1 bg-blue-100 text-blue-700 rounded text-xs">{{ pkg.industry }}</span>
              </td>
              <td class="px-6 py-4 text-sm text-gray-900">{{ pkg.keywords.toLocaleString() }}</td>
              <td class="px-6 py-4 text-sm text-gray-900">¥{{ pkg.cpc }}</td>
              <td class="px-6 py-4">
                <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                  pkg.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                  {{ pkg.status === 'active' ? '启用' : '暂停' }}
                </span>
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">{{ pkg.createdAt }}</td>
              <td class="px-6 py-4 text-sm">
<button @click="handleViewKeywords(pkg)" class="text-blue-600 hover:text-blue-800 mr-3">查看词</button>
                <button @click="handleSettings(pkg)" class="text-gray-600 hover:text-gray-800">设置</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
