<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 356 })
const dateRange = ref('7d')

const reportData = ref([
  { id: 'KR001', keyword: '智能手机', impressions: 256800, clicks: 7704, ctr: 3.0, cost: 12560, conversions: 256, cpa: 49.1 },
  { id: 'KR002', keyword: '手机配件', impressions: 189500, clicks: 5685, ctr: 3.0, cost: 8960, conversions: 189, cpa: 47.4 },
  { id: 'KR003', keyword: '无线耳机', impressions: 125600, clicks: 4398, ctr: 3.5, cost: 6580, conversions: 145, cpa: 45.4 },
  { id: 'KR004', keyword: '充电宝', impressions: 98600, clicks: 2958, ctr: 3.0, cost: 4560, conversions: 98, cpa: 46.5 }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleExport = () => {
  alert('导出关键词报表')
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '关键词' }, { name: '关键词报表' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">关键词报表</h1>
          <p class="mt-2 text-gray-600">分析关键词投放效果数据</p>
        </div>
        <div class="flex gap-3">
          <select v-model="dateRange" class="px-4 py-2 border border-gray-300 rounded-lg">
            <option value="7d">最近7天</option>
            <option value="30d">最近30天</option>
            <option value="90d">最近90天</option>
          </select>
<button @click="handleExport" class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">导出</button>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-5 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">关键词数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总展示</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">670万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总点击</p>
        <p class="text-2xl font-bold text-green-600 mt-1">20.7万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均CTR</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">3.1%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总消耗</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">¥32,660</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">关键词</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">展示</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">点击</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">CTR</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">消耗</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">转化</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">CPA</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="row in reportData" :key="row.id" class="hover:bg-gray-50">
            <td class="px-6 py-4 text-sm font-medium text-gray-900">{{ row.keyword }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ row.impressions.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ row.clicks.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm" :class="row.ctr >= 3 ? 'text-green-600' : 'text-yellow-600'">{{ row.ctr }}%</td>
            <td class="px-6 py-4 text-sm font-medium text-gray-900">¥{{ row.cost.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm font-medium text-green-600">{{ row.conversions }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">¥{{ row.cpa }}</td>
          </tr>
        </tbody>
      </table>
      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
