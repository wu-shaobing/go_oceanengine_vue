<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '素材管理' }, { name: '视频素材' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">视频素材</h1>
        <p class="text-gray-600 mt-1">管理视频广告素材</p>
      </div>
<button @click="handleUpload" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
        上传视频
      </button>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4">
        <input type="text" v-model="filters.keyword" placeholder="视频名称" class="border border-gray-300 rounded-lg px-3 py-2 w-48">
        <select v-model="filters.resolution" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部尺寸</option>
          <option value="9:16">9:16 竖版</option>
          <option value="16:9">16:9 横版</option>
          <option value="1:1">1:1 方形</option>
        </select>
        <select v-model="filters.duration" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部时长</option>
          <option value="0-15">15秒以内</option>
          <option value="15-30">15-30秒</option>
          <option value="30-60">30-60秒</option>
          <option value="60+">60秒以上</option>
        </select>
<button @click="handleSearch" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
      </div>
    </div>

    <!-- 视频列表 -->
    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left"><input type="checkbox" class="rounded"></th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">视频信息</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">时长</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">尺寸</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">大小</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">来源</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">使用次数</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">上传时间</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="video in videos" :key="video.id">
            <td class="px-4 py-3"><input type="checkbox" class="rounded"></td>
            <td class="px-4 py-3">
              <div class="flex items-center">
                <div class="w-20 h-12 bg-gray-100 rounded mr-3 relative overflow-hidden">
                  <img :src="video.cover" alt="" class="w-full h-full object-cover">
                  <div class="absolute inset-0 flex items-center justify-center">
                    <span class="text-white text-lg">▶</span>
                  </div>
                </div>
                <div>
                  <div class="font-medium">{{ video.name }}</div>
                  <div class="text-sm text-gray-500">ID: {{ video.id }}</div>
                </div>
              </div>
            </td>
            <td class="px-4 py-3 text-sm">{{ video.duration }}</td>
            <td class="px-4 py-3 text-sm">{{ video.resolution }}</td>
            <td class="px-4 py-3 text-sm">{{ video.size }}</td>
            <td class="px-4 py-3 text-sm">{{ video.source }}</td>
            <td class="px-4 py-3 text-sm">{{ video.useCount }}</td>
            <td class="px-4 py-3 text-sm">{{ video.createTime }}</td>
            <td class="px-4 py-3">
              <div class="flex space-x-2">
<button @click="handlePreview(video)" class="text-blue-600 hover:text-blue-800 text-sm">预览</button>
                <button @click="handleUse(video)" class="text-blue-600 hover:text-blue-800 text-sm">使用</button>
                <button @click="handleDelete(video)" class="text-red-600 hover:text-red-800 text-sm">删除</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="180" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const filters = ref({
  keyword: '',
  resolution: '',
  duration: ''
})

const videos = ref([
  { id: 'V001', name: '美妆直播精彩片段', cover: 'https://via.placeholder.com/160x90', duration: '00:15', resolution: '720x1280', size: '15.2MB', source: '直播录制', useCount: 12, createTime: '2024-03-15' },
  { id: 'V002', name: '新品发布宣传片', cover: 'https://via.placeholder.com/160x90', duration: '00:30', resolution: '1920x1080', size: '28.5MB', source: '本地上传', useCount: 8, createTime: '2024-03-14' },
  { id: 'V003', name: '达人口播推荐', cover: 'https://via.placeholder.com/160x90', duration: '00:20', resolution: '720x1280', size: '18.6MB', source: '抖音视频', useCount: 15, createTime: '2024-03-14' },
  { id: 'V004', name: '产品使用教程', cover: 'https://via.placeholder.com/160x90', duration: '00:45', resolution: '720x1280', size: '42.3MB', source: '本地上传', useCount: 6, createTime: '2024-03-13' },
  { id: 'V005', name: '用户好评合集', cover: 'https://via.placeholder.com/160x90', duration: '01:00', resolution: '1080x1080', size: '56.8MB', source: '本地上传', useCount: 4, createTime: '2024-03-12' }
])

const handleUpload = () => {
  alert('上传视频')
}

const handleSearch = () => {
  alert('搜索视频')
}

const handlePreview = (video: typeof videos.value[0]) => {
  alert(`预览视频: ${video.name}`)
}

const handleUse = (video: typeof videos.value[0]) => {
  alert(`使用视频: ${video.name}`)
}

const handleDelete = (video: typeof videos.value[0]) => {
  alert(`删除视频: ${video.name}`)
}
</script>
