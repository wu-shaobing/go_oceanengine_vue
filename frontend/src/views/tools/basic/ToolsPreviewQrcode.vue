<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const form = reactive({
  advertiserId: '',
  adId: ''
})

const qrcodeUrl = ref('')
const loading = ref(false)

const handleGenerate = () => {
  loading.value = true
  setTimeout(() => {
    qrcodeUrl.value = 'data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIyMDAiIGhlaWdodD0iMjAwIj48cmVjdCB3aWR0aD0iMjAwIiBoZWlnaHQ9IjIwMCIgZmlsbD0iI2ZmZiIvPjxyZWN0IHg9IjIwIiB5PSIyMCIgd2lkdGg9IjE2MCIgaGVpZ2h0PSIxNjAiIGZpbGw9IiMwMDAiIG9wYWNpdHk9IjAuMSIvPjx0ZXh0IHg9IjEwMCIgeT0iMTEwIiBmb250LXNpemU9IjE0IiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmaWxsPSIjNjY2Ij7lupPlkYrpooTop4g8L3RleHQ+PC9zdmc+'
    loading.value = false
  }, 1000)
}

const handleDownload = () => {
  alert('下载二维码')
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '基础工具' }, { name: '广告预览' }]" />
      <h1 class="text-3xl font-bold text-gray-900">广告预览二维码</h1>
      <p class="mt-2 text-gray-600">生成广告预览二维码，扫码查看广告投放效果</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="font-semibold text-gray-900 mb-4">生成预览二维码</h3>
        <form @submit.prevent="handleGenerate" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">广告主ID</label>
            <input v-model="form.advertiserId" type="text" placeholder="请输入广告主ID"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">广告计划ID</label>
            <input v-model="form.adId" type="text" placeholder="请输入广告计划ID"
                   class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          <button type="submit" :disabled="loading"
                  class="w-full px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors disabled:bg-gray-400">
            {{ loading ? '生成中...' : '生成二维码' }}
          </button>
        </form>
      </div>

      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h3 class="font-semibold text-gray-900 mb-4">预览二维码</h3>
        <div v-if="qrcodeUrl" class="text-center">
          <div class="inline-block p-4 bg-white border-2 border-gray-200 rounded-lg">
            <img :src="qrcodeUrl" alt="广告预览二维码" class="w-48 h-48">
          </div>
          <p class="mt-4 text-sm text-gray-500">使用抖音/今日头条APP扫码预览</p>
          <p class="mt-2 text-xs text-gray-400">二维码有效期：24小时</p>
<button @click="handleDownload" class="mt-4 px-4 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50">
            下载二维码
          </button>
        </div>
        <div v-else class="text-center py-12 text-gray-500">
          请填写信息生成预览二维码
        </div>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h3 class="font-semibold text-gray-900 mb-3">使用说明</h3>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="text-center">
          <div class="w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center mx-auto mb-3">
            <span class="text-blue-600 font-bold">1</span>
          </div>
          <p class="text-sm text-gray-600">输入广告主ID和计划ID</p>
        </div>
        <div class="text-center">
          <div class="w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center mx-auto mb-3">
            <span class="text-blue-600 font-bold">2</span>
          </div>
          <p class="text-sm text-gray-600">点击生成预览二维码</p>
        </div>
        <div class="text-center">
          <div class="w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center mx-auto mb-3">
            <span class="text-blue-600 font-bold">3</span>
          </div>
          <p class="text-sm text-gray-600">使用APP扫码预览广告效果</p>
        </div>
      </div>
    </div>
  </div>
</template>
