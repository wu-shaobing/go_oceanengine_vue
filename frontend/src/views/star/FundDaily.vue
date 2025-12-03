<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '星图', path: '/star' }, { name: '资金日流水' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">资金日流水</h1>
      <p class="text-gray-600 mt-1">查看每日资金变动情况</p>
    </div>

    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <input type="date" v-model="filters.startDate" class="border border-gray-300 rounded px-3 py-2">
        <span>至</span>
        <input type="date" v-model="filters.endDate" class="border border-gray-300 rounded px-3 py-2">
        <button class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">查询</button>
        <button class="px-4 py-2 border border-gray-300 rounded hover:bg-gray-50">导出</button>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">期间收入</div>
        <div class="text-xl font-bold text-green-600 mt-1">+¥{{ summary.income.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">期间支出</div>
        <div class="text-xl font-bold text-red-600 mt-1">-¥{{ summary.expense.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">期初余额</div>
        <div class="text-xl font-bold mt-1">¥{{ summary.startBalance.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">期末余额</div>
        <div class="text-xl font-bold mt-1">¥{{ summary.endBalance.toLocaleString() }}</div>
      </div>
    </div>

    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">日期</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">收入</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">支出</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">余额</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="item in dailyData" :key="item.date">
            <td class="px-4 py-3 font-medium">{{ item.date }}</td>
            <td class="px-4 py-3 text-right text-green-600">+¥{{ item.income.toLocaleString() }}</td>
            <td class="px-4 py-3 text-right text-red-600">-¥{{ item.expense.toLocaleString() }}</td>
            <td class="px-4 py-3 text-right font-medium">¥{{ item.balance.toLocaleString() }}</td>
            <td class="px-4 py-3">
              <router-link to="/star/fund/transaction" class="text-blue-600 hover:text-blue-800 text-sm">明细</router-link>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="30" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const filters = ref({ startDate: '', endDate: '' })

const summary = ref({ income: 156800, expense: 98500, startBalance: 200000, endBalance: 258300 })

const dailyData = ref([
  { date: '2024-06-18', income: 25000, expense: 18500, balance: 258300 },
  { date: '2024-06-17', income: 18000, expense: 12000, balance: 251800 },
  { date: '2024-06-16', income: 32000, expense: 25000, balance: 245800 },
  { date: '2024-06-15', income: 15000, expense: 8000, balance: 238800 },
  { date: '2024-06-14', income: 28000, expense: 15000, balance: 231800 }
])
</script>
