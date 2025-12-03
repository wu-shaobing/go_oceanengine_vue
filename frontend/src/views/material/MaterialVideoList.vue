<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 12, total: 156 })
const viewMode = ref<'grid' | 'list'>('grid')

const videos = ref([
  { id: 'V001', name: '品牌宣传片15s', thumbnail: '🎬', duration: 15, size: '12.5MB', resolution: '1080x1920', status: 'active', usedCount: 23, createdAt: '2025-10-01' },
  { id: 'V002', name: '产品展示30s', thumbnail: '📹', duration: 30, size: '25.8MB', resolution: '1080x1920', status: 'active', usedCount: 45, createdAt: '2025-10-05' },
  { id: 'V003', name: '促销活动视频', thumbnail: '🎥', duration: 20, size: '18.2MB', resolution: '720x1280', status: 'reviewing', usedCount: 0, createdAt: '2025-10-08' },
  { id: 'V004', name: '用户好评集锦', thumbnail: '📽️', duration: 45, size: '38.5MB', resolution: '1080x1920', status: 'active', usedCount: 12, createdAt: '2025-10-10' },
  { id: 'V005', name: '产品功能介绍', thumbnail: '🎞️', duration: 25, size: '22.1MB', resolution: '1080x1920', status: 'rejected', usedCount: 0, createdAt: '2025-10-12' },
  { id: 'V006', name: '双11预热视频', thumbnail: '🎬', duration: 18, size: '15.6MB', resolution: '1080x1920', status: 'active', usedCount: 67, createdAt: '2025-10-15' }
])

const getStatusConfig = (status: string) => {
  switch (status) {
    case 'active': return { label: '可用', class: 'bg-green-100 text-green-800' }
    case 'reviewing': return { label: '审核中', class: 'bg-yellow-100 text-yellow-800' }
    case 'rejected': return { label: '已拒绝', class: 'bg-red-100 text-red-800' }
    default: return { label: status, class: 'bg-gray-100 text-gray-800' }
  }
}

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '素材管理' }, { name: '视频素材' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">视频素材</h1>
          <p class="mt-2 text-gray-600">管理广告投放使用的视频素材</p>
        </div>
        <div class="flex items-center gap-4">
          <div class="flex border border-gray-300 rounded-lg overflow-hidden">
            <button @click="viewMode = 'grid'" :class="['px-3 py-1.5', viewMode === 'grid' ? 'bg-blue-50 text-blue-600' : 'bg-white text-gray-600']">
              网格
            </button>
            <button @click="viewMode = 'list'" :class="['px-3 py-1.5', viewMode === 'list' ? 'bg-blue-50 text-blue-600' : 'bg-white text-gray-600']">
              列表
            </button>
          </div>
          <router-link to="/material/video/upload" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
            上传视频
          </router-link>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总视频</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">可用</p>
        <p class="text-2xl font-bold text-green-600 mt-1">128</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">审核中</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">18</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总使用次数</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">1,256</p>
      </div>
    </div>

    <div v-if="viewMode === 'grid'" class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-6 gap-4">
      <div v-for="video in videos" :key="video.id" 
           class="bg-white rounded-lg border border-gray-200 overflow-hidden hover:shadow-md transition-shadow">
        <div class="aspect-[9/16] bg-gray-100 flex items-center justify-center text-6xl">
          {{ video.thumbnail }}
        </div>
        <div class="p-3">
          <h4 class="font-medium text-sm text-gray-900 truncate">{{ video.name }}</h4>
          <div class="flex items-center justify-between mt-2">
            <span class="text-xs text-gray-500">{{ video.duration }}s</span>
            <span :class="['px-1.5 py-0.5 rounded text-xs', getStatusConfig(video.status).class]">
              {{ getStatusConfig(video.status).label }}
            </span>
          </div>
        </div>
      </div>
    </div>

    <div v-else class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">视频</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">时长</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">大小</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">分辨率</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">使用次数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="video in videos" :key="video.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="flex items-center gap-3">
                <span class="text-2xl">{{ video.thumbnail }}</span>
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ video.name }}</div>
                  <div class="text-xs text-gray-500">{{ video.id }}</div>
                </div>
              </div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ video.duration }}s</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ video.size }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ video.resolution }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ video.usedCount }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium', getStatusConfig(video.status).class]">
                {{ getStatusConfig(video.status).label }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">预览</button>
              <button class="text-gray-600 hover:text-gray-800">删除</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div class="flex justify-center">
      <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
    </div>
  </div>
</template>
