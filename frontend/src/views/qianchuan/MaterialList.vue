<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '素材管理' }, { name: '素材库' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">素材库</h1>
        <p class="text-gray-600 mt-1">管理广告投放素材</p>
      </div>
      <div class="flex space-x-3">
        <router-link to="/qianchuan/material/video" class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">
          视频素材
        </router-link>
        <router-link to="/qianchuan/material/image" class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">
          图片素材
        </router-link>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          上传素材
        </button>
      </div>
    </div>

    <!-- 素材统计 -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">总素材数</div>
        <div class="text-2xl font-bold text-gray-900 mt-1">{{ stats.total }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">视频素材</div>
        <div class="text-2xl font-bold text-blue-600 mt-1">{{ stats.video }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">图片素材</div>
        <div class="text-2xl font-bold text-green-600 mt-1">{{ stats.image }}</div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">存储使用</div>
        <div class="text-2xl font-bold text-orange-600 mt-1">{{ stats.storage }}</div>
      </div>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4">
        <input type="text" v-model="filters.keyword" placeholder="素材名称" class="border border-gray-300 rounded-lg px-3 py-2 w-48">
        <select v-model="filters.type" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部类型</option>
          <option value="video">视频</option>
          <option value="image">图片</option>
        </select>
        <select v-model="filters.source" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部来源</option>
          <option value="upload">本地上传</option>
          <option value="live">直播间录制</option>
          <option value="aweme">抖音视频</option>
        </select>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
      </div>
    </div>

    <!-- 素材网格 -->
    <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4">
      <div v-for="material in materials" :key="material.id" class="bg-white rounded-lg shadow overflow-hidden group">
        <div class="aspect-square bg-gray-100 relative">
          <img :src="material.cover" alt="" class="w-full h-full object-cover">
          <div v-if="material.type === 'video'" class="absolute bottom-2 right-2 bg-black bg-opacity-60 text-white text-xs px-2 py-1 rounded">
            {{ material.duration }}
          </div>
          <div class="absolute inset-0 bg-black bg-opacity-0 group-hover:bg-opacity-40 transition-all flex items-center justify-center opacity-0 group-hover:opacity-100">
            <button class="px-3 py-1 bg-white rounded text-sm mr-2">预览</button>
            <button class="px-3 py-1 bg-blue-600 text-white rounded text-sm">使用</button>
          </div>
        </div>
        <div class="p-3">
          <div class="text-sm font-medium truncate">{{ material.name }}</div>
          <div class="text-xs text-gray-500 mt-1">{{ material.size }} · {{ material.createTime }}</div>
        </div>
      </div>
    </div>

    <div class="mt-6">
      <Pagination :current="1" :total="120" :page-size="18" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const stats = ref({
  total: 256,
  video: 180,
  image: 76,
  storage: '12.5 GB'
})

const filters = ref({
  keyword: '',
  type: '',
  source: ''
})

const materials = ref([
  { id: 1, name: '美妆直播片段', cover: 'https://via.placeholder.com/200', type: 'video', duration: '00:15', size: '15.2MB', createTime: '03-15' },
  { id: 2, name: '产品展示图', cover: 'https://via.placeholder.com/200', type: 'image', duration: '', size: '1.2MB', createTime: '03-15' },
  { id: 3, name: '达人口播视频', cover: 'https://via.placeholder.com/200', type: 'video', duration: '00:30', size: '28.5MB', createTime: '03-14' },
  { id: 4, name: '促销海报', cover: 'https://via.placeholder.com/200', type: 'image', duration: '', size: '0.8MB', createTime: '03-14' },
  { id: 5, name: '直播间引流', cover: 'https://via.placeholder.com/200', type: 'video', duration: '00:20', size: '18.6MB', createTime: '03-13' },
  { id: 6, name: '商品详情图', cover: 'https://via.placeholder.com/200', type: 'image', duration: '', size: '1.5MB', createTime: '03-13' },
  { id: 7, name: '用户好评视频', cover: 'https://via.placeholder.com/200', type: 'video', duration: '00:45', size: '42.3MB', createTime: '03-12' },
  { id: 8, name: '品牌Logo', cover: 'https://via.placeholder.com/200', type: 'image', duration: '', size: '0.3MB', createTime: '03-12' },
  { id: 9, name: '开箱测评', cover: 'https://via.placeholder.com/200', type: 'video', duration: '01:00', size: '56.8MB', createTime: '03-11' },
  { id: 10, name: '618活动图', cover: 'https://via.placeholder.com/200', type: 'image', duration: '', size: '2.1MB', createTime: '03-11' },
  { id: 11, name: '使用教程', cover: 'https://via.placeholder.com/200', type: 'video', duration: '00:25', size: '22.4MB', createTime: '03-10' },
  { id: 12, name: '产品对比图', cover: 'https://via.placeholder.com/200', type: 'image', duration: '', size: '1.8MB', createTime: '03-10' }
])
</script>
