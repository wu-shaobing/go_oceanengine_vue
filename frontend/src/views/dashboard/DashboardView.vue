<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { Line, Doughnut } from 'vue-chartjs'
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  ArcElement,
  Title,
  Tooltip,
  Legend,
  Filler
} from 'chart.js'
import StatsCard from '@/components/business/StatsCard.vue'
import ProductLineCard from '@/components/business/ProductLineCard.vue'

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, ArcElement, Title, Tooltip, Legend, Filler)

// 产品线入口数据
const productLines = [
  {
    title: '巨量千川',
    description: '电商广告投放平台，为电商商家提供精准流量运营服务',
    icon: 'shop',
    color: 'orange' as const,
    to: '/qianchuan'
  },
  {
    title: '本地推',
    description: '本地生活服务推广平台，助力门店获取更多到店客流',
    icon: 'location',
    color: 'green' as const,
    to: '/local'
  },
  {
    title: '企业号',
    description: '抹音企业号运营管理，内容运营和粉丝管理一站式服务',
    icon: 'building',
    color: 'purple' as const,
    to: '/enterprise'
  },
  {
    title: '巨量星图',
    description: '达人营销平台，连接品牌与达人，实现内容合作',
    icon: 'star',
    color: 'yellow' as const,
    to: '/star'
  },
  {
    title: '代理商管理',
    description: '代理商广告主管理、财务充值、数据报表综合服务',
    icon: 'agent',
    color: 'blue' as const,
    to: '/agent'
  },
  {
    title: '工具与服务',
    description: '基础工具、数据报表、素材管理等辅助功能',
    icon: 'tools',
    color: 'gray' as const,
    to: '/reports'
  }
]

const stats = ref({
  todayCost: 12845,
  todayImpressions: 245890,
  todayClicks: 18432,
  todayConversions: 1024,
  costTrend: 12.5,
  impressionsTrend: 8.3
})

const lineChartData = ref({
  labels: ['01/15', '01/16', '01/17', '01/18', '01/19', '01/20', '01/21'],
  datasets: [
    {
      label: '消耗',
      data: [9800, 11200, 10500, 13200, 10890, 11420, 12845],
      borderColor: 'rgb(59, 130, 246)',
      backgroundColor: 'rgba(59, 130, 246, 0.1)',
      fill: true,
      tension: 0.4
    },
    {
      label: '转化',
      data: [856, 920, 785, 1012, 965, 1156, 1024],
      borderColor: 'rgb(16, 185, 129)',
      backgroundColor: 'rgba(16, 185, 129, 0.1)',
      fill: true,
      tension: 0.4
    }
  ]
})

const doughnutChartData = ref({
  labels: ['抖音', '头条', '西瓜视频', '其他'],
  datasets: [{
    data: [45, 25, 18, 12],
    backgroundColor: ['#3b82f6', '#10b981', '#f59e0b', '#6b7280']
  }]
})

const chartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: { position: 'top' as const }
  }
}

const doughnutOptions = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: { position: 'right' as const }
  }
}

onMounted(() => {
  // Load dashboard data
})
</script>

<template>
  <div class="space-y-8">
    <!-- Hero Section -->
    <div class="bg-gradient-to-r from-blue-600 to-blue-700 rounded-2xl p-8 text-white">
      <div class="max-w-3xl">
        <h1 class="text-3xl font-bold">欢迎使用巨量引擎管理平台</h1>
        <p class="mt-3 text-blue-100 text-lg">一站式广告投放管理，涵盖千川、本地推、企业号、星图等产品线</p>
        <div class="mt-6 flex items-center gap-4">
          <router-link to="/qianchuan/campaign/create" class="inline-flex items-center gap-2 px-5 py-2.5 bg-white text-blue-600 font-medium rounded-lg hover:bg-blue-50 transition-colors">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
            </svg>
            创建广告计划
          </router-link>
          <router-link to="/reports" class="inline-flex items-center gap-2 px-5 py-2.5 border border-white/30 text-white font-medium rounded-lg hover:bg-white/10 transition-colors">
            查看数据报表
          </router-link>
        </div>
      </div>
    </div>

    <!-- Product Lines Grid -->
    <div>
      <h2 class="text-xl font-semibold text-gray-900 mb-4">产品线入口</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        <ProductLineCard
          v-for="product in productLines"
          :key="product.title"
          :title="product.title"
          :description="product.description"
          :icon="product.icon"
          :color="product.color"
          :to="product.to"
        />
      </div>
    </div>

    <!-- Stats Grid -->
    <div>
      <h2 class="text-xl font-semibold text-gray-900 mb-4">今日数据概览</h2>
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
      <StatsCard
        title="今日消耗"
        :value="`¥${stats.todayCost.toLocaleString()}`"
        color="blue"
        trend="up"
        :trendValue="`${stats.costTrend}%`"
      >
        <template #icon>
          <svg class="h-8 w-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
        </template>
      </StatsCard>

      <StatsCard
        title="今日展示"
        :value="stats.todayImpressions.toLocaleString()"
        color="green"
        trend="up"
        :trendValue="`${stats.impressionsTrend}%`"
      >
        <template #icon>
          <svg class="h-8 w-8 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/>
          </svg>
        </template>
      </StatsCard>

      <StatsCard
        title="今日点击"
        :value="stats.todayClicks.toLocaleString()"
        color="purple"
      >
        <template #icon>
          <svg class="h-8 w-8 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 15l-2 5L9 9l11 4-5 2zm0 0l5 5M7.188 2.239l.777 2.897M5.136 7.965l-2.898-.777M13.95 4.05l-2.122 2.122m-5.657 5.656l-2.12 2.122"/>
          </svg>
        </template>
      </StatsCard>

      <StatsCard
        title="今日转化"
        :value="stats.todayConversions.toLocaleString()"
        color="orange"
      >
        <template #icon>
          <svg class="h-8 w-8 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z"/>
          </svg>
        </template>
      </StatsCard>
      </div>
    </div>

    <!-- Charts Grid -->
    <div>
      <h2 class="text-xl font-semibold text-gray-900 mb-4">趋势分析</h2>
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">趋势分析</h3>
        <div style="height: 300px;">
          <Line :data="lineChartData" :options="chartOptions" />
        </div>
      </div>

      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">渠道分布</h3>
        <div style="height: 300px;">
          <Doughnut :data="doughnutChartData" :options="doughnutOptions" />
        </div>
        </div>
      </div>
    </div>

    <!-- Quick Actions -->
    <div class="bg-white rounded-xl border border-gray-200 p-6">
      <h2 class="text-xl font-semibold text-gray-900 mb-4">快捷操作</h2>
      <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
        <router-link to="/campaigns/create" class="flex flex-col items-center p-4 rounded-lg border border-gray-200 hover:border-blue-500 hover:shadow-md transition-all">
          <div class="p-3 rounded-lg bg-blue-50 text-blue-600 mb-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
            </svg>
          </div>
          <span class="text-sm font-medium text-gray-900">创建计划</span>
        </router-link>

        <router-link to="/reports" class="flex flex-col items-center p-4 rounded-lg border border-gray-200 hover:border-blue-500 hover:shadow-md transition-all">
          <div class="p-3 rounded-lg bg-green-50 text-green-600 mb-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/>
            </svg>
          </div>
          <span class="text-sm font-medium text-gray-900">查看报表</span>
        </router-link>

        <router-link to="/media" class="flex flex-col items-center p-4 rounded-lg border border-gray-200 hover:border-blue-500 hover:shadow-md transition-all">
          <div class="p-3 rounded-lg bg-purple-50 text-purple-600 mb-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"/>
            </svg>
          </div>
          <span class="text-sm font-medium text-gray-900">素材库</span>
        </router-link>

        <router-link to="/audiences" class="flex flex-col items-center p-4 rounded-lg border border-gray-200 hover:border-blue-500 hover:shadow-md transition-all">
          <div class="p-3 rounded-lg bg-orange-50 text-orange-600 mb-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"/>
            </svg>
          </div>
          <span class="text-sm font-medium text-gray-900">人群包</span>
        </router-link>
      </div>
    </div>
  </div>
</template>
