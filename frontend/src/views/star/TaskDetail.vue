<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '星图', path: '/star' }, { name: '任务管理', path: '/star/task' }, { name: '任务详情' }]" />
    
    <div class="mb-6 flex items-center justify-between">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">{{ task.name }}</h1>
        <p class="text-gray-600 mt-1">任务ID: {{ task.id }}</p>
      </div>
      <span :class="getStatusClass(task.status)" class="px-3 py-1 rounded text-sm">{{ getStatusText(task.status) }}</span>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6 mb-6">
      <div class="lg:col-span-2 space-y-6">
        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">任务信息</h3>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <span class="text-gray-500 text-sm">任务类型</span>
              <div class="font-medium">{{ task.type }}</div>
            </div>
            <div>
              <span class="text-gray-500 text-sm">创建时间</span>
              <div>{{ task.createTime }}</div>
            </div>
            <div>
              <span class="text-gray-500 text-sm">开始时间</span>
              <div>{{ task.startTime }}</div>
            </div>
            <div>
              <span class="text-gray-500 text-sm">结束时间</span>
              <div>{{ task.endTime }}</div>
            </div>
            <div class="col-span-2">
              <span class="text-gray-500 text-sm">任务描述</span>
              <div>{{ task.desc }}</div>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">达人列表</h3>
          <table class="min-w-full">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-4 py-2 text-left text-sm font-medium text-gray-500">达人信息</th>
                <th class="px-4 py-2 text-left text-sm font-medium text-gray-500">状态</th>
                <th class="px-4 py-2 text-right text-sm font-medium text-gray-500">费用</th>
                <th class="px-4 py-2 text-left text-sm font-medium text-gray-500">操作</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-200">
              <tr v-for="influencer in task.influencers" :key="influencer.id">
                <td class="px-4 py-3">
                  <div class="flex items-center">
                    <img src="https://via.placeholder.com/40" class="w-10 h-10 rounded-full mr-3" alt="">
                    <div>
                      <div class="font-medium">{{ influencer.name }}</div>
                      <div class="text-sm text-gray-500">粉丝: {{ (influencer.followers / 10000).toFixed(1) }}万</div>
                    </div>
                  </div>
                </td>
                <td class="px-4 py-3">
                  <span class="px-2 py-1 text-xs bg-blue-100 text-blue-800 rounded">{{ influencer.status }}</span>
                </td>
                <td class="px-4 py-3 text-right">¥{{ influencer.fee.toLocaleString() }}</td>
                <td class="px-4 py-3">
<button @click="handleViewContent(influencer)" class="text-blue-600 hover:text-blue-800 text-sm">查看内容</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <div class="space-y-6">
        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">费用概览</h3>
          <div class="space-y-3">
            <div class="flex justify-between">
              <span class="text-gray-500">预算</span>
              <span class="font-medium">¥{{ task.budget.toLocaleString() }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">已花费</span>
              <span class="text-orange-600">¥{{ task.spent.toLocaleString() }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">剩余</span>
              <span class="text-green-600">¥{{ (task.budget - task.spent).toLocaleString() }}</span>
            </div>
            <div class="mt-2 bg-gray-200 rounded-full h-2">
              <div class="bg-blue-600 h-2 rounded-full" :style="{width: (task.spent/task.budget*100)+'%'}"></div>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">效果数据</h3>
          <div class="space-y-3">
            <div class="flex justify-between">
              <span class="text-gray-500">总播放</span>
              <span class="font-medium">{{ (task.stats.views / 10000).toFixed(1) }}万</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">总点赞</span>
              <span>{{ (task.stats.likes / 10000).toFixed(1) }}万</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">总评论</span>
              <span>{{ task.stats.comments.toLocaleString() }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">总分享</span>
              <span>{{ task.stats.shares.toLocaleString() }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const task = ref({
  id: 'TASK202406001',
  name: '618年中大促-美妆种草',
  type: '视频推广',
  status: 'processing',
  desc: '618大促期间，通过达人种草视频推广美妆产品，提升品牌曝光和销售转化。',
  createTime: '2024-06-01 10:00',
  startTime: '2024-06-10 00:00',
  endTime: '2024-06-20 23:59',
  budget: 100000,
  spent: 68500,
  influencers: [
    { id: 1, name: '美妆达人小美', followers: 2560000, status: '已完成', fee: 25000 },
    { id: 2, name: '时尚博主Amy', followers: 1850000, status: '制作中', fee: 18000 },
    { id: 3, name: '护肤专家Lisa', followers: 980000, status: '待审核', fee: 12000 }
  ],
  stats: { views: 5680000, likes: 256000, comments: 18500, shares: 8600 }
})

const getStatusClass = (status: string) => {
  const classes: Record<string, string> = { pending: 'bg-orange-100 text-orange-800', processing: 'bg-blue-100 text-blue-800', completed: 'bg-green-100 text-green-800' }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts: Record<string, string> = { pending: '待开始', processing: '进行中', completed: '已完成' }
  return texts[status] || status
}

const handleViewContent = (influencer: typeof task.value.influencers[0]) => {
  alert(`查看达人内容: ${influencer.name}`)
}
</script>
