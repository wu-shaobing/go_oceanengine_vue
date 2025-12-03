<script setup lang="ts">
import { ref, onMounted, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import StatsCard from '@/components/business/StatsCard.vue'
import Pagination from '@/components/common/Pagination.vue'

const loading = ref(true)
const searchKeyword = ref('')
const typeFilter = ref('all')
const viewMode = ref<'grid' | 'list'>('grid')

const stats = reactive({
  total: 1256,
  images: 856,
  videos: 312,
  storage: '15.8 GB'
})

const medias = ref<any[]>([])
const pagination = reactive({
  current: 1,
  total: 1256,
  pageSize: 12
})

const fetchMedias = async () => {
  loading.value = true
  await new Promise(resolve => setTimeout(resolve, 500))
  
  medias.value = Array.from({ length: 12 }, (_, i) => ({
    id: `M${300001 + i + (pagination.current - 1) * 12}`,
    name: `素材_${i + 1 + (pagination.current - 1) * 12}.${Math.random() > 0.5 ? 'jpg' : 'mp4'}`,
    type: Math.random() > 0.3 ? 'image' : 'video',
    size: `${(Math.random() * 10 + 0.5).toFixed(1)} MB`,
    dimensions: Math.random() > 0.5 ? '1920x1080' : '1080x1920',
    uploadTime: '2024-01-15 14:30',
    thumbnail: `https://picsum.photos/seed/${i + pagination.current * 12}/200/150`
  }))
  
  loading.value = false
}

const handleSearch = () => {
  pagination.current = 1
  fetchMedias()
}

const handlePageChange = (page: number) => {
  pagination.current = page
  fetchMedias()
}

const selectedItems = ref<string[]>([])

const toggleSelect = (id: string) => {
  const index = selectedItems.value.indexOf(id)
  if (index > -1) {
    selectedItems.value.splice(index, 1)
  } else {
    selectedItems.value.push(id)
  }
}

const isSelected = (id: string) => selectedItems.value.includes(id)

onMounted(() => {
  fetchMedias()
})
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '素材库' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">素材库</h1>
          <p class="mt-2 text-gray-600">管理您的图片和视频素材</p>
        </div>
        <button
          class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center gap-2"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12"/>
          </svg>
          上传素材
        </button>
      </div>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-4 gap-6">
      <StatsCard title="素材总数" :value="stats.total" color="blue">
        <template #icon>
          <svg class="h-8 w-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10"/>
          </svg>
        </template>
      </StatsCard>

      <StatsCard title="图片素材" :value="stats.images" color="green">
        <template #icon>
          <svg class="h-8 w-8 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"/>
          </svg>
        </template>
      </StatsCard>

      <StatsCard title="视频素材" :value="stats.videos" color="purple">
        <template #icon>
          <svg class="h-8 w-8 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"/>
          </svg>
        </template>
      </StatsCard>

      <StatsCard title="已用空间" :value="stats.storage" color="orange">
        <template #icon>
          <svg class="h-8 w-8 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4m0 5c0 2.21-3.582 4-8 4s-8-1.79-8-4"/>
          </svg>
        </template>
      </StatsCard>
    </div>

    <!-- Filter Bar -->
    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="relative">
            <input
              v-model="searchKeyword"
              type="text"
              placeholder="搜索素材..."
              @keyup.enter="handleSearch"
              class="pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
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
            <option value="image">图片</option>
            <option value="video">视频</option>
          </select>
        </div>

        <div class="flex items-center gap-2">
          <span v-if="selectedItems.length > 0" class="text-sm text-gray-600 mr-2">
            已选择 {{ selectedItems.length }} 项
          </span>
          <button
            @click="viewMode = 'grid'"
            :class="[
              'p-2 rounded',
              viewMode === 'grid' ? 'bg-blue-100 text-blue-600' : 'text-gray-400 hover:text-gray-600'
            ]"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z"/>
            </svg>
          </button>
          <button
            @click="viewMode = 'list'"
            :class="[
              'p-2 rounded',
              viewMode === 'list' ? 'bg-blue-100 text-blue-600' : 'text-gray-400 hover:text-gray-600'
            ]"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"/>
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex justify-center py-12">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
    </div>

    <!-- Grid View -->
    <div v-else-if="viewMode === 'grid'" class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 xl:grid-cols-6 gap-4">
      <div
        v-for="media in medias"
        :key="media.id"
        @click="toggleSelect(media.id)"
        class="bg-white rounded-lg border border-gray-200 overflow-hidden cursor-pointer transition-all hover:shadow-lg"
        :class="{ 'ring-2 ring-blue-500': isSelected(media.id) }"
      >
        <div class="relative aspect-video bg-gray-100">
          <img :src="media.thumbnail" :alt="media.name" class="w-full h-full object-cover">
          <div v-if="media.type === 'video'" class="absolute inset-0 flex items-center justify-center">
            <div class="w-10 h-10 bg-black/50 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-white" fill="currentColor" viewBox="0 0 24 24">
                <path d="M8 5v14l11-7z"/>
              </svg>
            </div>
          </div>
          <div
            v-if="isSelected(media.id)"
            class="absolute top-2 right-2 w-6 h-6 bg-blue-600 rounded-full flex items-center justify-center"
          >
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
            </svg>
          </div>
        </div>
        <div class="p-3">
          <p class="text-sm font-medium text-gray-900 truncate">{{ media.name }}</p>
          <p class="text-xs text-gray-500 mt-1">{{ media.size }} · {{ media.dimensions }}</p>
        </div>
      </div>
    </div>

    <!-- List View -->
    <div v-else class="bg-white rounded-lg border border-gray-200 divide-y divide-gray-200">
      <div
        v-for="media in medias"
        :key="media.id"
        @click="toggleSelect(media.id)"
        class="flex items-center gap-4 p-4 cursor-pointer hover:bg-gray-50 transition-colors"
        :class="{ 'bg-blue-50': isSelected(media.id) }"
      >
        <div class="w-20 h-14 bg-gray-100 rounded overflow-hidden flex-shrink-0">
          <img :src="media.thumbnail" :alt="media.name" class="w-full h-full object-cover">
        </div>
        <div class="flex-1 min-w-0">
          <p class="font-medium text-gray-900 truncate">{{ media.name }}</p>
          <p class="text-sm text-gray-500">{{ media.size }} · {{ media.dimensions }}</p>
        </div>
        <div class="text-sm text-gray-500">{{ media.uploadTime }}</div>
        <div class="flex items-center gap-2">
          <button class="p-2 text-gray-400 hover:text-gray-600">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"/>
            </svg>
          </button>
          <button class="p-2 text-gray-400 hover:text-red-600">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/>
            </svg>
          </button>
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
