<script setup lang="ts">
import { ref, reactive } from 'vue'
import Breadcrumb from '@/components/common/Breadcrumb.vue'
import Pagination from '@/components/common/Pagination.vue'

const pagination = reactive({ page: 1, pageSize: 10, total: 24 })

const groups = ref([
  { id: 'GRP001', name: '双11活动专题', sitesCount: 12, totalPV: 156000, conversionRate: 5.2, status: 'active', createdAt: '2025-10-15' },
  { id: 'GRP002', name: '品牌旗舰店', sitesCount: 8, totalPV: 98000, conversionRate: 4.8, status: 'active', createdAt: '2025-09-20' },
  { id: 'GRP003', name: '新品上市', sitesCount: 5, totalPV: 45000, conversionRate: 3.6, status: 'active', createdAt: '2025-10-28' },
  { id: 'GRP004', name: '测试分组', sitesCount: 2, totalPV: 1200, conversionRate: 0, status: 'draft', createdAt: '2025-11-10' }
])

const formatNumber = (num: number) => {
  if (num >= 10000) return (num / 10000).toFixed(1) + '万'
  return num.toLocaleString()
}

const handlePageChange = (page: number) => {
  pagination.page = page
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <Breadcrumb :items="[{ name: '落地页' }, { name: '站点分组' }]" />
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">站点分组</h1>
          <p class="mt-2 text-gray-600">将落地页按活动或主题进行分组管理</p>
        </div>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
          创建分组
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总分组</p>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ pagination.total }}</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">启用中</p>
        <p class="text-2xl font-bold text-green-600 mt-1">18</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">总站点</p>
        <p class="text-2xl font-bold text-blue-600 mt-1">86</p>
      </div>
      <div class="bg-white rounded-lg border border-gray-200 p-4">
        <p class="text-sm text-gray-500">平均转化率</p>
        <p class="text-2xl font-bold text-purple-600 mt-1">4.5%</p>
      </div>
    </div>

    <div class="bg-white rounded-lg border border-gray-200">
      <table class="w-full">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">分组名称</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">站点数</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">总PV</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">转化率</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">状态</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">创建时间</th>
            <th class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase">操作</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="group in groups" :key="group.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">
              <div class="text-sm font-medium text-gray-900">{{ group.name }}</div>
              <div class="text-xs text-gray-500">{{ group.id }}</div>
            </td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ group.sitesCount }} 个</td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ formatNumber(group.totalPV) }}</td>
            <td class="px-6 py-4">
              <span :class="['text-sm font-medium', group.conversionRate >= 4 ? 'text-green-600' : 'text-gray-900']">
                {{ group.conversionRate }}%
              </span>
            </td>
            <td class="px-6 py-4">
              <span :class="['px-2 py-1 rounded-full text-xs font-medium',
                group.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800']">
                {{ group.status === 'active' ? '启用' : '草稿' }}
              </span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ group.createdAt }}</td>
            <td class="px-6 py-4 text-sm">
              <button class="text-blue-600 hover:text-blue-800 mr-3">管理</button>
              <button class="text-gray-600 hover:text-gray-800">编辑</button>
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
