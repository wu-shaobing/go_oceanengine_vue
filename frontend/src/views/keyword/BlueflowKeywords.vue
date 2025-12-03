<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 89 })

const keywords = ref([
  { id: 'BK001', keyword: '智能手表续航', searchVolume: 12500, competition: 'low', suggestBid: 1.2, trend: 'up', category: '数码' },
  { id: 'BK002', keyword: '运动手环防水', searchVolume: 8900, competition: 'low', suggestBid: 0.9, trend: 'up', category: '数码' },
  { id: 'BK003', keyword: '健康监测手环', searchVolume: 15600, competition: 'medium', suggestBid: 1.5, trend: 'stable', category: '数码' },
  { id: 'BK004', keyword: '智能穿戴送礼', searchVolume: 6800, competition: 'low', suggestBid: 0.8, trend: 'up', category: '数码' }
])

const getCompetitionConfig = (level: string) => {
  switch (level) {
    case 'low': return { label: '低竞争', class: 'bg-green-100 text-green-700' }
    case 'medium': return { label: '中竞争', class: 'bg-yellow-100 text-yellow-700' }
    default: return { label: '高竞争', class: 'bg-red-100 text-red-700' }
  }
}

const getTrendIcon = (trend: string) => {
  switch (trend) {
    case 'up': return '📈'
    case 'down': return '📉'
    default: return '➡️'
  }
}

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '关键词' }, { name: '蓝海词' }]" />
      <h1 class="text-3xl font-bold text-gray-900">蓝海词发现</h1>
      <p class="mt-2 text-gray-600">发现低竞争高潜力关键词</p>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">发现蓝海词</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">低竞争词</p>
        <p class="text-2xl font-bold text-green-600 mt-1">65</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">上升趋势</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">42</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均建议出价</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">¥1.10</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex flex-wrap gap-4">
        <select class="px-4 py-2 border border-gray-300 rounded-lg">
          <option value="">全部分类</option>
          <option value="digital">数码</option>
          <option value="fashion">服装</option>
          <option value="beauty">美妆</option>
        </select>
        <select class="px-4 py-2 border border-gray-300 rounded-lg">
          <option value="">竞争程度</option>
          <option value="low">低竞争</option>
          <option value="medium">中竞争</option>
        </select>
        <select class="px-4 py-2 border border-gray-300 rounded-lg">
          <option value="">趋势</option>
          <option value="up">上升</option>
          <option value="stable">稳定</option>
        </select>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">关键词</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">搜索量</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">竞争度</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">建议出价</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">趋势</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">分类</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="kw in keywords" :key="kw.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ kw.keyword }}</div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ kw.searchVolume.toLocaleString() }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded text-xs', getCompetitionConfig(kw.competition).class]">
                {{ getCompetitionConfig(kw.competition).label }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm font-medium text-green-600">¥{{ kw.suggestBid.toFixed(2) }}</td>
            <td class="px-6 py-4 text-sm">{{ getTrendIcon(kw.trend) }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ kw.category }}</td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">添加</button>
              <button class="text-gray-600 hover:text-gray-800">详情</button>
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
