<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 32 })

const shareLinks = ref([
  { id: 'SL001', appName: '购物APP', shortUrl: 'https://s.oe.com/abc123', clicks: 15600, installs: 2340, cvr: 15.0, createdAt: '2025-11-20' },
  { id: 'SL002', appName: '游戏A', shortUrl: 'https://s.oe.com/def456', clicks: 28900, installs: 5670, cvr: 19.6, createdAt: '2025-11-18' },
  { id: 'SL003', appName: '工具应用', shortUrl: 'https://s.oe.com/ghi789', clicks: 8900, installs: 890, cvr: 10.0, createdAt: '2025-11-25' }
])

const formatNumber = (num: number) => {
  if (num >= 10000) return (num / 10000).toFixed(1) + '万'
  return num.toLocaleString()
}

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '应用工具' }, { name: '应用分享' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">应用分享链接</h1>
          <p class="mt-2 text-gray-600">创建和管理应用分享短链</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建分享链接
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">链接数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总点击</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">5.3万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总安装</p>
        <p class="text-2xl font-bold text-green-600 mt-1">8,900</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均转化率</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">16.8%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">应用</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">短链接</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">点击数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">安装数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">转化率</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">创建时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="link in shareLinks" :key="link.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ link.appName }}</div>
              <div class="text-xs text-gray-500">{{ link.id }}</div>
            </td>
            <td class="px-6 py-4">
              <div class="flex items-center gap-2">
                <a :href="link.shortUrl" target="_blank" class="text-sm text-blue-600 hover:underline">{{ link.shortUrl }}</a>
                <button class="text-gray-400 hover:text-gray-600">📋</button>
              </div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ formatNumber(link.clicks) }}</td>
            <td class="px-6 py-4 text-sm font-medium text-green-600">{{ formatNumber(link.installs) }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ link.cvr }}%</td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ link.createdAt }}</td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">详情</button>
              <button class="text-gray-600 hover:text-gray-800">二维码</button>
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
