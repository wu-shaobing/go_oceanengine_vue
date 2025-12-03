<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const searchForm = reactive({
  awemeId: ''
})

const awemeInfo = ref<null | {
  id: string
  nickname: string
  avatar: string
  signature: string
  followerCount: number
  followingCount: number
  likesCount: number
  videoCount: number
  verified: boolean
  verifyName: string
  tags: string[]
}>(null)

const handleSearch = () => {
  awemeInfo.value = {
    id: searchForm.awemeId || '123456789',
    nickname: '品牌官方账号',
    avatar: '👤',
    signature: '分享生活美学，传递品质生活',
    followerCount: 1256000,
    followingCount: 128,
    likesCount: 8520000,
    videoCount: 356,
    verified: true,
    verifyName: '官方认证账号',
    tags: ['品牌', '生活', '时尚']
  }
}

const formatNumber = (num: number) => {
  if (num >= 10000000) return (num / 10000000).toFixed(1) + '千万'
  if (num >= 10000) return (num / 10000).toFixed(1) + '万'
  return num.toLocaleString()
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '抖音工具' }, { name: '账号信息' }]" />
      <h1 class="text-3xl font-bold text-gray-900">抖音账号信息</h1>
      <p class="mt-2 text-gray-600">查询抖音账号详细信息</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <form @submit.prevent="handleSearch" class="flex gap-4">
        <input v-model="searchForm.awemeId" type="text" placeholder="输入抖音号或账号ID"
               class="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        <button type="submit" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          查询
        </button>
      </form>
    </div>

    <div v-if="awemeInfo" class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <div class="text-center">
          <div class="w-24 h-24 mx-auto bg-gray-100 rounded-full flex items-center justify-center text-5xl">
            {{ awemeInfo.avatar }}
          </div>
          <h2 class="mt-4 text-xl font-bold text-gray-900">{{ awemeInfo.nickname }}</h2>
          <p class="text-sm text-gray-500 mt-1">抖音号: {{ awemeInfo.id }}</p>
          <div v-if="awemeInfo.verified" class="mt-2 inline-flex items-center gap-1 px-2 py-1 bg-blue-50 text-blue-600 rounded text-xs">
            <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"/>
            </svg>
            {{ awemeInfo.verifyName }}
          </div>
          <p class="mt-4 text-sm text-gray-600">{{ awemeInfo.signature }}</p>
          <div class="flex flex-wrap gap-2 justify-center mt-4">
            <span v-for="tag in awemeInfo.tags" :key="tag" 
                  class="px-2 py-1 bg-gray-100 text-gray-600 rounded text-xs">
              {{ tag }}
            </span>
          </div>
        </div>
      </div>

      <div class="lg:col-span-2 grid grid-cols-2 gap-4">
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <p class="text-sm text-gray-500">粉丝数</p>
          <p class="text-3xl font-bold text-gray-900 mt-2">{{ formatNumber(awemeInfo.followerCount) }}</p>
        </div>
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <p class="text-sm text-gray-500">关注数</p>
          <p class="text-3xl font-bold text-gray-900 mt-2">{{ formatNumber(awemeInfo.followingCount) }}</p>
        </div>
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <p class="text-sm text-gray-500">获赞数</p>
          <p class="text-3xl font-bold text-red-600 mt-2">{{ formatNumber(awemeInfo.likesCount) }}</p>
        </div>
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <p class="text-sm text-gray-500">作品数</p>
          <p class="text-3xl font-bold text-blue-600 mt-2">{{ awemeInfo.videoCount }}</p>
        </div>
      </div>
    </div>

    <div v-else class="bg-white rounded-lg border border-gray-200 p-12 text-center">
      <div class="text-6xl mb-4">🔍</div>
      <p class="text-gray-500">输入抖音号查询账号信息</p>
    </div>
  </div>
</template>
