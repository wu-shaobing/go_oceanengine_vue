<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '星图', path: '/star' }, { name: '资金管理' }, { name: '资金余额' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">资金管理</h1>
      <p class="text-gray-600 mt-1">查看和管理星图账户资金</p>
    </div>

    <!-- 余额概览 -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-6">
      <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-lg shadow p-6 text-white">
        <div class="text-sm opacity-80">账户余额</div>
        <div class="text-3xl font-bold mt-2">¥{{ balance.available.toLocaleString() }}</div>
        <div class="flex space-x-3 mt-4">
          <button class="px-4 py-1 bg-white text-blue-600 rounded text-sm font-medium">充值</button>
          <button class="px-4 py-1 bg-white bg-opacity-20 text-white rounded text-sm">提现</button>
        </div>
      </div>
      <div class="bg-white rounded-lg shadow p-6">
        <div class="text-sm text-gray-500">冻结金额</div>
        <div class="text-2xl font-bold text-orange-600 mt-2">¥{{ balance.frozen.toLocaleString() }}</div>
        <div class="text-sm text-gray-500 mt-2">进行中订单占用</div>
      </div>
      <div class="bg-white rounded-lg shadow p-6">
        <div class="text-sm text-gray-500">累计消耗</div>
        <div class="text-2xl font-bold text-gray-900 mt-2">¥{{ balance.totalSpent.toLocaleString() }}</div>
        <div class="text-sm text-gray-500 mt-2">本月消耗 ¥{{ balance.monthSpent.toLocaleString() }}</div>
      </div>
    </div>

    <!-- 收支记录 -->
    <div class="bg-white rounded-lg shadow">
      <div class="p-4 border-b flex justify-between items-center">
        <h3 class="text-lg font-medium">收支明细</h3>
        <div class="flex space-x-2">
          <select v-model="filters.type" class="border border-gray-300 rounded px-3 py-1 text-sm">
            <option value="">全部类型</option>
            <option value="recharge">充值</option>
            <option value="consume">消费</option>
            <option value="refund">退款</option>
          </select>
          <input type="date" v-model="filters.startDate" class="border border-gray-300 rounded px-3 py-1 text-sm">
          <input type="date" v-model="filters.endDate" class="border border-gray-300 rounded px-3 py-1 text-sm">
        </div>
      </div>
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">时间</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">类型</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">说明</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">金额</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">余额</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="record in records" :key="record.id">
            <td class="px-4 py-3 text-sm">{{ record.time }}</td>
            <td class="px-4 py-3">
              <span :class="getTypeClass(record.type)" class="px-2 py-1 text-xs rounded-full">
                {{ getTypeText(record.type) }}
              </span>
            </td>
            <td class="px-4 py-3 text-sm">{{ record.desc }}</td>
            <td class="px-4 py-3 text-sm text-right" :class="record.amount > 0 ? 'text-green-600' : 'text-red-600'">
              {{ record.amount > 0 ? '+' : '' }}¥{{ Math.abs(record.amount).toLocaleString() }}
            </td>
            <td class="px-4 py-3 text-sm text-right">¥{{ record.balance.toLocaleString() }}</td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="100" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const balance = ref({
  available: 256800,
  frozen: 45000,
  totalSpent: 1568000,
  monthSpent: 128000
})

const filters = ref({
  type: '',
  startDate: '',
  endDate: ''
})

const records = ref([
  { id: 1, time: '2024-06-15 14:30', type: 'consume', desc: '任务订单支付 - 美妆达人小美', amount: -15000, balance: 256800 },
  { id: 2, time: '2024-06-15 10:00', type: 'recharge', desc: '在线充值', amount: 50000, balance: 271800 },
  { id: 3, time: '2024-06-14 16:45', type: 'consume', desc: '任务订单支付 - 时尚博主Amy', amount: -12000, balance: 221800 },
  { id: 4, time: '2024-06-14 09:20', type: 'refund', desc: '订单退款 - ST20240610001', amount: 8000, balance: 233800 },
  { id: 5, time: '2024-06-13 15:30', type: 'consume', desc: '任务订单支付 - 生活家小王', amount: -8000, balance: 225800 }
])

const getTypeClass = (type: string) => {
  const classes: Record<string, string> = {
    recharge: 'bg-green-100 text-green-800',
    consume: 'bg-blue-100 text-blue-800',
    refund: 'bg-orange-100 text-orange-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}

const getTypeText = (type: string) => {
  const texts: Record<string, string> = {
    recharge: '充值',
    consume: '消费',
    refund: '退款'
  }
  return texts[type] || type
}
</script>
