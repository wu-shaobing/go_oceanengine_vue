<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const form = reactive({
  advertiserId: '',
  adIds: ''
})

const results = ref<Array<{
  adId: string
  qualityScore: number
  creativityScore: number
  relevanceScore: number
  landingScore: number
  suggestion: string
}>>([])

const handleQuery = () => {
  results.value = [
    { adId: '174812345678', qualityScore: 85, creativityScore: 90, relevanceScore: 80, landingScore: 85, suggestion: '创意质量较好，可适当提升出价' },
    { adId: '174812345679', qualityScore: 65, creativityScore: 60, relevanceScore: 70, landingScore: 65, suggestion: '建议优化创意素材，提升吸引力' },
    { adId: '174812345680', qualityScore: 92, creativityScore: 95, relevanceScore: 90, landingScore: 91, suggestion: '广告质量优秀，保持当前策略' }
  ]
}

const getScoreColor = (score: number) => {
  if (score >= 80) return 'text-green-600'
  if (score >= 60) return 'text-yellow-600'
  return 'text-red-600'
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '基础工具' }, { name: '广告质量度' }]" />
      <h1 class="text-3xl font-bold text-gray-900">广告质量度查询</h1>
      <p class="mt-2 text-gray-600">查询广告计划的质量度评分及优化建议</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <form @submit.prevent="handleQuery" class="flex flex-wrap gap-4">
        <div class="flex-1 min-w-[200px]">
          <input v-model="form.advertiserId" type="text" placeholder="广告主ID"
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        </div>
        <div class="flex-1 min-w-[300px]">
          <input v-model="form.adIds" type="text" placeholder="广告计划ID（多个用逗号分隔）"
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        </div>
        <button type="submit" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
          查询
        </button>
      </form>
    </div>

    <div v-if="results.length" class="space-y-4">
      <div v-for="item in results" :key="item.adId" class="bg-white rounded-lg border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <div>
            <h3 class="font-semibold text-gray-900">广告计划 {{ item.adId }}</h3>
            <p class="text-sm text-gray-500 mt-1">{{ item.suggestion }}</p>
          </div>
          <div class="text-right">
            <p class="text-sm text-gray-500">综合质量度</p>
            <p :class="['text-3xl font-bold', getScoreColor(item.qualityScore)]">{{ item.qualityScore }}</p>
          </div>
        </div>
        <div class="grid grid-cols-3 gap-4">
          <div class="text-center p-4 bg-gray-50 rounded-lg">
            <p class="text-sm text-gray-500 mb-1">创意质量</p>
            <p :class="['text-xl font-bold', getScoreColor(item.creativityScore)]">{{ item.creativityScore }}</p>
            <div class="w-full bg-gray-200 rounded-full h-2 mt-2">
              <div class="bg-blue-500 h-2 rounded-full" :style="{ width: item.creativityScore + '%' }"></div>
            </div>
          </div>
          <div class="text-center p-4 bg-gray-50 rounded-lg">
            <p class="text-sm text-gray-500 mb-1">相关性</p>
            <p :class="['text-xl font-bold', getScoreColor(item.relevanceScore)]">{{ item.relevanceScore }}</p>
            <div class="w-full bg-gray-200 rounded-full h-2 mt-2">
              <div class="bg-green-500 h-2 rounded-full" :style="{ width: item.relevanceScore + '%' }"></div>
            </div>
          </div>
          <div class="text-center p-4 bg-gray-50 rounded-lg">
            <p class="text-sm text-gray-500 mb-1">落地页体验</p>
            <p :class="['text-xl font-bold', getScoreColor(item.landingScore)]">{{ item.landingScore }}</p>
            <div class="w-full bg-gray-200 rounded-full h-2 mt-2">
              <div class="bg-purple-500 h-2 rounded-full" :style="{ width: item.landingScore + '%' }"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div v-else class="bg-white rounded-lg border border-gray-200 p-12 text-center text-gray-500">
      请输入广告主ID和广告计划ID进行查询
    </div>
  </div>
</template>
