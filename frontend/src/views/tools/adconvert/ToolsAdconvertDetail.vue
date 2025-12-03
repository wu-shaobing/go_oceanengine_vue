<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const convertInfo = ref({
  id: 'CV12345',
  name: '表单提交-产品咨询',
  type: 'form',
  source: '落地页',
  status: 'active',
  createdAt: '2025-11-15 10:30',
  totalConvert: 3256,
  todayConvert: 128,
  cost: 38.6
})

const dailyData = ref([
  { date: '11-22', converts: 98, cost: 42.5 },
  { date: '11-23', converts: 112, cost: 39.8 },
  { date: '11-24', converts: 95, cost: 41.2 },
  { date: '11-25', converts: 135, cost: 36.5 },
  { date: '11-26', converts: 128, cost: 38.9 },
  { date: '11-27', converts: 142, cost: 35.2 },
  { date: '11-28', converts: 128, cost: 38.6 }
])
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '转化追踪' }, { name: '转化详情' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">{{ convertInfo.name }}</h1>
          <p class="mt-2 text-gray-600">ID: {{ convertInfo.id }}</p>
        </div>
        <span :class="['px-3 py-1 rounded-full text-sm font-medium',
               convertInfo.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
          {{ convertInfo.status === 'active' ? '启用中' : '已停用' }}
        </span>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">累计转化</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ convertInfo.totalConvert.toLocaleString() }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日转化</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">{{ convertInfo.todayConvert }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均成本</p>
        <p class="text-2xl font-bold text-green-600 mt-1">¥{{ convertInfo.cost }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">创建时间</p>
        <p class="text-lg font-medium text-gray-900 mt-1">{{ convertInfo.createdAt }}</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h3 class="font-semibold text-gray-900 mb-4">近7日转化趋势</h3>
      <table class="w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-2 text-left text-xs font-semibold text-gray-600">日期</th>
            <th class="px-4 py-2 text-left text-xs font-semibold text-gray-600">转化数</th>
            <th class="px-4 py-2 text-left text-xs font-semibold text-gray-600">转化成本</th>
            <th class="px-4 py-2 text-left text-xs font-semibold text-gray-600">趋势</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="item in dailyData" :key="item.date" class="hover:bg-gray-50">
            <td class="px-4 py-3 text-sm text-gray-900">{{ item.date }}</td>
            <td class="px-4 py-3 text-sm font-medium text-blue-600">{{ item.converts }}</td>
            <td class="px-4 py-3 text-sm text-gray-900">¥{{ item.cost }}</td>
            <td class="px-4 py-3">
              <div class="w-24 h-2 bg-gray-200 rounded-full">
                <div class="h-full bg-blue-500 rounded-full" :style="{ width: (item.converts / 150 * 100) + '%' }"></div>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
