<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '星图', path: '/star' }, { name: '数据报表' }, { name: '数据概览' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">数据报表</h1>
      <p class="text-gray-600 mt-1">达人营销效果数据分析</p>
    </div>

    <!-- 时间筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex items-center space-x-4">
        <span class="text-gray-500">时间范围:</span>
        <div class="flex space-x-2">
          <button v-for="period in timePeriods" :key="period.value" 
            :class="filters.period === period.value ? 'bg-blue-600 text-white' : 'bg-gray-100 text-gray-700'"
            class="px-3 py-1 rounded text-sm" @click="filters.period = period.value">
            {{ period.label }}
          </button>
        </div>
        <div class="flex items-center space-x-2">
          <input type="date" v-model="filters.startDate" class="border border-gray-300 rounded px-3 py-1 text-sm">
          <span>-</span>
          <input type="date" v-model="filters.endDate" class="border border-gray-300 rounded px-3 py-1 text-sm">
        </div>
      </div>
    </div>

    <!-- 核心指标 -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
      <div v-for="stat in coreStats" :key="stat.label" class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">{{ stat.label }}</div>
        <div class="text-2xl font-bold text-gray-900 mt-1">{{ stat.value }}</div>
        <div :class="stat.trend > 0 ? 'text-green-500' : 'text-red-500'" class="text-sm mt-1">
          {{ stat.trend > 0 ? '+' : '' }}{{ stat.trend }}% 环比
        </div>
      </div>
    </div>

    <!-- 图表区域 -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">投放效果趋势</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
          <span class="text-gray-400">消耗/曝光/互动趋势图</span>
        </div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">达人类型分布</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
          <span class="text-gray-400">达人类型饼图</span>
        </div>
      </div>
    </div>

    <!-- 达人效果排行 -->
    <div class="bg-white rounded-lg shadow">
      <div class="p-4 border-b">
        <h3 class="text-lg font-medium">达人效果排行</h3>
      </div>
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">排名</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">达人</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">消耗</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">曝光</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">互动</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">互动率</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">CPM</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="(influencer, index) in topInfluencers" :key="influencer.id">
            <td class="px-4 py-3">
              <span :class="getRankClass(index + 1)" class="w-6 h-6 inline-flex items-center justify-center rounded-full text-sm font-bold">
                {{ index + 1 }}
              </span>
            </td>
            <td class="px-4 py-3">
              <div class="flex items-center">
                <img :src="influencer.avatar" class="w-10 h-10 rounded-full mr-3" alt="">
                <div>
                  <div class="font-medium">{{ influencer.name }}</div>
                  <div class="text-sm text-gray-500">{{ influencer.category }}</div>
                </div>
              </div>
            </td>
            <td class="px-4 py-3 text-sm text-right">¥{{ influencer.cost.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ influencer.impressions }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ influencer.engagement }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ influencer.engagementRate }}%</td>
            <td class="px-4 py-3 text-sm text-right">¥{{ influencer.cpm }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const timePeriods = [
  { label: '近7天', value: '7d' },
  { label: '近30天', value: '30d' },
  { label: '近90天', value: '90d' },
  { label: '自定义', value: 'custom' }
]

const filters = ref({
  period: '30d',
  startDate: '',
  endDate: ''
})

const coreStats = ref([
  { label: '总消耗', value: '¥256,800', trend: 18.5 },
  { label: '总曝光', value: '1,256万', trend: 12.3 },
  { label: '总互动', value: '86.5万', trend: 8.6 },
  { label: '平均CPM', value: '¥20.4', trend: -5.2 }
])

const topInfluencers = ref([
  { id: 1, name: '美妆达人小美', avatar: 'https://via.placeholder.com/40', category: '美妆护肤', cost: 15000, impressions: '125万', engagement: '8.6万', engagementRate: 6.88, cpm: 12.0 },
  { id: 2, name: '时尚博主Amy', avatar: 'https://via.placeholder.com/40', category: '时尚穿搭', cost: 12000, impressions: '98万', engagement: '6.2万', engagementRate: 6.33, cpm: 12.2 },
  { id: 3, name: '生活家小王', avatar: 'https://via.placeholder.com/40', category: '生活方式', cost: 8000, impressions: '76万', engagement: '5.1万', engagementRate: 6.71, cpm: 10.5 },
  { id: 4, name: '测评达人老李', avatar: 'https://via.placeholder.com/40', category: '数码科技', cost: 6500, impressions: '58万', engagement: '3.8万', engagementRate: 6.55, cpm: 11.2 },
  { id: 5, name: '美食家小陈', avatar: 'https://via.placeholder.com/40', category: '美食探店', cost: 5000, impressions: '45万', engagement: '3.2万', engagementRate: 7.11, cpm: 11.1 }
])

const getRankClass = (rank: number) => {
  if (rank === 1) return 'bg-yellow-400 text-white'
  if (rank === 2) return 'bg-gray-300 text-white'
  if (rank === 3) return 'bg-orange-400 text-white'
  return 'bg-gray-100 text-gray-600'
}
</script>
