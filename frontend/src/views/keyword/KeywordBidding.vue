<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 256 })

const biddings = ref([
  { id: 'KB001', keyword: '智能手机', currentBid: 5.5, suggestBid: 6.2, competitionLevel: 'high', impressions: 125600, clicks: 3560, ctr: 2.83 },
  { id: 'KB002', keyword: '手机壳', currentBid: 2.0, suggestBid: 2.5, competitionLevel: 'medium', impressions: 89500, clicks: 2680, ctr: 2.99 },
  { id: 'KB003', keyword: '充电器', currentBid: 1.8, suggestBid: 2.0, competitionLevel: 'low', impressions: 56800, clicks: 1890, ctr: 3.33 },
  { id: 'KB004', keyword: '耳机', currentBid: 3.5, suggestBid: 4.0, competitionLevel: 'high', impressions: 98600, clicks: 2960, ctr: 3.00 }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '关键词' }, { name: '关键词出价' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">关键词出价管理</h1>
          <p class="mt-2 text-gray-600">优化关键词出价，提升广告效果</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          批量调价
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">关键词总数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均出价</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">¥3.2</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">低于建议价</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">45</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">高竞争词</p>
        <p class="text-2xl font-bold text-red-600 mt-1">68</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">关键词</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">当前出价</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">建议出价</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">竞争程度</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">展示</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">点击</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">CTR</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="bid in biddings" :key="bid.id" class="hover:bg-gray-50">
            <td class="px-6 py-4 text-sm font-medium text-gray-900">{{ bid.keyword }}</td>
            <td class="px-6 py-4">
              <span class="text-sm" :class="bid.currentBid < bid.suggestBid ? 'text-yellow-600' : 'text-green-600'">
                ¥{{ bid.currentBid }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">¥{{ bid.suggestBid }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded text-xs',
                     bid.competitionLevel === 'high' ? 'bg-red-100 text-red-700' :
                     bid.competitionLevel === 'medium' ? 'bg-yellow-100 text-yellow-700' :
                     'bg-green-100 text-green-700']">
                {{ bid.competitionLevel === 'high' ? '高' : bid.competitionLevel === 'medium' ? '中' : '低' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ bid.impressions.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ bid.clicks.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm" :class="bid.ctr >= 3 ? 'text-green-600' : 'text-yellow-600'">{{ bid.ctr }}%</td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800">调价</button>
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
