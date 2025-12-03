<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '星图', path: '/star' }, { name: '需求管理' }, { name: '需求列表' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">需求管理</h1>
        <p class="text-gray-600 mt-1">发布和管理达人合作需求</p>
      </div>
      <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
        发布需求
      </button>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4">
        <input type="text" v-model="filters.keyword" placeholder="搜索需求标题" class="border border-gray-300 rounded-lg px-3 py-2 w-48">
        <select v-model="filters.status" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部状态</option>
          <option value="recruiting">招募中</option>
          <option value="executing">执行中</option>
          <option value="completed">已完成</option>
        </select>
        <select v-model="filters.type" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">需求类型</option>
          <option value="video">视频推广</option>
          <option value="live">直播带货</option>
          <option value="post">图文种草</option>
        </select>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
      </div>
    </div>

    <!-- 需求列表 -->
    <div class="space-y-4">
      <div v-for="demand in demands" :key="demand.id" class="bg-white rounded-lg shadow p-6">
        <div class="flex justify-between items-start">
          <div class="flex-1">
            <div class="flex items-center mb-2">
              <h3 class="text-lg font-medium">{{ demand.title }}</h3>
              <span :class="getStatusClass(demand.status)" class="ml-3 px-2 py-1 text-xs rounded-full">
                {{ getStatusText(demand.status) }}
              </span>
            </div>
            <p class="text-gray-600 text-sm mb-3">{{ demand.desc }}</p>
            <div class="flex flex-wrap gap-4 text-sm text-gray-500">
              <span>类型: {{ demand.type }}</span>
              <span>预算: ¥{{ demand.budget.toLocaleString() }}</span>
              <span>招募: {{ demand.recruited }}/{{ demand.targetCount }}人</span>
              <span>截止: {{ demand.deadline }}</span>
            </div>
          </div>
          <div class="text-right">
            <div class="text-2xl font-bold text-blue-600">{{ demand.applicants }}</div>
            <div class="text-sm text-gray-500">申请达人</div>
          </div>
        </div>
        <div class="mt-4 pt-4 border-t flex justify-between items-center">
          <div class="flex items-center space-x-4 text-sm">
            <span class="text-gray-500">创建时间: {{ demand.createTime }}</span>
          </div>
          <div class="flex space-x-2">
            <button class="px-4 py-1 text-blue-600 hover:bg-blue-50 rounded text-sm">查看详情</button>
            <button class="px-4 py-1 text-blue-600 hover:bg-blue-50 rounded text-sm">管理申请</button>
            <button v-if="demand.status === 'recruiting'" class="px-4 py-1 text-gray-600 hover:bg-gray-50 rounded text-sm">编辑</button>
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

const demands = ref([
  { id: 1, title: '618大促品牌种草视频招募', desc: '寻找美妆/护肤领域达人，拍摄产品种草视频，要求粉丝量5万+', type: '视频推广', status: 'recruiting', budget: 150000, targetCount: 30, recruited: 18, applicants: 86, deadline: '2024-06-10', createTime: '2024-05-20' },
  { id: 2, title: '新品发布直播带货合作', desc: '寻找带货能力强的达人进行新品首发直播，ROI要求1:5以上', type: '直播带货', status: 'executing', budget: 200000, targetCount: 5, recruited: 5, applicants: 42, deadline: '2024-06-05', createTime: '2024-05-15' },
  { id: 3, title: '品牌故事图文推广', desc: '寻找生活方式类达人发布品牌故事图文，传递品牌理念', type: '图文种草', status: 'completed', budget: 50000, targetCount: 20, recruited: 20, applicants: 56, deadline: '2024-05-30', createTime: '2024-05-10' },
  { id: 4, title: '夏日新品试用体验招募', desc: '招募美妆达人试用夏日新品，产出真实使用感受视频', type: '视频推广', status: 'recruiting', budget: 80000, targetCount: 15, recruited: 6, applicants: 35, deadline: '2024-06-20', createTime: '2024-06-01' }
])

const getStatusClass = (status: string) => {
  const classes: Record<string, string> = {
    recruiting: 'bg-blue-100 text-blue-800',
    executing: 'bg-orange-100 text-orange-800',
    completed: 'bg-green-100 text-green-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts: Record<string, string> = {
    recruiting: '招募中',
    executing: '执行中',
    completed: '已完成'
  }
  return texts[status] || status
}
</script>
