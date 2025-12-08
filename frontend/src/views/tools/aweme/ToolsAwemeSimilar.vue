<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const searchKeyword = ref('')

const similarAuthors = ref([
  { id: 'AU001', name: '科技小王子', avatar: '👨', followers: 2560000, category: '数码科技', similarity: 92, videos: 156 },
  { id: 'AU002', name: '数码测评君', avatar: '👤', followers: 1890000, category: '数码科技', similarity: 88, videos: 234 },
  { id: 'AU003', name: '极客玩家', avatar: '🧑', followers: 3200000, category: '数码科技', similarity: 85, videos: 312 },
  { id: 'AU004', name: '科技潮流', avatar: '👦', followers: 980000, category: '数码科技', similarity: 82, videos: 89 },
  { id: 'AU005', name: '智能生活家', avatar: '🧔', followers: 1560000, category: '智能家居', similarity: 78, videos: 178 },
  { id: 'AU006', name: '数码发烧友', avatar: '👱', followers: 2100000, category: '数码科技', similarity: 75, videos: 267 }
])

const formatNumber = (num: number) => {
  if (num >= 10000) return (num / 10000).toFixed(1) + '万'
  return num.toLocaleString()
}

const handleSearch = () => {
  alert(`搜索: ${searchKeyword.value}`)
}

const handleViewDetail = (author: typeof similarAuthors.value[0]) => {
  alert(`查看详情: ${author.name}`)
}

const handleAddTarget = (author: typeof similarAuthors.value[0]) => {
  alert(`添加定向: ${author.name}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '抖音工具' }, { name: '相似达人' }]" />
      <h1 class="text-3xl font-bold text-gray-900">相似达人发现</h1>
      <p class="mt-2 text-gray-600">发现与目标账号相似的抖音达人</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex gap-4">
        <input v-model="searchKeyword" type="text" placeholder="输入抖音号或达人名称"
               class="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
<button @click="handleSearch" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          搜索相似达人
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <div v-for="author in similarAuthors" :key="author.id"
           class="bg-white rounded-lg border border-gray-200 p-4 hover:shadow-lg transition-shadow">
        <div class="flex items-start gap-4">
          <div class="w-16 h-16 bg-gray-100 rounded-full flex items-center justify-center text-3xl">
            {{ author.avatar }}
          </div>
          <div class="flex-1">
            <div class="flex items-center justify-between">
              <h4 class="font-semibold text-gray-900">{{ author.name }}</h4>
              <span class="px-2 py-1 bg-green-100 text-green-700 rounded text-xs font-medium">
                {{ author.similarity }}% 相似
              </span>
            </div>
            <p class="text-sm text-gray-500 mt-1">{{ author.category }}</p>
            <div class="flex items-center gap-4 mt-3 text-sm">
              <span class="text-gray-600">
                <span class="font-medium text-gray-900">{{ formatNumber(author.followers) }}</span> 粉丝
              </span>
              <span class="text-gray-600">
                <span class="font-medium text-gray-900">{{ author.videos }}</span> 作品
              </span>
            </div>
          </div>
        </div>
<div class="flex gap-2 mt-4">
          <button @click="handleViewDetail(author)" class="flex-1 py-2 text-sm text-blue-600 border border-blue-300 rounded hover:bg-blue-50">
            查看详情
          </button>
          <button @click="handleAddTarget(author)" class="flex-1 py-2 text-sm text-white bg-blue-600 rounded hover:bg-blue-700">
            添加定向
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
