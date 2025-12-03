<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const quotaData = ref([
  { api: '/ad/get/', name: '获取计划列表', dailyLimit: 5000, used: 3256, remaining: 1744 },
  { api: '/ad/create/', name: '创建计划', dailyLimit: 1000, used: 128, remaining: 872 },
  { api: '/ad/update/', name: '更新计划', dailyLimit: 2000, used: 456, remaining: 1544 },
  { api: '/creative/get/', name: '获取创意列表', dailyLimit: 5000, used: 4821, remaining: 179 },
  { api: '/creative/create/', name: '创建创意', dailyLimit: 500, used: 89, remaining: 411 },
  { api: '/report/ad/get/', name: '广告数据报表', dailyLimit: 10000, used: 8965, remaining: 1035 },
  { api: '/report/advertiser/get/', name: '账户数据报表', dailyLimit: 5000, used: 1230, remaining: 3770 },
  { api: '/file/image/ad/', name: '上传图片', dailyLimit: 2000, used: 523, remaining: 1477 },
  { api: '/file/video/ad/', name: '上传视频', dailyLimit: 500, used: 67, remaining: 433 }
])

const getUsagePercent = (used: number, limit: number) => {
  return Math.round((used / limit) * 100)
}

const getStatusColor = (percent: number) => {
  if (percent >= 90) return 'bg-red-500'
  if (percent >= 70) return 'bg-yellow-500'
  return 'bg-green-500'
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '基础工具' }, { name: '接口配额' }]" />
      <h1 class="text-3xl font-bold text-gray-900">接口配额查询</h1>
      <p class="mt-2 text-gray-600">查看API接口调用配额及使用情况</p>
    </div>

    <!-- Summary Cards -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日总调用</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">19,535</p>
        <p class="text-xs text-green-600 mt-1">↑ 12.5% 较昨日</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总配额</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">31,000</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已使用配额</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">63%</p>
        <div class="w-full bg-gray-200 rounded-full h-2 mt-2">
          <div class="bg-blue-600 h-2 rounded-full" style="width: 63%"></div>
        </div>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">即将达限</p>
        <p class="text-2xl font-bold text-orange-600 mt-1">2</p>
        <p class="text-xs text-orange-600 mt-1">需要关注</p>
      </div>
    </div>

    <!-- Quota Table -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="p-4 border-b border-gray-200">
        <h3 class="font-semibold text-gray-900">接口配额详情</h3>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">接口路径</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">接口名称</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">日配额</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">已使用</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">剩余</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">使用率</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="item in quotaData" :key="item.api" class="hover:bg-gray-50">
              <td class="px-6 py-4 text-sm font-mono text-gray-900">{{ item.api }}</td>
              <td class="px-6 py-4 text-sm text-gray-700">{{ item.name }}</td>
              <td class="px-6 py-4 text-sm text-gray-900">{{ item.dailyLimit.toLocaleString() }}</td>
              <td class="px-6 py-4 text-sm text-gray-900">{{ item.used.toLocaleString() }}</td>
              <td class="px-6 py-4 text-sm" :class="item.remaining < 500 ? 'text-red-600 font-medium' : 'text-gray-900'">
                {{ item.remaining.toLocaleString() }}
              </td>
              <td class="px-6 py-4">
                <div class="flex items-center gap-3">
                  <div class="flex-1 w-24 bg-gray-200 rounded-full h-2">
                    <div :class="['h-2 rounded-full', getStatusColor(getUsagePercent(item.used, item.dailyLimit))]"
                         :style="{ width: getUsagePercent(item.used, item.dailyLimit) + '%' }"></div>
                  </div>
                  <span class="text-sm text-gray-600 w-12">{{ getUsagePercent(item.used, item.dailyLimit) }}%</span>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Notice -->
    <div class="bg-yellow-50 border border-yellow-200 rounded-lg p-4">
      <div class="flex gap-3">
        <svg class="w-5 h-5 text-yellow-500 flex-shrink-0 mt-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
        </svg>
        <div>
          <h4 class="font-medium text-yellow-800">配额重置说明</h4>
          <p class="text-sm text-yellow-700 mt-1">接口配额每日00:00(北京时间)自动重置。如需提升配额，请联系商务经理。</p>
        </div>
      </div>
    </div>
  </div>
</template>
