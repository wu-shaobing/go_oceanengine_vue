<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 45 })

const orders = ref([
  { id: 'DOU001', videoTitle: '产品测评视频', budget: 500, target: 'views', targetValue: 50000, actualValue: 52300, status: 'completed', createdAt: '2025-11-25' },
  { id: 'DOU002', videoTitle: '品牌故事分享', budget: 1000, target: 'followers', targetValue: 500, actualValue: 380, status: 'running', createdAt: '2025-11-26' },
  { id: 'DOU003', videoTitle: '限时优惠活动', budget: 2000, target: 'clicks', targetValue: 3000, actualValue: 0, status: 'pending', createdAt: '2025-11-27' }
])

const getTargetLabel = (target: string) => {
  switch (target) {
    case 'views': return '播放量'
    case 'followers': return '涨粉'
    case 'clicks': return '点击'
    default: return target
  }
}

const getStatusConfig = (status: string) => {
  switch (status) {
    case 'completed': return { label: '已完成', class: 'bg-green-100 text-green-800' }
    case 'running': return { label: '投放中', class: 'bg-blue-100 text-blue-800' }
    default: return { label: '待投放', class: 'bg-yellow-100 text-yellow-800' }
  }
}

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '其他功能' }, { name: 'Dou+管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">Dou+投放管理</h1>
          <p class="mt-2 text-gray-600">管理抖音Dou+投放订单</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建投放
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总订单</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">投放中</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">8</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总消耗</p>
        <p class="text-2xl font-bold text-green-600 mt-1">¥12.5万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总播放</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">258万</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">视频</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">预算</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">投放目标</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">目标值</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">实际值</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="order in orders" :key="order.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ order.videoTitle }}</div>
              <div class="text-xs text-gray-500">{{ order.id }}</div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">¥{{ order.budget }}</td>
            <td class="px-6 py-4">
              <span class="px-2 py-1 bg-blue-100 text-blue-700 rounded text-xs">{{ getTargetLabel(order.target) }}</span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ order.targetValue.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm font-medium" :class="order.actualValue >= order.targetValue ? 'text-green-600' : 'text-gray-900'">
              {{ order.actualValue.toLocaleString() }}
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium', getStatusConfig(order.status).class]">
                {{ getStatusConfig(order.status).label }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">详情</button>
              <button v-if="order.status === 'running'" class="text-red-600 hover:text-red-800">停止</button>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
