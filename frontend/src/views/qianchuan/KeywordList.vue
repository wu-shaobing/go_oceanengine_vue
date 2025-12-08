<template>
  <div class="p-6">
    <Breadcrumb :items="[{ name: '千川电商广告', path: '/qianchuan' }, { name: '工具' }, { name: '关键词管理' }]" />
    
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">关键词管理</h1>
        <p class="text-gray-600 mt-1">管理搜索广告关键词</p>
      </div>
      <div class="flex space-x-3">
        <router-link to="/qianchuan/keyword/recommend" class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">
          关键词推荐
        </router-link>
<button @click="handleAddKeyword" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          添加关键词
        </button>
      </div>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="flex flex-wrap gap-4">
        <input type="text" v-model="filters.keyword" placeholder="搜索关键词" class="border border-gray-300 rounded-lg px-3 py-2 w-48">
        <select v-model="filters.matchType" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部匹配方式</option>
          <option value="exact">精确匹配</option>
          <option value="phrase">短语匹配</option>
          <option value="broad">广泛匹配</option>
        </select>
        <select v-model="filters.status" class="border border-gray-300 rounded-lg px-3 py-2">
          <option value="">全部状态</option>
          <option value="enable">已启用</option>
          <option value="disable">已暂停</option>
        </select>
<button @click="handleSearch" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">搜索</button>
      </div>
    </div>

    <!-- 关键词列表 -->
    <div class="bg-white rounded-lg shadow">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left"><input type="checkbox" class="rounded"></th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">关键词</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">匹配方式</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">状态</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">出价</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">消耗</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">曝光</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">点击</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-500">转化</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-500">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="kw in keywords" :key="kw.id">
            <td class="px-4 py-3"><input type="checkbox" class="rounded"></td>
            <td class="px-4 py-3">
              <div class="font-medium">{{ kw.word }}</div>
              <div class="text-sm text-gray-500">质量分: {{ kw.qualityScore }}</div>
            </td>
            <td class="px-4 py-3 text-sm">{{ kw.matchType }}</td>
            <td class="px-4 py-3">
              <span :class="kw.status === 'enable' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'" 
                class="px-2 py-1 text-xs rounded-full">
                {{ kw.status === 'enable' ? '已启用' : '已暂停' }}
              </span>
            </td>
            <td class="px-4 py-3 text-sm text-right">¥{{ kw.bid }}</td>
            <td class="px-4 py-3 text-sm text-right">¥{{ kw.cost.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ kw.impression.toLocaleString() }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ kw.click }}</td>
            <td class="px-4 py-3 text-sm text-right">{{ kw.conversion }}</td>
            <td class="px-4 py-3">
              <div class="flex space-x-2">
<button @click="handleEdit(kw)" class="text-blue-600 hover:text-blue-800 text-sm">编辑</button>
                <button v-if="kw.status === 'enable'" @click="handlePause(kw)" class="text-orange-600 hover:text-orange-800 text-sm">暂停</button>
                <button v-else @click="handleEnable(kw)" class="text-green-600 hover:text-green-800 text-sm">启用</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="p-4 border-t">
        <Pagination :current="1" :total="100" :page-size="10" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const filters = ref({
  keyword: '',
  matchType: '',
  status: ''
})

const keywords = ref([
  { id: 1, word: '美白面霜', matchType: '精确匹配', status: 'enable', qualityScore: 9, bid: 2.5, cost: 1560, impression: 25600, click: 856, conversion: 45 },
  { id: 2, word: '保湿护肤品', matchType: '短语匹配', status: 'enable', qualityScore: 8, bid: 2.0, cost: 1280, impression: 32500, click: 1024, conversion: 38 },
  { id: 3, word: '抗衰老精华', matchType: '精确匹配', status: 'enable', qualityScore: 9, bid: 3.0, cost: 2100, impression: 18600, click: 686, conversion: 52 },
  { id: 4, word: '夏季连衣裙', matchType: '广泛匹配', status: 'disable', qualityScore: 7, bid: 1.8, cost: 860, impression: 15800, click: 456, conversion: 22 },
  { id: 5, word: '无线耳机', matchType: '短语匹配', status: 'enable', qualityScore: 8, bid: 2.2, cost: 1680, impression: 28900, click: 768, conversion: 35 }
])

const handleAddKeyword = () => {
  alert('添加关键词')
}

const handleSearch = () => {
  alert('搜索关键词')
}

const handleEdit = (kw: typeof keywords.value[0]) => {
  alert(`编辑关键词: ${kw.word}`)
}

const handlePause = (kw: typeof keywords.value[0]) => {
  kw.status = 'disable'
  alert(`已暂停: ${kw.word}`)
}

const handleEnable = (kw: typeof keywords.value[0]) => {
  kw.status = 'enable'
  alert(`已启用: ${kw.word}`)
}
</script>
