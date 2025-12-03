<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '工作台' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">千川工作台</h1>
      <p class="text-gray-600 mt-1">电商广告投放数据概览</p>
    </div>

    <!-- 核心指标 -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
      <div v-for="stat in coreStats" :key="stat.label" class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">{{ stat.label }}</div>
        <div class="text-2xl font-bold text-gray-900 mt-1">{{ stat.value }}</div>
        <div :class="stat.trend > 0 ? 'text-green-500' : 'text-red-500'" class="text-sm mt-1">
          {{ stat.trend > 0 ? '+' : '' }}{{ stat.trend }}% 较昨日
        </div>
      </div>
    </div>

    <!-- 投放概况 -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">消耗趋势</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
          <span class="text-gray-400">消耗趋势图表</span>
        </div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">转化趋势</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
          <span class="text-gray-400">转化趋势图表</span>
        </div>
      </div>
    </div>

    <!-- 店铺/达人数据 -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">店铺消耗TOP5</h3>
        <table class="min-w-full">
          <thead>
            <tr class="border-b">
              <th class="text-left py-2 text-sm font-medium text-gray-500">店铺名称</th>
              <th class="text-right py-2 text-sm font-medium text-gray-500">消耗</th>
              <th class="text-right py-2 text-sm font-medium text-gray-500">ROI</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="shop in topShops" :key="shop.id" class="border-b">
              <td class="py-2 text-sm">{{ shop.name }}</td>
              <td class="py-2 text-sm text-right">¥{{ shop.cost.toLocaleString() }}</td>
              <td class="py-2 text-sm text-right">{{ shop.roi }}</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">直播间实时数据</h3>
        <table class="min-w-full">
          <thead>
            <tr class="border-b">
              <th class="text-left py-2 text-sm font-medium text-gray-500">直播间</th>
              <th class="text-right py-2 text-sm font-medium text-gray-500">在线人数</th>
              <th class="text-right py-2 text-sm font-medium text-gray-500">GMV</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="room in liveRooms" :key="room.id" class="border-b">
              <td class="py-2 text-sm">{{ room.name }}</td>
              <td class="py-2 text-sm text-right">{{ room.online }}</td>
              <td class="py-2 text-sm text-right">¥{{ room.gmv.toLocaleString() }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- 快捷操作 -->
    <div class="bg-white rounded-lg shadow p-4">
      <h3 class="text-lg font-medium mb-4">快捷操作</h3>
      <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4">
        <router-link v-for="action in quickActions" :key="action.name" :to="action.path" 
          class="flex flex-col items-center p-4 rounded-lg hover:bg-gray-50 transition-colors">
          <div class="w-12 h-12 rounded-full bg-blue-100 flex items-center justify-center mb-2">
            <span class="text-blue-600">{{ action.icon }}</span>
          </div>
          <span class="text-sm text-gray-700">{{ action.name }}</span>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const coreStats = ref([
  { label: '今日消耗', value: '¥128,456', trend: 12.5 },
  { label: '今日GMV', value: '¥856,234', trend: 8.3 },
  { label: '整体ROI', value: '6.67', trend: 5.2 },
  { label: '千次曝光成本', value: '¥18.6', trend: -3.1 }
])

const topShops = ref([
  { id: 1, name: '品牌旗舰店', cost: 45680, roi: '7.2' },
  { id: 2, name: '官方专卖店', cost: 32450, roi: '6.8' },
  { id: 3, name: '直播专营店', cost: 28900, roi: '5.9' },
  { id: 4, name: '美妆专柜', cost: 18760, roi: '8.1' },
  { id: 5, name: '食品专区', cost: 12340, roi: '4.5' }
])

const liveRooms = ref([
  { id: 1, name: '主播A直播间', online: 12580, gmv: 158600 },
  { id: 2, name: '品牌直播间', online: 8900, gmv: 98500 },
  { id: 3, name: '达人B直播间', online: 6780, gmv: 76800 },
  { id: 4, name: '新品首发间', online: 4560, gmv: 45600 },
  { id: 5, name: '清仓特卖间', online: 3200, gmv: 32000 }
])

const quickActions = ref([
  { name: '创建广告', path: '/qianchuan/ad/create', icon: '📝' },
  { name: '全域推广', path: '/qianchuan/uni', icon: '🌐' },
  { name: '随心推', path: '/qianchuan/aweme-order', icon: '🚀' },
  { name: '数据报表', path: '/qianchuan/report', icon: '📊' },
  { name: '素材管理', path: '/qianchuan/material', icon: '🎬' },
  { name: '账户管理', path: '/qianchuan/account', icon: '👤' }
])
</script>
