<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const searchKeyword = ref('')

const keywords = ref([
  { id: 'IK001', keyword: '智能手机', category: '数码科技', coverage: 125000000, heat: 98 },
  { id: 'IK002', keyword: '健身减肥', category: '运动健康', coverage: 89000000, heat: 95 },
  { id: 'IK003', keyword: '美妆护肤', category: '时尚美妆', coverage: 112000000, heat: 96 },
  { id: 'IK004', keyword: '旅游攻略', category: '旅行出行', coverage: 78000000, heat: 89 },
  { id: 'IK005', keyword: '理财投资', category: '金融理财', coverage: 56000000, heat: 85 },
  { id: 'IK006', keyword: '亲子教育', category: '母婴育儿', coverage: 67000000, heat: 88 }
])

const formatNumber = (num: number) => {
  if (num >= 100000000) return (num / 100000000).toFixed(1) + '亿'
  if (num >= 10000) return (num / 10000).toFixed(0) + '万'
  return num.toLocaleString()
}

const handleSearch = () => {
  alert(`搜索: ${searchKeyword.value}`)
}

const handleSelect = (item: typeof keywords.value[0]) => {
  alert(`选择关键词: ${item.keyword}`)
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '兴趣定向' }, { name: '兴趣关键词' }]" />
      <h1 class="text-3xl font-bold text-gray-900">兴趣关键词</h1>
      <p class="mt-2 text-gray-600">搜索和选择用户兴趣关键词</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex gap-4">
        <input v-model="searchKeyword" type="text" placeholder="搜索关键词..."
               class="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
<button @click="handleSearch" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <div v-for="item in keywords" :key="item.id"
           class="bg-white rounded-lg border border-gray-200 p-4 hover:shadow-md transition-shadow">
        <div class="flex items-start justify-between">
          <div>
            <h4 class="font-semibold text-gray-900">{{ item.keyword }}</h4>
            <span class="text-xs text-gray-500">{{ item.category }}</span>
          </div>
<button @click="handleSelect(item)" class="px-3 py-1 text-sm text-blue-600 border border-blue-300 rounded hover:bg-blue-50">
            选择
          </button>
        </div>
        <div class="mt-4 flex items-center justify-between">
          <div>
            <p class="text-xs text-gray-500">覆盖人数</p>
            <p class="font-medium text-gray-900">{{ formatNumber(item.coverage) }}</p>
          </div>
          <div>
            <p class="text-xs text-gray-500">热度指数</p>
            <div class="flex items-center gap-2">
              <div class="w-20 h-1.5 bg-gray-200 rounded-full">
                <div class="h-full bg-orange-500 rounded-full" :style="{ width: item.heat + '%' }"></div>
              </div>
              <span class="text-sm font-medium text-orange-600">{{ item.heat }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
