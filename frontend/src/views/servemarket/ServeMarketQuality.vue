<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '服务市场', path: '/servemarket' }, { name: '质量报告' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">质量报告</h1>
      <p class="text-gray-600 mt-1">服务使用质量分析报告</p>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">服务使用次数</div>
        <div class="text-2xl font-bold mt-1">{{ stats.usageCount.toLocaleString() }}</div>
        <div class="text-xs text-green-600 mt-1">较上周 +12%</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">平均响应时间</div>
        <div class="text-2xl font-bold mt-1">{{ stats.avgResponse }}ms</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">成功率</div>
        <div class="text-2xl font-bold text-green-600 mt-1">{{ stats.successRate }}%</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">满意度</div>
        <div class="text-2xl font-bold text-blue-600 mt-1">{{ stats.satisfaction }}</div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">服务使用趋势</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
          <span class="text-gray-400">图表区域</span>
        </div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">服务性能分布</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
          <span class="text-gray-400">图表区域</span>
        </div>
      </div>
    </div>

    <div class="bg-white rounded-lg shadow p-4">
      <h3 class="text-lg font-medium mb-4">服务使用明细</h3>
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">服务名称</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">使用次数</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">成功率</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">平均响应</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">状态</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="item in details" :key="item.id">
            <td class="px-4 py-3 font-medium">{{ item.name }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ item.count.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ item.successRate }}%</td>
            <td class="px-4 py-3 text-sm text-right">{{ item.avgResponse }}ms</td>
            <td class="px-4 py-3">
              <span :class="item.status === 'good' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800'" class="px-2 py-1 text-xs rounded">
                {{ item.status === 'good' ? '正常' : '待优化' }}
              </span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const stats = ref({ usageCount: 15680, avgResponse: 128, successRate: 99.2, satisfaction: '4.8/5.0' })

const details = ref([
  { id: 1, name: '智能文案生成', count: 5680, successRate: 99.8, avgResponse: 85, status: 'good' },
  { id: 2, name: '素材智能分析', count: 3240, successRate: 98.5, avgResponse: 156, status: 'good' },
  { id: 3, name: '竞品分析报告', count: 1560, successRate: 97.2, avgResponse: 320, status: 'warning' },
  { id: 4, name: '智能投放建议', count: 5200, successRate: 99.5, avgResponse: 95, status: 'good' }
])
</script>
