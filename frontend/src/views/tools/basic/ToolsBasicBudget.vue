<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 28 })

const budgetItems = ref([
  { id: 'BU001', name: '双十一广告组', budget: 50000, spent: 38560, remaining: 11440, progress: 77, status: 'normal' },
  { id: 'BU002', name: '新品推广组', budget: 20000, spent: 18500, remaining: 1500, progress: 92, status: 'warning' },
  { id: 'BU003', name: '品牌曝光组', budget: 30000, spent: 12560, remaining: 17440, progress: 42, status: 'normal' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleBatchAdjust = () => {
  alert('批量调整预算')
}

const handleAdjust = (item: typeof budgetItems.value[0]) => {
  alert(`调整预算: ${item.name}`)
}

const handleDetail = (item: typeof budgetItems.value[0]) => {
  alert(`查看详情: ${item.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '基础工具' }, { name: '预算工具' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">预算工具</h1>
          <p class="mt-2 text-gray-600">管理广告预算分配</p>
        </div>
<button @click="handleBatchAdjust" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          批量调整
        </button>
      </div>
    </div>

    <div class="grid grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总预算</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">¥100,000</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已消耗</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">¥69,620</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">剩余预算</p>
        <p class="text-2xl font-bold text-green-600 mt-1">¥30,380</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">预警数量</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">5</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">预算</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">已消耗</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">剩余</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">消耗进度</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="item in budgetItems" :key="item.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ item.name }}</div>
              <div class="text-xs text-gray-500">{{ item.id }}</div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">¥{{ item.budget.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">¥{{ item.spent.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm" :class="item.remaining < 5000 ? 'text-red-600 font-medium' : 'text-gray-600'">
              ¥{{ item.remaining.toLocaleString() }}
            </td>
            <td class="px-6 py-4">
              <div class="flex items-center">
                <div class="w-24 h-2 bg-gray-200 rounded-full mr-2">
                  <div :class="['h-2 rounded-full',
                         item.progress >= 90 ? 'bg-red-500' : item.progress >= 70 ? 'bg-yellow-500' : 'bg-green-500']"
                       :style="{ width: item.progress + '%' }"></div>
                </div>
                <span class="text-sm text-gray-600">{{ item.progress }}%</span>
              </div>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     item.status === 'warning' ? 'bg-yellow-100 text-yellow-800' : 'bg-green-100 text-green-800']">
                {{ item.status === 'warning' ? '预警' : '正常' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
<button @click="handleAdjust(item)" class="text-blue-600 hover:text-blue-800 mr-3">调整</button>
              <button @click="handleDetail(item)" class="text-gray-600 hover:text-gray-800">详情</button>
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
