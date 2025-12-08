<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 56 })

const orders = ref([
  { id: 'DP001', videoTitle: '产品展示视频', amount: 500, target: 'plays', targetValue: 50000, actualValue: 52300, status: 'completed', createdAt: '2025-11-25' },
  { id: 'DP002', videoTitle: '品牌宣传片', amount: 1000, target: 'likes', targetValue: 10000, actualValue: 8500, status: 'running', createdAt: '2025-11-27' },
  { id: 'DP003', videoTitle: '新品发布会回顾', amount: 300, target: 'plays', targetValue: 30000, actualValue: 31200, status: 'completed', createdAt: '2025-11-20' },
  { id: 'DP004', videoTitle: '用户测评合集', amount: 200, target: 'fans', targetValue: 500, actualValue: 320, status: 'running', createdAt: '2025-11-28' }
])

const formatNumber = (num: number) => {
  if (num >= 10000) return (num / 10000).toFixed(1) + '万'
  return num.toLocaleString()
}

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleCreateOrder = () => {
  alert('创建抖+订单')
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '其他工具' }, { name: '抖+订单' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">抖+订单管理</h1>
          <p class="mt-2 text-gray-600">管理抖音视频推广订单</p>
        </div>
<button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleCreateOrder">
          创建抖+订单
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总订单数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">投放中</p>
        <p class="text-2xl font-bold text-green-600 mt-1">8</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总消耗</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">¥12,500</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总播放</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">125万</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">视频</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">投放金额</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">推广目标</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">目标值</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">实际值</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">创建时间</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="order in orders" :key="order.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ order.videoTitle }}</div>
              <div class="text-xs text-gray-500">{{ order.id }}</div>
            </td>
            <td class="px-6 py-4 text-sm font-medium text-gray-900">¥{{ order.amount }}</td>
            <td class="px-6 py-4">
              <span class="px-2 py-1 bg-gray-100 text-gray-700 rounded text-xs">
                {{ order.target === 'plays' ? '播放量' : order.target === 'likes' ? '点赞数' : '涨粉数' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ formatNumber(order.targetValue) }}</td>
            <td class="px-6 py-4 text-sm font-medium" :class="order.actualValue >= order.targetValue ? 'text-green-600' : 'text-blue-600'">
              {{ formatNumber(order.actualValue) }}
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     order.status === 'completed' ? 'bg-green-100 text-green-800' : 'bg-blue-100 text-blue-800']">
                {{ order.status === 'completed' ? '已完成' : '投放中' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ order.createdAt }}</td>
          </tr>
        </tbody>
      </table>
      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
