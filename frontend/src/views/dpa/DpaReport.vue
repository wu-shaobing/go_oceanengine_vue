<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const dateRange = ref('7d')

const topProducts = ref([
  { id: 'P001', name: '智能手机 Pro Max', impressions: 256800, clicks: 8960, cost: 12560, conversions: 256, revenue: 128000 },
  { id: 'P002', name: '无线蓝牙耳机', impressions: 189500, clicks: 6630, cost: 8960, conversions: 189, revenue: 37800 },
  { id: 'P003', name: '快充充电器', impressions: 125600, clicks: 4398, cost: 5680, conversions: 145, revenue: 14500 },
  { id: 'P004', name: '手机保护壳', impressions: 98600, clicks: 3452, cost: 3560, conversions: 125, revenue: 6250 }
])
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: 'DPA商品' }, { name: 'DPA报表' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">DPA投放报表</h1>
          <p class="mt-2 text-gray-600">分析DPA商品广告投放效果</p>
        </div>
        <div class="flex gap-3">
          <select v-model="dateRange" class="px-4 py-2 border border-gray-300 rounded-lg">
            <option value="7d">最近7天</option>
            <option value="30d">最近30天</option>
            <option value="90d">最近90天</option>
          </select>
          <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">导出报表</button>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-5 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总消耗</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">¥30,760</p>
        <p class="text-xs text-green-600 mt-1">↑ 12.5%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总展示</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">670万</p>
        <p class="text-xs text-green-600 mt-1">↑ 8.2%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总点击</p>
        <p class="text-2xl font-bold text-green-600 mt-1">23.4万</p>
        <p class="text-xs text-green-600 mt-1">↑ 10.5%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总转化</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">715</p>
        <p class="text-xs text-green-600 mt-1">↑ 15.8%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总收入</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">¥186,550</p>
        <p class="text-xs text-green-600 mt-1">↑ 18.2%</p>
      </div>
    </div>

    <div class="grid grid-cols-2 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">消耗趋势</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded-lg">
          <p class="text-gray-500">消耗趋势图表区域</p>
        </div>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">转化趋势</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded-lg">
          <p class="text-gray-500">转化趋势图表区域</p>
        </div>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">Top 商品表现</h3>
      </div>
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">商品名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">展示</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">点击</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">消耗</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">转化</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">收入</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">ROI</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="product in topProducts" :key="product.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ product.name }}</div>
              <div class="text-xs text-gray-500">{{ product.id }}</div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ product.impressions.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ product.clicks.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">¥{{ product.cost.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm font-medium text-green-600">{{ product.conversions }}</td>
            <td class="px-6 py-4 text-sm font-medium text-blue-600">¥{{ product.revenue.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm font-medium" :class="product.revenue / product.cost >= 5 ? 'text-green-600' : 'text-yellow-600'">
              {{ (product.revenue / product.cost).toFixed(1) }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
