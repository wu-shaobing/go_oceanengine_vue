<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '随心推' }, { name: '订单列表' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">随心推订单</h1>
        <p class="text-gray-600 mt-1">快速投放，灵活推广</p>
      </div>
      <router-link to="/qianchuan/aweme-order/create" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
        创建随心推
      </router-link>
    </div>

    <!-- 数据概览 -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">今日消耗</div>
        <div class="text-2xl font-bold text-gray-900 mt-1">¥{{ stats.todayCost.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">进行中订单</div>
        <div class="text-2xl font-bold text-blue-600 mt-1">{{ stats.runningOrders }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">今日完成</div>
        <div class="text-2xl font-bold text-green-600 mt-1">{{ stats.completedOrders }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">平均ROI</div>
        <div class="text-2xl font-bold text-orange-600 mt-1">{{ stats.avgRoi }}</div>
      </div>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4">
        <input type="text" v-model="filters.keyword" placeholder="订单ID/视频名称" class="border border-gray-300 rounded-lg px-3 py-2 w-48">
        <select v-model="filters.status" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部状态</option>
          <option value="running">投放中</option>
          <option value="completed">已完成</option>
          <option value="failed">投放失败</option>
        </select>
        <select v-model="filters.goal" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部目标</option>
          <option value="video">视频播放</option>
          <option value="interaction">互动提升</option>
          <option value="fans">涨粉</option>
          <option value="product">商品购买</option>
        </select>
<button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handleSearch">搜索</button>
      </div>
    </div>

    <!-- 订单列表 -->
    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">订单信息</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">推广目标</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">状态</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">预算</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">消耗</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">播放量</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">互动数</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="order in orders" :key="order.id">
            <td class="px-4 py-3">
              <div class="flex items-center">
                <img :src="order.cover" class="w-16 h-20 rounded mr-3 object-cover" alt="">
                <div>
                  <div class="font-medium line-clamp-2 w-40">{{ order.videoTitle }}</div>
                  <div class="text-sm text-gray-500">{{ order.id }}</div>
                </div>
              </div>
            </td>
            <td class="px-4 py-3 text-sm">{{ order.goal }}</td>
            <td class="px-4 py-3">
              <span :class="getStatusClass(order.status)" class="px-2 py-1 text-xs rounded-full">
                {{ getStatusText(order.status) }}
              </span>
            </td>
            <td class="px-4 py-3 text-sm">¥{{ order.budget }}</td>
            <td class="px-4 py-3 text-sm">¥{{ order.cost }}</td>
            <td class="px-4 py-3 text-sm">{{ order.plays.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm">{{ order.interactions.toLocaleString() }}</td>
            <td class="px-4 py-3">
              <div class="flex space-x-2">
                <router-link :to="`/qianchuan/aweme-order/${order.id}`" class="text-blue-600 hover:text-blue-800 text-sm">详情</router-link>
<button v-if="order.status === 'running'" class="text-orange-600 hover:text-orange-800 text-sm" @click="handleStop(order)">停止</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="200" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const stats = ref({
  todayCost: 15680,
  runningOrders: 12,
  completedOrders: 45,
  avgRoi: '3.8'
})

const filters = ref({
  keyword: '',
  status: '',
  goal: ''
})

const orders = ref([
  { id: 'AO001', videoTitle: '新品美妆测评分享', cover: 'https://via.placeholder.com/64x80', goal: '商品购买', status: 'running', budget: 500, cost: 380, plays: 25600, interactions: 1580 },
  { id: 'AO002', videoTitle: '直播间精彩片段', cover: 'https://via.placeholder.com/64x80', goal: '视频播放', status: 'running', budget: 300, cost: 285, plays: 45800, interactions: 2360 },
  { id: 'AO003', videoTitle: '穿搭分享日常', cover: 'https://via.placeholder.com/64x80', goal: '涨粉', status: 'completed', budget: 200, cost: 200, plays: 18600, interactions: 960 },
  { id: 'AO004', videoTitle: '美食探店vlog', cover: 'https://via.placeholder.com/64x80', goal: '互动提升', status: 'completed', budget: 100, cost: 100, plays: 12800, interactions: 680 },
  { id: 'AO005', videoTitle: '好物推荐合集', cover: 'https://via.placeholder.com/64x80', goal: '商品购买', status: 'failed', budget: 500, cost: 0, plays: 0, interactions: 0 }
])

const getStatusClass = (status: string) => {
  const classes: Record<string, string> = {
    running: 'bg-green-100 text-green-800',
    completed: 'bg-blue-100 text-blue-800',
    failed: 'bg-red-100 text-red-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts: Record<string, string> = {
    running: '投放中',
    completed: '已完成',
    failed: '投放失败'
  }
  return texts[status] || status
}

const handleSearch = () => {
  alert(`搜索: ${filters.value.keyword}`)
}

const handleStop = (order: typeof orders.value[0]) => {
  alert(`停止订单: ${order.id}`)
}
</script>
