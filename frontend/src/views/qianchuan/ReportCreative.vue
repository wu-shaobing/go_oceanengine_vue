<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川', path: '/qianchuan' }, { name: '数据报表' }, { name: '创意报表' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">创意报表</h1>
      <p class="text-gray-600 mt-1">查看创意维度的投放数据</p>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <input type="date" v-model="filters.startDate" class="border border-gray-300 rounded px-3 py-2">
        <span>至</span>
        <input type="date" v-model="filters.endDate" class="border border-gray-300 rounded px-3 py-2">
        <select v-model="filters.type" class="border border-gray-300 rounded px-3 py-2">
          <option value="">全部类型</option>
          <option value="video">视频</option>
          <option value="image">图片</option>
        </select>
        <button @click="handleSearch" class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">查询</button>
        <button @click="handleExport" class="px-4 py-2 border border-gray-300 rounded hover:bg-gray-50">导出</button>
      </div>
    </div>

    <!-- 创意列表 -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <div v-for="item in creatives" :key="item.id" class="bg-white rounded-lg shadow overflow-hidden">
        <div class="relative">
          <img :src="item.cover" class="w-full h-40 object-cover" alt="">
          <span class="absolute top-2 right-2 px-2 py-1 text-xs bg-black bg-opacity-60 text-white rounded">
            {{ item.type === 'video' ? '视频' : '图片' }}
          </span>
        </div>
        <div class="p-4">
          <div class="text-sm text-gray-500 mb-2">ID: {{ item.id }}</div>
          <div class="grid grid-cols-2 gap-2 text-sm">
            <div><span class="text-gray-500">消耗:</span> ¥{{ item.cost.toLocaleString() }}</div>
            <div><span class="text-gray-500">展示:</span> {{ item.show.toLocaleString() }}</div>
            <div><span class="text-gray-500">点击:</span> {{ item.click.toLocaleString() }}</div>
            <div><span class="text-gray-500">CTR:</span> {{ item.ctr }}%</div>
            <div><span class="text-gray-500">转化:</span> {{ item.convert }}</div>
            <div><span class="text-gray-500">转化率:</span> {{ item.cvr }}%</div>
          </div>
        </div>
      </div>
    </div>

    <div class="mt-6">
      <Pagination :current="1" :total="50" :page-size="9" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const filters = ref({ startDate: '', endDate: '', type: '' })
const loading = ref(false)

const handleSearch = async () => {
  loading.value = true
  try {
    await new Promise(r => setTimeout(r, 500))
    alert('数据已刷新')
  } finally {
    loading.value = false
  }
}

const handleExport = () => {
  alert('导出功能开发中...')
}

const creatives = ref([
  { id: 'C001', cover: 'https://via.placeholder.com/300x160', type: 'video', cost: 15680, show: 850000, click: 25800, ctr: 3.04, convert: 586, cvr: 2.27 },
  { id: 'C002', cover: 'https://via.placeholder.com/300x160', type: 'video', cost: 12560, show: 680000, click: 19800, ctr: 2.91, convert: 425, cvr: 2.15 },
  { id: 'C003', cover: 'https://via.placeholder.com/300x160', type: 'image', cost: 8960, show: 520000, click: 15200, ctr: 2.92, convert: 286, cvr: 1.88 },
  { id: 'C004', cover: 'https://via.placeholder.com/300x160', type: 'video', cost: 10250, show: 560000, click: 16800, ctr: 3.0, convert: 356, cvr: 2.12 },
  { id: 'C005', cover: 'https://via.placeholder.com/300x160', type: 'image', cost: 6580, show: 380000, click: 11200, ctr: 2.95, convert: 198, cvr: 1.77 },
  { id: 'C006', cover: 'https://via.placeholder.com/300x160', type: 'video', cost: 9860, show: 520000, click: 15600, ctr: 3.0, convert: 328, cvr: 2.1 }
])
</script>
