<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川', path: '/qianchuan' }, { name: '数据报表' }, { name: '全域推广报表' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">全域推广报表</h1>
      <p class="text-gray-600 mt-1">查看全域推广投放效果数据</p>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <input type="date" v-model="filters.startDate" class="border border-gray-300 rounded px-3 py-2">
        <span>至</span>
        <input type="date" v-model="filters.endDate" class="border border-gray-300 rounded px-3 py-2">
        <select v-model="filters.type" class="border border-gray-300 rounded px-3 py-2">
          <option value="">全部类型</option>
          <option value="live">直播全域</option>
          <option value="video">短视频全域</option>
        </select>
<button class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700" @click="handleQuery">查询</button>
        <button class="px-4 py-2 border border-gray-300 rounded hover:bg-gray-50" @click="handleExport">导出</button>
      </div>
    </div>

    <!-- 汇总 -->
    <div class="grid grid-cols-1 md:grid-cols-5 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">总消耗</div>
        <div class="text-xl font-bold mt-1">¥{{ summary.cost.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">总GMV</div>
        <div class="text-xl font-bold mt-1">¥{{ summary.gmv.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">成交订单</div>
        <div class="text-xl font-bold mt-1">{{ summary.orders.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">平均客单价</div>
        <div class="text-xl font-bold mt-1">¥{{ summary.avgPrice }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">整体ROI</div>
        <div class="text-xl font-bold text-green-600 mt-1">{{ summary.roi }}</div>
      </div>
    </div>

    <!-- 数据表格 -->
    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">全域推广</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">类型</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">消耗</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">GMV</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">订单数</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">客单价</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">ROI</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="item in reportData" :key="item.id">
            <td class="px-4 py-3">
              <div class="font-medium">{{ item.name }}</div>
              <div class="text-sm text-gray-500">ID: {{ item.id }}</div>
            </td>
            <td class="px-4 py-3 text-sm">
              <span :class="item.type === 'live' ? 'bg-red-100 text-red-800' : 'bg-blue-100 text-blue-800'" class="px-2 py-1 text-xs rounded">
                {{ item.type === 'live' ? '直播全域' : '短视频全域' }}
              </span>
            </td>
            <td class="px-4 py-3 text-sm text-right">¥{{ item.cost.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">¥{{ item.gmv.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ item.orders }}</td>
            <td class="px-4 py-3 text-sm text-right">¥{{ item.avgPrice }}</td>
            <td class="px-4 py-3 text-sm text-right text-green-600 font-medium">{{ item.roi }}</td>
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

const filters = ref({ startDate: '', endDate: '', type: '' })

const handleQuery = () => {
  alert('查询全域推广报表')
}

const handleExport = () => {
  alert('导出全域推广报表')
}

const summary = ref({
  cost: 256800,
  gmv: 1285600,
  orders: 8560,
  avgPrice: 150.2,
  roi: 5.0
})

const reportData = ref([
  { id: 'U001', name: '618大促直播全域', type: 'live', cost: 85600, gmv: 456800, orders: 2856, avgPrice: 160.0, roi: 5.33 },
  { id: 'U002', name: '夏日新品短视频', type: 'video', cost: 65800, gmv: 328000, orders: 2186, avgPrice: 150.0, roi: 4.98 },
  { id: 'U003', name: '品牌日直播全域', type: 'live', cost: 52600, gmv: 268500, orders: 1856, avgPrice: 144.7, roi: 5.1 },
  { id: 'U004', name: '爆款商品推广', type: 'video', cost: 32800, gmv: 156800, orders: 1086, avgPrice: 144.4, roi: 4.78 },
  { id: 'U005', name: '日常带货全域', type: 'live', cost: 20000, gmv: 75500, orders: 576, avgPrice: 131.1, roi: 3.78 }
])
</script>
