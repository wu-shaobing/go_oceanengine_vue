<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川', path: '/qianchuan' }, { name: '数据报表' }, { name: '直播间报表' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">直播间报表</h1>
      <p class="text-gray-600 mt-1">查看直播间维度的投放效果数据</p>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <input type="date" v-model="filters.startDate" class="border border-gray-300 rounded px-3 py-2">
        <span>至</span>
        <input type="date" v-model="filters.endDate" class="border border-gray-300 rounded px-3 py-2">
        <button class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">查询</button>
        <button class="px-4 py-2 border border-gray-300 rounded hover:bg-gray-50">导出</button>
      </div>
    </div>

    <!-- 直播间列表 -->
    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">直播间</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">消耗</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">曝光PV</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">进入直播间</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">商品点击</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">成交订单</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">成交金额</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">ROI</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="room in rooms" :key="room.id">
            <td class="px-4 py-3">
              <div class="flex items-center">
                <img :src="room.cover" class="w-12 h-12 rounded mr-3" alt="">
                <div>
                  <div class="font-medium">{{ room.name }}</div>
                  <div class="text-sm text-gray-500">{{ room.date }}</div>
                </div>
              </div>
            </td>
            <td class="px-4 py-3 text-sm text-right">¥{{ room.cost.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ room.pv.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ room.enter.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ room.productClick.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ room.orders }}</td>
            <td class="px-4 py-3 text-sm text-right">¥{{ room.gmv.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right text-green-600 font-medium">{{ room.roi }}</td>
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

const filters = ref({ startDate: '', endDate: '' })

const rooms = ref([
  { id: 1, name: '618大促专场直播', cover: 'https://via.placeholder.com/48', date: '2024-06-18 19:00', cost: 25680, pv: 580000, enter: 125000, productClick: 38500, orders: 856, gmv: 128560, roi: 5.0 },
  { id: 2, name: '夏日新品首发', cover: 'https://via.placeholder.com/48', date: '2024-06-17 20:00', cost: 18560, pv: 420000, enter: 95000, productClick: 28600, orders: 625, gmv: 93750, roi: 5.05 },
  { id: 3, name: '爆款秒杀专场', cover: 'https://via.placeholder.com/48', date: '2024-06-16 19:30', cost: 15230, pv: 356000, enter: 78000, productClick: 22500, orders: 486, gmv: 68040, roi: 4.47 },
  { id: 4, name: '品牌日直播', cover: 'https://via.placeholder.com/48', date: '2024-06-15 20:00', cost: 12860, pv: 298000, enter: 65000, productClick: 18900, orders: 385, gmv: 57750, roi: 4.49 },
  { id: 5, name: '日常带货直播', cover: 'https://via.placeholder.com/48', date: '2024-06-14 19:00', cost: 10580, pv: 245000, enter: 52000, productClick: 15600, orders: 298, gmv: 41720, roi: 3.94 }
])
</script>
