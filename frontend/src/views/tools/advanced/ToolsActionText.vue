<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 12, total: 86 })

const actionTexts = ref([
  { id: 'AT001', text: '立即购买', category: '电商', usedCount: 12560, ctr: 3.2 },
  { id: 'AT002', text: '免费领取', category: '活动', usedCount: 8900, ctr: 4.1 },
  { id: 'AT003', text: '限时抢购', category: '促销', usedCount: 7650, ctr: 3.8 },
  { id: 'AT004', text: '了解更多', category: '品牌', usedCount: 5600, ctr: 2.4 },
  { id: 'AT005', text: '立即下载', category: 'APP', usedCount: 9800, ctr: 3.5 },
  { id: 'AT006', text: '预约咨询', category: '服务', usedCount: 4500, ctr: 2.8 },
  { id: 'AT007', text: '马上体验', category: '游戏', usedCount: 6700, ctr: 3.6 },
  { id: 'AT008', text: '一键购买', category: '电商', usedCount: 3400, ctr: 3.0 }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '高级工具' }, { name: '行动号召' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">行动号召文案</h1>
          <p class="mt-2 text-gray-600">选择合适的行动号召提升点击率</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          自定义文案
        </button>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex flex-wrap gap-4">
        <select class="px-4 py-2 border border-gray-300 rounded-lg">
          <option value="">全部分类</option>
          <option value="ecommerce">电商</option>
          <option value="activity">活动</option>
          <option value="brand">品牌</option>
          <option value="app">APP</option>
          <option value="game">游戏</option>
        </select>
        <input type="text" placeholder="搜索文案..." class="flex-1 min-w-[200px] px-4 py-2 border border-gray-300 rounded-lg">
      </div>
    </div>

    <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
      <div v-for="item in actionTexts" :key="item.id"
           class="bg-white rounded-lg border border-gray-200 p-4 hover:shadow-lg transition-shadow cursor-pointer group">
        <div class="flex items-center justify-between mb-3">
          <span class="px-2 py-1 bg-blue-100 text-blue-700 rounded text-xs">{{ item.category }}</span>
          <span class="text-xs text-gray-400">{{ item.id }}</span>
        </div>
        <div class="py-4 text-center">
          <span class="px-6 py-2 bg-gradient-to-r from-blue-500 to-purple-500 text-white rounded-full font-medium">
            {{ item.text }}
          </span>
        </div>
        <div class="flex items-center justify-between mt-4 pt-3 border-t border-gray-100 text-xs text-gray-500">
          <span>使用 {{ (item.usedCount / 1000).toFixed(1) }}K</span>
          <span class="text-green-600">CTR {{ item.ctr }}%</span>
        </div>
        <button class="w-full mt-3 py-2 text-sm text-blue-600 border border-blue-300 rounded-lg opacity-0 group-hover:opacity-100 transition-opacity hover:bg-blue-50">
          使用此文案
        </button>
      </div>
    </div>

    <div class="flex justify-center">
      <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
    </div>
  </div>
</template>
