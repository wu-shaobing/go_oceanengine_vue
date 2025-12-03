<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '数据报表' }, { name: '广告主报表' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">广告主报表</h1>
        <p class="text-gray-600 mt-1">账户整体投放数据分析</p>
      </div>
      <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">
        导出报表
      </button>
    </div>

    <!-- 时间筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <div class="flex gap-2">
          <button v-for="period in timePeriods" :key="period.value"
            @click="filters.period = period.value"
            :class="filters.period === period.value ? 'bg-blue-600 text-white' : 'bg-gray-100'"
            class="px-3 py-1 rounded text-sm">
            {{ period.label }}
          </button>
        </div>
        <div class="flex items-center gap-2">
          <input type="date" v-model="filters.startDate" class="border border-gray-300 rounded-lg px-3 py-1.5 text-sm">
          <span>至</span>
          <input type="date" v-model="filters.endDate" class="border border-gray-300 rounded-lg px-3 py-1.5 text-sm">
        </div>
        <button class="px-4 py-1.5 bg-blue-600 text-white rounded-lg text-sm hover:bg-blue-700">查询</button>
      </div>
    </div>

    <!-- 核心指标 -->
    <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4 mb-6">
      <div v-for="stat in coreStats" :key="stat.label" class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">{{ stat.label }}</div>
        <div class="text-xl font-bold text-gray-900 mt-1">{{ stat.value }}</div>
        <div :class="stat.trend >= 0 ? 'text-green-500' : 'text-red-500'" class="text-sm">
          {{ stat.trend >= 0 ? '+' : '' }}{{ stat.trend }}% 环比
        </div>
      </div>
    </div>

    <!-- 趋势图 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-lg font-medium">数据趋势</h3>
        <div class="flex gap-2">
          <button v-for="metric in metrics" :key="metric.value"
            @click="selectedMetric = metric.value"
            :class="selectedMetric === metric.value ? 'bg-blue-100 text-blue-600' : ''"
            class="px-3 py-1 rounded text-sm hover:bg-gray-100">
            {{ metric.label }}
          </button>
        </div>
      </div>
      <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
        <span class="text-gray-400">{{ selectedMetric }}趋势图表</span>
      </div>
    </div>

    <!-- 详细数据表格 -->
    <div class="bg-white rounded-lg shadow">
      <div class="p-4 border-b">
        <h3 class="text-lg font-medium">日报数据</h3>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">日期</th>
              <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">消耗</th>
              <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">GMV</th>
              <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">ROI</th>
              <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">曝光</th>
              <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">点击</th>
              <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">转化</th>
              <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">转化成本</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="row in reportData" :key="row.date">
              <td class="px-4 py-3 text-sm">{{ row.date }}</td>
              <td class="px-4 py-3 text-sm text-right">¥{{ row.cost.toLocaleString() }}</td>
              <td class="px-4 py-3 text-sm text-right">¥{{ row.gmv.toLocaleString() }}</td>
              <td class="px-4 py-3 text-sm text-right font-medium">{{ row.roi }}</td>
              <td class="px-4 py-3 text-sm text-right">{{ row.impression.toLocaleString() }}</td>
              <td class="px-4 py-3 text-sm text-right">{{ row.click.toLocaleString() }}</td>
              <td class="px-4 py-3 text-sm text-right">{{ row.conversion }}</td>
              <td class="px-4 py-3 text-sm text-right">¥{{ row.conversionCost }}</td>
            </tr>
          </tbody>
        </table>
      </div>
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

const timePeriods = [
  { value: 'today', label: '今日' },
  { value: 'yesterday', label: '昨日' },
  { value: '7days', label: '近7天' },
  { value: '30days', label: '近30天' }
]

const metrics = [
  { value: '消耗', label: '消耗' },
  { value: 'GMV', label: 'GMV' },
  { value: 'ROI', label: 'ROI' },
  { value: '转化', label: '转化' }
]

const filters = ref({
  period: '7days',
  startDate: '',
  endDate: ''
})

const selectedMetric = ref('消耗')

const coreStats = ref([
  { label: '消耗', value: '¥128,456', trend: 12.5 },
  { label: 'GMV', value: '¥568,900', trend: 15.8 },
  { label: 'ROI', value: '4.43', trend: 3.2 },
  { label: '曝光', value: '5.6M', trend: 8.6 },
  { label: '点击', value: '186,500', trend: 10.3 },
  { label: '转化', value: '2,568', trend: 18.5 }
])

const reportData = ref([
  { date: '2024-03-15', cost: 18560, gmv: 86500, roi: '4.66', impression: 856000, click: 28600, conversion: 386, conversionCost: 48.1 },
  { date: '2024-03-14', cost: 16800, gmv: 72300, roi: '4.30', impression: 768000, click: 25400, conversion: 342, conversionCost: 49.1 },
  { date: '2024-03-13', cost: 19200, gmv: 89600, roi: '4.67', impression: 892000, click: 30200, conversion: 398, conversionCost: 48.2 },
  { date: '2024-03-12', cost: 17600, gmv: 76800, roi: '4.36', impression: 812000, click: 26800, conversion: 356, conversionCost: 49.4 },
  { date: '2024-03-11', cost: 18200, gmv: 82500, roi: '4.53', impression: 835000, click: 27600, conversion: 372, conversionCost: 48.9 }
])
</script>
