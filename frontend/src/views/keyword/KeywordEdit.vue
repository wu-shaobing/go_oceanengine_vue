<script setup lang="ts">
import { reactive } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const route = useRoute()
const router = useRouter()
const keywordId = route.params.id || 'KW001'

const form = reactive({
  id: keywordId,
  keyword: '智能手表',
  matchType: 'phrase',
  bid: 2.5,
  status: 'active',
  negativeKeywords: ['便宜', '山寨']
})

const matchTypes = [
  { value: 'exact', label: '精确匹配' },
  { value: 'phrase', label: '短语匹配' },
  { value: 'broad', label: '广泛匹配' }
]

const handleSubmit = () => {
  console.log('Update keyword:', form)
  router.push('/keyword/list')
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '关键词' }, { name: '编辑关键词' }]" />
      <h1 class="text-3xl font-bold text-gray-900">编辑关键词</h1>
      <p class="mt-2 text-gray-600">修改关键词配置</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <form @submit.prevent="handleSubmit" class="space-y-6 max-w-2xl">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">关键词</label>
          <input v-model="form.keyword" type="text" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">匹配类型</label>
          <div class="flex gap-4">
            <label v-for="type in matchTypes" :key="type.value" class="flex items-center gap-2 cursor-pointer">
              <input v-model="form.matchType" type="radio" :value="type.value" class="text-blue-600">
              <span class="text-gray-700">{{ type.label }}</span>
            </label>
          </div>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">出价 (元)</label>
          <input v-model.number="form.bid" type="number" step="0.1" min="0.1"
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">状态</label>
          <select v-model="form.status" class="w-full px-4 py-2 border border-gray-300 rounded-lg">
            <option value="active">启用</option>
            <option value="paused">暂停</option>
          </select>
        </div>
        <div class="flex justify-end gap-4 pt-4">
          <button type="button" @click="router.back()" class="px-6 py-2 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50">取消</button>
          <button type="submit" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">保存</button>
        </div>
      </form>
    </div>
  </div>
</template>
