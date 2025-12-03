<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const refreshInterval = ref(30)
const lastUpdated = ref('2025-11-28 09:30:15')

const realtimeStats = ref({
  cost: 125680,
  impressions: 2568000,
  clicks: 89560,
  conversions: 3256,
  ctr: 3.49,
  cvr: 3.64,
  cpc: 1.40
})

const hourlyData = ref([
  { hour: '00:00', cost: 2500, clicks: 1200 },
  { hour: '01:00', cost: 1800, clicks: 850 },
  { hour: '02:00', cost: 1200, clicks: 620 },
  { hour: '03:00', cost: 800, clicks: 380 },
  { hour: '04:00', cost: 600, clicks: 290 },
  { hour: '05:00', cost: 1500, clicks: 720 },
  { hour: '06:00', cost: 3500, clicks: 1680 },
  { hour: '07:00', cost: 8600, clicks: 4120 },
  { hour: '08:00', cost: 15800, clicks: 7580 },
  { hour: '09:00', cost: 18500, clicks: 8920 }
])

const formatNumber = (num: number) => {
  if (num >= 10000) return (num / 10000).toFixed(1) + '万'
  return num.toLocaleString()
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '报表中心' }, { name: '实时数据' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">实时数据</h1>
          <p class="mt-2 text-gray-600">查看今日广告投放实时效果</p>
        </div>
        <div class="flex items-center gap-4">
          <span class="text-sm text-gray-500">更新时间: {{ lastUpdated }}</span>
          <select v-model="refreshInterval" class="px-3 py-1.5 border border-gray-300 rounded text-sm">
            <option :value="30">30秒刷新</option>
            <option :value="60">1分钟刷新</option>
            <option :value="300">5分钟刷新</option>
          </select>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-7 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">消耗</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">¥{{ formatNumber(realtimeStats.cost) }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">展示</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">{{ formatNumber(realtimeStats.impressions) }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">点击</p>
        <p class="text-2xl font-bold text-green-600 mt-1">{{ formatNumber(realtimeStats.clicks) }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">转化</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">{{ realtimeStats.conversions }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">CTR</p>
        <p class="text-2xl font-bold text-orange-600 mt-1">{{ realtimeStats.ctr }}%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">CVR</p>
        <p class="text-2xl font-bold text-pink-600 mt-1">{{ realtimeStats.cvr }}%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">CPC</p>
        <p class="text-2xl font-bold text-indigo-600 mt-1">¥{{ realtimeStats.cpc }}</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h3 class="font-semibold text-gray-900 mb-4">今日分时数据</h3>
      <div class="h-64 flex items-end gap-1">
        <div v-for="item in hourlyData" :key="item.hour" class="flex-1 flex flex-col items-center">
          <div class="w-full bg-blue-500 rounded-t" :style="{ height: (item.cost / 200) + 'px' }"></div>
          <span class="text-xs text-gray-500 mt-2 rotate-45 origin-left">{{ item.hour }}</span>
        </div>
      </div>
    </div>
  </div>
</template>
