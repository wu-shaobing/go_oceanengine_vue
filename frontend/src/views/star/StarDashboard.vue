<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '星图', path: '/star' }, { name: '工作台' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">星图工作台</h1>
      <p class="text-gray-600 mt-1">达人营销平台数据概览</p>
    </div>

    <!-- 核心指标 -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
      <div v-for="stat in coreStats" :key="stat.label" class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">{{ stat.label }}</div>
        <div class="text-2xl font-bold text-gray-900 mt-1">{{ stat.value }}</div>
        <div :class="stat.trend > 0 ? 'text-green-500' : 'text-red-500'" class="text-sm mt-1">
          {{ stat.trend > 0 ? '+' : '' }}{{ stat.trend }}% 较上月
        </div>
      </div>
    </div>

    <!-- 数据图表 -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">投放消耗趋势</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
          <span class="text-gray-400">消耗趋势图表</span>
        </div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">任务完成情况</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
          <span class="text-gray-400">任务统计图表</span>
        </div>
      </div>
    </div>

    <!-- 任务/订单概览 -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-lg font-medium">进行中的任务</h3>
          <router-link to="/star/task" class="text-blue-600 text-sm">查看全部</router-link>
        </div>
        <table class="min-w-full">
          <thead>
            <tr class="border-b">
              <th class="text-left py-2 text-sm font-medium text-gray-500">任务名称</th>
              <th class="text-right py-2 text-sm font-medium text-gray-500">达人数</th>
              <th class="text-right py-2 text-sm font-medium text-gray-500">进度</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="task in activeTasks" :key="task.id" class="border-b">
              <td class="py-2 text-sm">{{ task.name }}</td>
              <td class="py-2 text-sm text-right">{{ task.influencers }}</td>
              <td class="py-2 text-sm text-right">{{ task.progress }}%</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-lg font-medium">最近订单</h3>
          <router-link to="/star/demand" class="text-blue-600 text-sm">查看全部</router-link>
        </div>
        <table class="min-w-full">
          <thead>
            <tr class="border-b">
              <th class="text-left py-2 text-sm font-medium text-gray-500">订单</th>
              <th class="text-left py-2 text-sm font-medium text-gray-500">达人</th>
              <th class="text-right py-2 text-sm font-medium text-gray-500">金额</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="order in recentOrders" :key="order.id" class="border-b">
              <td class="py-2 text-sm">{{ order.id }}</td>
              <td class="py-2 text-sm">{{ order.influencer }}</td>
              <td class="py-2 text-sm text-right">¥{{ order.amount.toLocaleString() }}</td>
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
          <div class="w-12 h-12 rounded-full bg-orange-100 flex items-center justify-center mb-2">
            <span class="text-orange-600">{{ action.icon }}</span>
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
  { label: '本月消耗', value: '¥256,800', trend: 18.5 },
  { label: '进行中任务', value: '12', trend: 25.0 },
  { label: '合作达人', value: '86', trend: 12.3 },
  { label: '内容产出', value: '156', trend: 8.6 }
])

const activeTasks = ref([
  { id: 1, name: '618大促种草任务', influencers: 25, progress: 68 },
  { id: 2, name: '新品体验官招募', influencers: 15, progress: 45 },
  { id: 3, name: '品牌故事传播', influencers: 10, progress: 82 },
  { id: 4, name: '直播带货合作', influencers: 8, progress: 30 }
])

const recentOrders = ref([
  { id: 'ST001', influencer: '美妆达人小美', amount: 15000 },
  { id: 'ST002', influencer: '时尚博主Amy', amount: 12000 },
  { id: 'ST003', influencer: '生活家小王', amount: 8000 },
  { id: 'ST004', influencer: '测评达人老李', amount: 6500 }
])

const quickActions = ref([
  { name: '创建任务', path: '/star/task/create', icon: '📝' },
  { name: '发布需求', path: '/star/demand/create', icon: '📢' },
  { name: '达人库', path: '/star/influencer', icon: '👥' },
  { name: '资金管理', path: '/star/fund', icon: '💰' },
  { name: '数据报表', path: '/star/report', icon: '📊' },
  { name: '账户设置', path: '/star/account', icon: '⚙️' }
])
</script>
