<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const category = ref('all')

const templates = ref([
  { id: 'TPL001', name: '电商促销模板', category: 'promo', preview: '🛍️', useCount: 1256, rating: 4.8 },
  { id: 'TPL002', name: '品牌展示模板', category: 'brand', preview: '🏢', useCount: 890, rating: 4.6 },
  { id: 'TPL003', name: '表单收集模板', category: 'form', preview: '📝', useCount: 2340, rating: 4.9 },
  { id: 'TPL004', name: '产品详情模板', category: 'product', preview: '📦', useCount: 756, rating: 4.5 },
  { id: 'TPL005', name: '活动报名模板', category: 'event', preview: '🎉', useCount: 567, rating: 4.7 },
  { id: 'TPL006', name: 'APP下载模板', category: 'app', preview: '📱', useCount: 1890, rating: 4.8 }
])

const categories = [
  { value: 'all', label: '全部' },
  { value: 'promo', label: '促销活动' },
  { value: 'brand', label: '品牌展示' },
  { value: 'form', label: '表单收集' },
  { value: 'product', label: '产品展示' }
]

const filteredTemplates = () => {
  if (category.value === 'all') return templates.value
  return templates.value.filter(t => t.category === category.value)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '落地页管理' }, { name: '模板库' }]" />
      <h1 class="text-3xl font-bold text-gray-900">落地页模板</h1>
      <p class="mt-2 text-gray-600">选择模板快速创建落地页</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex gap-2">
        <button v-for="cat in categories" :key="cat.value"
                :class="['px-4 py-2 rounded-lg', category === cat.value ? 'bg-blue-600 text-white' : 'bg-gray-100 text-gray-700']"
                @click="category = cat.value">
          {{ cat.label }}
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div v-for="tpl in filteredTemplates()" :key="tpl.id"
           class="bg-white rounded-lg border border-gray-200 overflow-hidden hover:shadow-lg transition-shadow">
        <div class="aspect-video bg-gradient-to-br from-blue-50 to-purple-50 flex items-center justify-center">
          <span class="text-6xl">{{ tpl.preview }}</span>
        </div>
        <div class="p-4">
          <h4 class="font-semibold text-gray-900">{{ tpl.name }}</h4>
          <div class="flex items-center gap-4 mt-2 text-sm text-gray-500">
            <span>使用 {{ tpl.useCount.toLocaleString() }} 次</span>
            <span class="flex items-center gap-1">
              ⭐ {{ tpl.rating }}
            </span>
          </div>
          <div class="flex gap-2 mt-4">
            <button class="flex-1 py-2 text-sm text-blue-600 border border-blue-300 rounded hover:bg-blue-50">
              预览
            </button>
            <button class="flex-1 py-2 text-sm text-white bg-blue-600 rounded hover:bg-blue-700">
              使用模板
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
