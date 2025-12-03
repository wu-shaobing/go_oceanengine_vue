<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 156 })

const excludeKeywords = ref([
  { id: 'EK001', keyword: '免费', matchType: 'exact', level: 'account', campaigns: 12, createdAt: '2025-11-15' },
  { id: 'EK002', keyword: '便宜', matchType: 'phrase', level: 'campaign', campaigns: 5, createdAt: '2025-11-18' },
  { id: 'EK003', keyword: '二手', matchType: 'exact', level: 'account', campaigns: 12, createdAt: '2025-11-20' },
  { id: 'EK004', keyword: '租赁', matchType: 'phrase', level: 'adgroup', campaigns: 3, createdAt: '2025-11-22' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '关键词' }, { name: '否定关键词' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">否定关键词</h1>
          <p class="mt-2 text-gray-600">管理广告否定关键词，排除无效流量</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          添加否定词
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">否定词总数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">账户级别</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">45</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">计划级别</p>
        <p class="text-2xl font-bold text-green-600 mt-1">78</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">单元级别</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">33</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">关键词</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">匹配方式</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">级别</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">关联计划</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">添加时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="kw in excludeKeywords" :key="kw.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <span class="text-sm font-medium text-gray-900">{{ kw.keyword }}</span>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded text-xs',
                     kw.matchType === 'exact' ? 'bg-blue-100 text-blue-700' : 'bg-green-100 text-green-700']">
                {{ kw.matchType === 'exact' ? '精确' : '短语' }}
              </span>
            </td>
            <td class="px-6 py-4">
              <span class="px-2 py-1 bg-gray-100 text-gray-700 rounded text-xs">
                {{ kw.level === 'account' ? '账户' : kw.level === 'campaign' ? '计划' : '单元' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ kw.campaigns }} 个</td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ kw.createdAt }}</td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">编辑</button>
              <button class="text-red-600 hover:text-red-800">删除</button>
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
