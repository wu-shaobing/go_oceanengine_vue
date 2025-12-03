<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '服务市场', path: '/servemarket' }, { name: '订单管理' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">订单管理</h1>
      <p class="text-gray-600 mt-1">管理服务市场订单</p>
    </div>

    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <input type="text" v-model="filters.keyword" placeholder="搜索订单" class="border border-gray-300 rounded px-3 py-2 w-48">
        <select v-model="filters.status" class="border border-gray-300 rounded px-3 py-2">
          <option value="">全部状态</option>
          <option value="pending">待处理</option>
          <option value="processing">进行中</option>
          <option value="completed">已完成</option>
        </select>
        <button class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">查询</button>
      </div>
    </div>

    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">订单信息</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">服务类型</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">金额</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">状态</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">创建时间</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="order in orders" :key="order.id">
            <td class="px-4 py-3">
              <div class="font-medium">{{ order.name }}</div>
              <div class="text-sm text-gray-500">{{ order.id }}</div>
            </td>
            <td class="px-4 py-3 text-sm">{{ order.type }}</td>
            <td class="px-4 py-3 text-sm text-right">¥{{ order.amount.toLocaleString() }}</td>
            <td class="px-4 py-3">
              <span :class="getStatusClass(order.status)" class="px-2 py-1 text-xs rounded">{{ getStatusText(order.status) }}</span>
            </td>
            <td class="px-4 py-3 text-sm text-gray-500">{{ order.createTime }}</td>
            <td class="px-4 py-3">
              <div class="flex space-x-2">
                <button class="text-blue-600 hover:text-blue-800 text-sm">详情</button>
                <button v-if="order.status === 'pending'" class="text-red-600 hover:text-red-800 text-sm">取消</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="50" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const filters = ref({ keyword: '', status: '' })

const orders = ref([
  { id: 'SM202406180001', name: '视频制作服务-618专题', type: '视频制作', amount: 2980, status: 'processing', createTime: '2024-06-18 10:30' },
  { id: 'SM202406150002', name: '账户诊断服务', type: '咨询服务', amount: 599, status: 'completed', createTime: '2024-06-15 14:20' },
  { id: 'SM202406100003', name: '代运营服务-月度', type: '代运营', amount: 9999, status: 'processing', createTime: '2024-06-10 09:00' },
  { id: 'SM202406080004', name: '智能投放助手-年度', type: '工具订阅', amount: 2988, status: 'completed', createTime: '2024-06-08 16:45' }
])

const getStatusClass = (status: string) => {
  const classes: Record<string, string> = { pending: 'bg-orange-100 text-orange-800', processing: 'bg-blue-100 text-blue-800', completed: 'bg-green-100 text-green-800' }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts: Record<string, string> = { pending: '待处理', processing: '进行中', completed: '已完成' }
  return texts[status] || status
}
</script>
