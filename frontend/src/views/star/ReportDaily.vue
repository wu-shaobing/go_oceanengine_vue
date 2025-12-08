<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '星图', path: '/star' }, { name: '每日报告' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">每日数据报告</h1>
      <p class="text-gray-600 mt-1">查看每日任务数据表现</p>
    </div>

    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <select v-model="filters.task" class="border border-gray-300 rounded px-3 py-2 w-48">
          <option value="">全部任务</option>
          <option value="1">618美妆种草任务</option>
          <option value="2">新品护肤推广</option>
        </select>
        <input type="date" v-model="filters.startDate" class="border border-gray-300 rounded px-3 py-2">
        <span>至</span>
        <input type="date" v-model="filters.endDate" class="border border-gray-300 rounded px-3 py-2">
<button @click="handleQuery" class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">查询</button>
        <button @click="handleExport" class="px-4 py-2 border border-gray-300 rounded hover:bg-gray-50">导出</button>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-5 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">总播放</div>
        <div class="text-xl font-bold mt-1">{{ (summary.views / 10000).toFixed(1) }}万</div>
        <div class="text-xs text-green-600 mt-1">+15.2%</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">总点赞</div>
        <div class="text-xl font-bold mt-1">{{ (summary.likes / 10000).toFixed(1) }}万</div>
        <div class="text-xs text-green-600 mt-1">+12.8%</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">总评论</div>
        <div class="text-xl font-bold mt-1">{{ summary.comments.toLocaleString() }}</div>
        <div class="text-xs text-green-600 mt-1">+8.5%</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">总分享</div>
        <div class="text-xl font-bold mt-1">{{ summary.shares.toLocaleString() }}</div>
        <div class="text-xs text-green-600 mt-1">+10.2%</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">转化数</div>
        <div class="text-xl font-bold text-blue-600 mt-1">{{ summary.conversions.toLocaleString() }}</div>
        <div class="text-xs text-green-600 mt-1">+18.6%</div>
      </div>
    </div>

    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <h3 class="text-lg font-medium mb-4">数据趋势</h3>
      <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
        <span class="text-gray-400">图表区域</span>
      </div>
    </div>

    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">日期</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">播放</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">点赞</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">评论</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">分享</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">转化</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">互动率</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="item in dailyData" :key="item.date">
            <td class="px-4 py-3 font-medium">{{ item.date }}</td>
            <td class="px-4 py-3 text-right">{{ (item.views / 10000).toFixed(1) }}万</td>
            <td class="px-4 py-3 text-right">{{ (item.likes / 10000).toFixed(1) }}万</td>
            <td class="px-4 py-3 text-right">{{ item.comments.toLocaleString() }}</td>
            <td class="px-4 py-3 text-right">{{ item.shares.toLocaleString() }}</td>
            <td class="px-4 py-3 text-right text-blue-600">{{ item.conversions }}</td>
            <td class="px-4 py-3 text-right">{{ item.engagementRate }}%</td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="30" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const filters = ref({ task: '', startDate: '', endDate: '' })

const summary = ref({ views: 5680000, likes: 256000, comments: 18500, shares: 8600, conversions: 3560 })

const dailyData = ref([
  { date: '2024-06-18', views: 856000, likes: 38500, comments: 2800, shares: 1280, conversions: 536, engagementRate: 5.0 },
  { date: '2024-06-17', views: 768000, likes: 35200, comments: 2560, shares: 1150, conversions: 485, engagementRate: 5.1 },
  { date: '2024-06-16', views: 925000, likes: 42800, comments: 3100, shares: 1420, conversions: 612, engagementRate: 5.1 },
  { date: '2024-06-15', views: 680000, likes: 31500, comments: 2280, shares: 1050, conversions: 425, engagementRate: 5.1 },
  { date: '2024-06-14', views: 745000, likes: 34200, comments: 2480, shares: 1120, conversions: 468, engagementRate: 5.1 }
])

const handleQuery = () => {
  alert('查询每日报告')
}

const handleExport = () => {
  alert('导出每日报告')
}
</script>
