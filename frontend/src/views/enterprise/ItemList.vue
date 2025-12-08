<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '企业号', path: '/enterprise' }, { name: '视频列表' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">视频列表</h1>
        <p class="text-gray-600 mt-1">管理企业号发布的视频</p>
      </div>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4">
        <input type="text" v-model="filters.keyword" placeholder="搜索视频标题" class="border border-gray-300 rounded-lg px-3 py-2 w-48">
        <select v-model="filters.status" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部状态</option>
          <option value="published">已发布</option>
          <option value="reviewing">审核中</option>
          <option value="rejected">已拒绝</option>
        </select>
        <button @click="handleSearch" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
      </div>
    </div>

    <!-- 视频列表 -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
      <div v-for="video in videos" :key="video.id" class="bg-white rounded-lg shadow overflow-hidden">
        <div class="relative">
          <img :src="video.cover" class="w-full h-48 object-cover" alt="">
          <div class="absolute bottom-2 right-2 bg-black bg-opacity-60 text-white text-xs px-2 py-1 rounded">
            {{ video.duration }}
          </div>
        </div>
        <div class="p-4">
          <h4 class="font-medium line-clamp-2 mb-2">{{ video.title }}</h4>
          <div class="flex items-center text-sm text-gray-500 mb-2">
            <span>{{ video.publishTime }}</span>
            <span :class="getStatusClass(video.status)" class="ml-2 px-2 py-0.5 text-xs rounded-full">
              {{ getStatusText(video.status) }}
            </span>
          </div>
          <div class="grid grid-cols-3 gap-2 text-center text-xs text-gray-500">
            <div>
              <div class="font-medium text-gray-900">{{ video.plays }}</div>
              <div>播放</div>
            </div>
            <div>
              <div class="font-medium text-gray-900">{{ video.likes }}</div>
              <div>点赞</div>
            </div>
            <div>
              <div class="font-medium text-gray-900">{{ video.comments }}</div>
              <div>评论</div>
            </div>
          </div>
          <div class="flex space-x-2 mt-3">
            <button @click="handleVideoDetail(video)" class="flex-1 px-3 py-1 text-sm text-blue-600 hover:bg-blue-50 rounded">详情</button>
            <button @click="handleVideoData(video)" class="flex-1 px-3 py-1 text-sm text-gray-600 hover:bg-gray-50 rounded">数据</button>
          </div>
        </div>
      </div>
    </div>

    <div class="mt-6">
      <Pagination :current="1" :total="120" :page-size="12" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'
import { enterpriseApi } from '@/api/enterprise'
import { useAdvertiserStore } from '@/stores/advertiser'

const advertiserStore = useAdvertiserStore()

const filters = ref({ keyword: '', status: '' })
const loading = ref(false)
const error = ref('')
const pagination = ref({ page: 1, pageSize: 12, total: 0 })

interface VideoItem {
  id: string | number
  title: string
  cover: string
  duration: string
  publishTime: string
  status: string
  plays: string
  likes: string
  comments: string
}

const videos = ref<VideoItem[]>([])

const formatCount = (num: number) => {
  if (num >= 10000) return (num / 10000).toFixed(1) + '万'
  return num.toLocaleString()
}

const fetchVideos = async () => {
  const accountId = String(advertiserStore.currentAdvertiserId || '')
  if (!accountId) {
    error.value = '请先选择账户'
    return
  }
  
  loading.value = true
  error.value = ''
  try {
    const params: { account_id: string; page: number; page_size: number; status?: string } = {
      account_id: accountId,
      page: pagination.value.page,
      page_size: pagination.value.pageSize
    }
    if (filters.value.status) {
      params.status = filters.value.status
    }
    
    const res = await enterpriseApi.getItemList(params)
    const data = res as any
    
    if (data?.list) {
      videos.value = data.list.map((item: any) => ({
        id: item.item_id,
        title: item.title,
        cover: item.cover_url || 'https://via.placeholder.com/300x200',
        duration: item.duration || '00:00',
        publishTime: item.create_time?.substring(0, 10) || '-',
        status: item.item_status || 'published',
        plays: formatCount(item.play_count || 0),
        likes: formatCount(item.digg_count || 0),
        comments: formatCount(item.comment_count || 0)
      }))
      pagination.value.total = data.page_info?.total_number || 0
    }
  } catch (e: any) {
    error.value = e.message || '加载失败'
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.value.page = 1
  fetchVideos()
}

const handlePageChange = (page: number) => {
  pagination.value.page = page
  fetchVideos()
}

const handleVideoDetail = (video: VideoItem) => {
  alert(`视频详情: ${video.title}`)
}

const handleVideoData = (video: VideoItem) => {
  alert(`视频数据: ${video.title}`)
}

onMounted(() => {
  fetchVideos()
})

defineExpose({ handleSearch, handlePageChange })

const getStatusClass = (status: string) => {
  const classes: Record<string, string> = {
    published: 'bg-green-100 text-green-800',
    reviewing: 'bg-yellow-100 text-yellow-800',
    rejected: 'bg-red-100 text-red-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts: Record<string, string> = {
    published: '已发布',
    reviewing: '审核中',
    rejected: '已拒绝'
  }
  return texts[status] || status
}
</script>
