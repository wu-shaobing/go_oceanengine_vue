<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '服务市场', path: '/servemarket' }, { name: '功能点管理' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">功能点管理</h1>
      <p class="text-gray-600 mt-1">查看和购买功能点</p>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
      <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-lg shadow p-6 text-white">
        <div class="text-sm opacity-80">功能点余额</div>
        <div class="text-3xl font-bold mt-2">{{ balance.available.toLocaleString() }}</div>
        <button @click="handleRecharge" class="mt-4 px-4 py-1 bg-white text-blue-600 rounded text-sm font-medium">充值</button>
      </div>
      <div class="bg-white rounded-lg shadow p-6">
        <div class="text-sm text-gray-500">本月消耗</div>
        <div class="text-2xl font-bold text-orange-600 mt-2">{{ balance.monthUsed.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-6">
        <div class="text-sm text-gray-500">累计消耗</div>
        <div class="text-2xl font-bold mt-2">{{ balance.totalUsed.toLocaleString() }}</div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">功能点套餐</h3>
        <div class="space-y-3">
          <div v-for="pkg in packages" :key="pkg.id" class="flex items-center justify-between p-4 border rounded-lg hover:border-blue-500">
            <div>
              <div class="font-medium">{{ pkg.name }}</div>
              <div class="text-sm text-gray-500">{{ pkg.points.toLocaleString() }} 功能点</div>
            </div>
            <div class="text-right">
              <div class="text-xl font-bold text-blue-600">¥{{ pkg.price }}</div>
              <button @click="handleBuyPackage(pkg)" class="mt-1 px-3 py-1 text-sm text-blue-600 border border-blue-600 rounded hover:bg-blue-50">购买</button>
            </div>
          </div>
        </div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">消耗记录</h3>
        <table class="min-w-full">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-3 py-2 text-left text-sm font-medium text-gray-500">时间</th>
              <th class="px-3 py-2 text-left text-sm font-medium text-gray-500">服务</th>
              <th class="px-3 py-2 text-right text-sm font-medium text-gray-500">功能点</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="record in records" :key="record.id">
              <td class="px-3 py-2 text-sm text-gray-500">{{ record.time }}</td>
              <td class="px-3 py-2 text-sm">{{ record.service }}</td>
              <td class="px-3 py-2 text-sm text-right text-red-600">-{{ record.points }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const balance = ref({ available: 12500, monthUsed: 3500, totalUsed: 85600 })

const handleRecharge = () => {
  alert('功能点充值')
}

const handleBuyPackage = (pkg: typeof packages.value[0]) => {
  alert(`购买套餐: ${pkg.name} - ¥${pkg.price}`)
}

const packages = ref([
  { id: 1, name: '基础套餐', points: 1000, price: 99 },
  { id: 2, name: '标准套餐', points: 5000, price: 399 },
  { id: 3, name: '专业套餐', points: 10000, price: 699 },
  { id: 4, name: '企业套餐', points: 50000, price: 2999 }
])

const records = ref([
  { id: 1, time: '2024-06-18 15:30', service: '智能文案生成', points: 50 },
  { id: 2, time: '2024-06-18 14:20', service: '素材智能分析', points: 100 },
  { id: 3, time: '2024-06-17 18:45', service: '竞品分析报告', points: 200 },
  { id: 4, time: '2024-06-17 16:30', service: '智能文案生成', points: 50 }
])
</script>
