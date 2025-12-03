<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 68 })

const audiences = ref([
  { id: 'AU001', name: '高价值用户包', source: '自定义上传', coverage: 1250000, matchRate: 85, usageCount: 45, status: 'ready' },
  { id: 'AU002', name: '历史购买用户', source: 'CRM同步', coverage: 850000, matchRate: 92, usageCount: 38, status: 'ready' },
  { id: 'AU003', name: '潜在兴趣用户', source: 'Lookalike扩展', coverage: 5800000, matchRate: 78, usageCount: 23, status: 'ready' },
  { id: 'AU004', name: '流失用户召回', source: '自定义上传', coverage: 320000, matchRate: 88, usageCount: 12, status: 'processing' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '工具' }, { name: '人群包管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">人群包管理</h1>
          <p class="mt-2 text-gray-600">创建和管理自定义人群包</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建人群包
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">人群包总数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">可用人群包</p>
        <p class="text-2xl font-bold text-green-600 mt-1">62</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总覆盖人数</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">822万</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均匹配率</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">85.8%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200 p-4">
      <div class="flex gap-4">
        <select class="px-4 py-2 border border-gray-300 rounded-lg">
          <option value="">全部来源</option>
          <option value="upload">自定义上传</option>
          <option value="crm">CRM同步</option>
          <option value="lookalike">Lookalike扩展</option>
        </select>
        <select class="px-4 py-2 border border-gray-300 rounded-lg">
          <option value="">全部状态</option>
          <option value="ready">可用</option>
          <option value="processing">处理中</option>
          <option value="expired">已过期</option>
        </select>
        <input type="text" placeholder="搜索人群包..." class="flex-1 px-4 py-2 border border-gray-300 rounded-lg">
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">人群包名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">来源</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">覆盖人数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">匹配率</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">使用次数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="audience in audiences" :key="audience.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ audience.name }}</div>
              <div class="text-xs text-gray-500">{{ audience.id }}</div>
            </td>
            <td class="px-6 py-4">
              <span class="px-2 py-1 bg-gray-100 text-gray-700 rounded text-xs">{{ audience.source }}</span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ (audience.coverage / 10000).toFixed(0) }}万</td>
            <td class="px-6 py-4">
              <div class="flex items-center gap-2">
                <div class="w-16 h-2 bg-gray-200 rounded-full">
                  <div class="h-full bg-blue-500 rounded-full" :style="{ width: `${audience.matchRate}%` }"></div>
                </div>
                <span class="text-sm text-gray-600">{{ audience.matchRate }}%</span>
              </div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-600">{{ audience.usageCount }}</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     audience.status === 'ready' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800']">
                {{ audience.status === 'ready' ? '可用' : '处理中' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">应用</button>
              <button class="text-gray-600 hover:text-gray-800">详情</button>
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
