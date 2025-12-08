<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '企业号', path: '/enterprise' }, { name: '账号信息' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">账号信息</h1>
      <p class="text-gray-600 mt-1">查看和管理企业号账号信息</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <div class="lg:col-span-2 space-y-6">
        <!-- 基本信息 -->
        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">基本信息</h3>
          <div class="flex items-start">
            <img :src="accountInfo.avatar" class="w-20 h-20 rounded-full mr-6" alt="">
            <div class="flex-1 space-y-3">
              <div class="flex items-center">
                <span class="text-gray-500 w-24">账号名称</span>
                <span class="font-medium">{{ accountInfo.name }}</span>
                <span class="ml-2 px-2 py-0.5 text-xs bg-blue-100 text-blue-800 rounded">已认证</span>
              </div>
              <div class="flex items-center">
                <span class="text-gray-500 w-24">抖音号</span>
                <span>{{ accountInfo.douyinId }}</span>
              </div>
              <div class="flex items-center">
                <span class="text-gray-500 w-24">账号类型</span>
                <span>{{ accountInfo.type }}</span>
              </div>
              <div class="flex items-center">
                <span class="text-gray-500 w-24">行业分类</span>
                <span>{{ accountInfo.industry }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- 认证信息 -->
        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">认证信息</h3>
          <div class="space-y-3">
            <div class="flex items-center">
              <span class="text-gray-500 w-24">企业名称</span>
              <span>{{ accountInfo.companyName }}</span>
            </div>
            <div class="flex items-center">
              <span class="text-gray-500 w-24">认证状态</span>
              <span class="px-2 py-0.5 text-xs bg-green-100 text-green-800 rounded">已认证</span>
            </div>
            <div class="flex items-center">
              <span class="text-gray-500 w-24">认证到期</span>
              <span>{{ accountInfo.certExpire }}</span>
            </div>
          </div>
        </div>

        <!-- 权限信息 -->
        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">权限信息</h3>
          <div class="grid grid-cols-2 md:grid-cols-3 gap-4">
            <div v-for="perm in permissions" :key="perm.name" class="flex items-center p-3 bg-gray-50 rounded">
              <span :class="perm.enabled ? 'text-green-500' : 'text-gray-300'" class="mr-2">●</span>
              <span class="text-sm">{{ perm.name }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- 数据概览 -->
      <div class="space-y-6">
        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">数据概览</h3>
          <div class="space-y-4">
            <div class="flex justify-between">
              <span class="text-gray-500">粉丝数</span>
              <span class="font-medium">{{ accountInfo.followers }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">获赞数</span>
              <span class="font-medium">{{ accountInfo.likes }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">视频数</span>
              <span class="font-medium">{{ accountInfo.videos }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">关注数</span>
              <span class="font-medium">{{ accountInfo.following }}</span>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <h3 class="text-lg font-medium mb-4">快捷操作</h3>
          <div class="space-y-2">
            <button @click="handleChangeAvatar" class="w-full text-left px-4 py-2 hover:bg-gray-50 rounded text-sm">修改头像</button>
            <button @click="handleChangeBio" class="w-full text-left px-4 py-2 hover:bg-gray-50 rounded text-sm">修改简介</button>
            <button @click="handleBindManage" class="w-full text-left px-4 py-2 hover:bg-gray-50 rounded text-sm">绑定管理</button>
            <button @click="handleUnbindAccount" class="w-full text-left px-4 py-2 hover:bg-gray-50 rounded text-sm text-red-600">解绑账号</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const handleChangeAvatar = () => {
  alert('修改头像')
}

const handleChangeBio = () => {
  alert('修改简介')
}

const handleBindManage = () => {
  alert('绑定管理')
}

const handleUnbindAccount = () => {
  if (confirm('确定解绑账号吗？')) {
    alert('已解绑账号')
  }
}

const accountInfo = ref({
  avatar: 'https://via.placeholder.com/80',
  name: '品牌官方旗舰店',
  douyinId: 'brand_official',
  type: '企业蓝V',
  industry: '电商 / 品牌零售',
  companyName: '某某科技有限公司',
  certExpire: '2025-12-31',
  followers: '125.6万',
  likes: '856.2万',
  videos: '286',
  following: '56'
})

const permissions = ref([
  { name: '视频发布', enabled: true },
  { name: '评论管理', enabled: true },
  { name: '私信管理', enabled: true },
  { name: '数据分析', enabled: true },
  { name: '粉丝管理', enabled: true },
  { name: '直播管理', enabled: true },
  { name: '商品橱窗', enabled: true },
  { name: '卡券管理', enabled: false },
  { name: '小程序', enabled: false }
])
</script>
