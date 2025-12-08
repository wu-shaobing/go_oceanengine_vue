<script setup lang="ts">
import { ref, onMounted, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import StatsCard from '@/components/business/StatsCard.vue'
import StatusBadge from '@/components/business/StatusBadge.vue'
import Pagination from '@/components/common/Pagination.vue'

const loading = ref(true)
const searchKeyword = ref('')
const typeFilter = ref('all')

const stats = reactive({
  total: 328,
  active: 215,
  paused: 78,
  pending: 35
})

const creatives = ref<any[]>([])
const pagination = reactive({
  current: 1,
  total: 328,
  pageSize: 8
})

const creativeTypes = ['image', 'video', 'carousel']
const creativeTypeLabels: Record<string, string> = {
  image: '图片广告',
  video: '视频广告',
  carousel: '轮播广告'
}

const fetchCreatives = async () => {
  loading.value = true
  await new Promise(resolve => setTimeout(resolve, 500))
  
  creatives.value = Array.from({ length: 8 }, (_, i) => ({
    id: `CR${400001 + i + (pagination.current - 1) * 8}`,
    name: `创意素材 ${i + 1 + (pagination.current - 1) * 8}`,
    type: creativeTypes[Math.floor(Math.random() * 3)],
    impressions: Math.floor(Math.random() * 150000) + 20000,
    clicks: Math.floor(Math.random() * 4000) + 500,
    ctr: (Math.random() * 3 + 1).toFixed(2),
    spend: Math.floor(Math.random() * 3000) + 500,
    status: ['active', 'paused', 'pending'][Math.floor(Math.random() * 3)],
    thumbnail: `https://picsum.photos/seed/${i + pagination.current * 8 + 100}/300/200`
  }))
  
  loading.value = false
}

const handleSearch = () => {
  pagination.current = 1
  fetchCreatives()
}

const handlePageChange = (page: number) => {
  pagination.current = page
  fetchCreatives()
}

const formatMoney = (value: number) => `¥${value.toLocaleString()}`
const formatNumber = (value: number) => value.toLocaleString()

const getStatusConfig = (status: string) => {
  const map: Record<string, { status: 'success' | 'warning' | 'info', text: string }> = {
    active: { status: 'success', text: '投放中' },
    paused: { status: 'warning', text: '已暂停' },
    pending: { status: 'info', text: '审核中' }
  }
  return map[status] || { status: 'info', text: status }
}

onMounted(() => {
  fetchCreatives()
})

const handleCreateCreative = () => {
  alert('创建创意')
}

const handleEditCreative = (creative: any) => {
  alert(`编辑创意: ${creative.name}`)
}

const handleCopyCreative = (creative: any) => {
  alert(`复制创意: ${creative.name}`)
}

const handleMoreActions = (creative: any) => {
  alert(`更多操作: ${creative.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '创意管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">创意管理</h1>
          <p class="mt-2 text-gray-600">管理您的广告创意素材</p>
        </div>
<button
          class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center gap-2"
          @click="handleCreateCreative"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
          </svg>
          创建创意
        </button>
      </div>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-4 gap-6">
      <StatsCard title="创意总数" :value="stats.total" color="blue">
        <template #icon>
          <svg class="h-8 w-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zm0 0h12a2 2 0 002-2v-4a2 2 0 00-2-2h-2.343M11 7.343l1.657-1.657a2 2 0 012.828 0l2.829 2.829a2 2 0 010 2.828l-8.486 8.485M7 17h.01"/>
          </svg>
        </template>
      </StatsCard>

      <StatsCard title="投放中" :value="stats.active" color="green">
        <template #icon>
          <svg class="h-8 w-8 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z"/>
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
        </template>
      </StatsCard>

      <StatsCard title="已暂停" :value="stats.paused" color="orange">
        <template #icon>
          <svg class="h-8 w-8 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 9v6m4-6v6m7-3a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
        </template>
      </StatsCard>

      <StatsCard title="审核中" :value="stats.pending" color="purple">
        <template #icon>
          <svg class="h-8 w-8 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
        </template>
      </StatsCard>
    </div>

    <!-- Filter Bar -->
    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex items-center gap-3">
        <div class="relative flex-1 max-w-md">
          <input
            v-model="searchKeyword"
            type="text"
            placeholder="搜索创意..."
            @keyup.enter="handleSearch"
            class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
          <svg class="absolute left-3 top-2.5 h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/>
          </svg>
        </div>
        <select
          v-model="typeFilter"
          @change="handleSearch"
          class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        >
          <option value="all">全部类型</option>
          <option value="image">图片广告</option>
          <option value="video">视频广告</option>
          <option value="carousel">轮播广告</option>
        </select>
      </div>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex justify-center py-12">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
    </div>

    <!-- Creative Cards -->
    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
      <div
        v-for="creative in creatives"
        :key="creative.id"
        class="bg-white rounded-lg border border-gray-200 overflow-hidden hover:shadow-lg transition-shadow"
      >
        <div class="relative aspect-video bg-gray-100">
          <img :src="creative.thumbnail" :alt="creative.name" class="w-full h-full object-cover">
          <div class="absolute top-2 right-2">
            <StatusBadge
              :status="getStatusConfig(creative.status).status"
              :text="getStatusConfig(creative.status).text"
            />
          </div>
          <div class="absolute bottom-2 left-2 px-2 py-1 bg-black/60 text-white text-xs rounded">
            {{ creativeTypeLabels[creative.type] }}
          </div>
        </div>
        <div class="p-4">
          <h4 class="font-medium text-gray-900 truncate">{{ creative.name }}</h4>
          <p class="text-sm text-gray-500 mt-1">ID: {{ creative.id }}</p>
          
          <div class="grid grid-cols-2 gap-2 mt-4 text-sm">
            <div>
              <p class="text-gray-500">展示</p>
              <p class="font-medium text-gray-900">{{ formatNumber(creative.impressions) }}</p>
            </div>
            <div>
              <p class="text-gray-500">点击</p>
              <p class="font-medium text-gray-900">{{ formatNumber(creative.clicks) }}</p>
            </div>
            <div>
              <p class="text-gray-500">点击率</p>
              <p class="font-medium text-gray-900">{{ creative.ctr }}%</p>
            </div>
            <div>
              <p class="text-gray-500">消耗</p>
              <p class="font-medium text-gray-900">{{ formatMoney(creative.spend) }}</p>
            </div>
          </div>

<div class="flex items-center gap-2 mt-4 pt-4 border-t border-gray-100">
            <button class="flex-1 px-3 py-1.5 text-sm text-blue-600 hover:bg-blue-50 rounded transition-colors" @click="handleEditCreative(creative)">
              编辑
            </button>
            <button class="flex-1 px-3 py-1.5 text-sm text-gray-600 hover:bg-gray-50 rounded transition-colors" @click="handleCopyCreative(creative)">
              复制
            </button>
            <button class="px-3 py-1.5 text-sm text-gray-400 hover:text-gray-600 rounded transition-colors" @click="handleMoreActions(creative)">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z"/>
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Pagination -->
    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <Pagination
        :current="pagination.current"
        :total="pagination.total"
        :page-size="pagination.pageSize"
        @change="handlePageChange"
      />
    </div>
  </div>
</template>
