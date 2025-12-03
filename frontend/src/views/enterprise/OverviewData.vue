<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '企业号', path: '/enterprise' }, { name: '数据分析' }, { name: '数据概览' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">数据概览</h1>
      <p class="text-gray-600 mt-1">企业号经营数据分析</p>
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
      </div>
    </div>

    <!-- 核心指标 -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-4 mb-6">
      <div v-for="stat in coreStats" :key="stat.label" class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">{{ stat.label }}</div>
        <div class="text-2xl font-bold text-gray-900 mt-1">{{ stat.value }}</div>
        <div :class="stat.trend > 0 ? 'text-green-500' : 'text-red-500'" class="text-sm mt-1">
          {{ stat.trend > 0 ? '+' : '' }}{{ stat.trend }}%
        </div>
      </div>
    </div>

    <!-- 图表 -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">粉丝增长趋势</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
          <span class="text-gray-400">粉丝增长折线图</span>
        </div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">内容互动趋势</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
          <span class="text-gray-400">播放/点赞/评论趋势</span>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">粉丝画像</h3>
        <div class="space-y-4">
          <div>
            <div class="flex justify-between text-sm mb-1">
              <span>男性</span><span>42%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div class="bg-blue-600 h-2 rounded-full" style="width: 42%"></div>
            </div>
          </div>
          <div>
            <div class="flex justify-between text-sm mb-1">
              <span>女性</span><span>58%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div class="bg-pink-500 h-2 rounded-full" style="width: 58%"></div>
            </div>
          </div>
          <div class="pt-4 border-t">
            <div class="text-sm text-gray-500 mb-2">年龄分布</div>
            <div class="grid grid-cols-2 gap-2 text-sm">
              <div class="flex justify-between"><span>18-24岁</span><span>35%</span></div>
              <div class="flex justify-between"><span>25-34岁</span><span>42%</span></div>
              <div class="flex justify-between"><span>35-44岁</span><span>15%</span></div>
              <div class="flex justify-between"><span>45岁以上</span><span>8%</span></div>
            </div>
          </div>
        </div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">地域分布TOP5</h3>
        <div class="space-y-3">
          <div v-for="(region, index) in topRegions" :key="region.name" class="flex items-center">
            <span class="w-6 text-gray-400">{{ index + 1 }}</span>
            <span class="flex-1">{{ region.name }}</span>
            <span class="text-gray-600">{{ region.percent }}%</span>
          </div>
        </div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">活跃时段</h3>
        <div class="h-48 flex items-center justify-center bg-gray-50 rounded">
          <span class="text-gray-400">活跃时段热力图</span>
        </div>
        <div class="mt-4 text-sm text-gray-500">
          最佳发布时间: 19:00-21:00
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const timePeriods = [
  { label: '近7天', value: '7d' },
  { label: '近30天', value: '30d' },
  { label: '近90天', value: '90d' }
]

const filters = ref({ period: '7d' })

const coreStats = ref([
  { label: '新增粉丝', value: '12,560', trend: 15.3 },
  { label: '总播放量', value: '856万', trend: 8.2 },
  { label: '总点赞数', value: '56.8万', trend: 12.6 },
  { label: '总评论数', value: '2.8万', trend: 5.4 },
  { label: '总分享数', value: '8,560', trend: -2.1 }
])

const topRegions = ref([
  { name: '广东省', percent: 18.5 },
  { name: '江苏省', percent: 12.3 },
  { name: '浙江省', percent: 10.8 },
  { name: '北京市', percent: 8.6 },
  { name: '上海市', percent: 7.2 }
])
</script>
