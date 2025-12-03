<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const form = reactive({
  seedKeyword: '',
  industry: '',
  region: ''
})

const suggestions = ref<Array<{
  keyword: string
  searchVolume: number
  competition: string
  suggestedBid: number
  selected: boolean
}>>([])

const handleQuery = () => {
  suggestions.value = [
    { keyword: '网络营销培训', searchVolume: 15000, competition: '高', suggestedBid: 8.5, selected: false },
    { keyword: '数字营销课程', searchVolume: 12000, competition: '中', suggestedBid: 6.2, selected: false },
    { keyword: '在线广告投放', searchVolume: 9800, competition: '高', suggestedBid: 9.8, selected: false },
    { keyword: '互联网推广方案', searchVolume: 7500, competition: '低', suggestedBid: 4.5, selected: false },
    { keyword: '品牌营销策略', searchVolume: 6200, competition: '中', suggestedBid: 5.8, selected: false }
  ]
}

const selectedCount = () => suggestions.value.filter(s => s.selected).length
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '关键词工具' }, { name: '推荐关键词' }]" />
      <h1 class="text-3xl font-bold text-gray-900">推荐关键词</h1>
      <p class="mt-2 text-gray-600">基于种子词获取相关关键词推荐</p>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <form @submit.prevent="handleQuery" class="flex flex-wrap gap-4">
        <div class="flex-1 min-w-[200px]">
          <input v-model="form.seedKeyword" type="text" placeholder="输入种子关键词"
                 class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
        </div>
        <select v-model="form.industry" class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          <option value="">选择行业</option>
          <option value="100">电商零售</option>
          <option value="200">游戏</option>
          <option value="300">教育培训</option>
        </select>
        <select v-model="form.region" class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          <option value="">选择地域</option>
          <option value="110000">北京</option>
          <option value="310000">上海</option>
          <option value="440000">广东</option>
        </select>
        <button type="submit" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          获取推荐
        </button>
      </form>
    </div>

    <div v-if="suggestions.length" class="bg-white rounded-lg border border-gray-200">
      <div class="p-4 border-b border-gray-200 flex items-center justify-between">
        <h3 class="font-semibold text-gray-900">推荐结果 ({{ suggestions.length }})</h3>
        <div class="flex items-center gap-4">
          <span class="text-sm text-gray-500">已选 {{ selectedCount() }} 个</span>
          <button v-if="selectedCount() > 0" class="px-4 py-1.5 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700">
            添加到计划
          </button>
        </div>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-left">
                <input type="checkbox" class="w-4 h-4 rounded border-gray-300">
              </th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">关键词</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">搜索量</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">竞争度</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">建议出价</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="item in suggestions" :key="item.keyword" class="hover:bg-gray-50">
              <td class="px-6 py-4">
                <input v-model="item.selected" type="checkbox" class="w-4 h-4 rounded border-gray-300">
              </td>
              <td class="px-6 py-4 text-sm font-medium text-gray-900">{{ item.keyword }}</td>
              <td class="px-6 py-4 text-sm text-gray-700">{{ item.searchVolume.toLocaleString() }}</td>
              <td class="px-6 py-4">
                <span :class="['px-2 py-1 rounded text-xs font-medium',
                  item.competition === '高' ? 'bg-red-100 text-red-700' :
                  item.competition === '中' ? 'bg-yellow-100 text-yellow-700' : 'bg-green-100 text-green-700']">
                  {{ item.competition }}
                </span>
              </td>
              <td class="px-6 py-4 text-sm text-gray-900">¥{{ item.suggestedBid.toFixed(2) }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <div v-else class="bg-white rounded-lg border border-gray-200 p-12 text-center text-gray-500">
      输入种子关键词后点击"获取推荐"
    </div>
  </div>
</template>
