<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const categories = ref([
  { id: 'C01', name: '生活', icon: '🏠', count: 15600, children: ['美食', '旅行', '家居', '宠物'] },
  { id: 'C02', name: '娱乐', icon: '🎬', count: 23400, children: ['搞笑', '剧情', '音乐', '舞蹈'] },
  { id: 'C03', name: '知识', icon: '📚', count: 8900, children: ['科普', '教育', '财经', '职场'] },
  { id: 'C04', name: '科技', icon: '📱', count: 6700, children: ['数码', '汽车', '游戏', '互联网'] },
  { id: 'C05', name: '时尚', icon: '👗', count: 12300, children: ['穿搭', '美妆', '护肤', '发型'] },
  { id: 'C06', name: '体育', icon: '⚽', count: 5600, children: ['健身', '篮球', '足球', '户外'] }
])

const expandedCategories = ref<string[]>(['C01', 'C04'])

const toggleExpand = (id: string) => {
  const idx = expandedCategories.value.indexOf(id)
  if (idx > -1) {
    expandedCategories.value.splice(idx, 1)
  } else {
    expandedCategories.value.push(id)
  }
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '抖音工具' }, { name: '内容分类' }]" />
      <h1 class="text-3xl font-bold text-gray-900">抖音内容分类</h1>
      <p class="mt-2 text-gray-600">查询抖音内容分类标签</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <input type="text" placeholder="搜索分类..."
             class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <div v-for="cat in categories" :key="cat.id"
           class="bg-white rounded-lg border border-gray-200 overflow-hidden">
        <div class="p-4 bg-gray-50 border-b border-gray-200 cursor-pointer hover:bg-gray-100"
             @click="toggleExpand(cat.id)">
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-3">
              <span class="text-3xl">{{ cat.icon }}</span>
              <div>
                <h4 class="font-semibold text-gray-900">{{ cat.name }}</h4>
                <p class="text-xs text-gray-500">{{ cat.count.toLocaleString() }} 个标签</p>
              </div>
            </div>
            <svg :class="['w-5 h-5 text-gray-400 transition-transform',
                         expandedCategories.includes(cat.id) ? 'rotate-180' : '']"
                 fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
            </svg>
          </div>
        </div>
        <div v-if="expandedCategories.includes(cat.id)" class="p-4">
          <div class="flex flex-wrap gap-2">
            <span v-for="child in cat.children" :key="child"
                  class="px-3 py-1.5 bg-blue-50 text-blue-700 rounded-full text-sm cursor-pointer hover:bg-blue-100">
              {{ child }}
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
