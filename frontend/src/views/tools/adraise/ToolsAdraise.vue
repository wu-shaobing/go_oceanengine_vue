<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 18 })

const raiseRecords = ref([
  { id: 'AR001', adId: 'AD10001', adName: '品牌推广计划A', originalBid: 35, raisedBid: 45, raisePercent: 28.6, status: 'active', effect: 'good', createdAt: '2025-11-25' },
  { id: 'AR002', adId: 'AD10002', adName: '双11促销计划', originalBid: 50, raisedBid: 65, raisePercent: 30, status: 'active', effect: 'excellent', createdAt: '2025-11-24' },
  { id: 'AR003', adId: 'AD10003', adName: '新品上市推广', originalBid: 40, raisedBid: 52, raisePercent: 30, status: 'paused', effect: 'normal', createdAt: '2025-11-23' }
])

const getEffectConfig = (effect: string) => {
  switch (effect) {
    case 'excellent': return { label: '效果显著', class: 'bg-green-100 text-green-800' }
    case 'good': return { label: '效果良好', class: 'bg-blue-100 text-blue-800' }
    default: return { label: '效果一般', class: 'bg-gray-100 text-gray-800' }
  }
}

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '提价工具' }, { name: '智能提价' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">智能提价</h1>
          <p class="mt-2 text-gray-600">智能优化出价，提升广告竞争力</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建提价策略
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">提价计划</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">生效中</p>
        <p class="text-2xl font-bold text-green-600 mt-1">12</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均提价幅度</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">25.8%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">转化提升</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">+32%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">广告计划</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">原出价</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">提价后</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">提价幅度</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">效果评估</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="record in raiseRecords" :key="record.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ record.adName }}</div>
              <div class="text-xs text-gray-500">{{ record.adId }}</div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">¥{{ record.originalBid }}</td>
            <td class="px-6 py-4 text-sm font-medium text-green-600">¥{{ record.raisedBid }}</td>
            <td class="px-6 py-4">
              <span class="px-2 py-1 bg-orange-100 text-orange-700 rounded text-xs">+{{ record.raisePercent }}%</span>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded text-xs', getEffectConfig(record.effect).class]">
                {{ getEffectConfig(record.effect).label }}
              </span>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                record.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                {{ record.status === 'active' ? '生效中' : '已暂停' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">调整</button>
              <button class="text-gray-600 hover:text-gray-800">取消</button>
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
