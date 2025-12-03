<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const dateRange = ref('7d')

const statsData = ref([
  { name: '官网落地页', pv: 125600, uv: 89500, avgTime: '2:35', bounceRate: 35.2, convRate: 12.5 },
  { name: '活动页A', pv: 89500, uv: 65800, avgTime: '3:12', bounceRate: 28.5, convRate: 15.8 },
  { name: '品牌页', pv: 56800, uv: 42500, avgTime: '1:58', bounceRate: 42.1, convRate: 8.5 },
  { name: '产品详情页', pv: 45600, uv: 35200, avgTime: '2:45', bounceRate: 32.8, convRate: 10.2 }
])
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '落地页' }, { name: '数据统计' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">落地页数据统计</h1>
          <p class="mt-2 text-gray-600">分析落地页访问和转化数据</p>
        </div>
        <select v-model="dateRange" class="px-4 py-2 border border-gray-300 rounded-lg">
          <option value="7d">最近7天</option>
          <option value="30d">最近30天</option>
          <option value="90d">最近90天</option>
        </select>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-5 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总PV</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">31.7万</p>
        <p class="text-xs text-green-600 mt-1">↑ 8.5%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总UV</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">23.3万</p>
        <p class="text-xs text-green-600 mt-1">↑ 6.2%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均停留</p>
        <p class="text-2xl font-bold text-green-600 mt-1">2:37</p>
        <p class="text-xs text-green-600 mt-1">↑ 0:12</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">跳出率</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">34.7%</p>
        <p class="text-xs text-red-600 mt-1">↑ 2.1%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均转化率</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">11.8%</p>
        <p class="text-xs text-green-600 mt-1">↑ 1.5%</p>
      </div>
    </div>

    <div class="grid grid-cols-2 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">访问趋势</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded-lg">
          <p class="text-gray-500">访问趋势图表区域</p>
        </div>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">转化漏斗</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded-lg">
          <p class="text-gray-500">转化漏斗图表区域</p>
        </div>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">各落地页数据</h3>
      </div>
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">页面名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">PV</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">UV</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">平均停留</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">跳出率</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">转化率</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="stat in statsData" :key="stat.name" class="hover:bg-gray-50">
            <td class="px-6 py-4 text-sm font-medium text-gray-900">{{ stat.name }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ stat.pv.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ stat.uv.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ stat.avgTime }}</td>
            <td class="px-6 py-4 text-sm" :class="stat.bounceRate <= 35 ? 'text-green-600' : 'text-yellow-600'">{{ stat.bounceRate }}%</td>
            <td class="px-6 py-4 text-sm font-medium" :class="stat.convRate >= 10 ? 'text-green-600' : 'text-yellow-600'">{{ stat.convRate }}%</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
