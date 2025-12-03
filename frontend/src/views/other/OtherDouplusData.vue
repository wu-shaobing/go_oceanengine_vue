<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 156 })
const dateRange = ref('7d')

const orders = ref([
  { id: 'DP001', videoTitle: '产品种草视频#1', amount: 500, views: 125600, likes: 3560, comments: 256, shares: 89, roi: 2.5 },
  { id: 'DP002', videoTitle: '品牌故事视频', amount: 1000, views: 268900, likes: 8920, comments: 658, shares: 234, roi: 3.2 },
  { id: 'DP003', videoTitle: '用户评测视频', amount: 300, views: 68500, likes: 1890, comments: 125, shares: 45, roi: 2.1 },
  { id: 'DP004', videoTitle: '限时优惠活动', amount: 800, views: 189600, likes: 5620, comments: 389, shares: 156, roi: 2.8 }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '其他' }, { name: 'DOU+数据' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">DOU+数据分析</h1>
          <p class="mt-2 text-gray-600">分析DOU+投放效果数据</p>
        </div>
        <select v-model="dateRange" class="px-4 py-2 border border-gray-300 rounded-lg">
          <option value="7d">最近7天</option>
          <option value="30d">最近30天</option>
          <option value="90d">最近90天</option>
        </select>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-5 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总投放</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">¥15,680</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总播放</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">652.6万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总点赞</p>
        <p class="text-2xl font-bold text-red-500 mt-1">19.9万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总评论</p>
        <p class="text-2xl font-bold text-green-600 mt-1">1.4万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均ROI</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">2.65</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">投放效果趋势</h3>
      <div class="h-64 flex items-center justify-center bg-gray-50 rounded-lg">
        <p class="text-gray-500">效果趋势图表区域</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">投放明细</h3>
      </div>
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">视频标题</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">投放金额</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">播放量</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">点赞</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">评论</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">分享</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">ROI</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="order in orders" :key="order.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ order.videoTitle }}</div>
              <div class="text-xs text-gray-500">{{ order.id }}</div>
            </td>
            <td class="px-6 py-4 text-sm font-medium text-gray-900">¥{{ order.amount }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ (order.views / 10000).toFixed(1) }}万</td>
            <td class="px-6 py-4 text-sm text-red-500">{{ order.likes.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ order.comments }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ order.shares }}</td>
            <td class="px-6 py-4">
              <span :class="['text-sm font-medium', order.roi >= 2.5 ? 'text-green-600' : 'text-yellow-600']">
                {{ order.roi }}
              </span>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
