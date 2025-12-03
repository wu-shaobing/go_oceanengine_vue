<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '本地推', path: '/local' }, { name: '推广管理', path: '/local/promotion' }, { name: '推广详情' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">{{ promotion.name }}</h1>
        <p class="text-gray-600 mt-1">ID: {{ promotion.id }}</p>
      </div>
      <div class="flex space-x-3">
        <button v-if="promotion.status === 'running'" @click="pausePromotion" class="px-4 py-2 border border-orange-600 text-orange-600 rounded hover:bg-orange-50">暂停推广</button>
        <button v-else-if="promotion.status === 'paused'" @click="resumePromotion" class="px-4 py-2 border border-green-600 text-green-600 rounded hover:bg-green-50">启动推广</button>
        <router-link :to="`/local/promotion/edit/${promotion.id}`" class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">编辑</router-link>
      </div>
    </div>

    <!-- 状态卡片 -->
    <div class="grid grid-cols-1 md:grid-cols-5 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">状态</div>
        <div class="mt-1">
          <span :class="getStatusClass(promotion.status)" class="px-2 py-1 text-sm rounded-full">
            {{ getStatusText(promotion.status) }}
          </span>
        </div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">今日消耗</div>
        <div class="text-xl font-bold text-gray-900 mt-1">¥{{ promotion.todayCost.toLocaleString() }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">今日线索</div>
        <div class="text-xl font-bold text-gray-900 mt-1">{{ promotion.todayClues }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">线索成本</div>
        <div class="text-xl font-bold text-gray-900 mt-1">¥{{ promotion.clueCost }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">总消耗</div>
        <div class="text-xl font-bold text-gray-900 mt-1">¥{{ promotion.totalCost.toLocaleString() }}</div>
      </div>
    </div>

    <!-- 详细信息 -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
      <!-- 基本信息 -->
      <div class="bg-white rounded-lg shadow p-6">
        <h3 class="text-lg font-medium mb-4">基本信息</h3>
        <dl class="space-y-3">
          <div class="flex">
            <dt class="w-24 text-gray-500">所属项目</dt>
            <dd class="text-gray-900">{{ promotion.project }}</dd>
          </div>
          <div class="flex">
            <dt class="w-24 text-gray-500">推广目标</dt>
            <dd class="text-gray-900">{{ getGoalText(promotion.goal) }}</dd>
          </div>
          <div class="flex">
            <dt class="w-24 text-gray-500">投放时间</dt>
            <dd class="text-gray-900">{{ promotion.startDate }} ~ {{ promotion.endDate || '长期' }}</dd>
          </div>
          <div class="flex">
            <dt class="w-24 text-gray-500">日预算</dt>
            <dd class="text-gray-900">¥{{ promotion.dailyBudget }}</dd>
          </div>
          <div class="flex">
            <dt class="w-24 text-gray-500">出价方式</dt>
            <dd class="text-gray-900">{{ getBidTypeText(promotion.bidType) }}</dd>
          </div>
          <div class="flex">
            <dt class="w-24 text-gray-500">创建时间</dt>
            <dd class="text-gray-900">{{ promotion.createdAt }}</dd>
          </div>
        </dl>
      </div>

      <!-- 定向设置 -->
      <div class="bg-white rounded-lg shadow p-6">
        <h3 class="text-lg font-medium mb-4">定向设置</h3>
        <dl class="space-y-3">
          <div class="flex">
            <dt class="w-24 text-gray-500">投放范围</dt>
            <dd class="text-gray-900">门店周边{{ promotion.locationRadius }}公里</dd>
          </div>
          <div class="flex">
            <dt class="w-24 text-gray-500">年龄定向</dt>
            <dd class="text-gray-900">{{ promotion.ageMin }}-{{ promotion.ageMax }}岁</dd>
          </div>
          <div class="flex">
            <dt class="w-24 text-gray-500">性别定向</dt>
            <dd class="text-gray-900">{{ getGenderText(promotion.gender) }}</dd>
          </div>
        </dl>
      </div>
    </div>

    <!-- 数据趋势 -->
    <div class="bg-white rounded-lg shadow p-6 mb-6">
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-lg font-medium">数据趋势</h3>
        <div class="flex space-x-2">
          <button v-for="range in dateRanges" :key="range.value" @click="selectedRange = range.value" :class="selectedRange === range.value ? 'bg-blue-600 text-white' : 'bg-gray-100 text-gray-700'" class="px-3 py-1 text-sm rounded">
            {{ range.label }}
          </button>
        </div>
      </div>
      <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
        <span class="text-gray-400">数据趋势图表</span>
      </div>
    </div>

    <!-- 创意预览 -->
    <div class="bg-white rounded-lg shadow p-6">
      <h3 class="text-lg font-medium mb-4">创意预览</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="border rounded-lg overflow-hidden">
          <div class="aspect-video bg-gray-100 flex items-center justify-center">
            <span class="text-gray-400">{{ promotion.creativeType === 'video' ? '视频预览' : '图片预览' }}</span>
          </div>
          <div class="p-4">
            <div class="font-medium">{{ promotion.title }}</div>
            <div class="text-sm text-gray-500 mt-1">{{ promotion.creativeType === 'video' ? '短视频' : '图片' }}</div>
          </div>
        </div>
        <div class="border rounded-lg p-4">
          <h4 class="font-medium mb-3">创意数据</h4>
          <dl class="space-y-2 text-sm">
            <div class="flex justify-between">
              <dt class="text-gray-500">曝光量</dt>
              <dd class="text-gray-900">{{ promotion.stats.impressions.toLocaleString() }}</dd>
            </div>
            <div class="flex justify-between">
              <dt class="text-gray-500">点击量</dt>
              <dd class="text-gray-900">{{ promotion.stats.clicks.toLocaleString() }}</dd>
            </div>
            <div class="flex justify-between">
              <dt class="text-gray-500">点击率</dt>
              <dd class="text-gray-900">{{ promotion.stats.ctr }}%</dd>
            </div>
            <div class="flex justify-between">
              <dt class="text-gray-500">转化量</dt>
              <dd class="text-gray-900">{{ promotion.stats.conversions }}</dd>
            </div>
            <div class="flex justify-between">
              <dt class="text-gray-500">转化率</dt>
              <dd class="text-gray-900">{{ promotion.stats.cvr }}%</dd>
            </div>
          </dl>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const selectedRange = ref('7d')
const dateRanges = [
  { value: 'today', label: '今日' },
  { value: '7d', label: '近7天' },
  { value: '30d', label: '近30天' }
]

const promotion = ref({
  id: 'PM001',
  name: '618限时优惠推广',
  project: '北京朝阳店',
  status: 'running',
  goal: 'clue',
  todayCost: 1580,
  todayClues: 28,
  clueCost: 56.4,
  totalCost: 28560,
  startDate: '2024-06-01',
  endDate: '',
  dailyBudget: 2000,
  bidType: 'ocpm',
  locationRadius: 5,
  ageMin: 18,
  ageMax: 45,
  gender: 'all',
  creativeType: 'video',
  title: '618大促来袭！限时特惠等你来',
  createdAt: '2024-05-28 10:30:00',
  stats: {
    impressions: 125680,
    clicks: 8560,
    ctr: 6.81,
    conversions: 285,
    cvr: 3.33
  }
})

const getStatusClass = (status: string) => {
  const classes: Record<string, string> = {
    running: 'bg-green-100 text-green-800',
    paused: 'bg-orange-100 text-orange-800',
    ended: 'bg-gray-100 text-gray-800',
    reviewing: 'bg-blue-100 text-blue-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts: Record<string, string> = {
    running: '投放中',
    paused: '已暂停',
    ended: '已结束',
    reviewing: '审核中'
  }
  return texts[status] || status
}

const getGoalText = (goal: string) => {
  const texts: Record<string, string> = {
    clue: '获取线索',
    visit: '到店访问',
    awareness: '品牌曝光'
  }
  return texts[goal] || goal
}

const getBidTypeText = (bidType: string) => {
  const texts: Record<string, string> = {
    ocpm: 'OCPM智能出价',
    cpc: 'CPC点击出价',
    cpm: 'CPM展示出价'
  }
  return texts[bidType] || bidType
}

const getGenderText = (gender: string) => {
  const texts: Record<string, string> = {
    all: '不限',
    male: '男',
    female: '女'
  }
  return texts[gender] || gender
}

const pausePromotion = () => {
  promotion.value.status = 'paused'
}

const resumePromotion = () => {
  promotion.value.status = 'running'
}
</script>
