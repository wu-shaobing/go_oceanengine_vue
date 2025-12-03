<script setup lang="ts">
import { ref, onMounted, reactive } from 'vue'
import { Line, Bar, Doughnut } from 'vue-chartjs'
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  BarElement,
  ArcElement,
  Title,
  Tooltip,
  Legend,
  Filler
} from 'chart.js'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import StatsCard from '@/components/business/StatsCard.vue'
import DataTable from '@/components/common/DataTable.vue'

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  BarElement,
  ArcElement,
  Title,
  Tooltip,
  Legend,
  Filler
)

const loading = ref(true)
const dateRange = ref('7d')

const stats = reactive({
  totalSpend: 1258600,
  impressions: 15800000,
  clicks: 526000,
  conversions: 21500,
  ctr: 3.33,
  cvr: 4.09
})

const dateRanges = [
  { value: '1d', label: '今日' },
  { value: '7d', label: '近7天' },
  { value: '30d', label: '近30天' },
  { value: 'custom', label: '自定义' }
]

// 趋势图数据
const trendChartData = ref({
  labels: ['01-09', '01-10', '01-11', '01-12', '01-13', '01-14', '01-15'],
  datasets: [
    {
      label: '消耗',
      data: [150000, 168000, 145000, 192000, 178000, 205000, 220600],
      borderColor: '#3B82F6',
      backgroundColor: 'rgba(59, 130, 246, 0.1)',
      fill: true,
      tension: 0.4
    },
    {
      label: '转化',
      data: [2800, 3100, 2650, 3500, 3200, 3800, 4450],
      borderColor: '#10B981',
      backgroundColor: 'rgba(16, 185, 129, 0.1)',
      fill: true,
      tension: 0.4,
      yAxisID: 'y1'
    }
  ]
})

const trendChartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  interaction: {
    mode: 'index' as const,
    intersect: false
  },
  plugins: {
    legend: {
      position: 'top' as const
    }
  },
  scales: {
    y: {
      type: 'linear' as const,
      display: true,
      position: 'left' as const,
      ticks: {
        callback: function(value: string | number) {
          const numValue = typeof value === 'string' ? parseFloat(value) : value
          return '¥' + (numValue / 10000).toFixed(0) + '万'
        }
      }
    },
    y1: {
      type: 'linear' as const,
      display: true,
      position: 'right' as const,
      grid: {
        drawOnChartArea: false
      }
    }
  }
}

// 渠道分布数据
const channelChartData = ref({
  labels: ['信息流', '短视频', '搜索广告', '开屏广告'],
  datasets: [{
    data: [45, 30, 15, 10],
    backgroundColor: ['#3B82F6', '#10B981', '#F59E0B', '#8B5CF6']
  }]
})

const channelChartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: {
      position: 'right' as const
    }
  }
}

// 广告系列消耗排行数据
const campaignBarData = ref({
  labels: ['系列A', '系列B', '系列C', '系列D', '系列E'],
  datasets: [{
    label: '消耗 (万元)',
    data: [32.5, 28.6, 22.1, 18.8, 15.2],
    backgroundColor: '#3B82F6'
  }]
})

const campaignBarOptions = {
  responsive: true,
  maintainAspectRatio: false,
  indexAxis: 'y' as const,
  plugins: {
    legend: {
      display: false
    }
  }
}

// 表格数据
const tableColumns = [
  { key: 'name', title: '广告系列' },
  { key: 'spend', title: '消耗', align: 'right' as const },
  { key: 'impressions', title: '展示', align: 'right' as const },
  { key: 'clicks', title: '点击', align: 'right' as const },
  { key: 'ctr', title: 'CTR', align: 'right' as const },
  { key: 'conversions', title: '转化', align: 'right' as const },
  { key: 'cvr', title: 'CVR', align: 'right' as const }
]

const tableData = ref<any[]>([])

const fetchReportData = async () => {
  loading.value = true
  await new Promise(resolve => setTimeout(resolve, 500))
  
  tableData.value = Array.from({ length: 10 }, (_, i) => ({
    name: `广告系列 ${i + 1}`,
    spend: Math.floor(Math.random() * 100000) + 50000,
    impressions: Math.floor(Math.random() * 2000000) + 500000,
    clicks: Math.floor(Math.random() * 50000) + 10000,
    ctr: (Math.random() * 3 + 1).toFixed(2),
    conversions: Math.floor(Math.random() * 3000) + 500,
    cvr: (Math.random() * 5 + 2).toFixed(2)
  }))
  
  loading.value = false
}

const formatMoney = (value: number) => `¥${value.toLocaleString()}`
const formatNumber = (value: number) => value.toLocaleString()

onMounted(() => {
  fetchReportData()
})
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '数据报告' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">数据报告</h1>
          <p class="mt-2 text-gray-600">查看广告投放数据分析</p>
        </div>
        <div class="flex items-center gap-3">
          <div class="flex bg-gray-100 rounded-lg p-1">
            <button
              v-for="range in dateRanges"
              :key="range.value"
              @click="dateRange = range.value"
              class="px-4 py-2 text-sm rounded-md transition-colors"
              :class="{
                'bg-white shadow text-blue-600': dateRange === range.value,
                'text-gray-600 hover:text-gray-900': dateRange !== range.value
              }"
            >
              {{ range.label }}
            </button>
          </div>
          <button class="px-4 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors flex items-center gap-2">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"/>
            </svg>
            导出报告
          </button>
        </div>
      </div>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-6 gap-4">
      <StatsCard title="总消耗" :value="formatMoney(stats.totalSpend)" color="blue" />
      <StatsCard title="总展示" :value="formatNumber(stats.impressions)" color="purple" />
      <StatsCard title="总点击" :value="formatNumber(stats.clicks)" color="green" />
      <StatsCard title="总转化" :value="formatNumber(stats.conversions)" color="orange" />
      <StatsCard title="点击率" :value="`${stats.ctr}%`" color="blue" />
      <StatsCard title="转化率" :value="`${stats.cvr}%`" color="green" />
    </div>

    <!-- Charts Row -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- Trend Chart -->
      <div class="lg:col-span-2 bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">投放趋势</h3>
        <div class="h-80">
          <Line :data="trendChartData" :options="trendChartOptions" />
        </div>
      </div>

      <!-- Channel Distribution -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">渠道分布</h3>
        <div class="h-80">
          <Doughnut :data="channelChartData" :options="channelChartOptions" />
        </div>
      </div>
    </div>

    <!-- Campaign Performance -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- Top Campaigns -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">消耗排行 TOP5</h3>
        <div class="h-64">
          <Bar :data="campaignBarData" :options="campaignBarOptions" />
        </div>
      </div>

      <!-- Key Metrics -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">核心指标</h3>
        <div class="space-y-4">
          <div class="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
            <div>
              <p class="text-sm text-gray-500">平均点击成本 (CPC)</p>
              <p class="text-2xl font-bold text-gray-900">¥2.39</p>
            </div>
            <div class="text-green-600 flex items-center gap-1">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 10l7-7m0 0l7 7m-7-7v18"/>
              </svg>
              <span>-5.2%</span>
            </div>
          </div>
          <div class="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
            <div>
              <p class="text-sm text-gray-500">平均转化成本 (CPA)</p>
              <p class="text-2xl font-bold text-gray-900">¥58.54</p>
            </div>
            <div class="text-red-600 flex items-center gap-1">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3"/>
              </svg>
              <span>+2.8%</span>
            </div>
          </div>
          <div class="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
            <div>
              <p class="text-sm text-gray-500">投资回报率 (ROI)</p>
              <p class="text-2xl font-bold text-gray-900">285%</p>
            </div>
            <div class="text-green-600 flex items-center gap-1">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 10l7-7m0 0l7 7m-7-7v18"/>
              </svg>
              <span>+12.5%</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Data Table -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="p-6 border-b border-gray-200 flex items-center justify-between">
        <h3 class="text-lg font-semibold text-gray-900">广告系列明细</h3>
        <div class="flex items-center gap-3">
          <button class="px-3 py-1.5 text-sm text-gray-600 hover:text-gray-900 rounded border border-gray-300">
            筛选
          </button>
          <button class="px-3 py-1.5 text-sm text-gray-600 hover:text-gray-900 rounded border border-gray-300">
            列设置
          </button>
        </div>
      </div>

      <DataTable :columns="tableColumns" :data="tableData" :loading="loading">
        <template #name="{ value }">
          <span class="text-blue-600 hover:text-blue-800 cursor-pointer">{{ value }}</span>
        </template>

        <template #spend="{ value }">
          <span class="font-medium text-gray-900">{{ formatMoney(value) }}</span>
        </template>

        <template #impressions="{ value }">
          <span class="text-gray-600">{{ formatNumber(value) }}</span>
        </template>

        <template #clicks="{ value }">
          <span class="text-gray-600">{{ formatNumber(value) }}</span>
        </template>

        <template #ctr="{ value }">
          <span class="text-gray-600">{{ value }}%</span>
        </template>

        <template #conversions="{ value }">
          <span class="text-gray-600">{{ formatNumber(value) }}</span>
        </template>

        <template #cvr="{ value }">
          <span class="text-gray-600">{{ value }}%</span>
        </template>
      </DataTable>
    </div>
  </div>
</template>
