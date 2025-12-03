<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '素材管理' }, { name: '图片素材' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">图片素材</h1>
        <p class="text-gray-600 mt-1">管理图片广告素材</p>
      </div>
      <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
        上传图片
      </button>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4">
        <input type="text" v-model="filters.keyword" placeholder="图片名称" class="border border-gray-300 rounded-lg px-3 py-2 w-48">
        <select v-model="filters.type" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部类型</option>
          <option value="product">商品图</option>
          <option value="banner">Banner</option>
          <option value="poster">海报</option>
        </select>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
      </div>
    </div>

    <!-- 图片网格 -->
    <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4">
      <div v-for="image in images" :key="image.id" class="bg-white rounded-lg shadow overflow-hidden group">
        <div class="aspect-square bg-gray-100 relative">
          <img :src="image.url" alt="" class="w-full h-full object-cover">
          <div class="absolute inset-0 bg-black bg-opacity-0 group-hover:bg-opacity-40 transition-all flex items-center justify-center opacity-0 group-hover:opacity-100">
            <button class="px-3 py-1 bg-white rounded text-sm mr-2">预览</button>
            <button class="px-3 py-1 bg-blue-600 text-white rounded text-sm">使用</button>
          </div>
        </div>
        <div class="p-3">
          <div class="text-sm font-medium truncate">{{ image.name }}</div>
          <div class="text-xs text-gray-500 mt-1">{{ image.resolution }} · {{ image.size }}</div>
        </div>
      </div>
    </div>

    <div class="mt-6">
      <Pagination :current="1" :total="76" :page-size="18" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const filters = ref({
  keyword: '',
  type: ''
})

const images = ref([
  { id: 1, name: '产品展示图1', url: 'https://via.placeholder.com/200', resolution: '800x800', size: '1.2MB' },
  { id: 2, name: '促销海报', url: 'https://via.placeholder.com/200', resolution: '750x1334', size: '2.1MB' },
  { id: 3, name: '618活动图', url: 'https://via.placeholder.com/200', resolution: '1920x1080', size: '1.8MB' },
  { id: 4, name: '商品详情图', url: 'https://via.placeholder.com/200', resolution: '800x1200', size: '1.5MB' },
  { id: 5, name: '品牌Logo', url: 'https://via.placeholder.com/200', resolution: '500x500', size: '0.3MB' },
  { id: 6, name: '产品对比图', url: 'https://via.placeholder.com/200', resolution: '1200x800', size: '1.8MB' },
  { id: 7, name: '直播间封面', url: 'https://via.placeholder.com/200', resolution: '720x1280', size: '0.9MB' },
  { id: 8, name: '优惠券设计', url: 'https://via.placeholder.com/200', resolution: '600x300', size: '0.5MB' },
  { id: 9, name: '新品首发图', url: 'https://via.placeholder.com/200', resolution: '800x800', size: '1.1MB' },
  { id: 10, name: '限时特惠', url: 'https://via.placeholder.com/200', resolution: '750x400', size: '0.7MB' },
  { id: 11, name: '用户好评截图', url: 'https://via.placeholder.com/200', resolution: '720x1280', size: '0.4MB' },
  { id: 12, name: '店铺Banner', url: 'https://via.placeholder.com/200', resolution: '1920x600', size: '2.5MB' }
])
</script>
