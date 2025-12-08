<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 12, total: 156 })
const filterAccount = ref('')

const videos = ref([
  { id: 'V001', title: '新品发布会精彩回顾', account: '品牌官方号', plays: 125600, likes: 8560, comments: 356, shares: 189, time: '2025-11-25' },
  { id: 'V002', title: '产品使用教程', account: '产品测评号', plays: 89500, likes: 5230, comments: 215, shares: 98, time: '2025-11-26' },
  { id: 'V003', title: '双十一活动预告', account: '活动推广号', plays: 156800, likes: 12560, comments: 568, shares: 325, time: '2025-11-27' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}

const handlePublish = () => {
  alert('发布视频')
}

const handleView = (video: typeof videos.value[0]) => {
  alert(`查看视频: ${video.title}`)
}

const handleAnalyze = (video: typeof videos.value[0]) => {
  alert(`分析视频: ${video.title}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '抖音工具' }, { name: '视频管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">抖音视频管理</h1>
          <p class="mt-2 text-gray-600">管理抖音账号发布的视频</p>
        </div>
        <div class="flex gap-3">
          <select v-model="filterAccount" class="px-4 py-2 border border-gray-300 rounded-lg">
            <option value="">全部账号</option>
            <option value="AWM001">品牌官方号</option>
            <option value="AWM002">产品测评号</option>
          </select>
<button @click="handlePublish" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
            发布视频
          </button>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-5 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总视频数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总播放量</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">371.9万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总点赞数</p>
        <p class="text-2xl font-bold text-red-500 mt-1">26,350</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总评论数</p>
        <p class="text-2xl font-bold text-green-600 mt-1">1,139</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总分享数</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">612</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">视频信息</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">发布账号</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">播放量</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">点赞</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">评论</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">分享</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">发布时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="video in videos" :key="video.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="flex items-center">
                <div class="w-16 h-9 bg-gray-200 rounded flex items-center justify-center mr-3">
                  <span class="text-xl">🎬</span>
                </div>
                <div>
                  <div class="text-sm font-medium text-gray-900 max-w-xs truncate">{{ video.title }}</div>
                  <div class="text-xs text-gray-500">{{ video.id }}</div>
                </div>
              </div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ video.account }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ video.plays.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm text-red-500">{{ video.likes.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ video.comments }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ video.shares }}</td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ video.time }}</td>
            <td class="px-6 py-4 text-sm">
<button @click="handleView(video)" class="text-blue-600 hover:text-blue-800 mr-3">查看</button>
              <button @click="handleAnalyze(video)" class="text-gray-600 hover:text-gray-800">分析</button>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
