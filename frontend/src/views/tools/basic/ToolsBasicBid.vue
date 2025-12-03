<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 35 })

const bidSuggestions = ref([
  { id: 'B001', adName: '促销广告A', currentBid: 25.00, suggestBid: 32.00, change: '+28%', status: 'low' },
  { id: 'B002', adName: '推广广告B', currentBid: 45.00, suggestBid: 42.00, change: '-7%', status: 'high' },
  { id: 'B003', adName: '品牌广告C', currentBid: 30.00, suggestBid: 30.00, change: '0%', status: 'optimal' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '基础工具' }, { name: '出价工具' }]" />
      <h1 class="text-3xl font-bold text-gray-900">出价工具</h1>
      <p class="mt-2 text-gray-600">智能出价建议与批量调整</p>
    </div>

    <div class="grid grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">广告数量</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">出价偏低</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">12</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">出价偏高</p>
        <p class="text-2xl font-bold text-red-600 mt-1">8</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">最优出价</p>
        <p class="text-2xl font-bold text-green-600 mt-1">15</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">广告名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">当前出价</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">建议出价</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">调整幅度</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="bid in bidSuggestions" :key="bid.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ bid.adName }}</div>
              <div class="text-xs text-gray-500">{{ bid.id }}</div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">¥{{ bid.currentBid.toFixed(2) }}</td>
            <td class="px-6 py-4 text-sm font-medium text-blue-600">¥{{ bid.suggestBid.toFixed(2) }}</td>
            <td class="px-6 py-4 text-sm" :class="bid.change.startsWith('+') ? 'text-green-600' : bid.change.startsWith('-') ? 'text-red-600' : 'text-gray-600'">
              {{ bid.change }}
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     bid.status === 'low' ? 'bg-yellow-100 text-yellow-800' :
                     bid.status === 'high' ? 'bg-red-100 text-red-800' : 'bg-green-100 text-green-800']">
                {{ bid.status === 'low' ? '偏低' : bid.status === 'high' ? '偏高' : '最优' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button v-if="bid.status !== 'optimal'" class="text-blue-600 hover:text-blue-800 mr-3">应用建议</button>
              <button class="text-gray-600 hover:text-gray-800">详情</button>
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
