<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川', path: '/qianchuan' }, { name: '财务管理' }, { name: '流水明细' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">流水明细</h1>
      <p class="text-gray-600 mt-1">查看账户资金流水记录</p>
    </div>

    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <input type="date" v-model="filters.startDate" class="border border-gray-300 rounded px-3 py-2">
        <span>至</span>
        <input type="date" v-model="filters.endDate" class="border border-gray-300 rounded px-3 py-2">
        <select v-model="filters.type" class="border border-gray-300 rounded px-3 py-2">
          <option value="">全部类型</option>
          <option value="consume">消耗</option>
          <option value="recharge">充值</option>
          <option value="transfer">转账</option>
        </select>
<button @click="handleQuery" class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">查询</button>
        <button @click="handleExport" class="px-4 py-2 border border-gray-300 rounded hover:bg-gray-50">导出</button>
      </div>
    </div>

    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">流水号</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">时间</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">类型</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">说明</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">金额</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">余额</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="item in records" :key="item.id">
            <td class="px-4 py-3 text-sm font-mono">{{ item.id }}</td>
            <td class="px-4 py-3 text-sm text-gray-500">{{ item.time }}</td>
            <td class="px-4 py-3">
              <span :class="getTypeClass(item.type)" class="px-2 py-1 text-xs rounded">{{ getTypeText(item.type) }}</span>
            </td>
            <td class="px-4 py-3 text-sm">{{ item.desc }}</td>
            <td class="px-4 py-3 text-sm text-right" :class="item.amount > 0 ? 'text-green-600' : 'text-red-600'">
              {{ item.amount > 0 ? '+' : '' }}¥{{ Math.abs(item.amount).toLocaleString() }}
            </td>
            <td class="px-4 py-3 text-sm text-right">¥{{ item.balance.toLocaleString() }}</td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="200" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const filters = ref({ startDate: '', endDate: '', type: '' })

const records = ref([
  { id: 'FD20240618001', time: '2024-06-18 15:30', type: 'consume', desc: '广告投放消耗', amount: -2580, balance: 125680 },
  { id: 'FD20240618002', time: '2024-06-18 10:00', type: 'recharge', desc: '在线充值', amount: 50000, balance: 128260 },
  { id: 'FD20240617001', time: '2024-06-17 18:00', type: 'consume', desc: '广告投放消耗', amount: -3250, balance: 78260 },
  { id: 'FD20240617002', time: '2024-06-17 14:30', type: 'transfer', desc: '账户间转账', amount: 10000, balance: 81510 },
  { id: 'FD20240616001', time: '2024-06-16 20:00', type: 'consume', desc: '广告投放消耗', amount: -1860, balance: 71510 }
])

const getTypeClass = (type: string) => {
  const classes: Record<string, string> = { consume: 'bg-red-100 text-red-800', recharge: 'bg-green-100 text-green-800', transfer: 'bg-blue-100 text-blue-800' }
  return classes[type] || 'bg-gray-100 text-gray-800'
}

const getTypeText = (type: string) => {
  const texts: Record<string, string> = { consume: '消耗', recharge: '充值', transfer: '转账' }
  return texts[type] || type
}

const handleQuery = () => {
  alert('查询流水明细')
}

const handleExport = () => {
  alert('导出流水明细')
}
</script>
