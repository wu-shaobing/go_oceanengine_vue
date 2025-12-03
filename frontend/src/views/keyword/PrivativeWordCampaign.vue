<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const form = reactive({
  campaignId: '',
  negativeWords: ''
})

const existingWords = ref([
  { id: 1, word: '免费', matchType: '精确' },
  { id: 2, word: '便宜', matchType: '短语' },
  { id: 3, word: '下载', matchType: '广泛' }
])

const handleSubmit = () => {
  console.log('Submit negative words:', form)
}

const handleDelete = (id: number) => {
  existingWords.value = existingWords.value.filter(w => w.id !== id)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '关键词工具' }, { name: '组否定词' }]" />
      <h1 class="text-3xl font-bold text-gray-900">广告组否定词管理</h1>
      <p class="mt-2 text-gray-600">管理广告组级别的否定关键词</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="font-semibold text-gray-900 mb-4">添加否定词</h3>
        <form @submit.prevent="handleSubmit" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">广告组ID</label>
            <input v-model="form.campaignId" type="text" placeholder="请输入广告组ID"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">否定关键词</label>
            <textarea v-model="form.negativeWords" rows="6" placeholder="每行输入一个否定词"
                      class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 font-mono text-sm"></textarea>
          </div>
          <button type="submit" class="w-full px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
            添加否定词
          </button>
        </form>
      </div>

      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="font-semibold text-gray-900 mb-4">已有否定词 ({{ existingWords.length }})</h3>
        <div class="space-y-2 max-h-80 overflow-y-auto">
          <div v-for="word in existingWords" :key="word.id" 
               class="flex items-center justify-between px-4 py-3 bg-gray-50 rounded-lg">
            <div class="flex items-center gap-3">
              <span class="text-gray-900">{{ word.word }}</span>
              <span class="px-2 py-0.5 bg-blue-100 text-blue-700 rounded text-xs">{{ word.matchType }}</span>
            </div>
            <button @click="handleDelete(word.id)" class="text-red-500 hover:text-red-700">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/>
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
