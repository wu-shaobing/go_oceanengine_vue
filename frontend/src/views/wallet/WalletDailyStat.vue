<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { Line, Doughnut } from 'vue-chartjs'
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  ArcElement,
  Title,
  Tooltip,
  Legend,
  Filler
} from 'chart.js'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
// StatsCard - using inline display instead

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, ArcElement, Title, Tooltip, Legend, Filler)

const stats = reactive({
  totalBalance: 5680000,
  todayIncome: 128450,
  todayExpense: 82340,
  netInflow: 46110
})

const dailyData = ref([
  { date: '2025-11-11', income: 12845, expense: 8234, net: 4611, count: 156 },
  { date: '2025-11-10', income: 11420, expense: 7820, net: 3600, count: 133 },
  { date: '2025-11-09', income: 10890, expense: 9100, net: 1790, count: 145 },
  { date: '2025-11-08', income: 13200, expense: 8500, net: 4700, count: 168 },
  { date: '2025-11-07', income: 10500, expense: 7800, net: 2700, count: 128 }
])

const trendChartData = ref({
  labels: ['11/05', '11/06', '11/07', '11/08', '11/09', '11/10', '11/11'],
  datasets: [
    {
      label: '收入',
      data: [9800, 11200, 10500, 13200, 10890, 11420, 12845],
      borderColor: '#10b981',
      backgroundColor: 'rgba(16, 185, 129, 0.1)',
      fill: true,
      tension: 0.4
    },
    {
      label: '支出',
      data: [7500, 8100, 7800, 8500, 9100, 7820, 8234],
      borderColor: '#ef4444',
      backgroundColor: 'rgba(239, 68, 68, 0.1)',
      fill: true,
      tension: 0.4
    }
  ]
})

const pieChartData = ref({
  labels: ['广告消耗', '服务费', '转账', '其他'],
  datasets: [{
    data: [65, 15, 12, 8],
    backgroundColor: ['#3b82f6', '#10b981', '#f59e0b', '#6b7280']
  }]
})

const chartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: { legend: { position: 'top' as const } },
  scales: { y: { beginAtZero: true } }
}

const pieOptions = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: { legend: { position: 'right' as const } }
}

onMounted(() => {
  // Load data
})
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '财务管理' }, { name: '共享钱包日流水' }]" />
      <h1 class="text-3xl font-bold text-gray-900">共享钱包日流水</h1>
      <p class="mt-2 text-gray-600">查看共享钱包每日收支统计</p>
    </div>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-600">钱包余额</p>
            <p class="text-2xl font-bold text-gray-900 mt-1">¥{{ (stats.totalBalance / 10000).toFixed(0) }}万</p>
          </div>
          <div class="p-3 rounded-lg bg-blue-50">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"/>
            </svg>
          </div>
        </div>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-600">今日收入</p>
            <p class="text-2xl font-bold text-green-600 mt-1">+¥{{ stats.todayIncome.toLocaleString() }}</p>
          </div>
          <div class="p-3 rounded-lg bg-green-50">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11l5-5m0 0l5 5m-5-5v12"/>
            </svg>
          </div>
        </div>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-600">今日支出</p>
            <p class="text-2xl font-bold text-red-600 mt-1">-¥{{ stats.todayExpense.toLocaleString() }}</p>
          </div>
          <div class="p-3 rounded-lg bg-red-50">
            <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 13l-5 5m0 0l-5-5m5 5V6"/>
            </svg>
          </div>
        </div>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-600">净流入</p>
            <p class="text-2xl font-bold text-purple-600 mt-1">¥{{ stats.netInflow.toLocaleString() }}</p>
          </div>
          <div class="p-3 rounded-lg bg-purple-50">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- Charts -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">收支趋势</h3>
        <div style="height: 300px;">
          <Line :data="trendChartData" :options="chartOptions" />
        </div>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">支出分布</h3>
        <div style="height: 300px;">
          <Doughnut :data="pieChartData" :options="pieOptions" />
        </div>
      </div>
    </div>

    <!-- Daily Table -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">日流水明细</h3>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">日期</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">收入</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">支出</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">净流入</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">交易笔数</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="item in dailyData" :key="item.date" class="hover:bg-gray-50">
              <td class="px-6 py-4 text-sm text-gray-900">{{ item.date }}</td>
              <td class="px-6 py-4 text-sm text-green-600">+¥{{ item.income.toLocaleString() }}</td>
              <td class="px-6 py-4 text-sm text-red-600">-¥{{ item.expense.toLocaleString() }}</td>
              <td class="px-6 py-4 text-sm font-medium text-gray-900">¥{{ item.net.toLocaleString() }}</td>
              <td class="px-6 py-4 text-sm text-gray-500">{{ item.count }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>
