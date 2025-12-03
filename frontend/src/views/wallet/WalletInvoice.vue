<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 35 })

const invoices = ref([
  { id: 'INV001', month: '2025年11月', amount: 125600, type: '增值税专用发票', status: 'issued', applyTime: '2025-11-25' },
  { id: 'INV002', month: '2025年10月', amount: 98560, type: '增值税专用发票', status: 'issued', applyTime: '2025-10-28' },
  { id: 'INV003', month: '2025年9月', amount: 76890, type: '增值税普通发票', status: 'pending', applyTime: '2025-11-20' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '财务中心' }, { name: '发票管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">发票管理</h1>
          <p class="mt-2 text-gray-600">申请和管理发票</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          申请发票
        </button>
      </div>
    </div>

    <div class="grid grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">可开票金额</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">¥125,600</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已开票金额</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">¥301,050</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">处理中</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">1</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">已完成</p>
        <p class="text-2xl font-bold text-green-600 mt-1">34</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">发票编号</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">账期</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">金额</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">发票类型</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">申请时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="invoice in invoices" :key="invoice.id" class="hover:bg-gray-50">
            <td class="px-6 py-4 text-sm font-medium text-gray-900">{{ invoice.id }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ invoice.month }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">¥{{ invoice.amount.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ invoice.type }}</td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ invoice.applyTime }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     invoice.status === 'issued' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800']">
                {{ invoice.status === 'issued' ? '已开具' : '处理中' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button v-if="invoice.status === 'issued'" class="text-blue-600 hover:text-blue-800 mr-3">下载</button>
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
