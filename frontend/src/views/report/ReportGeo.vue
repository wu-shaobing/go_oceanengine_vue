<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const dateRange = ref('7d')

const geoData = ref([
  { id: 1, name: '广东省', impressions: 2560000, clicks: 89600, cost: 35680, conversions: 1256, ctr: 3.5 },
  { id: 2, name: '江苏省', impressions: 1980000, clicks: 69300, cost: 28560, conversions: 986, ctr: 3.5 },
  { id: 3, name: '浙江省', impressions: 1650000, clicks: 57750, cost: 23890, conversions: 825, ctr: 3.5 },
  { id: 4, name: '北京市', impressions: 1280000, clicks: 51200, cost: 19560, conversions: 685, ctr: 4.0 },
  { id: 5, name: '上海市', impressions: 1150000, clicks: 46000, cost: 17890, conversions: 612, ctr: 4.0 }
])
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '报表' }, { name: '地域报表' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">地域报表</h1>
          <p class="mt-2 text-gray-600">按地域维度分析广告投放效果</p>
        </div>
        <div class="flex gap-3">
          <select v-model="dateRange" class="px-4 py-2 border border-gray-300 rounded-lg">
            <option value="7d">最近7天</option>
            <option value="30d">最近30天</option>
            <option value="90d">最近90天</option>
          </select>
          <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">导出</button>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">覆盖省份</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">31</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">覆盖城市</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">326</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">Top5消耗占比</p>
        <p class="text-2xl font-bold text-green-600 mt-1">68.2%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">最高CTR地域</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">北京 4.0%</p>
      </div>
    </div>

    <div class="grid grid-cols-2 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">地域分布地图</h3>
        <div class="h-80 flex items-center justify-center bg-gray-50 rounded-lg">
          <p class="text-gray-500">中国地图热力图区域</p>
        </div>
      </div>

      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">Top 5 地域</h3>
        <div class="space-y-4">
          <div v-for="(geo, index) in geoData" :key="geo.id" class="flex items-center gap-4">
            <span class="w-6 h-6 flex items-center justify-center rounded-full text-sm font-medium"
                  :class="index < 3 ? 'bg-blue-100 text-blue-700' : 'bg-gray-100 text-gray-600'">
              {{ index + 1 }}
            </span>
            <div class="flex-1">
              <div class="flex items-center justify-between mb-1">
                <span class="text-sm font-medium text-gray-900">{{ geo.name }}</span>
                <span class="text-sm text-gray-600">¥{{ (geo.cost / 1000).toFixed(1) }}K</span>
              </div>
              <div class="w-full h-2 bg-gray-200 rounded-full">
                <div class="h-full bg-blue-500 rounded-full" :style="{ width: `${(geo.cost / 35680) * 100}%` }"></div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">地域数据明细</h3>
      </div>
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">地域</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">展示</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">点击</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">CTR</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">消耗</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">转化</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="geo in geoData" :key="geo.id" class="hover:bg-gray-50">
            <td class="px-6 py-4 text-sm font-medium text-gray-900">{{ geo.name }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ (geo.impressions / 10000).toFixed(0) }}万</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ (geo.clicks / 10000).toFixed(1) }}万</td>
            <td class="px-6 py-4 text-sm" :class="geo.ctr >= 3.5 ? 'text-green-600' : 'text-yellow-600'">{{ geo.ctr }}%</td>
            <td class="px-6 py-4 text-sm font-medium text-gray-900">¥{{ geo.cost.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm font-medium text-green-600">{{ geo.conversions }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
