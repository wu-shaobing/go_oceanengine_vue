<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川', path: '/qianchuan' }, { name: '直播间详情' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div class="flex items-center">
        <img src="https://via.placeholder.com/64" class="w-16 h-16 rounded-full mr-4" alt="">
        <div>
          <h1 class="text-2xl font-bold text-gray-900">{{ room.name }}</h1>
          <p class="text-gray-600">开播时间: {{ room.startTime }}</p>
        </div>
      </div>
      <span :class="room.status === 'live' ? 'bg-red-100 text-red-800' : 'bg-gray-100 text-gray-800'" class="px-3 py-1 rounded-full">
        {{ room.status === 'live' ? '直播中' : '已结束' }}
      </span>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-5 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">消耗</div>
        <div class="text-xl font-bold mt-1">¥{{ room.cost.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">曝光PV</div>
        <div class="text-xl font-bold mt-1">{{ room.pv.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">进入直播间</div>
        <div class="text-xl font-bold mt-1">{{ room.enter.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">成交金额</div>
        <div class="text-xl font-bold mt-1">¥{{ room.gmv.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">ROI</div>
        <div class="text-xl font-bold text-green-600 mt-1">{{ room.roi }}</div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">实时数据趋势</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
          <span class="text-gray-400">实时数据图表</span>
        </div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">转化漏斗</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
          <span class="text-gray-400">漏斗图表</span>
        </div>
      </div>
    </div>

    <div class="bg-white rounded-lg shadow p-4">
      <h3 class="text-lg font-medium mb-4">热销商品</h3>
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">商品</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">点击</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">成交</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">销售额</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="p in products" :key="p.id">
            <td class="px-4 py-3">
              <div class="flex items-center">
                <img :src="p.image" class="w-12 h-12 rounded mr-3" alt="">
                <span class="text-sm">{{ p.name }}</span>
              </div>
            </td>
            <td class="px-4 py-3 text-sm text-right">{{ p.clicks.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ p.orders }}</td>
            <td class="px-4 py-3 text-sm text-right">¥{{ p.gmv.toLocaleString() }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const room = ref({
  name: '618大促专场直播',
  startTime: '2024-06-18 19:00',
  status: 'ended',
  cost: 25680,
  pv: 580000,
  enter: 125000,
  gmv: 128560,
  roi: 5.0
})

const products = ref([
  { id: 1, name: '夏季新款连衣裙', image: 'https://via.placeholder.com/48', clicks: 15600, orders: 286, gmv: 42900 },
  { id: 2, name: '清凉防晒衣', image: 'https://via.placeholder.com/48', clicks: 12800, orders: 215, gmv: 32250 },
  { id: 3, name: '百搭休闲T恤', image: 'https://via.placeholder.com/48', clicks: 9500, orders: 168, gmv: 16800 }
])
</script>
