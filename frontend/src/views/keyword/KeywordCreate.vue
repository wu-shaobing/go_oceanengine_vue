<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const form = reactive({
  advertiserId: '',
  campaignId: '',
  keywords: '',
  matchType: 'phrase',
  bid: ''
})

const keywordList = ref<string[]>([])

const parseKeywords = () => {
  keywordList.value = form.keywords.split('\n').filter(k => k.trim())
}

const handleSubmit = () => {
  console.log('Submit keywords:', form, keywordList.value)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '关键词工具' }, { name: '创建关键词' }]" />
      <h1 class="text-3xl font-bold text-gray-900">创建关键词</h1>
      <p class="mt-2 text-gray-600">批量创建搜索广告关键词</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <div class="lg:col-span-2 bg-white rounded-lg border border-gray-200 p-6">
        <form @submit.prevent="handleSubmit" class="space-y-6">
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">广告主ID</label>
              <input v-model="form.advertiserId" type="text" placeholder="请输入广告主ID"
                     class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">广告组ID</label>
              <input v-model="form.campaignId" type="text" placeholder="请输入广告组ID"
                     class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">关键词列表</label>
            <textarea v-model="form.keywords" @blur="parseKeywords" rows="8" placeholder="每行输入一个关键词"
                      class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 font-mono text-sm"></textarea>
            <p class="mt-1 text-sm text-gray-500">已输入 {{ keywordList.length }} 个关键词</p>
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">匹配类型</label>
              <select v-model="form.matchType" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
                <option value="exact">精确匹配</option>
                <option value="phrase">短语匹配</option>
                <option value="broad">广泛匹配</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">出价 (元)</label>
              <input v-model="form.bid" type="number" step="0.01" placeholder="请输入出价"
                     class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
            </div>
          </div>
          <div class="flex gap-4 pt-4">
            <button type="submit" class="flex-1 px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
              创建关键词
            </button>
            <button type="button" class="px-6 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50">
              取消
            </button>
          </div>
        </form>
      </div>

      <div class="space-y-4">
        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h3 class="font-medium text-gray-900 mb-3">关键词预览</h3>
          <div v-if="keywordList.length" class="max-h-80 overflow-y-auto space-y-2">
            <div v-for="(kw, idx) in keywordList" :key="idx" class="flex items-center justify-between px-3 py-2 bg-gray-50 rounded">
              <span class="text-sm text-gray-700">{{ kw }}</span>
              <span class="text-xs text-gray-400">#{{ idx + 1 }}</span>
            </div>
          </div>
          <div v-else class="text-center py-8 text-gray-400 text-sm">
            在左侧输入关键词后显示预览
          </div>
        </div>
        <div class="bg-blue-50 border border-blue-200 rounded-lg p-4">
          <h4 class="font-medium text-blue-800 mb-2">匹配类型说明</h4>
          <ul class="text-sm text-blue-700 space-y-1">
            <li>• 精确匹配：仅匹配完全相同的搜索词</li>
            <li>• 短语匹配：匹配包含关键词的搜索</li>
            <li>• 广泛匹配：匹配相关的搜索词</li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>
