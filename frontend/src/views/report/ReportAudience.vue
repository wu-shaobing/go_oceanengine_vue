<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const dateRange = ref('7d')

const ageData = ref([
  { range: '18-23', impressions: 2560000, clicks: 102400, ctr: 4.0, conversions: 1580, ratio: 28 },
  { range: '24-30', impressions: 3180000, clicks: 111300, ctr: 3.5, conversions: 1890, ratio: 35 },
  { range: '31-40', impressions: 1980000, clicks: 69300, ctr: 3.5, conversions: 1125, ratio: 22 },
  { range: '41-50', impressions: 890000, clicks: 26700, ctr: 3.0, conversions: 456, ratio: 10 },
  { range: '50+', impressions: 450000, clicks: 11250, ctr: 2.5, conversions: 189, ratio: 5 }
])

const genderData = ref([
  { gender: '男性', impressions: 4560000, clicks: 159600, ctr: 3.5, ratio: 52 },
  { gender: '女性', impressions: 4200000, clicks: 168000, ctr: 4.0, ratio: 48 }
])
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '报表' }, { name: '受众报表' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">受众报表</h1>
          <p class="mt-2 text-gray-600">分析广告受众人群特征</p>
        </div>
        <select v-model="dateRange" class="px-4 py-2 border border-gray-300 rounded-lg">
          <option value="7d">最近7天</option>
          <option value="30d">最近30天</option>
          <option value="90d">最近90天</option>
        </select>
      </div>
    </div>

    <div class="grid grid-cols-2 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">性别分布</h3>
        <div class="flex items-center gap-8">
          <div class="w-32 h-32 rounded-full bg-gradient-to-r from-blue-500 to-pink-500 flex items-center justify-center">
            <div class="w-24 h-24 rounded-full bg-white flex items-center justify-center">
              <span class="text-sm font-medium text-gray-600">总计</span>
            </div>
          </div>
          <div class="space-y-4">
            <div v-for="g in genderData" :key="g.gender" class="flex items-center gap-3">
              <span :class="['w-3 h-3 rounded-full', g.gender === '男性' ? 'bg-blue-500' : 'bg-pink-500']"></span>
              <span class="text-sm text-gray-600">{{ g.gender }}</span>
              <span class="text-sm font-medium text-gray-900">{{ g.ratio }}%</span>
            </div>
          </div>
        </div>
        <table class="w-full mt-6">
          <thead>
            <tr class="text-xs text-gray-500 border-b">
              <th class="py-2 text-left">性别</th>
              <th class="py-2 text-left">展示</th>
              <th class="py-2 text-left">点击</th>
              <th class="py-2 text-left">CTR</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="g in genderData" :key="g.gender" class="text-sm">
              <td class="py-2 font-medium text-gray-900">{{ g.gender }}</td>
              <td class="py-2 text-gray-600">{{ (g.impressions / 10000).toFixed(0) }}万</td>
              <td class="py-2 text-gray-600">{{ (g.clicks / 10000).toFixed(1) }}万</td>
              <td class="py-2" :class="g.ctr >= 3.5 ? 'text-green-600' : 'text-yellow-600'">{{ g.ctr }}%</td>
            </tr>
          </tbody>
        </table>
      </div>

      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">年龄分布</h3>
        <div class="space-y-3">
          <div v-for="age in ageData" :key="age.range" class="flex items-center gap-4">
            <span class="w-16 text-sm text-gray-600">{{ age.range }}岁</span>
            <div class="flex-1 h-6 bg-gray-100 rounded-full overflow-hidden">
              <div class="h-full bg-blue-500 rounded-full flex items-center justify-end pr-2"
                   :style="{ width: `${age.ratio}%` }">
                <span v-if="age.ratio >= 15" class="text-xs text-white font-medium">{{ age.ratio }}%</span>
              </div>
            </div>
            <span v-if="age.ratio < 15" class="text-xs text-gray-500">{{ age.ratio }}%</span>
          </div>
        </div>
        <table class="w-full mt-6">
          <thead>
            <tr class="text-xs text-gray-500 border-b">
              <th class="py-2 text-left">年龄段</th>
              <th class="py-2 text-left">展示</th>
              <th class="py-2 text-left">CTR</th>
              <th class="py-2 text-left">转化</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="age in ageData" :key="age.range" class="text-sm">
              <td class="py-2 font-medium text-gray-900">{{ age.range }}岁</td>
              <td class="py-2 text-gray-600">{{ (age.impressions / 10000).toFixed(0) }}万</td>
              <td class="py-2" :class="age.ctr >= 3.5 ? 'text-green-600' : 'text-yellow-600'">{{ age.ctr }}%</td>
              <td class="py-2 font-medium text-green-600">{{ age.conversions }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">兴趣分布 Top 10</h3>
      <div class="grid grid-cols-5 gap-4">
        <div v-for="i in 10" :key="i" class="text-center p-3 bg-gray-50 rounded-lg">
          <span class="text-2xl">{{ ['🛍️', '📱', '🎮', '🏃', '🍔', '✈️', '💄', '📚', '🎵', '🎬'][i - 1] }}</span>
          <p class="text-sm font-medium text-gray-900 mt-2">兴趣{{ i }}</p>
          <p class="text-xs text-gray-500">{{ (100 - i * 5).toFixed(0) }}%</p>
        </div>
      </div>
    </div>
  </div>
</template>
