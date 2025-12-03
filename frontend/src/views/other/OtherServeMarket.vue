<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const activeCategory = ref('all')

const categories = [
  { key: 'all', label: '全部' },
  { key: 'creative', label: '创意服务' },
  { key: 'data', label: '数据服务' },
  { key: 'operation', label: '运营服务' },
  { key: 'tools', label: '工具服务' }
]

const services = ref([
  { id: 'SVC001', name: '智能创意生成', category: 'creative', provider: 'AI创意工坊', price: 299, priceUnit: '月', rating: 4.8, usersCount: 1256, icon: '🎨', description: 'AI驱动的广告创意素材生成工具' },
  { id: 'SVC002', name: '人群洞察报告', category: 'data', provider: '数据魔方', price: 599, priceUnit: '次', rating: 4.6, usersCount: 892, icon: '📊', description: '深度分析目标人群画像与行为特征' },
  { id: 'SVC003', name: '广告代运营', category: 'operation', provider: '效果营销团队', price: 5000, priceUnit: '月', rating: 4.9, usersCount: 456, icon: '🚀', description: '专业团队全程管理广告投放' },
  { id: 'SVC004', name: '批量素材处理', category: 'tools', provider: '效率工具集', price: 99, priceUnit: '月', rating: 4.5, usersCount: 2350, icon: '⚡', description: '一键批量处理图片视频素材' },
  { id: 'SVC005', name: '竞品分析服务', category: 'data', provider: '市场情报站', price: 999, priceUnit: '次', rating: 4.7, usersCount: 678, icon: '🔍', description: '全方位竞品广告投放策略分析' },
  { id: 'SVC006', name: '视频剪辑服务', category: 'creative', provider: '视觉工坊', price: 200, priceUnit: '条', rating: 4.4, usersCount: 1890, icon: '🎬', description: '专业视频剪辑与特效制作' }
])

const filteredServices = ref(services.value)

const filterByCategory = (category: string) => {
  activeCategory.value = category
  if (category === 'all') {
    filteredServices.value = services.value
  } else {
    filteredServices.value = services.value.filter(s => s.category === category)
  }
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '其他功能' }, { name: '服务市场' }]" />
      <h1 class="text-3xl font-bold text-gray-900">服务市场</h1>
      <p class="mt-2 text-gray-600">发现优质广告营销服务</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex items-center gap-4">
        <button v-for="cat in categories" :key="cat.key"
                :class="['px-4 py-2 rounded-lg text-sm transition-colors',
                  activeCategory === cat.key ? 'bg-blue-600 text-white' : 'bg-gray-100 text-gray-700 hover:bg-gray-200']"
                @click="filterByCategory(cat.key)">
          {{ cat.label }}
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div v-for="service in filteredServices" :key="service.id"
           class="bg-white rounded-lg border border-gray-200 overflow-hidden hover:shadow-lg transition-shadow">
        <div class="p-6">
          <div class="flex items-start gap-4">
            <div class="w-14 h-14 bg-blue-50 rounded-xl flex items-center justify-center text-3xl">
              {{ service.icon }}
            </div>
            <div class="flex-1">
              <h3 class="font-semibold text-gray-900">{{ service.name }}</h3>
              <p class="text-sm text-gray-500">{{ service.provider }}</p>
            </div>
          </div>
          <p class="mt-4 text-sm text-gray-600 line-clamp-2">{{ service.description }}</p>
          <div class="mt-4 flex items-center gap-4">
            <div class="flex items-center gap-1">
              <span class="text-yellow-400">★</span>
              <span class="text-sm font-medium text-gray-900">{{ service.rating }}</span>
            </div>
            <span class="text-sm text-gray-500">{{ service.usersCount }}+ 用户</span>
          </div>
        </div>
        <div class="px-6 py-4 bg-gray-50 border-t border-gray-100 flex items-center justify-between">
          <div>
            <span class="text-xl font-bold text-blue-600">¥{{ service.price }}</span>
            <span class="text-sm text-gray-500">/{{ service.priceUnit }}</span>
          </div>
          <button class="px-4 py-2 bg-blue-600 text-white rounded-lg text-sm hover:bg-blue-700">
            立即使用
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
