<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 45 })

const pages = ref([
  { id: 'LP001', name: '双十一促销页', url: 'https://example.com/sale', pv: 125600, bounce: 32, conversion: 5.2, status: 'active' },
  { id: 'LP002', name: '新品发布页', url: 'https://example.com/new', pv: 89500, bounce: 28, conversion: 4.8, status: 'active' },
  { id: 'LP003', name: '品牌介绍页', url: 'https://example.com/brand', pv: 56800, bounce: 45, conversion: 2.1, status: 'inactive' }
])

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '落地页' }, { name: '页面管理' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">落地页管理</h1>
          <p class="mt-2 text-gray-600">管理广告落地页</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建落地页
        </button>
      </div>
    </div>

    <div class="grid grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">落地页数</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日PV</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">271,900</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均跳出率</p>
        <p class="text-2xl font-bold text-yellow-600 mt-1">35%</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均转化率</p>
        <p class="text-2xl font-bold text-green-600 mt-1">4.0%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">页面名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">URL</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">今日PV</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">跳出率</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">转化率</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="page in pages" :key="page.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ page.name }}</div>
              <div class="text-xs text-gray-500">{{ page.id }}</div>
            </td>
            <td class="px-6 py-4">
              <a :href="page.url" target="_blank" class="text-sm text-blue-600 hover:underline truncate block max-w-xs">{{ page.url }}</a>
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ page.pv.toLocaleString() }}</td>
            <td class="px-6 py-4 text-sm" :class="page.bounce > 40 ? 'text-red-600' : 'text-gray-600'">{{ page.bounce }}%</td>
            <td class="px-6 py-4 text-sm font-medium text-green-600">{{ page.conversion }}%</td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                     page.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                {{ page.status === 'active' ? '已上线' : '未上线' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">编辑</button>
              <button class="text-gray-600 hover:text-gray-800 mr-3">预览</button>
              <button class="text-gray-600 hover:text-gray-800">数据</button>
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
