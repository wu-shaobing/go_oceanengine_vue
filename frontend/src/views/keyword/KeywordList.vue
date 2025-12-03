<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 156 })

const keywords = ref([
  { id: 'KW001', keyword: '智能手表', matchType: 'phrase', bid: 2.5, impressions: 125600, clicks: 3256, ctr: 2.59, status: 'active' },
  { id: 'KW002', keyword: '运动手环', matchType: 'broad', bid: 1.8, impressions: 89000, clicks: 2134, ctr: 2.40, status: 'active' },
  { id: 'KW003', keyword: '健康监测', matchType: 'exact', bid: 3.2, impressions: 56000, clicks: 1890, ctr: 3.38, status: 'active' },
  { id: 'KW004', keyword: '智能穿戴', matchType: 'phrase', bid: 2.0, impressions: 45000, clicks: 1234, ctr: 2.74, status: 'paused' }
])

const getMatchTypeLabel = (type: string) => {
  const types: Record<string, string> = { exact: '精确', phrase: '短语', broad: '广泛' }
  return types[type] || type
}

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '关键词' }, { name: '关键词列表' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">关键词管理</h1>
          <p class="mt-2 text-gray-600">管理广告投放关键词</p>
        </div>
        <router-link to="/keyword/create" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          添加关键词
        </router-link>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总关键词</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">启用中</p>
        <p class="text-2xl font-bold text-green-600 mt-1">128</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总展示</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">315.6K</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均CTR</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">2.78%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">关键词</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">匹配类型</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">出价</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">展示</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">点击</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">CTR</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="kw in keywords" :key="kw.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ kw.keyword }}</div>
              <div class="text-xs text-gray-500">{{ kw.id }}</div>
            </td>
            <td class="px-6 py-4">
              <span class="px-2 py-1 bg-blue-100 text-blue-700 rounded text-xs">{{ getMatchTypeLabel(kw.matchType) }}</span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">¥{{ kw.bid.toFixed(2) }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ kw.impressions.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ kw.clicks.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ kw.ctr }}%</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                kw.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                {{ kw.status === 'active' ? '启用' : '暂停' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <router-link :to="`/keyword/edit/${kw.id}`" class="text-blue-600 hover:text-blue-800 mr-3">编辑</router-link>
              <button class="text-gray-600 hover:text-gray-800">删除</button>
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
