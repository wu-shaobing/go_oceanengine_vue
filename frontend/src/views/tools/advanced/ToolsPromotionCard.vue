<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 36 })

const cards = ref([
  { id: 'PC001', name: '商品卡片', type: 'product', adsCount: 45, clicks: 125600, ctr: 3.8, status: 'active', preview: '🛍️' },
  { id: 'PC002', name: '应用下载卡', type: 'app', adsCount: 28, clicks: 89000, ctr: 4.2, status: 'active', preview: '📱' },
  { id: 'PC003', name: '表单收集卡', type: 'form', adsCount: 32, clicks: 56000, ctr: 2.9, status: 'active', preview: '📝' },
  { id: 'PC004', name: '电话咨询卡', type: 'phone', adsCount: 18, clicks: 34500, ctr: 3.1, status: 'active', preview: '📞' }
])

const getTypeConfig = (type: string) => {
  switch (type) {
    case 'product': return { label: '商品', class: 'bg-blue-100 text-blue-700' }
    case 'app': return { label: '应用', class: 'bg-green-100 text-green-700' }
    case 'form': return { label: '表单', class: 'bg-purple-100 text-purple-700' }
    default: return { label: '电话', class: 'bg-orange-100 text-orange-700' }
  }
}

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '高级工具' }, { name: '推广卡片' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">推广卡片管理</h1>
          <p class="mt-2 text-gray-600">创建和管理广告推广卡片</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建卡片
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总卡片</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">关联广告</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">123</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总点击</p>
        <p class="text-2xl font-bold text-green-600 mt-1">305K</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均CTR</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">3.5%</p>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
      <div v-for="card in cards" :key="card.id"
           class="bg-white rounded-lg border border-gray-200 overflow-hidden hover:shadow-lg transition-shadow">
        <div class="aspect-video bg-gradient-to-br from-gray-100 to-gray-200 flex items-center justify-center text-6xl">
          {{ card.preview }}
        </div>
        <div class="p-4">
          <div class="flex items-center justify-between mb-2">
            <h4 class="font-medium text-gray-900">{{ card.name }}</h4>
            <span :class="['px-2 py-0.5 rounded text-xs', getTypeConfig(card.type).class]">
              {{ getTypeConfig(card.type).label }}
            </span>
          </div>
          <div class="grid grid-cols-3 gap-2 text-center text-sm mt-4">
            <div>
              <p class="text-gray-500">广告</p>
              <p class="font-semibold text-gray-900">{{ card.adsCount }}</p>
            </div>
            <div>
              <p class="text-gray-500">点击</p>
              <p class="font-semibold text-gray-900">{{ (card.clicks / 1000).toFixed(0) }}K</p>
            </div>
            <div>
              <p class="text-gray-500">CTR</p>
              <p class="font-semibold text-green-600">{{ card.ctr }}%</p>
            </div>
          </div>
          <div class="flex gap-2 mt-4">
            <button class="flex-1 py-2 text-sm text-blue-600 border border-blue-300 rounded hover:bg-blue-50">编辑</button>
            <button class="flex-1 py-2 text-sm text-gray-600 border border-gray-300 rounded hover:bg-gray-50">数据</button>
          </div>
        </div>
      </div>
    </div>

    <div class="flex justify-center">
      <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
    </div>
  </div>
</template>
