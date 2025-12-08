<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川', path: '/qianchuan' }, { name: '工具' }, { name: '关键词推荐' }]" />
    
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">关键词推荐</h1>
      <p class="text-gray-600 mt-1">获取智能推荐的投放关键词</p>
    </div>

    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4 items-center">
        <input type="text" v-model="seedWord" placeholder="输入种子词获取推荐" class="border border-gray-300 rounded px-3 py-2 w-64">
        <select v-model="filters.category" class="border border-gray-300 rounded px-3 py-2">
          <option value="">选择品类</option>
          <option value="fashion">服饰鞋包</option>
          <option value="beauty">美妆护肤</option>
          <option value="food">食品饮料</option>
        </select>
<button class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700" @click="handleGetRecommend">获取推荐</button>
      </div>
    </div>

    <div class="bg-white rounded-lg shadow">
      <div class="p-4 border-b flex justify-between items-center">
        <h3 class="font-medium">推荐关键词</h3>
<button class="text-blue-600 text-sm" @click="handleBatchAdd">批量添加</button>
      </div>
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 w-10"><input type="checkbox" class="rounded"></th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">关键词</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">搜索热度</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">竞争度</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">建议出价</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="kw in keywords" :key="kw.word">
            <td class="px-4 py-3"><input type="checkbox" class="rounded"></td>
            <td class="px-4 py-3 font-medium">{{ kw.word }}</td>
            <td class="px-4 py-3 text-sm text-right">
              <div class="flex items-center justify-end">
                <div class="w-20 bg-gray-200 rounded-full h-2 mr-2">
                  <div class="bg-orange-500 h-2 rounded-full" :style="{ width: kw.heat + '%' }"></div>
                </div>
                {{ kw.heat }}
              </div>
            </td>
            <td class="px-4 py-3 text-sm text-right">
              <span :class="kw.competition === '高' ? 'text-red-600' : kw.competition === '中' ? 'text-orange-600' : 'text-green-600'">{{ kw.competition }}</span>
            </td>
            <td class="px-4 py-3 text-sm text-right">¥{{ kw.bidPrice }}</td>
            <td class="px-4 py-3">
<button class="text-blue-600 hover:text-blue-800 text-sm" @click="handleAddKeyword(kw)">添加</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'

const seedWord = ref('')
const filters = ref({ category: '' })

const keywords = ref([
  { word: '618大促', heat: 95, competition: '高', bidPrice: 2.5 },
  { word: '直播间优惠', heat: 88, competition: '中', bidPrice: 1.8 },
  { word: '夏季新款', heat: 82, competition: '中', bidPrice: 1.5 },
  { word: '爆款推荐', heat: 78, competition: '低', bidPrice: 1.2 },
  { word: '限时折扣', heat: 75, competition: '中', bidPrice: 1.6 },
  { word: '品牌特卖', heat: 70, competition: '低', bidPrice: 1.0 }
])

const handleGetRecommend = () => {
  alert(`获取推荐: ${seedWord.value}`)
}

const handleBatchAdd = () => {
  alert('批量添加关键词')
}

const handleAddKeyword = (kw: typeof keywords.value[0]) => {
  alert(`添加关键词: ${kw.word}`)
}
</script>
