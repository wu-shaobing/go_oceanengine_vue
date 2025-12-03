<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '本地推', path: '/local' }, { name: '工作台' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">本地推工作台</h1>
      <p class="text-gray-600 mt-1">本地生活服务推广数据概览</p>
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

    <!-- 项目概览 -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">消耗趋势</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
          <span class="text-gray-400">消耗趋势图表</span>
        </div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">线索转化</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
          <span class="text-gray-400">线索转化图表</span>
        </div>
      </div>
    </div>

    <!-- 项目排行 -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">消耗TOP5项目</h3>
        <table class="min-w-full">
          <thead>
            <tr class="border-b">
              <th class="text-left py-2 text-sm font-medium text-gray-500">项目名称</th>
              <th class="text-right py-2 text-sm font-medium text-gray-500">消耗</th>
              <th class="text-right py-2 text-sm font-medium text-gray-500">线索数</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="project in topProjects" :key="project.id" class="border-b">
              <td class="py-2 text-sm">{{ project.name }}</td>
              <td class="py-2 text-sm text-right">¥{{ project.cost.toLocaleString() }}</td>
              <td class="py-2 text-sm text-right">{{ project.leads }}</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">门店数据</h3>
        <table class="min-w-full">
          <thead>
            <tr class="border-b">
              <th class="text-left py-2 text-sm font-medium text-gray-500">门店名称</th>
              <th class="text-right py-2 text-sm font-medium text-gray-500">曝光</th>
              <th class="text-right py-2 text-sm font-medium text-gray-500">到店</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="store in stores" :key="store.id" class="border-b">
              <td class="py-2 text-sm">{{ store.name }}</td>
              <td class="py-2 text-sm text-right">{{ store.exposure.toLocaleString() }}</td>
              <td class="py-2 text-sm text-right">{{ store.visits }}</td>
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
          <div class="w-12 h-12 rounded-full bg-green-100 flex items-center justify-center mb-2">
            <span class="text-green-600">{{ action.icon }}</span>
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
  { label: '今日消耗', value: '¥45,680', trend: 8.5 },
  { label: '今日线索', value: '568', trend: 12.3 },
  { label: '线索成本', value: '¥80.4', trend: -5.2 },
  { label: '到店人数', value: '126', trend: 15.8 }
])

const topProjects = ref([
  { id: 1, name: '餐饮推广项目', cost: 18560, leads: 186 },
  { id: 2, name: '美容美发推广', cost: 12800, leads: 142 },
  { id: 3, name: '教育培训招生', cost: 8600, leads: 98 },
  { id: 4, name: '汽车服务推广', cost: 3560, leads: 86 },
  { id: 5, name: '健身房获客', cost: 2160, leads: 56 }
])

const stores = ref([
  { id: 1, name: '中央大街店', exposure: 156000, visits: 58 },
  { id: 2, name: '万达广场店', exposure: 128000, visits: 42 },
  { id: 3, name: '大学城店', exposure: 98000, visits: 26 }
])

const quickActions = ref([
  { name: '创建项目', path: '/local/project/create', icon: '📝' },
  { name: '创建广告', path: '/local/promotion/create', icon: '📢' },
  { name: '线索管理', path: '/local/clue', icon: '📋' },
  { name: '数据报表', path: '/local/report', icon: '📊' },
  { name: '素材管理', path: '/local/file', icon: '🎬' },
  { name: '门店管理', path: '/local/store', icon: '🏪' }
])
</script>
