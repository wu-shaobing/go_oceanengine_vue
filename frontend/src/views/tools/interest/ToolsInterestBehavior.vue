<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const selectedBehaviors = ref<string[]>([])

const behaviors = ref([
  { id: 'B01', name: '电商购物', icon: '🛒', coverage: 350000000, children: ['网购达人', '促销敏感', '品质追求'] },
  { id: 'B02', name: '短视频观看', icon: '📱', coverage: 680000000, children: ['重度用户', '内容创作者', '互动活跃'] },
  { id: 'B03', name: '游戏娱乐', icon: '🎮', coverage: 280000000, children: ['手游玩家', '主机玩家', '电竞爱好者'] },
  { id: 'B04', name: '阅读学习', icon: '📚', coverage: 150000000, children: ['知识付费', '在线教育', '自我提升'] },
  { id: 'B05', name: '出行旅游', icon: '✈️', coverage: 120000000, children: ['商务出行', '休闲度假', '自驾游'] },
  { id: 'B06', name: '本地生活', icon: '🏠', coverage: 420000000, children: ['外卖用户', '到店消费', '生活服务'] }
])

const formatNumber = (num: number) => {
  if (num >= 100000000) return (num / 100000000).toFixed(1) + '亿'
  if (num >= 10000) return (num / 10000).toFixed(0) + '万'
  return num.toLocaleString()
}

const toggleBehavior = (id: string) => {
  const idx = selectedBehaviors.value.indexOf(id)
  if (idx > -1) selectedBehaviors.value.splice(idx, 1)
  else selectedBehaviors.value.push(id)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '兴趣定向' }, { name: '行为兴趣' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">行为兴趣定向</h1>
          <p class="mt-2 text-gray-600">基于用户行为进行人群定向</p>
        </div>
        <div class="text-sm text-gray-500">
          已选择 <span class="font-medium text-blue-600">{{ selectedBehaviors.length }}</span> 项
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <div v-for="item in behaviors" :key="item.id"
           :class="['bg-white rounded-lg border-2 p-4 cursor-pointer transition-all',
                    selectedBehaviors.includes(item.id) ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300']"
           @click="toggleBehavior(item.id)">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-3">
            <span class="text-3xl">{{ item.icon }}</span>
            <div>
              <h4 class="font-semibold text-gray-900">{{ item.name }}</h4>
              <p class="text-xs text-gray-500">覆盖 {{ formatNumber(item.coverage) }}</p>
            </div>
          </div>
          <div :class="['w-5 h-5 rounded border-2 flex items-center justify-center',
                        selectedBehaviors.includes(item.id) ? 'bg-blue-600 border-blue-600' : 'border-gray-300']">
            <svg v-if="selectedBehaviors.includes(item.id)" class="w-3 h-3 text-white" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"/>
            </svg>
          </div>
        </div>
        <div class="mt-3 flex flex-wrap gap-1">
          <span v-for="child in item.children" :key="child"
                class="px-2 py-0.5 bg-gray-100 text-gray-600 rounded text-xs">
            {{ child }}
          </span>
        </div>
      </div>
    </div>

    <div v-if="selectedBehaviors.length > 0" class="sticky bottom-4">
      <div class="bg-white rounded-lg border border-gray-200 shadow-lg p-4 flex items-center justify-between">
        <span class="text-gray-600">已选择 {{ selectedBehaviors.length }} 个行为标签</span>
        <button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          确认选择
        </button>
      </div>
    </div>
  </div>
</template>
