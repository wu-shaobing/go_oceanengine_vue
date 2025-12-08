<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '资金管理' }, { name: '钱包' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">资金钱包</h1>
      <p class="text-gray-600 mt-1">账户资金管理</p>
    </div>

    <!-- 余额卡片 -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-6">
      <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-lg shadow-lg p-6 text-white">
        <div class="text-sm opacity-80">账户余额</div>
        <div class="text-3xl font-bold mt-2">¥{{ walletInfo.balance.toLocaleString() }}</div>
        <div class="mt-4 flex space-x-3">
<button @click="handleRecharge" class="px-4 py-1.5 bg-white text-blue-600 rounded text-sm font-medium">充值</button>
          <button @click="handleWithdraw" class="px-4 py-1.5 bg-blue-400 rounded text-sm">提现</button>
        </div>
      </div>
      <div class="bg-white rounded-lg shadow p-6">
        <div class="text-sm text-gray-500">可用余额</div>
        <div class="text-2xl font-bold text-gray-900 mt-2">¥{{ walletInfo.available.toLocaleString() }}</div>
        <div class="text-sm text-gray-500 mt-2">可用于投放消耗</div>
      </div>
      <div class="bg-white rounded-lg shadow p-6">
        <div class="text-sm text-gray-500">冻结金额</div>
        <div class="text-2xl font-bold text-orange-600 mt-2">¥{{ walletInfo.frozen.toLocaleString() }}</div>
        <div class="text-sm text-gray-500 mt-2">投放中冻结</div>
      </div>
    </div>

    <!-- 近期消耗 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <h3 class="text-lg font-medium mb-4">近7天消耗趋势</h3>
      <div class="h-48 flex items-center justify-center bg-gray-50 rounded">
        <span class="text-gray-400">消耗趋势图表</span>
      </div>
    </div>

    <!-- 交易记录 -->
    <div class="bg-white rounded-lg shadow">
      <div class="p-4 border-b flex justify-between items-center">
        <h3 class="text-lg font-medium">交易记录</h3>
        <router-link to="/qianchuan/finance/detail" class="text-blue-600 hover:text-blue-800 text-sm">查看全部</router-link>
      </div>
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">交易时间</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">交易类型</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">交易说明</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">金额</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">余额</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="record in recentRecords" :key="record.id">
            <td class="px-4 py-3 text-sm">{{ record.time }}</td>
            <td class="px-4 py-3 text-sm">{{ record.type }}</td>
            <td class="px-4 py-3 text-sm">{{ record.description }}</td>
            <td class="px-4 py-3 text-sm text-right" :class="record.amount > 0 ? 'text-green-600' : 'text-red-600'">
              {{ record.amount > 0 ? '+' : '' }}¥{{ Math.abs(record.amount).toLocaleString() }}
            </td>
            <td class="px-4 py-3 text-sm text-right">¥{{ record.balance.toLocaleString() }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const walletInfo = ref({
  balance: 156800,
  available: 148500,
  frozen: 8300
})

const recentRecords = ref([
  { id: 1, time: '2024-03-15 18:30', type: '广告消耗', description: '千川广告投放消耗', amount: -2860, balance: 156800 },
  { id: 2, time: '2024-03-15 14:20', type: '广告消耗', description: '千川广告投放消耗', amount: -1580, balance: 159660 },
  { id: 3, time: '2024-03-15 10:00', type: '充值', description: '支付宝充值', amount: 50000, balance: 161240 },
  { id: 4, time: '2024-03-14 22:30', type: '广告消耗', description: '千川广告投放消耗', amount: -3260, balance: 111240 },
  { id: 5, time: '2024-03-14 16:45', type: '广告消耗', description: '千川广告投放消耗', amount: -2180, balance: 114500 }
])

const handleRecharge = () => {
  alert('账户充值')
}

const handleWithdraw = () => {
  alert('账户提现')
}
</script>
