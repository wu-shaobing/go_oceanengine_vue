<script setup lang="ts">
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const route = useRoute()
const siteId = route.params.id || 'SITE001'
const previewMode = ref<'mobile' | 'desktop'>('mobile')

const site = ref({
  id: siteId,
  name: '双11促销活动页',
  url: 'https://landing.example.com/promo',
  status: 'published',
  lastUpdated: '2025-11-25 14:30'
})

const handlePublish = () => {
  alert('发布页面')
}

const handleCopyLink = () => {
  alert(`已复制链接: ${site.value.url}`)
}

const handleViewData = () => {
  alert('查看数据')
}

const handleGenerateQR = () => {
  alert('生成二维码')
}

const handleExportPage = () => {
  alert('导出页面')
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '落地页' }, { name: '页面预览' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">{{ site.name }}</h1>
          <p class="mt-2 text-gray-600">预览落地页效果</p>
        </div>
        <div class="flex items-center gap-4">
          <div class="flex border border-gray-300 rounded-lg overflow-hidden">
            <button @click="previewMode = 'mobile'" 
                    :class="['px-4 py-2 text-sm', previewMode === 'mobile' ? 'bg-blue-50 text-blue-600' : 'bg-white text-gray-600']">
              📱 手机
            </button>
            <button @click="previewMode = 'desktop'" 
                    :class="['px-4 py-2 text-sm', previewMode === 'desktop' ? 'bg-blue-50 text-blue-600' : 'bg-white text-gray-600']">
              🖥️ 电脑
            </button>
          </div>
          <router-link :to="`/site/edit/${site.id}`" class="px-4 py-2 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50">
            编辑
          </router-link>
          <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="handlePublish">
            发布
          </button>
        </div>
      </div>
    </div>

    <div class="flex gap-6">
      <div class="flex-1 bg-gray-100 rounded-lg p-8 flex justify-center">
        <div :class="[
          'bg-white rounded-lg shadow-lg overflow-hidden transition-all',
          previewMode === 'mobile' ? 'w-[375px]' : 'w-full max-w-4xl'
        ]">
          <div class="bg-gray-800 px-4 py-2 flex items-center gap-2">
            <div class="flex gap-1.5">
              <div class="w-3 h-3 rounded-full bg-red-500"></div>
              <div class="w-3 h-3 rounded-full bg-yellow-500"></div>
              <div class="w-3 h-3 rounded-full bg-green-500"></div>
            </div>
            <div class="flex-1 bg-gray-700 rounded px-3 py-1 text-xs text-gray-300 truncate">
              {{ site.url }}
            </div>
          </div>
          <div :class="['bg-gradient-to-b from-blue-500 to-purple-600', previewMode === 'mobile' ? 'h-[667px]' : 'h-[500px]']">
            <div class="p-6 text-white text-center">
              <h2 class="text-2xl font-bold mt-8">🎉 双11狂欢节</h2>
              <p class="mt-2 text-white/80">限时特惠 低至5折</p>
              <div class="mt-8 bg-white/20 rounded-lg p-6 backdrop-blur">
                <p class="text-lg">倒计时</p>
                <p class="text-4xl font-bold mt-2">02:15:36</p>
              </div>
              <button class="mt-8 px-8 py-3 bg-yellow-400 text-gray-900 rounded-full font-bold">
                立即抢购
              </button>
            </div>
          </div>
        </div>
      </div>

      <div class="w-80 space-y-4">
        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h3 class="font-semibold text-gray-900 mb-4">页面信息</h3>
          <div class="space-y-3 text-sm">
            <div class="flex justify-between">
              <span class="text-gray-500">页面ID</span>
              <span class="text-gray-900">{{ site.id }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">状态</span>
              <span class="px-2 py-0.5 bg-green-100 text-green-700 rounded text-xs">已发布</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-500">最后更新</span>
              <span class="text-gray-900">{{ site.lastUpdated }}</span>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg border border-gray-200 p-4">
          <h3 class="font-semibold text-gray-900 mb-4">快捷操作</h3>
          <div class="space-y-2">
            <button class="w-full px-4 py-2 text-left text-sm text-gray-700 hover:bg-gray-50 rounded" @click="handleCopyLink">📋 复制链接</button>
            <button class="w-full px-4 py-2 text-left text-sm text-gray-700 hover:bg-gray-50 rounded" @click="handleViewData">📊 查看数据</button>
            <button class="w-full px-4 py-2 text-left text-sm text-gray-700 hover:bg-gray-50 rounded" @click="handleGenerateQR">🔄 生成二维码</button>
            <button class="w-full px-4 py-2 text-left text-sm text-gray-700 hover:bg-gray-50 rounded" @click="handleExportPage">📥 导出页面</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
