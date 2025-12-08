<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '投放管理' }, { name: '创意管理' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">创意管理</h1>
        <p class="text-gray-600 mt-1">管理广告创意素材</p>
      </div>
<button @click="handleUpload" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
        上传创意
      </button>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4">
        <input type="text" v-model="filters.keyword" placeholder="创意名称/ID" class="border border-gray-300 rounded-lg px-3 py-2 w-48">
        <select v-model="filters.status" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部状态</option>
          <option value="pass">审核通过</option>
          <option value="audit">审核中</option>
          <option value="reject">审核拒绝</option>
        </select>
        <select v-model="filters.type" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部类型</option>
          <option value="video">视频</option>
          <option value="image">图片</option>
        </select>
<button @click="handleSearch" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
      </div>
    </div>

    <!-- 创意列表 -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
      <div v-for="creative in creatives" :key="creative.id" class="bg-white rounded-lg shadow overflow-hidden">
        <div class="aspect-video bg-gray-100 relative">
          <img :src="creative.cover" alt="" class="w-full h-full object-cover">
          <div v-if="creative.type === 'video'" class="absolute bottom-2 right-2 bg-black bg-opacity-60 text-white text-xs px-2 py-1 rounded">
            {{ creative.duration }}
          </div>
        </div>
        <div class="p-4">
          <div class="font-medium truncate">{{ creative.name }}</div>
          <div class="text-sm text-gray-500 mt-1">ID: {{ creative.id }}</div>
          <div class="flex items-center justify-between mt-2">
            <span :class="getStatusClass(creative.status)" class="px-2 py-1 text-xs rounded-full">
              {{ getStatusText(creative.status) }}
            </span>
            <span class="text-sm text-gray-500">{{ creative.type === 'video' ? '视频' : '图片' }}</span>
          </div>
          <div class="mt-3 pt-3 border-t text-sm">
            <div class="flex justify-between">
              <span class="text-gray-500">使用次数</span>
              <span>{{ creative.useCount }}</span>
            </div>
            <div class="flex justify-between mt-1">
              <span class="text-gray-500">创建时间</span>
              <span>{{ creative.createTime }}</span>
            </div>
          </div>
          <div class="flex justify-between mt-3 pt-3 border-t">
<button @click="handlePreview(creative)" class="text-blue-600 hover:text-blue-800 text-sm">预览</button>
            <button @click="handleEditCreative(creative)" class="text-blue-600 hover:text-blue-800 text-sm">编辑</button>
            <button @click="handleDelete(creative)" class="text-red-600 hover:text-red-800 text-sm">删除</button>
          </div>
        </div>
      </div>
    </div>

    <div class="mt-6">
      <Pagination :current="1" :total="50" :page-size="12" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'
import { qianchuanApi } from '@/api/qianchuan'
import { useAdvertiserStore } from '@/stores/advertiser'

const advertiserStore = useAdvertiserStore()

const filters = ref({
  keyword: '',
  status: '',
  type: ''
})

const loading = ref(false)
const error = ref('')
const pagination = ref({ page: 1, pageSize: 12, total: 0 })

interface CreativeItem {
  id: string
  name: string
  cover: string
  type: string
  duration: string
  status: string
  useCount: number
  createTime: string
}

const creatives = ref<CreativeItem[]>([])

const fetchCreatives = async () => {
  const advertiserId = advertiserStore.currentAdvertiserId
  if (!advertiserId) {
    error.value = '请先选择广告主'
    return
  }
  
  loading.value = true
  error.value = ''
  try {
    const res = await qianchuanApi.getCreativeList({
      advertiser_id: advertiserId,
      page: pagination.value.page,
      page_size: pagination.value.pageSize
    })
    const data = res as any
    
    if (data?.list) {
      creatives.value = data.list.map((item: any) => ({
        id: String(item.creative_id),
        name: item.title || '未命名创意',
        cover: item.image_url || 'https://via.placeholder.com/320x180',
        type: item.creative_material_mode === 'VIDEO' ? 'video' : 'image',
        duration: item.duration || '',
        status: item.status || 'unknown',
        useCount: item.use_count || 0,
        createTime: item.create_time || '-'
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
  fetchCreatives()
}

const handlePageChange = (page: number) => {
  pagination.value.page = page
  fetchCreatives()
}

onMounted(() => {
  fetchCreatives()
})

defineExpose({ handleSearch, handlePageChange })

const getStatusClass = (status: string) => {
  const classes: Record<string, string> = {
    pass: 'bg-green-100 text-green-800',
    audit: 'bg-yellow-100 text-yellow-800',
    reject: 'bg-red-100 text-red-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts: Record<string, string> = {
    pass: '审核通过',
    audit: '审核中',
    reject: '审核拒绝'
  }
  return texts[status] || status
}

const handleUpload = () => {
  alert('上传创意')
}

const handlePreview = (creative: CreativeItem) => {
  alert(`预览创意: ${creative.name}`)
}

const handleEditCreative = (creative: CreativeItem) => {
  alert(`编辑创意: ${creative.name}`)
}

const handleDelete = (creative: CreativeItem) => {
  alert(`删除创意: ${creative.name}`)
}
</script>
