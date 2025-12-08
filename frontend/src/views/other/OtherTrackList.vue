<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 78 })
const searchKeyword = ref('')

const trackLinks = ref([
  { id: 'TL001', name: '双11落地页追踪', url: 'https://track.example.com/c/abc123', clicks: 125600, conversions: 3560, cvr: 2.83, createdAt: '2025-11-01' },
  { id: 'TL002', name: '品牌官网追踪', url: 'https://track.example.com/c/def456', clicks: 89500, conversions: 2180, cvr: 2.44, createdAt: '2025-10-15' },
  { id: 'TL003', name: 'APP下载追踪', url: 'https://track.example.com/c/ghi789', clicks: 56800, conversions: 1890, cvr: 3.33, createdAt: '2025-10-20' },
  { id: 'TL004', name: '活动页追踪', url: 'https://track.example.com/c/jkl012', clicks: 34500, conversions: 980, cvr: 2.84, createdAt: '2025-11-10' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleCreateTrack = () => {
  alert('创建追踪链接')
}

const handleSearch = () => {
  alert('搜索链接')
}

const handleCopyLink = (link: typeof trackLinks.value[0]) => {
  alert(`已复制链接: ${link.url}`)
}

const handleViewDetail = (link: typeof trackLinks.value[0]) => {
  alert(`查看详情: ${link.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '其他' }, { name: '追踪链接' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">追踪链接管理</h1>
          <p class="mt-2 text-gray-600">管理广告追踪链接</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleCreateTrack">
          创建追踪链接
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">追踪链接</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总点击</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">30.6万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总转化</p>
        <p class="text-2xl font-bold text-green-600 mt-1">8,610</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均转化率</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">2.81%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex gap-4">
        <input v-model="searchKeyword" type="text" placeholder="搜索链接名称..." class="flex-1 px-4 py-2 border border-gray-300 rounded-lg">
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleSearch">搜索</button>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">链接名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">追踪URL</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">点击数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">转化数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">转化率</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">创建时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="link in trackLinks" :key="link.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ link.name }}</div>
              <div class="text-xs text-gray-500">{{ link.id }}</div>
            </td>
            <td class="px-6 py-4">
              <code class="text-xs bg-gray-100 px-2 py-1 rounded">{{ link.url.substring(0, 35) }}...</code>
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ link.clicks.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm font-medium text-green-600">{{ link.conversions.toLocaleString() }}</td>
            <td class="px-6 py-4">
              <span :class="['text-sm font-medium', link.cvr >= 3 ? 'text-green-600' : 'text-yellow-600']">
                {{ link.cvr }}%
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ link.createdAt }}</td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3" @click="handleCopyLink(link)">复制</button>
              <button class="text-gray-600 hover:text-gray-800" @click="handleViewDetail(link)">详情</button>
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
