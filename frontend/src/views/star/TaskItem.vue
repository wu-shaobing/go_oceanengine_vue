<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '星图', path: '/star' }, { name: '任务列表' }]" />
    
    <div class="mb-6 flex justify-between items-center">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">任务管理</h1>
        <p class="text-gray-600 mt-1">管理达人合作任务</p>
      </div>
      <button class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">创建任务</button>
    </div>

    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <input type="text" v-model="filters.keyword" placeholder="搜索任务" class="border border-gray-300 rounded px-3 py-2 w-48">
        <select v-model="filters.status" class="border border-gray-300 rounded px-3 py-2">
          <option value="">全部状态</option>
          <option value="pending">待开始</option>
          <option value="processing">进行中</option>
          <option value="completed">已完成</option>
        </select>
        <select v-model="filters.type" class="border border-gray-300 rounded px-3 py-2">
          <option value="">全部类型</option>
          <option value="video">视频推广</option>
          <option value="live">直播推广</option>
          <option value="article">图文推广</option>
        </select>
        <button class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">查询</button>
      </div>
    </div>

    <div class="space-y-4">
      <div v-for="task in tasks" :key="task.id" class="bg-white rounded-lg shadow p-4">
        <div class="flex items-start justify-between">
          <div class="flex-1">
            <div class="flex items-center mb-2">
              <h3 class="text-lg font-medium">{{ task.name }}</h3>
              <span :class="getStatusClass(task.status)" class="ml-3 px-2 py-1 text-xs rounded">{{ getStatusText(task.status) }}</span>
            </div>
            <p class="text-gray-500 text-sm mb-3">{{ task.desc }}</p>
            <div class="flex items-center text-sm text-gray-500 space-x-6">
              <span>类型: {{ task.type }}</span>
              <span>达人数: {{ task.influencerCount }}</span>
              <span>时间: {{ task.startTime }} ~ {{ task.endTime }}</span>
            </div>
          </div>
          <div class="text-right">
            <div class="text-sm text-gray-500">预算</div>
            <div class="text-xl font-bold text-blue-600">¥{{ task.budget.toLocaleString() }}</div>
          </div>
        </div>
        <div class="mt-4 pt-4 border-t flex items-center justify-between">
          <div class="flex items-center space-x-8 text-sm">
            <div><span class="text-gray-500">播放: </span><span class="font-medium">{{ (task.stats.views / 10000).toFixed(1) }}万</span></div>
            <div><span class="text-gray-500">互动: </span><span class="font-medium">{{ (task.stats.interactions / 10000).toFixed(1) }}万</span></div>
            <div><span class="text-gray-500">转化: </span><span class="font-medium">{{ task.stats.conversions }}</span></div>
          </div>
          <div class="flex space-x-2">
            <router-link :to="`/star/task/${task.id}`" class="px-3 py-1 text-sm text-blue-600 border border-blue-600 rounded hover:bg-blue-50">详情</router-link>
            <button v-if="task.status === 'pending'" class="px-3 py-1 text-sm text-gray-600 border border-gray-300 rounded hover:bg-gray-50">编辑</button>
          </div>
        </div>
      </div>
    </div>

    <div class="mt-6">
      <Pagination :current="1" :total="50" :page-size="10" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const filters = ref({ keyword: '', status: '', type: '' })

const tasks = ref([
  { id: 1, name: '618年中大促-美妆种草', desc: '618期间美妆产品达人种草推广', type: '视频推广', status: 'processing', influencerCount: 15, budget: 100000, startTime: '2024-06-10', endTime: '2024-06-20', stats: { views: 5680000, interactions: 286000, conversions: 3560 } },
  { id: 2, name: '新品发布-护肤系列', desc: '新品护肤系列达人测评推广', type: '视频推广', status: 'pending', influencerCount: 8, budget: 50000, startTime: '2024-06-25', endTime: '2024-07-05', stats: { views: 0, interactions: 0, conversions: 0 } },
  { id: 3, name: '品牌直播专场', desc: '品牌日直播带货专场', type: '直播推广', status: 'completed', influencerCount: 5, budget: 200000, startTime: '2024-06-01', endTime: '2024-06-05', stats: { views: 12500000, interactions: 856000, conversions: 15680 } }
])

const getStatusClass = (status: string) => {
  const classes: Record<string, string> = { pending: 'bg-orange-100 text-orange-800', processing: 'bg-blue-100 text-blue-800', completed: 'bg-green-100 text-green-800' }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts: Record<string, string> = { pending: '待开始', processing: '进行中', completed: '已完成' }
  return texts[status] || status
}
</script>
