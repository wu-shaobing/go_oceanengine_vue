<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 28 })

const converts = ref([
  { id: 'CV001', name: 'APP下载', type: 'download', adsCount: 45, status: 'active', createdAt: '2025-10-01' },
  { id: 'CV002', name: '表单提交', type: 'form', adsCount: 32, status: 'active', createdAt: '2025-10-05' },
  { id: 'CV003', name: '商品购买', type: 'purchase', adsCount: 28, status: 'active', createdAt: '2025-10-08' },
  { id: 'CV004', name: '电话咨询', type: 'call', adsCount: 15, status: 'active', createdAt: '2025-10-10' },
  { id: 'CV005', name: '页面浏览', type: 'pageview', adsCount: 0, status: 'inactive', createdAt: '2025-10-15' }
])

const getTypeLabel = (type: string) => {
  const types: Record<string, string> = {
    download: '下载',
    form: '表单',
    purchase: '购买',
    call: '电话',
    pageview: '浏览'
  }
  return types[type] || type
}

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '转化工具' }, { name: '转化目标' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">转化目标列表</h1>
          <p class="mt-2 text-gray-600">管理广告投放的转化目标</p>
        </div>
        <router-link to="/tools/adconvert/create" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建目标
        </router-link>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总目标</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">启用中</p>
        <p class="text-2xl font-bold text-green-600 mt-1">24</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">关联计划</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">120</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">今日转化</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">1,256</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-200">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">转化目标</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">类型</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">关联计划</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">创建时间</th>
              <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="item in converts" :key="item.id" class="hover:bg-gray-50">
              <td class="px-6 py-4">
                <div class="text-sm font-medium text-gray-900">{{ item.name }}</div>
                <div class="text-xs text-gray-500">ID: {{ item.id }}</div>
              </td>
              <td class="px-6 py-4">
                <span class="px-2 py-1 bg-blue-100 text-blue-700 rounded text-xs">{{ getTypeLabel(item.type) }}</span>
              </td>
              <td class="px-6 py-4 text-sm text-gray-900">{{ item.adsCount }} 个</td>
              <td class="px-6 py-4">
                <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                  item.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                  {{ item.status === 'active' ? '启用' : '停用' }}
                </span>
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">{{ item.createdAt }}</td>
              <td class="px-6 py-4 text-sm">
                <button class="text-blue-600 hover:text-blue-800 mr-3">编辑</button>
                <button class="text-gray-600 hover:text-gray-800">详情</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="px-6 py-4 border-t border-gray-200">
        <Pagination :current="pagination.page" :total="pagination.total" :page-size="pagination.pageSize" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>
