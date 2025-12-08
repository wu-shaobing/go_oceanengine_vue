<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 15 })

const contracts = ref([
  { id: 'CT001', name: '2025年度广告投放合同', amount: 500000, startDate: '2025-01-01', endDate: '2025-12-31', status: 'active' },
  { id: 'CT002', name: '2024年度广告投放合同', amount: 350000, startDate: '2024-01-01', endDate: '2024-12-31', status: 'expired' },
  { id: 'CT003', name: 'Q4季度推广合同', amount: 150000, startDate: '2025-10-01', endDate: '2025-12-31', status: 'active' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handleCreateContract = () => {
  alert('新建合同')
}

const handleViewContract = (contract: any) => {
  alert(`查看合同: ${contract.name}`)
}

const handleDownloadContract = (contract: any) => {
  alert(`下载合同: ${contract.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '财务中心' }, { name: '合同管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">合同管理</h1>
          <p class="mt-2 text-gray-600">管理广告投放合同</p>
        </div>
<button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleCreateContract">
          新建合同
        </button>
      </div>
    </div>

    <div class="grid grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">合同总数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">生效中</p>
        <p class="text-2xl font-bold text-green-600 mt-1">8</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">合同金额</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">¥1,000,000</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">即将到期</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">2</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">合同编号</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">合同名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">合同金额</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">有效期</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="contract in contracts" :key="contract.id" class="hover:bg-gray-50">
            <td class="px-6 py-4 text-sm font-medium text-gray-900">{{ contract.id }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ contract.name }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">¥{{ contract.amount.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ contract.startDate }} ~ {{ contract.endDate }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     contract.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                {{ contract.status === 'active' ? '生效中' : '已过期' }}
              </span>
            </td>
<td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3" @click="handleViewContract(contract)">查看</button>
              <button class="text-gray-600 hover:text-gray-800" @click="handleDownloadContract(contract)">下载</button>
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
