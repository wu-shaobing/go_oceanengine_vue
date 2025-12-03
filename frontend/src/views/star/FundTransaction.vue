<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '星图', path: '/star' }, { name: '交易明细' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">交易明细</h1>
      <p class="text-gray-600 mt-1">查看详细交易记录</p>
    </div>

    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <input type="date" v-model="filters.startDate" class="border border-gray-300 rounded px-3 py-2">
        <span>至</span>
        <input type="date" v-model="filters.endDate" class="border border-gray-300 rounded px-3 py-2">
        <select v-model="filters.type" class="border border-gray-300 rounded px-3 py-2">
          <option value="">全部类型</option>
          <option value="income">收入</option>
          <option value="expense">支出</option>
        </select>
        <button class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">查询</button>
        <button class="px-4 py-2 border border-gray-300 rounded hover:bg-gray-50">导出</button>
      </div>
    </div>

    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">交易时间</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">交易类型</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">描述</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">金额</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">余额</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="item in transactions" :key="item.id">
            <td class="px-4 py-3 text-sm text-gray-500">{{ item.time }}</td>
            <td class="px-4 py-3">
              <span :class="item.type === 'income' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" class="px-2 py-1 text-xs rounded">
                {{ item.type === 'income' ? '收入' : '支出' }}
              </span>
            </td>
            <td class="px-4 py-3">
              <div class="font-medium text-sm">{{ item.desc }}</div>
              <div class="text-xs text-gray-500">{{ item.remark }}</div>
            </td>
            <td class="px-4 py-3 text-right" :class="item.type === 'income' ? 'text-green-600' : 'text-red-600'">
              {{ item.type === 'income' ? '+' : '-' }}¥{{ item.amount.toLocaleString() }}
            </td>
            <td class="px-4 py-3 text-right font-medium">¥{{ item.balance.toLocaleString() }}</td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="100" :page-size="20" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const filters = ref({ startDate: '', endDate: '', type: '' })

const transactions = ref([
  { id: 1, time: '2024-06-18 15:30:25', type: 'income', desc: '任务结算', remark: '618美妆种草任务-达人小美', amount: 25000, balance: 258300 },
  { id: 2, time: '2024-06-18 14:20:10', type: 'expense', desc: '达人佣金', remark: '视频推广任务-时尚博主Amy', amount: 18500, balance: 233300 },
  { id: 3, time: '2024-06-18 10:15:30', type: 'income', desc: '充值', remark: '账户充值', amount: 50000, balance: 251800 },
  { id: 4, time: '2024-06-17 18:45:00', type: 'expense', desc: '平台服务费', remark: '6月服务费', amount: 12000, balance: 201800 },
  { id: 5, time: '2024-06-17 16:30:15', type: 'income', desc: '任务结算', remark: '护肤测评任务-Lisa', amount: 18000, balance: 213800 }
])
</script>
