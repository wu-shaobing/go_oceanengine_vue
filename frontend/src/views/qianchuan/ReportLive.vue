<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '数据报表' }, { name: '直播间报表' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">直播间报表</h1>
        <p class="text-gray-600 mt-1">直播投放数据分析</p>
      </div>
      <button class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">
        导出报表
      </button>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <select v-model="filters.liveRoom" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部直播间</option>
          <option value="L001">主播A直播间</option>
          <option value="L002">品牌直播间</option>
        </select>
        <div class="flex items-center gap-2">
          <input type="date" v-model="filters.startDate" class="border border-gray-300 rounded-lg px-3 py-1.5 text-sm">
          <span>至</span>
          <input type="date" v-model="filters.endDate" class="border border-gray-300 rounded-lg px-3 py-1.5 text-sm">
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">查询</button>
      </div>
    </div>

    <!-- 核心指标 -->
    <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4 mb-6">
      <div v-for="stat in coreStats" :key="stat.label" class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">{{ stat.label }}</div>
        <div class="text-xl font-bold text-gray-900 mt-1">{{ stat.value }}</div>
      </div>
    </div>

    <!-- 直播间列表 -->
    <div class="bg-white rounded-lg shadow">
      <div class="p-4 border-b">
        <h3 class="text-lg font-medium">直播间数据</h3>
      </div>
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">直播间</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">消耗</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">GMV</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">ROI</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">进入人数</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">停留时长</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">商品点击</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">成交人数</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="room in liveRooms" :key="room.id">
            <td class="px-4 py-3">
              <div class="flex items-center">
                <img :src="room.cover" class="w-12 h-12 rounded-lg mr-3 object-cover" alt="">
                <div>
                  <div class="font-medium">{{ room.name }}</div>
                  <div class="text-sm text-gray-500">{{ room.date }}</div>
                </div>
              </div>
            </td>
            <td class="px-4 py-3 text-sm text-right">¥{{ room.cost.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">¥{{ room.gmv.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right font-medium">{{ room.roi }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ room.enter.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ room.duration }}s</td>
            <td class="px-4 py-3 text-sm text-right">{{ room.productClick.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ room.buyers }}</td>
            <td class="px-4 py-3">
              <button class="text-blue-600 hover:text-blue-800 text-sm">详情</button>
            </td>
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

const filters = ref({
  liveRoom: '',
  startDate: '',
  endDate: ''
})

const coreStats = ref([
  { label: '总消耗', value: '¥86,500' },
  { label: '总GMV', value: '¥386,800' },
  { label: '平均ROI', value: '4.47' },
  { label: '进入人数', value: '125,600' },
  { label: '平均停留', value: '45s' },
  { label: '成交人数', value: '2,860' }
])

const liveRooms = ref([
  { id: 1, name: '主播A直播间', cover: 'https://via.placeholder.com/48', date: '2024-03-15', cost: 28600, gmv: 136800, roi: '4.78', enter: 45600, duration: 52, productClick: 12800, buyers: 986 },
  { id: 2, name: '品牌直播间', cover: 'https://via.placeholder.com/48', date: '2024-03-15', cost: 32500, gmv: 148600, roi: '4.57', enter: 52800, duration: 48, productClick: 15600, buyers: 1128 },
  { id: 3, name: '达人B直播间', cover: 'https://via.placeholder.com/48', date: '2024-03-15', cost: 18600, gmv: 76800, roi: '4.13', enter: 28600, duration: 38, productClick: 8200, buyers: 568 },
  { id: 4, name: '新品首发间', cover: 'https://via.placeholder.com/48', date: '2024-03-14', cost: 6800, gmv: 24600, roi: '3.62', enter: 12800, duration: 32, productClick: 3600, buyers: 178 }
])
</script>
