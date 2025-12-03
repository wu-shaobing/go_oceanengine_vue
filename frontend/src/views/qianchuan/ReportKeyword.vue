<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川', path: '/qianchuan' }, { name: '数据报表' }, { name: '关键词报表' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">关键词报表</h1>
      <p class="text-gray-600 mt-1">查看搜索关键词的投放效果数据</p>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <input type="date" v-model="filters.startDate" class="border border-gray-300 rounded px-3 py-2">
        <span>至</span>
        <input type="date" v-model="filters.endDate" class="border border-gray-300 rounded px-3 py-2">
        <input type="text" v-model="filters.keyword" placeholder="搜索关键词" class="border border-gray-300 rounded px-3 py-2 w-48">
        <button class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">查询</button>
        <button class="px-4 py-2 border border-gray-300 rounded hover:bg-gray-50">导出</button>
      </div>
    </div>

    <!-- 关键词表格 -->
    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">关键词</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">匹配类型</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">消耗</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">展示</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">点击</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">CTR</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">转化</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">转化成本</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">ROI</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="item in keywords" :key="item.id">
            <td class="px-4 py-3 text-sm font-medium">{{ item.keyword }}</td>
            <td class="px-4 py-3 text-sm">
              <span :class="getMatchClass(item.matchType)" class="px-2 py-1 text-xs rounded">
                {{ item.matchType }}
              </span>
            </td>
            <td class="px-4 py-3 text-sm text-right">¥{{ item.cost.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ item.show.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ item.click.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ item.ctr }}%</td>
            <td class="px-4 py-3 text-sm text-right">{{ item.convert }}</td>
            <td class="px-4 py-3 text-sm text-right">¥{{ item.convertCost }}</td>
            <td class="px-4 py-3 text-sm text-right text-green-600">{{ item.roi }}</td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="200" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const filters = ref({ startDate: '', endDate: '', keyword: '' })

const keywords = ref([
  { id: 1, keyword: '618大促', matchType: '精准匹配', cost: 8560, show: 450000, click: 13800, ctr: 3.07, convert: 286, convertCost: 29.9, roi: 3.5 },
  { id: 2, keyword: '直播间', matchType: '短语匹配', cost: 6850, show: 380000, click: 11200, ctr: 2.95, convert: 215, convertCost: 31.9, roi: 3.2 },
  { id: 3, keyword: '女装新款', matchType: '精准匹配', cost: 5680, show: 320000, click: 9600, ctr: 3.0, convert: 186, convertCost: 30.5, roi: 3.8 },
  { id: 4, keyword: '夏季连衣裙', matchType: '智能匹配', cost: 4560, show: 250000, click: 7500, ctr: 3.0, convert: 145, convertCost: 31.4, roi: 3.1 },
  { id: 5, keyword: '爆款推荐', matchType: '短语匹配', cost: 3850, show: 210000, click: 6300, ctr: 3.0, convert: 118, convertCost: 32.6, roi: 2.9 }
])

const getMatchClass = (type: string) => {
  const classes: Record<string, string> = {
    '精准匹配': 'bg-blue-100 text-blue-800',
    '短语匹配': 'bg-green-100 text-green-800',
    '智能匹配': 'bg-orange-100 text-orange-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}
</script>
