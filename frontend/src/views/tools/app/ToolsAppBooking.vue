<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 15 })

const bookings = ref([
  { id: 'BK001', appName: '新游戏A', bookingCount: 12560, targetCount: 50000, progress: 25, endDate: '2025-12-15', status: 'active' },
  { id: 'BK002', appName: '电商APP', bookingCount: 38900, targetCount: 50000, progress: 78, endDate: '2025-12-01', status: 'active' },
  { id: 'BK003', appName: '工具应用B', bookingCount: 50000, targetCount: 50000, progress: 100, endDate: '2025-11-25', status: 'completed' },
  { id: 'BK004', appName: '社交APP', bookingCount: 8900, targetCount: 30000, progress: 30, endDate: '2025-12-20', status: 'active' }
])

const formatNumber = (num: number) => {
  if (num >= 10000) return (num / 10000).toFixed(1) + '万'
  return num.toLocaleString()
}

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '应用工具' }, { name: '应用预约' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">应用预约</h1>
          <p class="mt-2 text-gray-600">管理应用预约推广活动</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建预约活动
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">活动数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">进行中</p>
        <p class="text-2xl font-bold text-green-600 mt-1">8</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总预约数</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">11.0万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">转化率</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">68%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">应用名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">预约人数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">目标人数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">进度</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">截止日期</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="booking in bookings" :key="booking.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ booking.appName }}</div>
              <div class="text-xs text-gray-500">{{ booking.id }}</div>
            </td>
            <td class="px-6 py-4 text-sm font-medium text-blue-600">{{ formatNumber(booking.bookingCount) }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ formatNumber(booking.targetCount) }}</td>
            <td class="px-6 py-4">
              <div class="flex items-center gap-2">
                <div class="w-24 h-2 bg-gray-200 rounded-full">
                  <div :class="['h-full rounded-full', booking.progress >= 100 ? 'bg-green-500' : 'bg-blue-500']"
                       :style="{ width: Math.min(booking.progress, 100) + '%' }"></div>
                </div>
                <span class="text-sm text-gray-600">{{ booking.progress }}%</span>
              </div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ booking.endDate }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     booking.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-blue-100 text-blue-800']">
                {{ booking.status === 'active' ? '进行中' : '已完成' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">详情</button>
              <button class="text-gray-600 hover:text-gray-800">推广</button>
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
