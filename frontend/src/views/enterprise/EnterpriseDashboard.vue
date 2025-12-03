<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '企业号', path: '/enterprise' }, { name: '工作台' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">企业号工作台</h1>
      <p class="text-gray-600 mt-1">企业号经营数据概览</p>
    </div>

    <!-- 核心指标 -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
      <div v-for="stat in coreStats" :key="stat.label" class="bg-white rounded-lg shadow p-4">
        <div class="text-sm text-gray-500">{{ stat.label }}</div>
        <div class="text-2xl font-bold text-gray-900 mt-1">{{ stat.value }}</div>
        <div :class="stat.trend > 0 ? 'text-green-500' : 'text-red-500'" class="text-sm mt-1">
          {{ stat.trend > 0 ? '+' : '' }}{{ stat.trend }}% 较昨日
        </div>
      </div>
    </div>

    <!-- 数据趋势 -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">粉丝增长趋势</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
          <span class="text-gray-400">粉丝增长图表</span>
        </div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-medium mb-4">内容互动趋势</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
          <span class="text-gray-400">互动趋势图表</span>
        </div>
      </div>
    </div>

    <!-- 内容表现 -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-lg font-medium">热门视频TOP5</h3>
          <router-link to="/enterprise/item" class="text-blue-600 text-sm">查看全部</router-link>
        </div>
        <div class="space-y-3">
          <div v-for="video in topVideos" :key="video.id" class="flex items-center p-2 hover:bg-gray-50 rounded">
            <img :src="video.cover" class="w-16 h-20 rounded object-cover mr-3" alt="">
            <div class="flex-1">
              <div class="font-medium line-clamp-2 text-sm">{{ video.title }}</div>
              <div class="text-xs text-gray-500 mt-1">播放 {{ video.plays }} · 点赞 {{ video.likes }}</div>
            </div>
          </div>
        </div>
      </div>
      <div class="bg-white rounded-lg shadow p-4">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-lg font-medium">待处理评论</h3>
          <router-link to="/enterprise/comment" class="text-blue-600 text-sm">查看全部</router-link>
        </div>
        <div class="space-y-3">
          <div v-for="comment in pendingComments" :key="comment.id" class="p-3 bg-gray-50 rounded">
            <div class="flex items-center mb-2">
              <img :src="comment.avatar" class="w-8 h-8 rounded-full mr-2" alt="">
              <span class="font-medium text-sm">{{ comment.user }}</span>
              <span class="text-xs text-gray-400 ml-2">{{ comment.time }}</span>
            </div>
            <div class="text-sm text-gray-700">{{ comment.content }}</div>
            <div class="flex space-x-2 mt-2">
              <button class="text-blue-600 text-xs">回复</button>
              <button class="text-gray-500 text-xs">隐藏</button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 快捷操作 -->
    <div class="bg-white rounded-lg shadow p-4">
      <h3 class="text-lg font-medium mb-4">快捷操作</h3>
      <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4">
        <router-link v-for="action in quickActions" :key="action.name" :to="action.path" 
          class="flex flex-col items-center p-4 rounded-lg hover:bg-gray-50 transition-colors">
          <div class="w-12 h-12 rounded-full bg-purple-100 flex items-center justify-center mb-2">
            <span class="text-purple-600">{{ action.icon }}</span>
          </div>
          <span class="text-sm text-gray-700">{{ action.name }}</span>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const coreStats = ref([
  { label: '总粉丝', value: '125.6w', trend: 2.5 },
  { label: '今日新增粉丝', value: '1,256', trend: 15.3 },
  { label: '视频总播放', value: '8,560w', trend: 8.2 },
  { label: '待处理评论', value: '68', trend: -12.5 }
])

const topVideos = ref([
  { id: 1, title: '新品发布会精彩回顾', cover: 'https://via.placeholder.com/64x80', plays: '125.6w', likes: '8.6w' },
  { id: 2, title: '产品使用教程合集', cover: 'https://via.placeholder.com/64x80', plays: '98.2w', likes: '6.2w' },
  { id: 3, title: '幕后花絮大公开', cover: 'https://via.placeholder.com/64x80', plays: '76.5w', likes: '5.1w' },
  { id: 4, title: '用户故事分享', cover: 'https://via.placeholder.com/64x80', plays: '65.8w', likes: '4.3w' },
  { id: 5, title: '限时优惠活动', cover: 'https://via.placeholder.com/64x80', plays: '52.3w', likes: '3.8w' }
])

const pendingComments = ref([
  { id: 1, user: '小明', avatar: 'https://via.placeholder.com/32', content: '产品质量真的很好，已经回购三次了！', time: '5分钟前' },
  { id: 2, user: '用户A', avatar: 'https://via.placeholder.com/32', content: '请问这款什么时候补货？', time: '15分钟前' },
  { id: 3, user: '小红', avatar: 'https://via.placeholder.com/32', content: '客服态度很好，点赞！', time: '30分钟前' }
])

const quickActions = ref([
  { name: '账号信息', path: '/enterprise/info', icon: '👤' },
  { name: '评论管理', path: '/enterprise/comment', icon: '💬' },
  { name: '视频列表', path: '/enterprise/item', icon: '🎬' },
  { name: '数据分析', path: '/enterprise/overview', icon: '📊' },
  { name: '绑定管理', path: '/enterprise/bind', icon: '🔗' },
  { name: '操作日志', path: '/enterprise/log', icon: '📋' }
])
</script>
