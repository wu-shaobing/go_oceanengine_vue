<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const dateRange = ref('7d')
const reportData = ref([
  { date: '2025-11-28', impressions: 6260000, clicks: 125200, revenue: 118930, ecpm: 19.0 },
  { date: '2025-11-27', impressions: 5890000, clicks: 117800, revenue: 112560, ecpm: 19.1 },
  { date: '2025-11-26', impressions: 6120000, clicks: 122400, revenue: 115890, ecpm: 18.9 },
  { date: '2025-11-25', impressions: 5650000, clicks: 113000, revenue: 108750, ecpm: 19.2 },
  { date: '2025-11-24', impressions: 5980000, clicks: 119600, revenue: 113620, ecpm: 19.0 }
])
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '工具' }, { name: '穿山甲收益报表' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">收益报表</h1>
          <p class="mt-2 text-gray-600">查看穿山甲广告收益数据</p>
        </div>
        <div class="flex gap-3">
          <select v-model="dateRange" class="px-4 py-2 border border-gray-300 rounded-lg">
            <option value="7d">最近7天</option>
            <option value="30d">最近30天</option>
            <option value="90d">最近90天</option>
          </select>
          <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">
            导出报表
          </button>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-5 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总展示</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">2,990万</p>
        <p class="text-xs text-green-600 mt-1">↑ 5.2%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总点击</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">59.8万</p>
        <p class="text-xs text-green-600 mt-1">↑ 3.8%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">点击率</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">2.0%</p>
        <p class="text-xs text-red-600 mt-1">↓ 0.1%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总收入</p>
        <p class="text-2xl font-bold text-green-600 mt-1">¥569,750</p>
        <p class="text-xs text-green-600 mt-1">↑ 8.5%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均eCPM</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">¥19.0</p>
        <p class="text-xs text-green-600 mt-1">↑ 2.1%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">收益趋势</h3>
      <div class="h-64 flex items-center justify-center bg-gray-50 rounded-lg">
        <p class="text-gray-500">收益趋势图表区域</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">每日数据明细</h3>
      </div>
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">日期</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">展示次数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">点击次数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">收入</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">eCPM</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="row in reportData" :key="row.date" class="hover:bg-gray-50">
            <td class="px-6 py-4 text-sm font-medium text-gray-900">{{ row.date }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ (row.impressions / 10000).toFixed(0) }}万</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ (row.clicks / 10000).toFixed(1) }}万</td>
            <td class="px-6 py-4 text-sm font-medium text-green-600">¥{{ row.revenue.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">¥{{ row.ecpm }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
