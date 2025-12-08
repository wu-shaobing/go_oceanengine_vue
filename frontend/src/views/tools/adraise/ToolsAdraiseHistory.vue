<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 56 })

const histories = ref([
  { id: 'H001', adId: 'AD001', adName: '双十一促销广告', startTime: '2025-11-20', endTime: '2025-11-27', target: 10000, result: 12560, status: 'success' },
  { id: 'H002', adId: 'AD002', adName: '新品推广广告', startTime: '2025-11-15', endTime: '2025-11-22', target: 5000, result: 4850, status: 'partial' },
  { id: 'H003', adId: 'AD003', adName: '品牌曝光广告', startTime: '2025-11-10', endTime: '2025-11-17', target: 8000, result: 2560, status: 'failed' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleDetail = (history: typeof histories.value[0]) => {
  alert(`查看详情: ${history.adName}`)
}

const handleReuse = (history: typeof histories.value[0]) => {
  alert(`复用策略: ${history.adName}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '起量工具' }, { name: '起量历史' }]" />
      <h1 class="text-3xl font-bold text-gray-900">起量历史</h1>
      <p class="mt-2 text-gray-600">查看历史起量记录</p>
    </div>

    <div class="grid grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">历史记录</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">成功起量</p>
        <p class="text-2xl font-bold text-green-600 mt-1">38</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">部分成功</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">12</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">未达标</p>
        <p class="text-2xl font-bold text-red-600 mt-1">6</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">广告信息</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">起量周期</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">目标消耗</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">实际消耗</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">完成率</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="history in histories" :key="history.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ history.adName }}</div>
              <div class="text-xs text-gray-500">{{ history.adId }}</div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-500">
              {{ history.startTime }} ~ {{ history.endTime }}
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">¥{{ history.target.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm text-gray-900 font-medium">¥{{ history.result.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm">
              <span :class="history.result >= history.target ? 'text-green-600' : 'text-yellow-600'">
                {{ Math.round(history.result / history.target * 100) }}%
              </span>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     history.status === 'success' ? 'bg-green-100 text-green-800' :
                     history.status === 'partial' ? 'bg-yellow-100 text-yellow-800' : 'bg-red-100 text-red-800']">
                {{ history.status === 'success' ? '成功' : history.status === 'partial' ? '部分成功' : '未达标' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
<button @click="handleDetail(history)" class="text-blue-600 hover:text-blue-800 mr-3">详情</button>
              <button @click="handleReuse(history)" class="text-gray-600 hover:text-gray-800">复用</button>
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
