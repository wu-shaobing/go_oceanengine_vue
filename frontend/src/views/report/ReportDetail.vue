<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import StatsCard from '@/components/business/StatsCard.vue'
import StatusBadge from '@/components/business/StatusBadge.vue'

const route = useRoute()
const router = useRouter()
const loading = ref(true)

interface ReportDetail {
  id: number
  name: string
  type: string
  status: string
  dateRange: string
  createdAt: string
  createdBy: string
  metrics: {
    impressions: number
    clicks: number
    conversions: number
    cost: number
    ctr: number
    cvr: number
    cpc: number
    cpa: number
  }
  dimensions: string[]
  filters: Record<string, string>
}

const report = ref<ReportDetail | null>(null)

const formatNumber = (num: number): string => {
  if (num >= 100000000) {
    return (num / 100000000).toFixed(2) + '亿'
  }
  if (num >= 10000) {
    return (num / 10000).toFixed(2) + '万'
  }
  return num.toLocaleString()
}

const formatPercent = (num: number): string => {
  return (num * 100).toFixed(2) + '%'
}

const formatCurrency = (num: number): string => {
  return '¥' + num.toLocaleString('zh-CN', { minimumFractionDigits: 2 })
}

const fetchReport = async () => {
  loading.value = true
  await new Promise(resolve => setTimeout(resolve, 500))
  
  report.value = {
    id: Number(route.params.id) || 1,
    name: '2024年Q1效果报表',
    type: 'performance',
    status: 'completed',
    dateRange: '2024-01-01 至 2024-03-31',
    createdAt: '2024-04-01 10:30:00',
    createdBy: '张三',
    metrics: {
      impressions: 125000000,
      clicks: 3750000,
      conversions: 187500,
      cost: 2250000,
      ctr: 0.03,
      cvr: 0.05,
      cpc: 0.6,
      cpa: 12
    },
    dimensions: ['日期', '广告主', '广告计划', '创意'],
    filters: {
      '广告主': '全部',
      '投放状态': '投放中',
      '出价方式': 'OCPM'
    }
  }
  
  loading.value = false
}

const reportTypeText = computed(() => {
  const types: Record<string, string> = {
    performance: '效果报表',
    creative: '创意报表',
    audience: '人群报表',
    conversion: '转化报表'
  }
  return types[report.value?.type || ''] || '未知类型'
})

const goBack = () => {
  router.push('/report')
}

const exportReport = () => {
  alert(`导出报表: ${report.value?.name}`)
}

const handleRefresh = () => {
  alert('刷新报表数据')
  fetchReport()
}

onMounted(fetchReport)
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '数据报表', path: '/report' }, { name: '报表详情' }]" />
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-4">
          <button
            @click="goBack"
            class="p-2 text-gray-500 hover:text-gray-700 hover:bg-gray-100 rounded-lg transition-colors"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"/>
            </svg>
          </button>
          <div>
            <h1 class="text-3xl font-bold text-gray-900">{{ report?.name || '报表详情' }}</h1>
            <p class="mt-2 text-gray-600">{{ reportTypeText }} · {{ report?.dateRange }}</p>
          </div>
        </div>
        <div class="flex items-center gap-3">
          <button
            @click="exportReport"
            class="px-4 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors flex items-center gap-2"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"/>
            </svg>
            导出报表
          </button>
<button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center gap-2" @click="handleRefresh">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/>
            </svg>
            刷新数据
          </button>
        </div>
      </div>
    </div>

    <!-- 加载状态 -->
    <div v-if="loading" class="flex items-center justify-center h-64">
      <div class="text-gray-500">加载中...</div>
    </div>

    <template v-else-if="report">
      <!-- 报表信息 -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">报表信息</h3>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <div>
            <p class="text-sm text-gray-500">报表类型</p>
            <p class="mt-1 font-medium text-gray-900">{{ reportTypeText }}</p>
          </div>
          <div>
            <p class="text-sm text-gray-500">状态</p>
            <div class="mt-1">
              <StatusBadge
                :status="report.status === 'completed' ? 'success' : 'warning'"
                :text="report.status === 'completed' ? '已完成' : '生成中'"
              />
            </div>
          </div>
          <div>
            <p class="text-sm text-gray-500">创建人</p>
            <p class="mt-1 font-medium text-gray-900">{{ report.createdBy }}</p>
          </div>
          <div>
            <p class="text-sm text-gray-500">创建时间</p>
            <p class="mt-1 font-medium text-gray-900">{{ report.createdAt }}</p>
          </div>
        </div>
      </div>

      <!-- 核心指标 -->
      <div>
        <h3 class="text-lg font-semibold text-gray-900 mb-4">核心指标</h3>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <StatsCard
            title="展示量"
            :value="formatNumber(report.metrics.impressions)"
            trend="up"
            trendValue="12.5%"
          />
          <StatsCard
            title="点击量"
            :value="formatNumber(report.metrics.clicks)"
            trend="up"
            trendValue="8.3%"
          />
          <StatsCard
            title="转化量"
            :value="formatNumber(report.metrics.conversions)"
            trend="up"
            trendValue="15.2%"
          />
          <StatsCard
            title="消耗"
            :value="formatCurrency(report.metrics.cost)"
            trend="down"
            trendValue="3.1%"
          />
        </div>
      </div>

      <!-- 效果指标 -->
      <div>
        <h3 class="text-lg font-semibold text-gray-900 mb-4">效果指标</h3>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <StatsCard
            title="点击率 (CTR)"
            :value="formatPercent(report.metrics.ctr)"
          />
          <StatsCard
            title="转化率 (CVR)"
            :value="formatPercent(report.metrics.cvr)"
          />
          <StatsCard
            title="点击成本 (CPC)"
            :value="formatCurrency(report.metrics.cpc)"
          />
          <StatsCard
            title="转化成本 (CPA)"
            :value="formatCurrency(report.metrics.cpa)"
          />
        </div>
      </div>

      <!-- 报表维度和筛选条件 -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h3 class="text-lg font-semibold text-gray-900 mb-4">报表维度</h3>
          <div class="flex flex-wrap gap-2">
            <span
              v-for="dim in report.dimensions"
              :key="dim"
              class="px-3 py-1 bg-blue-100 text-blue-800 text-sm rounded-full"
            >
              {{ dim }}
            </span>
          </div>
        </div>
        
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h3 class="text-lg font-semibold text-gray-900 mb-4">筛选条件</h3>
          <div class="space-y-2">
            <div
              v-for="(value, key) in report.filters"
              :key="key"
              class="flex items-center justify-between"
            >
              <span class="text-sm text-gray-500">{{ key }}</span>
              <span class="text-sm font-medium text-gray-900">{{ value }}</span>
            </div>
          </div>
        </div>
      </div>
    </template>
  </div>
</template>
